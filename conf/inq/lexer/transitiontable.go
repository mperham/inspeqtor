
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
			case r == 97 : // ['a','a']
				return 6
			case r == 98 : // ['b','b']
				return 5
			case r == 99 : // ['c','c']
				return 7
			case 100 <= r && r <= 101 : // ['d','e']
				return 5
			case r == 102 : // ['f','f']
				return 8
			case r == 103 : // ['g','g']
				return 9
			case r == 104 : // ['h','h']
				return 10
			case r == 105 : // ['i','i']
				return 11
			case r == 106 : // ['j','j']
				return 5
			case r == 107 : // ['k','k']
				return 9
			case r == 108 : // ['l','l']
				return 12
			case r == 109 : // ['m','m']
				return 9
			case 110 <= r && r <= 111 : // ['n','o']
				return 5
			case r == 112 : // ['p','p']
				return 13
			case r == 113 : // ['q','q']
				return 5
			case r == 114 : // ['r','r']
				return 14
			case r == 115 : // ['s','s']
				return 15
			case r == 116 : // ['t','t']
				return 16
			case r == 117 : // ['u','u']
				return 17
			case 118 <= r && r <= 122 : // ['v','z']
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
				return 18
			
			
			default:
				return 2
			}
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 3
			
			
			
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
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 107 : // ['a','k']
				return 21
			case r == 108 : // ['l','l']
				return 22
			case 109 <= r && r <= 122 : // ['m','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 103 : // ['a','g']
				return 21
			case r == 104 : // ['h','h']
				return 23
			case 105 <= r && r <= 120 : // ['i','x']
				return 21
			case r == 121 : // ['y','y']
				return 24
			case r == 122 : // ['z','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 110 : // ['a','n']
				return 21
			case r == 111 : // ['o','o']
				return 25
			case 112 <= r && r <= 122 : // ['p','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 110 : // ['a','n']
				return 21
			case r == 111 : // ['o','o']
				return 26
			case 112 <= r && r <= 122 : // ['p','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 101 : // ['a','e']
				return 21
			case r == 102 : // ['f','f']
				return 27
			case 103 <= r && r <= 109 : // ['g','m']
				return 21
			case r == 110 : // ['n','n']
				return 28
			case 111 <= r && r <= 122 : // ['o','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case r == 97 : // ['a','a']
				return 29
			case 98 <= r && r <= 122 : // ['b','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 113 : // ['a','q']
				return 21
			case r == 114 : // ['r','r']
				return 30
			case 115 <= r && r <= 122 : // ['s','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 31
			case 102 <= r && r <= 116 : // ['f','t']
				return 21
			case r == 117 : // ['u','u']
				return 32
			case 118 <= r && r <= 122 : // ['v','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 33
			case 102 <= r && r <= 120 : // ['f','x']
				return 21
			case r == 121 : // ['y','y']
				return 34
			case r == 122 : // ['z','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 103 : // ['a','g']
				return 21
			case r == 104 : // ['h','h']
				return 35
			case 105 <= r && r <= 122 : // ['i','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 111 : // ['a','o']
				return 21
			case r == 112 : // ['p','p']
				return 36
			case 113 <= r && r <= 122 : // ['q','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 37
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 38
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 98 : // ['a','b']
				return 21
			case r == 99 : // ['c','c']
				return 39
			case 100 <= r && r <= 122 : // ['d','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 113 : // ['a','q']
				return 21
			case r == 114 : // ['r','r']
				return 40
			case 115 <= r && r <= 122 : // ['s','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 41
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 104 : // ['a','h']
				return 21
			case r == 105 : // ['i','i']
				return 42
			case 106 <= r && r <= 114 : // ['j','r']
				return 21
			case r == 115 : // ['s','s']
				return 43
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 116 : // ['a','t']
				return 21
			case r == 117 : // ['u','u']
				return 44
			case 118 <= r && r <= 122 : // ['v','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 110 : // ['a','n']
				return 21
			case r == 111 : // ['o','o']
				return 45
			case 112 <= r && r <= 122 : // ['p','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 46
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 109 : // ['a','m']
				return 21
			case r == 110 : // ['n','n']
				return 47
			case 111 <= r && r <= 122 : // ['o','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 113 : // ['a','q']
				return 21
			case r == 114 : // ['r','r']
				return 48
			case 115 <= r && r <= 122 : // ['s','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 49
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 50
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 51
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 113 : // ['a','q']
				return 21
			case r == 114 : // ['r','r']
				return 52
			case 115 <= r && r <= 122 : // ['s','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 98 : // ['a','b']
				return 21
			case r == 99 : // ['c','c']
				return 53
			case 100 <= r && r <= 122 : // ['d','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 107 : // ['a','k']
				return 21
			case r == 108 : // ['l','l']
				return 54
			case 109 <= r && r <= 122 : // ['m','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 55
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 56
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 111 : // ['a','o']
				return 21
			case r == 112 : // ['p','p']
				return 57
			case 113 <= r && r <= 122 : // ['q','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 109 : // ['a','m']
				return 21
			case r == 110 : // ['n','n']
				return 58
			case 111 <= r && r <= 122 : // ['o','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 98 : // ['a','b']
				return 21
			case r == 99 : // ['c','c']
				return 59
			case 100 <= r && r <= 122 : // ['d','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 60
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 104 : // ['a','h']
				return 21
			case r == 105 : // ['i','i']
				return 61
			case 106 <= r && r <= 122 : // ['j','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 117 : // ['a','u']
				return 21
			case r == 118 : // ['v','v']
				return 62
			case 119 <= r && r <= 122 : // ['w','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 63
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 109 : // ['a','m']
				return 21
			case r == 110 : // ['n','n']
				return 64
			case 111 <= r && r <= 122 : // ['o','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S51
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 65
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S52
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 66
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S53
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 106 : // ['a','j']
				return 21
			case r == 107 : // ['k','k']
				return 67
			case 108 <= r && r <= 122 : // ['l','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S54
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 68
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S55
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S56
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case r == 46 : // ['.','.']
				return 69
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S57
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 70
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S58
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 98 : // ['a','b']
				return 21
			case r == 99 : // ['c','c']
				return 71
			case 100 <= r && r <= 122 : // ['d','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S59
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 72
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S60
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case r == 97 : // ['a','a']
				return 73
			case 98 <= r && r <= 122 : // ['b','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S61
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 74
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S62
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 104 : // ['a','h']
				return 21
			case r == 105 : // ['i','i']
				return 75
			case 106 <= r && r <= 122 : // ['j','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S63
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 76
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S64
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S65
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case r == 97 : // ['a','a']
				return 77
			case 98 <= r && r <= 122 : // ['b','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S66
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S67
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S68
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 78
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S69
		func(r rune) int {
			switch {
			case r == 100 : // ['d','d']
				return 79
			
			
			
			}
			return NoState
			
		},
	
		// S70
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 98 : // ['a','b']
				return 21
			case r == 99 : // ['c','c']
				return 80
			case 100 <= r && r <= 122 : // ['d','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S71
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 103 : // ['a','g']
				return 21
			case r == 104 : // ['h','h']
				return 81
			case 105 <= r && r <= 122 : // ['i','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S72
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 82
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S73
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 113 : // ['a','q']
				return 21
			case r == 114 : // ['r','r']
				return 83
			case 115 <= r && r <= 122 : // ['s','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S74
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S75
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 98 : // ['a','b']
				return 21
			case r == 99 : // ['c','c']
				return 84
			case 100 <= r && r <= 122 : // ['d','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S76
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 108 : // ['a','l']
				return 21
			case r == 109 : // ['m','m']
				return 85
			case 110 <= r && r <= 122 : // ['n','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S77
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 113 : // ['a','q']
				return 21
			case r == 114 : // ['r','r']
				return 86
			case 115 <= r && r <= 122 : // ['s','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S78
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S79
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S80
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 87
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S81
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 98 : // ['a','b']
				return 21
			case r == 99 : // ['c','c']
				return 88
			case 100 <= r && r <= 122 : // ['d','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S82
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 89
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S83
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 90
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S84
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 91
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S85
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 99 : // ['a','c']
				return 21
			case r == 100 : // ['d','d']
				return 92
			case 101 <= r && r <= 122 : // ['e','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S86
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 93
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S87
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S88
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 115 : // ['a','s']
				return 21
			case r == 116 : // ['t','t']
				return 94
			case 117 <= r && r <= 122 : // ['u','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S89
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S90
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S91
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S92
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S93
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S94
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 107 : // ['a','k']
				return 21
			case r == 108 : // ['l','l']
				return 95
			case 109 <= r && r <= 122 : // ['m','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S95
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
}
