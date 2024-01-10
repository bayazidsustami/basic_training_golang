package belajargolanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "bay", GetName[Manager](&MyManager{Name: "bay"}))
	assert.Equal(t, "bayazid", GetName[VicePresident](&MyVicePresident{Name: "bayazid"}))
}
