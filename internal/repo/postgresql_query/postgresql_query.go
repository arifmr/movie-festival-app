package postgresql_query

import _ "embed"

//nolint:gochecknoglobals
var (
	// file format => {SCHEMA}.{TABLE}--{COMMAND}[.{EXTRA}].sql
	///////////////////////////////////////////////////////////////////////////.

	// MOVIES TABLE
	//
	// ----------------------------------------------------------------------------.

	//go:embed movies/movies--insert.sql
	InsertMovies string
)
