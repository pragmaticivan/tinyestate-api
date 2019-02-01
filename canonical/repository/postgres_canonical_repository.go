package repository

import (
	"context"
	"database/sql"

	"github.com/pragmaticivan/tinyestate-api/canonical"
	"github.com/pragmaticivan/tinyestate-api/domain"
	log "github.com/sirupsen/logrus"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type postgresCanonicalRepository struct {
	Conn *sql.DB
}

// NewPostgresCanonicalRepository will create an object that represent the city.Repository interface
func NewPostgresCanonicalRepository(Conn *sql.DB) canonical.Repository {
	return &postgresCanonicalRepository{Conn}
}

func (m *postgresCanonicalRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*domain.Canonical, error) {

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
	result := make([]*domain.Canonical, 0)
	for rows.Next() {
		t := new(domain.Canonical)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Canonical,
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

func (m *postgresCanonicalRepository) Fetch(ctx context.Context) ([]*domain.Canonical, error) {
	query := `SELECT id, name, canonical, allows_on_wheels, allows_on_foundation, requires_care_giver, created_at, updated_at FROM canonicals ORDER BY created_at`
	res, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (m *postgresCanonicalRepository) FetchByID(ctx context.Context, id int64) (*domain.Canonical, error) {
	query := `SELECT id, name, canonical, allows_on_wheels, allows_on_foundation, requires_care_giver, created_at, updated_at FROM canonicals WHERE id = ?`
	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	a := &domain.Canonical{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, domain.ErrNotFound
	}
	return a, nil
}

func (m *postgresCanonicalRepository) FetchByCanonical(ctx context.Context, canonical string) (*domain.Canonical, error) {
	query := `SELECT id, name, canonical, allows_on_wheels, allows_on_foundation, requires_care_giver, created_at, updated_at FROM canonicals WHERE canonical = ?`
	list, err := m.fetch(ctx, query, canonical)
	if err != nil {
		return nil, err
	}
	a := &domain.Canonical{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, domain.ErrNotFound
	}
	return a, nil
}

func (m *postgresCanonicalRepository) Create(ctx context.Context, c *domain.Canonical) (*domain.Canonical, error) {
	query := `INSERT INTO canonicals (name, canonical, allows_on_wheels, allows_on_foundation, requires_care_giver, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {

		return nil, err
	}

	res, err := stmt.ExecContext(ctx, c.Name, c.Canonical, c.AllowsOnWheels, c.AllowsOnFoundation, c.RequiresCareGiver, c.UpdatedAt, c.CreatedAt)
	if err != nil {

		return nil, err
	}
	LastInsertId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	c.ID = LastInsertId
	return c, nil
}
