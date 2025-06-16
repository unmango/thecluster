package result

import (
	"time"

	"github.com/unmango/thecluster/pkg/contract"
)

type Result struct {
	after time.Time
}

func (r Result) RequeueAfter() time.Time {
	return r.after
}

func RequeueAfter(t time.Time) contract.Result {
	return Result{t}
}

func Done() contract.Result {
	return Result{}
}
