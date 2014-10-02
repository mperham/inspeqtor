
package lexer

import (
	
	// "fmt"
	// "github.com/mperham/redacted/conf/global/util"
	
	"io/ioutil"
	"unicode/utf8"
	"github.com/mperham/redacted/conf/global/token"
)

const(
	NoState = -1
	NumStates = 33
	NumSymbols = 62
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
0: 's'
1: 'e'
2: 'n'
3: 'd'
4: 'a'
5: 'l'
6: 'e'
7: 'r'
8: 't'
9: 's'
10: 'v'
11: 'i'
12: 'a'
13: 'w'
14: 'i'
15: 't'
16: 'h'
17: 't'
18: 'o'
19: ','
20: 's'
21: 'e'
22: 't'
23: '#'
24: '\n'
25: '!'
26: '#'
27: '$'
28: '%'
29: '&'
30: '''
31: '*'
32: '+'
33: '-'
34: '/'
35: '='
36: '?'
37: '^'
38: '_'
39: '`'
40: '{'
41: '|'
42: '}'
43: '~'
44: '@'
45: '.'
46: ':'
47: ';'
48: '<'
49: '>'
50: '\'
51: '"'
52: '"'
53: ' '
54: '\t'
55: '\n'
56: '\r'
57: 'A'-'Z'
58: 'a'-'z'
59: '0'-'9'
60: \u0100-\U0010ffff
61: .

*/
