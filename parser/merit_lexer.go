// Code generated from /Users/christopherpoulsen/dev/git/merideum-go/Merit.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 73, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 43, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 48, 10,
	6, 12, 6, 14, 6, 51, 11, 6, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 57, 10, 8, 12,
	8, 14, 8, 60, 11, 8, 3, 8, 5, 8, 63, 10, 8, 3, 9, 3, 9, 3, 10, 6, 10, 68,
	10, 10, 13, 10, 14, 10, 69, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 2, 19, 10, 3, 2, 6, 4, 2, 67, 92, 99, 124,
	3, 2, 51, 59, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 78, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3,
	21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 30, 3, 2, 2, 2, 9, 36, 3, 2, 2, 2,
	11, 42, 3, 2, 2, 2, 13, 52, 3, 2, 2, 2, 15, 62, 3, 2, 2, 2, 17, 64, 3,
	2, 2, 2, 19, 67, 3, 2, 2, 2, 21, 22, 7, 63, 2, 2, 22, 4, 3, 2, 2, 2, 23,
	24, 7, 113, 2, 2, 24, 25, 7, 119, 2, 2, 25, 26, 7, 118, 2, 2, 26, 27, 7,
	114, 2, 2, 27, 28, 7, 119, 2, 2, 28, 29, 7, 118, 2, 2, 29, 6, 3, 2, 2,
	2, 30, 31, 7, 101, 2, 2, 31, 32, 7, 113, 2, 2, 32, 33, 7, 112, 2, 2, 33,
	34, 7, 117, 2, 2, 34, 35, 7, 118, 2, 2, 35, 8, 3, 2, 2, 2, 36, 37, 7, 120,
	2, 2, 37, 38, 7, 99, 2, 2, 38, 39, 7, 116, 2, 2, 39, 10, 3, 2, 2, 2, 40,
	43, 5, 13, 7, 2, 41, 43, 7, 97, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2,
	2, 2, 43, 49, 3, 2, 2, 2, 44, 48, 5, 13, 7, 2, 45, 48, 7, 97, 2, 2, 46,
	48, 5, 17, 9, 2, 47, 44, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2,
	2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 12,
	3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 9, 2, 2, 2, 53, 14, 3, 2, 2, 2,
	54, 58, 9, 3, 2, 2, 55, 57, 5, 17, 9, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3,
	2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 63, 3, 2, 2, 2, 60,
	58, 3, 2, 2, 2, 61, 63, 7, 50, 2, 2, 62, 54, 3, 2, 2, 2, 62, 61, 3, 2,
	2, 2, 63, 16, 3, 2, 2, 2, 64, 65, 9, 4, 2, 2, 65, 18, 3, 2, 2, 2, 66, 68,
	9, 5, 2, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2,
	69, 70, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 8, 10, 2, 2, 72, 20, 3,
	2, 2, 2, 9, 2, 42, 47, 49, 58, 62, 69, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'output'", "'const'", "'var'",
}

var lexerSymbolicNames = []string{
	"", "ASSIGN", "OUTPUT", "CONST", "VAR", "IDENTIFIER", "LETTER", "INTEGER",
	"WS",
}

var lexerRuleNames = []string{
	"ASSIGN", "OUTPUT", "CONST", "VAR", "IDENTIFIER", "LETTER", "INTEGER",
	"DIGIT", "WS",
}

type MeritLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMeritLexer(input antlr.CharStream) *MeritLexer {

	l := new(MeritLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Merit.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MeritLexer tokens.
const (
	MeritLexerASSIGN     = 1
	MeritLexerOUTPUT     = 2
	MeritLexerCONST      = 3
	MeritLexerVAR        = 4
	MeritLexerIDENTIFIER = 5
	MeritLexerLETTER     = 6
	MeritLexerINTEGER    = 7
	MeritLexerWS         = 8
)
