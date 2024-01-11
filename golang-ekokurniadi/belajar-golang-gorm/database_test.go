package belajargolanggorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}
