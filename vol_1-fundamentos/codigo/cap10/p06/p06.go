package main

import (
	"fmt"
	"strconv"
)

type Repository interface {
	Save(string) error
	Load(id string) (string, error)
}

// --------- MemoryRepository -------------
type MemoryRepository struct {
	Data   map[string]string
	LastId string
}

func (m *MemoryRepository) Save(value string) error {
	if m.Data == nil {
		m.Data = make(map[string]string)
	}

	if m.LastId == "" {
		m.LastId = "001"
	} else {
		id, err := strconv.Atoi(m.LastId)
		if err != nil {
			return err
		}
		m.LastId = fmt.Sprintf("%03d", id+1)
	}

	m.Data[m.LastId] = value
	return nil
}

func (m MemoryRepository) Load(id string) (string, error) {
	if value, ok := m.Data[id]; ok {
		return value, nil
	}

	return "", fmt.Errorf("índice %s não encontrado", id)
}

// --------- MockRepository -------------
type MockRepository struct {
	Data string
}

func (m *MockRepository) Save(value string) error {
	m.Data = value
	fmt.Println("Salvo...")
	return nil
}

func (m MockRepository) Load(id string) (string, error) {
	fmt.Println("Lido...")
	return m.Data, nil
}

// -------------------------------------
func Save(r Repository, value string) error {
	return r.Save(value)
}

func Load(r Repository, id string) (string, error) {
	return r.Load(id)
}

func main() {
	m := &MemoryRepository{}
	c := &MockRepository{}

	Save(m, "Ir ao médico")
	Save(m, "Fazer compras")

	ids := []string{"000", "001", "002"}
	for _, id := range ids {

		v, err := m.Load(id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}

	fmt.Println(Load(c, "123"))
	Save(c, "Estou aqui...")
	fmt.Println(Load(c, "123"))
}
