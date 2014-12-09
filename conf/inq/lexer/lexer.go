package lexer

import (

	// "fmt"
	// "github.com/mperham/inspeqtor/conf/inq/util"

	"io/ioutil"
	"unicode/utf8"

	"github.com/mperham/inspeqtor/conf/inq/token"
)

const (
	NoState    = -1
	NumStates  = 72
	NumSymbols = 112
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {

	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)

	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {

		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)

		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '>'
1: '<'
2: 'c'
3: 'h'
4: 'e'
5: 'c'
6: 'k'
7: 's'
8: 'e'
9: 'r'
10: 'v'
11: 'i'
12: 'c'
13: 'e'
14: 'h'
15: 'o'
16: 's'
17: 't'
18: 'e'
19: 'x'
20: 'p'
21: 'o'
22: 's'
23: 'e'
24: ','
25: 'r'
26: 'e'
27: 's'
28: 't'
29: 'a'
30: 'r'
31: 't'
32: 'r'
33: 'e'
34: 'l'
35: 'o'
36: 'a'
37: 'd'
38: 'a'
39: 'l'
40: 'e'
41: 'r'
42: 't'
43: 'w'
44: 'i'
45: 't'
46: 'h'
47: ':'
48: '('
49: ')'
50: 'i'
51: 'f'
52: 't'
53: 'h'
54: 'e'
55: 'n'
56: 'f'
57: 'o'
58: 'r'
59: 'c'
60: 'y'
61: 'c'
62: 'l'
63: 'e'
64: 's'
65: '#'
66: '\n'
67: '_'
68: '-'
69: '.'
70: '/'
71: 'k'
72: 'm'
73: 'g'
74: 't'
75: 'p'
76: '%'
77: '!'
78: '$'
79: '%'
80: '&'
81: '''
82: '*'
83: '+'
84: '-'
85: '/'
86: '='
87: '?'
88: '^'
89: '_'
90: '`'
91: '{'
92: '|'
93: '}'
94: '~'
95: '.'
96: '@'
97: '\'
98: '"'
99: '"'
100: ' '
101: '\t'
102: '\n'
103: '\r'
104: 'a'-'z'
105: 'A'-'Z'
106: '0'-'9'
107: 'A'-'Z'
108: 'a'-'z'
109: '0'-'9'
110: \u0100-\U0010ffff
111: .

*/
