package repository_test

import (
	"regexp"
	"testing"

	"github.com/manakuro/golang-clean-architecture/testutil"

	"github.com/manakuro/golang-clean-architecture/interface/repository"

	"github.com/stretchr/testify/assert"

	"github.com/manakuro/golang-clean-architecture/domain/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepository_FindAll(t *testing.T) {
	db, mock, _ := testutil.NewDBMock()

	defer db.Close()

	r := repository.NewUserRepository(db)

	// user
	u1 := model.User{ID: 1, Name: "manato", Age: "20"}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(u1.ID, u1.Name, u1.Age))

	// credit card
	c1 := model.CreditCard{
		ID:     1,
		UserID: 1,
		Number: "111111",
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "number"}).AddRow(c1.ID, c1.UserID, c1.Number))

	u, err := r.FindAll([]*model.User{})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(u))
	assert.NotEmpty(t, u[0].CreditCards)
}
