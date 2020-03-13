module github.com/snowzach/qvspot

require (
	github.com/blendle/zapdriver v1.3.1
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-chi/render v1.0.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.2
	github.com/olivere/elastic/v7 v7.0.12
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/snowzach/certtools v1.0.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v0.0.6
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.2
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.14.0
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	google.golang.org/genproto v0.0.0-20200309141739-5b75447e413d
	google.golang.org/grpc v1.27.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/ini.v1 v1.54.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

go 1.12

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1
