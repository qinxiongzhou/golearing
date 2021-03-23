package remote

import (
	"github.com/easierway/concurrent_map"
	"testing"
)
func TestConcurrentMap(t *testing.T) {
	m := concurrent_map.CreateConcurrentMap(10)
	m.Set(concurrent_map.StrKey("key"),10)
	t.Log(m.Get(concurrent_map.StrKey("key")))
}
