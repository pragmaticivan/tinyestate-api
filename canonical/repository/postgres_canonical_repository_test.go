package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/pragmaticivan/tinyestate-api/canonical/repository"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Info(err)
		}
	}()

	rows := sqlmock.NewRows([]string{"id", "name", "canonical", "allows_on_wheels", "allows_on_foundation", "requires_care_giver", "created_at", "updated_at"}).
		AddRow(1, "Fresno", "fresno-ca", true, true, false, time.Now(), time.Now())

	query := "SELECT id, name, canonical, allows_on_wheels, allows_on_foundation, requires_care_giver, created_at, updated_at FROM cities ORDER BY created_at"

	mock.ExpectQuery(query).WillReturnRows(rows)

	a := repository.NewPostgresCityRepository(db)

	dataCities, err := a.Fetch(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, dataCities)
}
