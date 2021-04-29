package exporter

import (
	"github.com/glightfoot/ripeatlas/measurement"
	"github.com/glightfoot/atlas_exporter/probe"
)

// DefaultResultValidator is the validator used by many measurement types
type DefaultResultValidator struct {
}

// IsValid returns whether an result is valid or not (e.g. IPv6 measurement and Probe does not support IPv6)
func (m *DefaultResultValidator) IsValid(res *measurement.Result, probe *probe.Probe) bool {
	return probe.ASNForIPVersion(res.Af()) > 0
}
