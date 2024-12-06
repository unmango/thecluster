package progress

import (
	"fmt"
	"io"

	"github.com/unmango/go/rx"
	"github.com/unmango/thecluster/pkg"
)

type writer struct{ io.Writer }

// OnComplete implements rx.Observer.
func (w writer) OnComplete() {
	fmt.Fprintln(w, "Done")
}

// OnError implements rx.Observer.
func (w writer) OnError(err error) {
	fmt.Fprintln(w, err)
}

// OnNext implements rx.Observer.
func (w writer) OnNext(e pkg.ProgressEvent) {
	fmt.Fprint(w, e.Message)
}

func WriteTo(obs pkg.Observable, w io.Writer) rx.Subscription {
	return obs.Subscribe(writer{w})
}
