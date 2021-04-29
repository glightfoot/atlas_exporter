package exporter

import (
	"github.com/glightfoot/ripeatlas/measurement"
	"github.com/glightfoot/atlas_exporter/probe"
	"github.com/prometheus/client_golang/prometheus"
)

// Exporter defines a set of metrics for an ATLAS measurement type
type Exporter interface {

	// Export exports a prometheus metric
	Export(res *measurement.Result, probe *probe.Probe, ch chan<- prometheus.Metric)

	// Describes metrics exported for this measurement type
	Describe(ch chan<- *prometheus.Desc)
}
