package memory

import (
	"testing"

	"github.com/elliotpeele/image/pkg/blobinfocache/internal/test"
	"github.com/elliotpeele/image/types"
)

func newTestCache(t *testing.T) (types.BlobInfoCache, func(t *testing.T)) {
	return New(), func(t *testing.T) {}
}

func TestNew(t *testing.T) {
	test.GenericCache(t, newTestCache)
}
