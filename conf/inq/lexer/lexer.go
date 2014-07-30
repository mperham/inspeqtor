
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
	NumStates = 69
	NumSymbols = 80
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
5: '%'
6: '>'
7: '<'
8: 'h'
9: 'o'
10: 's'
11: 't'
12: 'c'
13: 'h'
14: 'e'
15: 'c'
16: 'k'
17: 'i'
18: 'n'
19: 's'
20: 'p'
21: 'e'
22: 'c'
23: 't'
24: 'p'
25: 'r'
26: 'o'
27: 'c'
28: 'e'
29: 's'
30: 's'
31: 's'
32: 'e'
33: 'r'
34: 'v'
35: 'i'
36: 'c'
37: 'e'
38: 'r'
39: 'e'
40: 's'
41: 't'
42: 'a'
43: 'r'
44: 't'
45: 'a'
46: 'l'
47: 'e'
48: 'r'
49: 't'
50: '('
51: ')'
52: 'i'
53: 'f'
54: 't'
55: 'h'
56: 'e'
57: 'n'
58: 'f'
59: 'o'
60: 'r'
61: 'c'
62: 'y'
63: 'c'
64: 'l'
65: 'e'
66: 's'
67: '#'
68: '\n'
69: '_'
70: '-'
71: '.'
72: ' '
73: '\t'
74: '\n'
75: '\r'
76: 'a'-'z'
77: 'A'-'Z'
78: '0'-'9'
79: .

*/
