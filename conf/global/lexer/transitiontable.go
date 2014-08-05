
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 35 : // ['#','#']
				return 2
			case r == 44 : // [',',',']
				return 3
			case 48 <= r && r <= 57 : // ['0','9']
				return 4
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 6
			case r == 97 : // ['a','a']
				return 7
			case 98 <= r && r <= 114 : // ['b','r']
				return 5
			case r == 115 : // ['s','s']
				return 8
			case r == 116 : // ['t','t']
				return 9
			case r == 117 : // ['u','u']
				return 5
			case r == 118 : // ['v','v']
				return 10
			case r == 119 : // ['w','w']
				return 11
			case 120 <= r && r <= 122 : // ['x','z']
				return 5
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 12
			
			
			default:
				return 2
			}
			
		},
	
		// S3
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 107 : // ['a','k']
				return 15
			case r == 108 : // ['l','l']
				return 17
			case 109 <= r && r <= 122 : // ['m','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 15
			case r == 101 : // ['e','e']
				return 18
			case 102 <= r && r <= 122 : // ['f','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 110 : // ['a','n']
				return 15
			case r == 111 : // ['o','o']
				return 19
			case 112 <= r && r <= 122 : // ['p','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 104 : // ['a','h']
				return 15
			case r == 105 : // ['i','i']
				return 20
			case 106 <= r && r <= 122 : // ['j','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 104 : // ['a','h']
				return 15
			case r == 105 : // ['i','i']
				return 21
			case 106 <= r && r <= 122 : // ['j','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 15
			case r == 101 : // ['e','e']
				return 22
			case 102 <= r && r <= 122 : // ['f','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 15
			case r == 110 : // ['n','n']
				return 23
			case 111 <= r && r <= 115 : // ['o','s']
				return 15
			case r == 116 : // ['t','t']
				return 24
			case 117 <= r && r <= 122 : // ['u','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case r == 97 : // ['a','a']
				return 25
			case 98 <= r && r <= 122 : // ['b','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 115 : // ['a','s']
				return 15
			case r == 116 : // ['t','t']
				return 26
			case 117 <= r && r <= 122 : // ['u','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 113 : // ['a','q']
				return 15
			case r == 114 : // ['r','r']
				return 27
			case 115 <= r && r <= 122 : // ['s','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 99 : // ['a','c']
				return 15
			case r == 100 : // ['d','d']
				return 28
			case 101 <= r && r <= 122 : // ['e','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 103 : // ['a','g']
				return 15
			case r == 104 : // ['h','h']
				return 29
			case 105 <= r && r <= 122 : // ['i','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 115 : // ['a','s']
				return 15
			case r == 116 : // ['t','t']
				return 30
			case 117 <= r && r <= 122 : // ['u','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 114 : // ['a','r']
				return 15
			case r == 115 : // ['s','s']
				return 31
			case 116 <= r && r <= 122 : // ['t','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 13
			case r == 34 : // ['"','"']
				return 13
			case r == 35 : // ['#','#']
				return 13
			case r == 36 : // ['$','$']
				return 13
			case r == 37 : // ['%','%']
				return 13
			case r == 38 : // ['&','&']
				return 13
			case r == 39 : // [''',''']
				return 13
			case r == 42 : // ['*','*']
				return 13
			case r == 43 : // ['+','+']
				return 13
			case r == 45 : // ['-','-']
				return 13
			case r == 46 : // ['.','.']
				return 13
			case r == 47 : // ['/','/']
				return 13
			case 48 <= r && r <= 57 : // ['0','9']
				return 14
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 13
			case r == 61 : // ['=','=']
				return 13
			case r == 63 : // ['?','?']
				return 13
			case r == 64 : // ['@','@']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 15
			case r == 94 : // ['^','^']
				return 13
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 15
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 13
			case r == 125 : // ['}','}']
				return 13
			case r == 126 : // ['~','~']
				return 13
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 13
			case r == 8216 : // [\u2018,\u2018]
				return 13
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 13
			
			
			
			}
			return NoState
			
		},
	
}
