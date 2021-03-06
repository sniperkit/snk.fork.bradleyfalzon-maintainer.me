/*
Sniperkit-Bot
- Status: analyzed
*/

package notifier

import (
	"fmt"
	"io"

	"github.com/sniperkit/snk.fork.bradleyfalzon-maintainer.me/events"
)

// Writer is a Notifier that writes the event to the supplied writer.
type Writer struct {
	Writer io.Writer
}

var _ events.Notifier = &Writer{}

// Notify implements the Notifier interface.
func (w *Writer) Notify(event *events.Event) error {
	_, err := fmt.Fprintf(w.Writer, "NOTIFY: %q\n", event.String())
	return err
}
