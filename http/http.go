package http

import (
	"github.com/glightfoot/atlas_exporter/config"
	"github.com/glightfoot/atlas_exporter/exporter"
)

const (
	ns  = "atlas"
	sub = "http"
)

// NewMeasurement returns a new instance of `exorter.Measurement` for a HTTP measurement
func NewMeasurement(id, ipVersion string, cfg *config.Config) *exporter.Measurement {
	opts := []exporter.MeasurementOpt{
		exporter.WithHistograms(newRttHistogram(id, ipVersion, cfg.HistogramBuckets.HTTP.Rtt)),
	}

	if cfg.FilterInvalidResults {
		opts = append(opts, exporter.WithValidator(&exporter.DefaultResultValidator{}))
	}

	return exporter.NewMeasurement(&httpExporter{id}, opts...)
}
