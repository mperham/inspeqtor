package daemon

import (
	"io/ioutil"
	"os/exec"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/util"
)

/*
 Daemon-specific metrics are process-specific metrics which can be pushed into Inspeqtor
 and used as rule checks.

 check service redis with hostname [localhost], port [6379], socket [nil], password [nil]
   if redis(latest_fork_usec) > 10000 then alert ops
*/

var (
	c    = metrics.Counter
	g    = metrics.Gauge
	inMB = metrics.DisplayInMB
)

type executor func(string, []string, []byte) ([]byte, error)

/*
func (s *Store) Collect(pid int) error {
	err := s.Store.Collect(pid)
	if err != nil {
		return err
	}
	for _, ds := range s.DaemonSpecific {
		util.Debug("Collecting %s metrics", ds.Name())
		hash, err := ds.Capture()
		if err != nil {
			return err
		}
		for k, v := range hash {
			s.Store.Save(ds.Name(), k, v)
		}
	}
	return nil
}
*/

func execCmd(command string, args []string, stdin []byte) ([]byte, error) {
	cmd := exec.Command(command, args...)
	if stdin != nil {
		in, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		_, err = in.Write(stdin)
		in.Close()
		if err != nil {
			return nil, err
		}
	}

	return util.SafeRun(cmd)
}

func testExec(path string) func(string, []string, []byte) ([]byte, error) {
	return func(command string, args []string, stdin []byte) ([]byte, error) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}
