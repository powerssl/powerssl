module powerssl.dev/tools/dev-runner

go 1.15

replace powerssl.dev/common => ../../internal/common

require (
	github.com/daviddengcn/go-colortext v1.0.0
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/improbable-eng/grpc-web v0.14.0
	github.com/lib/pq v1.9.0
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.35.0 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
