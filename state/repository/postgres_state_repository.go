package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/pragmaticivan/tinyestate-api/domain"
	"github.com/pragmaticivan/tinyestate-api/state"
	log "github.com/sirupsen/logrus"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type postgresStateRepository struct {
	Conn *sql.DB
}

// NewPostgresStateRepository will create an object that represent the state.Repository interface
func NewPostgresStateRepository(Conn *sql.DB) state.Repository {
	return &postgresStateRepository{Conn}
}

func (m *postgresStateRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*domain.State, error) {

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
	result := make([]*domain.State, 0)
	for rows.Next() {
		t := new(domain.State)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Abbreviation,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		t.Cities = make([]domain.City, 0)
		result = append(result, t)
	}

	return result, nil
}

func (m *postgresStateRepository) Fetch(ctx context.Context) ([]*domain.State, error) {
	query := `SELECT id, name, abbreviation, updated_at, created_at FROM states ORDER BY created_at`

	res, err := m.fetch(ctx, query)
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
