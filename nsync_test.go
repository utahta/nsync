package nsync

import (
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestMutex_Lock(t *testing.T) {
	var (
		nmux Mutex
		eg   errgroup.Group
		a    int
		b    int
		c    int
	)

	nameStr := []string{"a", "b", "a", "a", "c", "c", "b"}
	for _, name := range nameStr {
		name := name
		eg.Go(func() error {
			nmux.Lock(name)
			defer nmux.Unlock(name)

			switch name {
			case "a":
				a++
			case "b":
				b++
			case "c":
				c++
			}
			return nil
		})
	}
	nameInt := []int{1, 2, 3, 2, 1, 1}
	for _, name := range nameInt {
		name := name
		eg.Go(func() error {
			nmux.Lock(name)
			defer nmux.Unlock(name)
			return nil
		})
	}

	eg.Wait()

	if a != 3 {
		t.Errorf("Expected a:3, got %v", a)
	}
	if b != 2 {
		t.Errorf("Expected b:2, got %v", a)
	}
	if c != 2 {
		t.Errorf("Expected c:2, got %v", a)
	}
}

func TestMutex_Unlock(t *testing.T) {
	var mux Mutex
	mux.Unlock("a")
}
