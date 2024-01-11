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

type Sample struct {
	Id   string
	Name string
}

func TestRawSql(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", sample.Id)
	assert.Equal(t, "sambo", sample.Name)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}
