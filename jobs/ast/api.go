package ast

import (
	"strconv"
	"time"

	"github.com/mperham/inspeqtor/jobs/token"
)

type Job struct {
	Name     string
	Interval time.Duration
}

type Content struct {
	Jobs       []Job
	Parameters map[string]string
}

func NewJobCheck(j interface{}, params interface{}) *Content {
	jobs := j.([]Job)
	hash := params.(map[string]string)
	return &Content{jobs, hash}
}

func AddParam(key interface{}, val interface{}, hash interface{}) (map[string]string, error) {
	k := string(key.(*token.Token).Lit)
	v := string(val.(*token.Token).Lit)

	// remove quotes from quoted strings
	if v[0] == '"' {
		val, err := strconv.Unquote(v)
		if err != nil {
			return nil, err
		}
		v = val
	}

	var h map[string]string

	if hash == nil {
		h = map[string]string{}
	} else {
		h = hash.(map[string]string)
	}
	h[k] = v

	return h, nil
}

func NewJobList(job interface{}) []Job {
	return []Job{job.(Job)} // Job
}

func AppendJob(list interface{}, rule interface{}) []Job {
	return append(list.([]Job), rule.(Job))
}

func NewJob(name interface{}, amount interface{}, period interface{}) Job {
	return Job{
		string(name.(*token.Token).Lit),
		ToSecs(amount.(int64), string(period.(*token.Token).Lit)),
	}
}

func ToSecs(amt int64, period string) time.Duration {
	var pd time.Duration

	switch period {
	case "day", "days":
		pd = 24 * 60 * 60 * time.Second
	case "hour", "hours":
		pd = time.Hour
	case "minute", "minutes":
		pd = time.Minute
	}

	return time.Duration(amt) * pd
}

func ToInt64(v interface{}) (int64, error) {
	raw := string(v.(*token.Token).Lit)
	parsed, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, err
	}
	return parsed, nil
}
