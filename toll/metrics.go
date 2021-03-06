package toll

import (
	"runtime"
)

const (
	unknown = "unknown"
)

type Metrics []*Metric

func (metrics Metrics) len() int {
	return len(metrics)
}

func (metrics *Metrics) push(m *Metric) {
	*metrics = append(*metrics, m)
}

func (metrics Metrics) peek() *Metric {
	l := metrics.len()
	if l == 0 {
		return nil
	}

	return metrics[l-1]
}

func (metrics *Metrics) pop() *Metric {
	l := metrics.len()
	if l == 0 {
		return nil
	}

	metric := (*metrics)[l-1]
	*metrics = (*metrics)[:l-1]
	return metric
}

func (metrics *Metrics) new(t string, params []interface{}) *Metric {
	name := unknown
	if pc, _, _, ok := runtime.Caller(2); ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			name = fn.Name()
		}
	}

	metric := newMetric(t, name, params)
	peek := metrics.peek()
	if peek != nil {
		peek.Metrics.push(metric)
	}

	metrics.push(metric)
	return metric
}

func (metrics *Metrics) Controller(params ...interface{}) *Metric {
	return metrics.new("controller", params)
}

func (metrics *Metrics) IO(params ...interface{}) *Metric {
	return metrics.new("io", params)
}

func (metrics *Metrics) Service(params ...interface{}) *Metric {
	return metrics.new("service", params)
}

func (metrics *Metrics) Event(params ...interface{}) *Metric {
	return metrics.new("event", params)
}

func (metrics *Metrics) Logic(params ...interface{}) *Metric {
	return metrics.new("logic", params)
}

func (metrics *Metrics) Data(params ...interface{}) *Metric {
	return metrics.new("data", params)
}

func (metrics *Metrics) Call(params ...interface{}) *Metric {
	return metrics.new("call", params)
}
