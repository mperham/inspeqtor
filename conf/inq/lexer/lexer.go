
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
	NumStates = 66
	NumSymbols = 76
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
6: 'h'
7: 'o'
8: 's'
9: 't'
10: 'c'
11: 'h'
12: 'e'
13: 'c'
14: 'k'
15: 'i'
16: 'n'
17: 's'
18: 'p'
19: 'e'
20: 'c'
21: 't'
22: 'p'
23: 'r'
24: 'o'
25: 'c'
26: 'e'
27: 's'
28: 's'
29: 's'
30: 'e'
31: 'r'
32: 'v'
33: 'i'
34: 'c'
35: 'e'
36: 'r'
37: 'e'
38: 's'
39: 't'
40: 'a'
41: 'r'
42: 't'
43: 'a'
44: 'l'
45: 'e'
46: 'r'
47: 't'
48: 'i'
49: 'f'
50: 't'
51: 'h'
52: 'e'
53: 'n'
54: 'f'
55: 'o'
56: 'r'
57: 'c'
58: 'y'
59: 'c'
60: 'l'
61: 'e'
62: 's'
63: '#'
64: '\n'
65: '_'
66: '-'
67: '.'
68: ' '
69: '\t'
70: '\n'
71: '\r'
72: 'a'-'z'
73: 'A'-'Z'
74: '0'-'9'
75: .

*/
