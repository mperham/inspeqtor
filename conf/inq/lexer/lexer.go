
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
	NumStates = 67
	NumSymbols = 79
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
11: 'p'
12: 'r'
13: 'o'
14: 'c'
15: 'e'
16: 's'
17: 's'
18: 's'
19: 'e'
20: 'r'
21: 'v'
22: 'i'
23: 'c'
24: 'e'
25: 'r'
26: 'e'
27: 's'
28: 't'
29: 'a'
30: 'r'
31: 't'
32: 'a'
33: 'l'
34: 'e'
35: 'r'
36: 't'
37: 'w'
38: 'i'
39: 't'
40: 'h'
41: ','
42: '('
43: ')'
44: 'i'
45: 'f'
46: 't'
47: 'h'
48: 'e'
49: 'n'
50: 'f'
51: 'o'
52: 'r'
53: 'c'
54: 'y'
55: 'c'
56: 'l'
57: 'e'
58: 's'
59: '#'
60: '\n'
61: '_'
62: '-'
63: '.'
64: '/'
65: 'k'
66: 'm'
67: 'g'
68: 't'
69: 'p'
70: '%'
71: ' '
72: '\t'
73: '\n'
74: '\r'
75: 'a'-'z'
76: 'A'-'Z'
77: '0'-'9'
78: .

*/
