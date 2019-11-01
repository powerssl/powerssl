module powerssl.dev/powerssl

go 1.13

require (
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/ahmetb/govvv v0.2.0
	github.com/arschles/assert v1.0.0 // indirect
	github.com/arschles/go-bindata-html-template v0.0.0-20170123182818-839a6918b9ff
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/cloudflare/cfssl v1.4.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/daviddengcn/go-colortext v0.0.0-20180409174941-186a3d44e920
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/eggsampler/acme/v2 v2.0.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml v1.0.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-kit/kit v0.9.0
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/gogo/googleapis v1.3.0
	github.com/gogo/protobuf v1.3.1
	github.com/gogo/status v1.1.0
	github.com/golang/protobuf v1.3.2
	github.com/golangplus/bytes v0.0.0-20160111154220-45c989fe5450 // indirect
	github.com/golangplus/fmt v0.0.0-20150411045040-2a5d6d7d2995 // indirect
	github.com/golangplus/testing v0.0.0-20180327235837-af21d9c3145e // indirect
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/hashicorp/vault v1.1.1
	github.com/improbable-eng/grpc-web v0.11.0
	github.com/jinzhu/gorm v1.9.11
	github.com/jinzhu/inflection v1.0.0
	github.com/johanbrandhorst/certify v1.6.0
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/looplab/fsm v0.1.0
	github.com/mattn/go-runewidth v0.0.5 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mwitkow/grpc-proxy v0.0.0-20181017164139-0f1106ef9c76 // indirect
	github.com/olekukonko/tablewriter v0.0.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pborman/uuid v1.2.0
	github.com/prometheus/client_golang v1.2.1
	github.com/qor/qor v0.0.0-20191022064424-b3deff729f68 // indirect
	github.com/qor/validations v0.0.0-20171228122639-f364bca61b46
	github.com/rs/cors v1.7.0 // indirect
	github.com/smacker/opentracing-gorm v0.0.0-20181207094635-cd4974441042
	github.com/spf13/cobra v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.5.0
	github.com/uber-go/atomic v0.0.0-00010101000000-000000000000 // indirect
	github.com/uber/jaeger-client-go v2.19.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/tools v0.0.0-20191029041327-9cc4af7d6b2c
	google.golang.org/grpc v1.24.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.0
	gopkg.in/square/go-jose.v2 v2.4.0
	gopkg.in/yaml.v2 v2.2.4
)

replace github.com/spf13/cobra => github.com/spf13/cobra v0.0.6-0.20191019221741-77e4d5aecc4d

replace github.com/uber-go/atomic => go.uber.org/atomic v1.5.0
