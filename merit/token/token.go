package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENTIFIER = "IDENT" // add, foobar, x, y, ...
	INTEGER    = "INT"   // 1343456
	// Operators
	ASSIGN = "="
	// PLUS   = "+"
	// // Delimiters
	// COMMA     = ","
	// SEMICOLON = ";"
	// LPAREN    = "("
	// RPAREN    = ")"
	// LBRACE    = "{"
	// RBRACE    = "}"
	// Keywords
	CONST = "CONST"
)

var keywords = map[string]TokenType{"const": CONST}

func LookupIdentifier(word string) TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return IDENTIFIER
}
