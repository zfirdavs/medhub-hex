package database

import sq "github.com/Masterminds/squirrel"

type Squirrel struct {
	Builder sq.StatementBuilderType
}

func NewSquirrel() *Squirrel {
	return &Squirrel{sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (s *Squirrel) Equal(key string, value interface{}) sq.Eq {
	return sq.Eq{key: value}
}

func (s *Squirrel) ILike(key string, value interface{}) sq.ILike {
	return sq.ILike{key: value}
}

func (s *Squirrel) NotEqual(key string, value interface{}) sq.NotEq {
	return sq.NotEq{key: value}
}

func (s *Squirrel) Or(cond ...sq.Sqlizer) sq.Or {
	sl := make([]sq.Sqlizer, 0, len(cond))
	for _, val := range cond {
		sl = append(sl, val)
	}
	return sl
}

func (s *Squirrel) Alias(expr sq.Sqlizer, alias string) sq.Sqlizer {
	return sq.Alias(expr, alias)
}

func (s *Squirrel) EqualMany(clauses map[string]interface{}) sq.Eq {
	eqMany := make(sq.Eq, len(clauses))
	for key, value := range clauses {
		eqMany[key] = value
	}
	return eqMany
}
