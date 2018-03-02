package jobs

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/jobs/ast"
	"github.com/mperham/inspeqtor/jobs/lexer"
	"github.com/mperham/inspeqtor/jobs/parser"
	"github.com/mperham/inspeqtor/util"
)

// Job is a recurring job which we want to verify executes
// every interval.
type Job struct {
	JobName    string
	Interval   time.Duration
	Parameters map[string]string
	LastRun    time.Time
	alerter    inspeqtor.Action
	state      inspeqtor.RuleState
}

var (
	JobOverdue inspeqtor.EventType = "JobOverdue"
	JobRan     inspeqtor.EventType = "JobRan"

	// global with our parsed jobs so the stateless command handler can find
	// the right job.
	jobs        = map[string]*Job{}
	runNotifier chan string
)

func init() {
	inspeqtor.Events = append(inspeqtor.Events, JobOverdue)
	inspeqtor.Events = append(inspeqtor.Events, JobRan)
	runNotifier = make(chan string, 1)
}

func Parse(global *inspeqtor.ConfigFile, confDir string) (map[string]*Job, error) {
	inspeqtor.CommandHandlers["job_done"] = jobDone

	parsedJobs, err := parseJobs(global, confDir)
	if err != nil {
		return nil, err
	}
	if len(parsedJobs) == 0 {
		return nil, nil
	}
	jobs = parsedJobs

	util.Info("Watching for %d recurring jobs", len(parsedJobs))
	return parsedJobs, nil
}

func Watch(i *inspeqtor.Inspeqtor, jobs map[string]*Job) {
	util.Debug("Starting recurring job watcher")
	go func() {
		for {
			untilNext := check(jobs)
			select {
			case <-i.Stopping:
				// reloading inspeqtor
				util.Debug("Shutting down recurring job watcher")
				return
			case <-runNotifier:
				// we just got notified a job ran,
				// verify we don't need to fire JobRan
			case <-time.After(untilNext + time.Minute):
				// a job is due at this point in time.
				// add an extra minute to allow for race conditions
				// and slow performance
			}
		}
	}()
}

func New(name string, interval time.Duration, params map[string]string) *Job {
	return &Job{name, interval, params, time.Now(), nil, inspeqtor.Ok}
}

func (j *Job) Name() string {
	return j.JobName
}

func (j *Job) Parameter(key string) string {
	return j.Parameters[key]
}

func (j *Job) alert(et inspeqtor.EventType) error {
	return j.alerter.Trigger(&inspeqtor.Event{Type: et, Eventable: j, Rule: nil})
}

func check(jobs map[string]*Job) time.Duration {
	min := time.Hour

	for _, j := range jobs {
		now := time.Now()
		due := j.LastRun.Add(j.Interval)
		if due.After(now) && min > due.Sub(now) {
			// calculate the delay time until the next job check
			min = due.Sub(now)
		}

		if due.Before(now) && j.state == inspeqtor.Ok {
			util.Warn("Recurring job \"%s\" is overdue", j.JobName)
			j.state = inspeqtor.Triggered
			err := j.alert(JobOverdue)
			if err != nil {
				util.Warn(fmt.Sprintf("Error firing cron job alert: %s", err.Error()))
			}
		}
		if !due.Before(now) && j.state == inspeqtor.Triggered {
			util.Info("Recurring job \"%s\" has recovered", j.JobName)
			err := j.alert(JobRan)
			if err != nil {
				util.Warn(fmt.Sprintf("Error firing cron job alert: %s", err.Error()))
			}
			j.state = inspeqtor.Ok
		}
	}
	return min
}

func jobDone(_ *inspeqtor.Inspeqtor, args []string, resp io.Writer) {
	if len(args) != 1 {
		io.WriteString(resp, "Usage: job_done <name>\n")
		return
	}

	job := jobs[args[0]]
	if job == nil {
		io.WriteString(resp, fmt.Sprintf("Error: no such job \"%s\", verify spelling or reload Inspeqtor\n", args[0]))
		return
	}

	job.LastRun = time.Now()
	runNotifier <- args[0]
	io.WriteString(resp, "OK\n")
}

func parseJobs(global *inspeqtor.ConfigFile, confDir string) (map[string]*Job, error) {
	util.Debug("Parsing jobs in " + confDir)
	files, err := filepath.Glob(confDir + "/jobs.d/*.inq")
	if err != nil {
		return nil, err
	}

	jobs := map[string]*Job{}

	for _, filename := range files {
		util.DebugDebug("Parsing " + filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		s := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		obj, err := p.Parse(s)
		if err != nil {
			util.Warn("Unable to parse " + filename + ": " + err.Error())
			continue
		}

		astcontent := obj.(*ast.Content)
		for _, astjob := range astcontent.Jobs {
			if _, ok := jobs[astjob.Name]; ok {
				return nil, fmt.Errorf("Duplicate job %s", astjob.Name)
			}

			j := New(astjob.Name, astjob.Interval, astcontent.Parameters)

			owner := j.Parameters["owner"]
			route := global.AlertRoutes[owner]
			if owner == "" && route == nil {
				return nil, fmt.Errorf("No default alert route configured!")
			}
			if route == nil {
				return nil, fmt.Errorf("No such alert route: %s", owner)
			}
			alert, err := inspeqtor.Actions["alert"](j, route)
			if err != nil {
				return nil, err
			}
			j.alerter = alert
			jobs[astjob.Name] = j
		}
	}

	return jobs, nil
}
