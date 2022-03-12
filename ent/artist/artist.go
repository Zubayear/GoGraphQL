// Code generated by entc, DO NOT EDIT.

package artist

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the artist type in the database.
	Label = "artist"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// EdgeSongs holds the string denoting the songs edge name in mutations.
	EdgeSongs = "songs"
	// Table holds the table name of the artist in the database.
	Table = "artists"
	// SongsTable is the table that holds the songs relation/edge.
	SongsTable = "artists"
	// SongsInverseTable is the table name for the Song entity.
	// It exists in this package in order to avoid circular dependency with the "song" package.
	SongsInverseTable = "songs"
	// SongsColumn is the table column denoting the songs relation/edge.
	SongsColumn = "song_artists"
)

// Columns holds all SQL columns for artist fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "artists"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"song_artists",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)