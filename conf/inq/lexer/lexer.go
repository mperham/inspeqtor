
package lexer

import (
	
	// "fmt"
	// "inspeqtor/conf/inq/util"
	
	"io/ioutil"
	"unicode/utf8"
	"inspeqtor/conf/inq/token"
)

const(
	NoState = -1
	NumStates = 62
	NumSymbols = 100
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
35: '('
36: ')'
37: 'i'
38: 'f'
39: 't'
40: 'h'
41: 'e'
42: 'n'
43: 'f'
44: 'o'
45: 'r'
46: 'c'
47: 'y'
48: 'c'
49: 'l'
50: 'e'
51: 's'
52: '#'
53: '\n'
54: '_'
55: '-'
56: '.'
57: '/'
58: 'k'
59: 'm'
60: 'g'
61: 't'
62: 'p'
63: '%'
64: '!'
65: '#'
66: '$'
67: '%'
68: '&'
69: '''
70: '*'
71: '+'
72: '-'
73: '/'
74: '='
75: '?'
76: '^'
77: '_'
78: '`'
79: '{'
80: '|'
81: '}'
82: '~'
83: '.'
84: '@'
85: '\'
86: '"'
87: '"'
88: ' '
89: '\t'
90: '\n'
91: '\r'
92: 'a'-'z'
93: 'A'-'Z'
94: '0'-'9'
95: 'A'-'Z'
96: 'a'-'z'
97: '0'-'9'
98: \u0100-\U0010ffff
99: .

*/
