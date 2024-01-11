package belajargolanggorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSql(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values(?,?)", "1", "sambo").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?,?)", "2", "lacukka ulu").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?,?)", "3", "ambo nai").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?,?)", "4", "beddu").Error
	assert.Nil(t, err)
}
