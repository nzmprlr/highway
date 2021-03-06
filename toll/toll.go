package toll

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/nzmprlr/idpp"
)

// TODO push tolls in a thread-safe queue,
// flush the queue within interval or threshold count
// send to collector(kafka, direct service, etc.)
// process the metrics and logs there.
type Toll struct {
	ID        string                 `json:"id"`
	Status    uint16                 `json:"status"`
	StartedAt string                 `json:"startedAt"`
	Source    string                 `json:"source,omitempty"`
	Labels    map[string]interface{} `json:"labels"`
	Metrics   *Metrics               `json:"metrics"`
}

func (toll *Toll) Metric(metric *Metric) {
	if toll.Metrics.len() > 1 { // dont pop root metric
		toll.Metrics.pop()
	}

	metric.measure()
}

func (toll *Toll) End() {
	if c.Quiet {
		return
	}

	var b []byte
	var err error
	if c.Indent != "" {
		b, err = json.MarshalIndent(toll, "", c.Indent)
	} else {
		b, err = json.Marshal(toll)
	}

	if err != nil {
		toll.Log(err.Error())
		return
	}

	fmt.Println(string(b))
}

func (toll *Toll) Label(k string, v interface{}) {
	toll.Labels[k] = v
}

func (toll *Toll) Log(format string, a ...interface{}) {
	peek := toll.Metrics.peek()
	if peek != nil {
		peek.log(fmt.Sprintf(format, a...))
	}
}

func (toll *Toll) SetStatus(status uint16) {
	peek := toll.Metrics.peek()
	if peek != nil {
		peek.Status = status
	}

	toll.Status = status
}

func (toll *Toll) Fork() *Toll {
	source := unknown
	peek := toll.Metrics.peek()
	if peek != nil {
		source = peek.Name
	}

	t := newToll()
	t.ID = toll.ID
	t.Source = source
	return t
}

func New() *Toll {
	t := newToll()

	t.ID = idpp.NewID12().String()
	return t
}

func NewWithID(ids string) *Toll {
	t := newToll()

	if ids != "" {
		if id, err := idpp.ParseID12(ids); err == nil {
			t.ID = id.String()
		}
	} else {
		t.ID = idpp.NewID12().String()
	}

	return t
}

func newToll() *Toll {
	return &Toll{
		Status:    http.StatusOK,
		StartedAt: strconv.FormatInt(time.Now().UnixNano(), 10),
		Labels:    map[string]interface{}{},
		Metrics:   new(Metrics),
	}
}
