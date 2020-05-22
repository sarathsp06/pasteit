package store

import (
	"context"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Memory acts as in-memory paste store
type Memory struct {
	data  map[string]Paste
	mutex *sync.RWMutex
}

// NewMemory returns a new instance of sore memory implementation
func NewMemory() *Memory {
	return &Memory{
		data:  make(map[string]Paste),
		mutex: new(sync.RWMutex),
	}
}

// SavePaste stores the paste in memmory
func (m *Memory) SavePaste(ctx context.Context, p Paste) (Paste, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[p.ID.String()] = p
	return p, nil
}

// paste return paste given id, it ignores the expiry
func (m *Memory) paste(ctx context.Context, id uuid.UUID) (Paste, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := m.data[id.String()]
	if !ok {
		return Paste{}, ErrNotFound
	}
	return v, nil
}

// paste return paste given id, it ignores the expiry
func (m *Memory) deletePaste(id uuid.UUID) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, id.String())
}

// GetPaste return paste given id
func (m *Memory) GetPaste(ctx context.Context, id uuid.UUID) (Paste, error) {
	p, err := m.paste(ctx, id)
	if err != nil {
		return Paste{}, err
	}
	if p.Expiry.Before(time.Now()) {
		go m.deletePaste(id)
		return Paste{}, ErrNotFound
	}
	return p, nil
}
