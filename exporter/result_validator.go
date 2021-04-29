package exporter

import (
	"github.com/glightfoot/ripeatlas/measurement"
	"github.com/glightfoot/atlas_exporter/probe"
)

// ResultValidator validates results for measurements
type ResultValidator interface {
	// IsValid returns if a meaurement result is valid (can be filtered when needed)
	IsValid(res *measurement.Result, probe *probe.Probe) bool
}
