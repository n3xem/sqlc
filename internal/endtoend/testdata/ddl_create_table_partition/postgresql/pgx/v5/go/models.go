// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.1

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Foo struct {
	ID      pgtype.UUID
	OtherID pgtype.UUID
}

type Foo1 struct {
	ID      pgtype.UUID
	OtherID pgtype.UUID
}