package nsync

import "sync"

type (
	namedMutex map[interface{}]*sync.Mutex

	// A Mutex is a named mutual exclusion lock
	Mutex struct {
		namedMux namedMutex
		mux      sync.Mutex
	}

	// A Locker represents an object that can be locked and unlocked.
	Locker interface {
		Lock(interface{})
		Unlock(interface{})
	}
)

var _ Locker = (*Mutex)(nil)

// Lock locks m given name
func (m *Mutex) Lock(name interface{}) {
	m.mux.Lock()
	v, ok := m.namedMux[name]
	if !ok {
		if m.namedMux == nil {
			m.namedMux = namedMutex{}
		}
		v = &sync.Mutex{}
		m.namedMux[name] = v
	}
	m.mux.Unlock()

	v.Lock()
}

// Unlock unlocks m given name
func (m *Mutex) Unlock(name interface{}) {
	m.mux.Lock()
	v, ok := m.namedMux[name]
	m.mux.Unlock()
	if !ok {
		return
	}

	v.Unlock()
}
