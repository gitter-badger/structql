package lexing

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	LITERAL

	ASTERISK
	COMMA

	SELECT
	FROM
	WHERE
	AND
	OR
	LIMIT
	OFFSET
)
