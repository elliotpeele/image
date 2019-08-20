// +build containers_image_ostree,linux

package alltransports

import (
	// Register the ostree transport
	_ "github.com/elliotpeele/image/ostree"
)
