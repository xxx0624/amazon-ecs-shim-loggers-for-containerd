//go:generate protoc -I . --gogofast_out=import_path=github.com/aws/shim-loggers-for-containerd/docker/api/types/swarm/runtime:. plugin.proto

package runtime // import "github.com/docker/docker/api/types/swarm/runtime"
