package fhir

import (
	"context"
	"errors"
	"fmt"

	schema "github.com/MedHubUz/fhirschema"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/medhub-hex/pkg/database"
)

type postgresRepository struct {
	dbpool   *pgxpool.Pool
	squirrel *database.Squirrel
}

// NewPostresRepository provides postgres database implementation
func NewPostresRepository(pool *pgxpool.Pool, sq *database.Squirrel) *postgresRepository {
	return &postgresRepository{pool, sq}
}

func (r *postgresRepository) Read(ctx context.Context, resource *Resource) (interface{}, error) {
	var dataStr string
	conn, err := r.dbpool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query := r.squirrel.Builder.
		Select("r.practitioner_id", "r.data", "r.created_at").
		From("resource AS r").
		Where(r.squirrel.Equal("r.resource_type", resource.ResourceType))

	if resource.ResourceType == "Patient" && len(resource.ResourceID) > 36 {
		query = query.LeftJoin("resource_id_hash AS h").
			JoinClause("ON (h.resource_id = r.resource_id)").
			Where(r.squirrel.Equal("h.hash", resource.ResourceID))
	} else {
		query = query.Where(r.squirrel.Equal("r.resource_id", resource.ResourceID))
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error during resource read sql model: %w", err)
	}

	err = conn.QueryRow(ctx, sqlStr, args...).Scan(
		&resource.PractitionerID,
		&dataStr,
		&resource.CreatedAt,
	)

	switch {
	case err == pgx.ErrNoRows:
		return nil, errors.New("no rows returned")
	case err != nil:
		return nil, err
	default:
		return resource, nil
		// default:
		// return r.DecodePatientByStringData(dataStr)
	}
}

func (r *postgresRepository) ReadAll(ctx context.Context, resourceType string) (*schema.Bundle, error) {
	return nil, nil
}

func (r *postgresRepository) Create(ctx context.Context, resource *Resource) (*Resource, error) {
	conn, err := r.dbpool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	clauses := map[string]interface{}{
		"resource_id":     resource.ResourceID,
		"practitioner_id": resource.PractitionerID,
		"resource_type":   resource.ResourceType,
		"clinic":          "clinic",
		"data":            resource.Data,
		"created_at":      resource.CreatedAt,
	}

	sqlStr, args, err := r.squirrel.Builder.
		Insert("resource").
		SetMap(clauses).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error during resource create model sql: %w", err)
	}

	// sqlStr := `
	// 			INSERT INTO resource
	// 			(resource_id, practitioner_id, resource_type, clinic, data, created_at)
	// 			VALUES
	// 			($1, $2, $3, $4, $5, $6);

	// `

	if _, err = conn.Exec(ctx, sqlStr, args...); err != nil {
		return nil, err
	}

	return resource, nil
}

func (r *postgresRepository) Update(ctx context.Context, resourceType, resourceID string) (interface{}, error) {
	return nil, nil
}

func (r *postgresRepository) DecodePatient(ctx context.Context, resource *Resource) (interface{}, error) {
	return nil, nil
}
