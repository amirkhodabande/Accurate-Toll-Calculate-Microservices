package store

import (
	"aggregator/types"
	"fmt"
)

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStore) Insert(d types.Distance) error {
	m.data[d.OBUID] += d.Value
	fmt.Println(m.data)

	return nil
}

func (m *MemoryStore) Get(id int) (float64, error) {
	dist, ok := m.data[id]
	fmt.Printf("%+v", m.data)
	if !ok {
		return 0.0, fmt.Errorf("the given id is invalid %d", id)
	}

	return dist, nil
}
