// +build containers_image_docker_daemon_stub

package alltransports

import "github.com/elliotpeele/image/transports"

func init() {
	transports.Register(transports.NewStubTransport("docker-daemon"))
}
