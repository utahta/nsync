package nsync

import (
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestMutex_Lock(t *testing.T) {
	nameList := []string{"a", "a", "b", "c", "c", "c", "d"}
	var eg errgroup.Group
	var mux Mutex
	for _, name := range nameList {
		name := name
		eg.Go(func() error {
			mux.Lock(name)
			mux.Unlock(name)
			return nil
		})
	}
	eg.Wait()
}

func TestMutex_Unlock(t *testing.T) {
	var mux Mutex
	mux.Unlock("a")
}
