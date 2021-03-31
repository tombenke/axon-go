module github.com/tombenke/axon-go

go 1.15

require (
	github.com/gizak/termui/v3 v3.1.0
	github.com/influxdata/influxdb-client-go/v2 v2.2.2
	github.com/robertkrimen/otto v0.0.0-20191219234010-c382bd3c16ff
	github.com/robfig/cron v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/tombenke/axon-go-common v1.6.1
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

//replace github.com/tombenke/axon-go-common => ../axon-go-common
