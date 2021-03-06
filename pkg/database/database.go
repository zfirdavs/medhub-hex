package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type exists struct {
	tableName string
	clauses   map[string]interface{}
}

// NewPool creates a new pgx interface database pool
func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}

// func IsExists(strct exists) (bool, error) {
// 	var isExists bool
// 	sqlStr, args, err := NewPostgresQuery().
// 		Select("1").
// 		Prefix("SELECT EXISTS (").
// 		From(strct.tableName).
// 		Where(EqualMany(strct.clauses)).
// 		Suffix(")").
// 		ToSql()
// 	if err != nil {
// 		return false, fmt.Errorf("error during %s check exists sql model: %w", strct.tableName, err)
// 	}

// 	if err = Get(&isExists, sqlStr, args...); err != nil {
// 		return false, fmt.Errorf("error during %s model exists get: %w", strct.tableName, err)
// 	}
// 	return isExists, nil
// }

// func Get(dest interface{}, query string, args ...interface{}) error {
// 	return db.Get(dest, query, args...)
// }
