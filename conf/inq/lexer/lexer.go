
package lexer

import (
	
	// "fmt"
	// "github.com/mperham/redacted/conf/inq/util"
	
	"io/ioutil"
	"unicode/utf8"
	"github.com/mperham/redacted/conf/inq/token"
)

const(
	NoState = -1
	NumStates = 63
	NumSymbols = 101
) 

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
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
2: 'h'
3: 'o'
4: 's'
5: 't'
6: 'c'
7: 'h'
8: 'e'
9: 'c'
10: 'k'
11: 's'
12: 'e'
13: 'r'
14: 'v'
15: 'i'
16: 'c'
17: 'e'
18: ','
19: 'r'
20: 'e'
21: 's'
22: 't'
23: 'a'
24: 'r'
25: 't'
26: 'a'
27: 'l'
28: 'e'
29: 'r'
30: 't'
31: 'w'
32: 'i'
33: 't'
34: 'h'
35: ':'
36: '('
37: ')'
38: 'i'
39: 'f'
40: 't'
41: 'h'
42: 'e'
43: 'n'
44: 'f'
45: 'o'
46: 'r'
47: 'c'
48: 'y'
49: 'c'
50: 'l'
51: 'e'
52: 's'
53: '#'
54: '\n'
55: '_'
56: '-'
57: '.'
58: '/'
59: 'k'
60: 'm'
61: 'g'
62: 't'
63: 'p'
64: '%'
65: '!'
66: '#'
67: '$'
68: '%'
69: '&'
70: '''
71: '*'
72: '+'
73: '-'
74: '/'
75: '='
76: '?'
77: '^'
78: '_'
79: '`'
80: '{'
81: '|'
82: '}'
83: '~'
84: '.'
85: '@'
86: '\'
87: '"'
88: '"'
89: ' '
90: '\t'
91: '\n'
92: '\r'
93: 'a'-'z'
94: 'A'-'Z'
95: '0'-'9'
96: 'A'-'Z'
97: 'a'-'z'
98: '0'-'9'
99: \u0100-\U0010ffff
100: .

*/
