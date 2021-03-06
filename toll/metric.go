package toll

import (
	"fmt"
	"net/http"
	"time"
)

type Metric struct {
	Type          string        `json:"type"`
	Status        uint16        `json:"status"`
	Name          string        `json:"name"`
	Params        string        `json:"params"`
	Measure       time.Duration `json:"measure"`
	MeasureString string        `json:"measureStr"`
	Logs          []string      `json:"logs"`
	Metrics       *Metrics      `json:"metrics"`

	startedAt time.Time
}

func (metric *Metric) measure() {
	since := time.Since(metric.startedAt)

	metric.Measure = since
	metric.MeasureString = since.String()
}

func (metric *Metric) log(l string) {
	metric.Logs = append(metric.Logs, l)
}

func newMetric(t, name string, params []interface{}) *Metric {
	return &Metric{
		Type:      t,
		Name:      name,
		Status:    http.StatusOK,
		Params:    fmt.Sprintf("%+v", params),
		Logs:      []string{},
		Metrics:   new(Metrics),
		startedAt: time.Now(),
	}
}
