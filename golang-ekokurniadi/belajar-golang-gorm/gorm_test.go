package belajargolanggorm

import (
	"strconv"
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

func TestRowsSql(t *testing.T) {
	var samples []Sample
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}

	assert.Equal(t, 4, len(samples))
}

func TestScanRowsSql(t *testing.T) {
	var samples []Sample
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := User{
		ID:       "1",
		Password: "rahasia",
		Name: Name{
			FirstName:  "bayazid",
			MiddleName: "sustami",
			LastName:   "Mohammad Nasir",
		},
		Information: "ini akan di ignore",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "rahasia",
			Name: Name{
				FirstName: "User" + strconv.Itoa(i),
			},
		})
	}

	result := db.Create(users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))
}
