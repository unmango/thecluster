package progress

import (
	"github.com/unmango/go/rx/observable"
	"github.com/unmango/thecluster/pkg"
)

func Observe(w pkg.Workspace) pkg.Observable {
	if obs, ok := w.(pkg.Observable); ok {
		return obs
	} else {
		return observable.New[pkg.ProgressEvent]()
	}
}
