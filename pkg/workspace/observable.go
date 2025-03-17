package workspace

import (
	"github.com/charmbracelet/log"
	"github.com/unmango/go/rx/observable"
	"github.com/unmango/thecluster/pkg"
)

func Observe(w pkg.Workspace) pkg.Observable {
	if obs, ok := w.(pkg.Observable); ok {
		log.Debug("observing workspace")
		return obs
	} else {
		log.Debug("not an observable workspace")
		return observable.New[pkg.ProgressEvent]()
	}
}
