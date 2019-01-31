package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/pragmaticivan/tinyestate-api/city"
	"github.com/pragmaticivan/tinyestate-api/domain"
	log "github.com/sirupsen/logrus"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type postgresCityRepository struct {
	Conn *sql.DB
}

// NewPostgresCityRepository will create an object that represent the city.Repository interface
func NewPostgresCityRepository(Conn *sql.DB) city.Repository {
	return &postgresCityRepository{Conn}
}

func (m *postgresCityRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*domain.City, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	result := make([]*domain.City, 0)
	for rows.Next() {
		t := new(domain.City)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.AllowsOnWheels,
			&t.AllowsOnFoundation,
			&t.RequiresCareGiver,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *postgresCityRepository) Fetch(ctx context.Context) ([]*domain.City, error) {
	query := `SELECT id, name, allows_on_wheels, allows_on_foundation, requires_care_giver, created_at, updated_at FROM cities ORDER BY created_at`
	res, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (m *postgresCityRepository) GetByStateID(ctx context.Context, id int64) ([]*domain.City, error) {
	query := `SELECT id, name, allows_on_wheels, allows_on_foundation, requires_care_giver, updated_at, created_at FROM cities WHERE state_id=$1 ORDER BY created_at`
	res, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	return res, err
}

// DecodeCursor -
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor -
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
