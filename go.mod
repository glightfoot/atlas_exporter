module github.com/glightfoot/atlas_exporter

go 1.16

require (
	github.com/glightfoot/ripeatlas v0.0.0-20210429151918-5da82a40f47b
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/common v0.18.0
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/glightfoot/golang-socketio v0.0.0-20210429151747-49c2ab51dc00 => ../golang-socketio

replace github.com/glightfoot/ripeatlas v0.0.0-20210429151918-5da82a40f47b => ../ripeatlas
