package q

/////////////
// Actions //
/////////////

const (
	ACTION_QUERY_ONE = iota
	ACTION_QUERY_MULTI
	ACTION_INSERT
	ACTION_UPDATE
	ACTION_DELETE
	ACTION_CREATE_TABLE
)

////////////////
// Conditions //
////////////////

const (
	EQUAL = iota
	UNEQUAL
	GREATER
	LESSER
	GREATER_EQ
	LESSER_EQ
	AND
	OR
)

///////////
// Order //
///////////

const (
	ASC = iota
	DESC
)

/////////////////
// Constraints //
/////////////////

const (
	NONE = iota
	UNIQUE
	NOT_NULL
	NOT_NULL_UNIQUE
	PRIMARY
)

///////////
// Types //
///////////

const (
	BOOLEAN = iota
	UUID
	VARCHAR
	CHAR
	TEXT
	SMALLINT
	INT
	BIGINT
	REAL
	DOUBLE
	TIMESTAMP
)
