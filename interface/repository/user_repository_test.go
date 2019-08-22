package datastore

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/manakuro/golang-clean-architecture/domain/model"

	"github.com/jinzhu/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

func newDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}

func TestUserRepository_FindAll(t *testing.T) {
	db, mock, _ := newDBMock()

	defer db.Close()
	db.LogMode(true)

	r := NewUserRepository(db)

	// mock
	u1 := model.User{ID: 1, Name: "manato", Age: "20"}
	col := []string{"id", "name", "age"}
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT *`)).
		WillReturnRows(sqlmock.NewRows(col).AddRow(u1.ID, u1.Name, u1.Age))

	var u []*model.User
	u, _ = r.FindAll(u)

	assert.Equal(t, 1, len(u))
}
