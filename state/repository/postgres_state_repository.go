package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/pragmaticivan/tinyestate-api/model"
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

func (m *postgresStateRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.State, error) {

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
	result := make([]*model.State, 0)
	for rows.Next() {
		t := new(model.State)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Acronym,
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

func (m *postgresStateRepository) Fetch(ctx context.Context, cursor string, num int64) ([]*model.State, string, error) {

	query := `SELECT id,name,acronym, updated_at, created_at
  						FROM states WHERE created_at > ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", model.ErrBadParamInput
	}
	res, err := m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""
	if len(res) == int(num) {
		nextCursor = EncodeCursor(res[len(res)-1].CreatedAt)
	}
	return res, nextCursor, err

}

func (m *postgresStateRepository) GetByID(ctx context.Context, id int64) (*model.State, error) {
	return nil, nil
}

func (m *postgresStateRepository) Save(ctx context.Context, a *model.State) error {
	return nil
}

func (m *postgresStateRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
func (m *postgresStateRepository) Update(ctx context.Context, ar *model.State) error {
	return nil
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
