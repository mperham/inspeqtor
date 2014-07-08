
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
			case 48 <= r && r <= 57 : // ['0','9']
				return 3
			case r == 62 : // ['>','>']
				return 4
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case 97 <= r && r <= 98 : // ['a','b']
				return 5
			case r == 99 : // ['c','c']
				return 6
			case 100 <= r && r <= 104 : // ['d','h']
				return 5
			case r == 105 : // ['i','i']
				return 7
			case 106 <= r && r <= 115 : // ['j','s']
				return 5
			case r == 116 : // ['t','t']
				return 8
			case 117 <= r && r <= 122 : // ['u','z']
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
				return 9
			
			
			default:
				return 2
			}
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 3
			case 65 <= r && r <= 90 : // ['A','Z']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 10
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 103 : // ['a','g']
				return 13
			case r == 104 : // ['h','h']
				return 14
			case 105 <= r && r <= 122 : // ['i','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 101 : // ['a','e']
				return 13
			case r == 102 : // ['f','f']
				return 15
			case 103 <= r && r <= 122 : // ['g','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 103 : // ['a','g']
				return 13
			case r == 104 : // ['h','h']
				return 16
			case 105 <= r && r <= 122 : // ['i','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 10
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 100 : // ['a','d']
				return 13
			case r == 101 : // ['e','e']
				return 17
			case 102 <= r && r <= 122 : // ['f','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 100 : // ['a','d']
				return 13
			case r == 101 : // ['e','e']
				return 18
			case 102 <= r && r <= 122 : // ['f','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 98 : // ['a','b']
				return 13
			case r == 99 : // ['c','c']
				return 19
			case 100 <= r && r <= 122 : // ['d','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 109 : // ['a','m']
				return 13
			case r == 110 : // ['n','n']
				return 20
			case 111 <= r && r <= 122 : // ['o','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 106 : // ['a','j']
				return 13
			case r == 107 : // ['k','k']
				return 21
			case 108 <= r && r <= 122 : // ['l','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 13
			
			
			
			}
			return NoState
			
		},
	
}
