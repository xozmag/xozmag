package entities

import "database/sql"

func NullString(uuid string) sql.NullString {
	return sql.NullString{String: uuid, Valid: true}
}