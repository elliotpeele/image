// +build !containers_image_storage_stub

package alltransports

import (
	// Register the storage transport
	_ "github.com/elliotpeele/image/storage"
)
