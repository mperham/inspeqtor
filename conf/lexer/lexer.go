
package lexer

import (
	
	// "fmt"
	// "inspeqtor/conf/util"
	
	"io/ioutil"
	"unicode/utf8"
	"inspeqtor/conf/token"
)

const(
	NoState = -1
	NumStates = 91
	NumSymbols = 102
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
0: 'k'
1: 'm'
2: 'g'
3: 't'
4: 'p'
5: '>'
6: 'c'
7: 'h'
8: 'e'
9: 'c'
10: 'k'
11: 'h'
12: 'o'
13: 's'
14: 't'
15: 'p'
16: 'r'
17: 'o'
18: 'c'
19: 'e'
20: 's'
21: 's'
22: 's'
23: 'e'
24: 'r'
25: 'v'
26: 'i'
27: 'c'
28: 'e'
29: 'r'
30: 'e'
31: 's'
32: 't'
33: 'a'
34: 'r'
35: 't'
36: 'a'
37: 'l'
38: 'e'
39: 'r'
40: 't'
41: 'l'
42: 'a'
43: 'u'
44: 'n'
45: 'c'
46: 'h'
47: 'c'
48: 't'
49: 'l'
50: 'u'
51: 'p'
52: 's'
53: 't'
54: 'a'
55: 'r'
56: 't'
57: 'r'
58: 'u'
59: 'n'
60: 'i'
61: 't'
62: 's'
63: 'y'
64: 's'
65: 't'
66: 'e'
67: 'm'
68: 'd'
69: 'i'
70: 'n'
71: 'i'
72: 't'
73: '.'
74: 'd'
75: 'i'
76: 'f'
77: 't'
78: 'h'
79: 'e'
80: 'n'
81: 'f'
82: 'o'
83: 'r'
84: 'c'
85: 'y'
86: 'c'
87: 'l'
88: 'e'
89: 's'
90: '#'
91: '\n'
92: '_'
93: '-'
94: ' '
95: '\t'
96: '\n'
97: '\r'
98: 'a'-'z'
99: 'A'-'Z'
100: '0'-'9'
101: .

*/
