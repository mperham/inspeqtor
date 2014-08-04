
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
			case r == 40 : // ['(','(']
				return 3
			case r == 41 : // [')',')']
				return 4
			case r == 44 : // [',',',']
				return 5
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case r == 60 : // ['<','<']
				return 8
			case r == 62 : // ['>','>']
				return 8
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case r == 97 : // ['a','a']
				return 10
			case r == 98 : // ['b','b']
				return 9
			case r == 99 : // ['c','c']
				return 11
			case 100 <= r && r <= 101 : // ['d','e']
				return 9
			case r == 102 : // ['f','f']
				return 12
			case r == 103 : // ['g','g']
				return 9
			case r == 104 : // ['h','h']
				return 13
			case r == 105 : // ['i','i']
				return 14
			case 106 <= r && r <= 111 : // ['j','o']
				return 9
			case r == 112 : // ['p','p']
				return 15
			case r == 113 : // ['q','q']
				return 9
			case r == 114 : // ['r','r']
				return 16
			case r == 115 : // ['s','s']
				return 17
			case r == 116 : // ['t','t']
				return 18
			case 117 <= r && r <= 118 : // ['u','v']
				return 9
			case r == 119 : // ['w','w']
				return 19
			case 120 <= r && r <= 122 : // ['x','z']
				return 9
			
			
			
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
				return 20
			
			
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
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 37 : // ['%','%']
				return 22
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 102 : // ['a','f']
				return 9
			case r == 103 : // ['g','g']
				return 23
			case 104 <= r && r <= 106 : // ['h','j']
				return 9
			case r == 107 : // ['k','k']
				return 23
			case r == 108 : // ['l','l']
				return 9
			case r == 109 : // ['m','m']
				return 23
			case 110 <= r && r <= 111 : // ['n','o']
				return 9
			case r == 112 : // ['p','p']
				return 23
			case 113 <= r && r <= 115 : // ['q','s']
				return 9
			case r == 116 : // ['t','t']
				return 23
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 107 : // ['a','k']
				return 9
			case r == 108 : // ['l','l']
				return 24
			case 109 <= r && r <= 122 : // ['m','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 103 : // ['a','g']
				return 9
			case r == 104 : // ['h','h']
				return 25
			case 105 <= r && r <= 120 : // ['i','x']
				return 9
			case r == 121 : // ['y','y']
				return 26
			case r == 122 : // ['z','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 110 : // ['a','n']
				return 9
			case r == 111 : // ['o','o']
				return 27
			case 112 <= r && r <= 122 : // ['p','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 110 : // ['a','n']
				return 9
			case r == 111 : // ['o','o']
				return 28
			case 112 <= r && r <= 122 : // ['p','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 101 : // ['a','e']
				return 9
			case r == 102 : // ['f','f']
				return 29
			case 103 <= r && r <= 122 : // ['g','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 113 : // ['a','q']
				return 9
			case r == 114 : // ['r','r']
				return 30
			case 115 <= r && r <= 122 : // ['s','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 31
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 32
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 103 : // ['a','g']
				return 9
			case r == 104 : // ['h','h']
				return 33
			case 105 <= r && r <= 122 : // ['i','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 104 : // ['a','h']
				return 9
			case r == 105 : // ['i','i']
				return 34
			case 106 <= r && r <= 122 : // ['j','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 35
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 36
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 98 : // ['a','b']
				return 9
			case r == 99 : // ['c','c']
				return 37
			case 100 <= r && r <= 122 : // ['d','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 113 : // ['a','q']
				return 9
			case r == 114 : // ['r','r']
				return 38
			case 115 <= r && r <= 122 : // ['s','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 114 : // ['a','r']
				return 9
			case r == 115 : // ['s','s']
				return 39
			case 116 <= r && r <= 122 : // ['t','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 110 : // ['a','n']
				return 9
			case r == 111 : // ['o','o']
				return 40
			case 112 <= r && r <= 122 : // ['p','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 114 : // ['a','r']
				return 9
			case r == 115 : // ['s','s']
				return 41
			case 116 <= r && r <= 122 : // ['t','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 113 : // ['a','q']
				return 9
			case r == 114 : // ['r','r']
				return 42
			case 115 <= r && r <= 122 : // ['s','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 43
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 115 : // ['a','s']
				return 9
			case r == 116 : // ['t','t']
				return 44
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 113 : // ['a','q']
				return 9
			case r == 114 : // ['r','r']
				return 45
			case 115 <= r && r <= 122 : // ['s','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 98 : // ['a','b']
				return 9
			case r == 99 : // ['c','c']
				return 46
			case 100 <= r && r <= 122 : // ['d','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 107 : // ['a','k']
				return 9
			case r == 108 : // ['l','l']
				return 47
			case 109 <= r && r <= 122 : // ['m','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 115 : // ['a','s']
				return 9
			case r == 116 : // ['t','t']
				return 48
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 98 : // ['a','b']
				return 9
			case r == 99 : // ['c','c']
				return 49
			case 100 <= r && r <= 122 : // ['d','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 115 : // ['a','s']
				return 9
			case r == 116 : // ['t','t']
				return 50
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 117 : // ['a','u']
				return 9
			case r == 118 : // ['v','v']
				return 51
			case 119 <= r && r <= 122 : // ['w','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 109 : // ['a','m']
				return 9
			case r == 110 : // ['n','n']
				return 52
			case 111 <= r && r <= 122 : // ['o','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 103 : // ['a','g']
				return 9
			case r == 104 : // ['h','h']
				return 53
			case 105 <= r && r <= 122 : // ['i','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 115 : // ['a','s']
				return 9
			case r == 116 : // ['t','t']
				return 54
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 106 : // ['a','j']
				return 9
			case r == 107 : // ['k','k']
				return 55
			case 108 <= r && r <= 122 : // ['l','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 56
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 57
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case r == 97 : // ['a','a']
				return 58
			case 98 <= r && r <= 122 : // ['b','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S51
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 104 : // ['a','h']
				return 9
			case r == 105 : // ['i','i']
				return 59
			case 106 <= r && r <= 122 : // ['j','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S52
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S53
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S54
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S55
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S56
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 114 : // ['a','r']
				return 9
			case r == 115 : // ['s','s']
				return 60
			case 116 <= r && r <= 122 : // ['t','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S57
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 114 : // ['a','r']
				return 9
			case r == 115 : // ['s','s']
				return 61
			case 116 <= r && r <= 122 : // ['t','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S58
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 113 : // ['a','q']
				return 9
			case r == 114 : // ['r','r']
				return 62
			case 115 <= r && r <= 122 : // ['s','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S59
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 98 : // ['a','b']
				return 9
			case r == 99 : // ['c','c']
				return 63
			case 100 <= r && r <= 122 : // ['d','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S60
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S61
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 114 : // ['a','r']
				return 9
			case r == 115 : // ['s','s']
				return 64
			case 116 <= r && r <= 122 : // ['t','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S62
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 115 : // ['a','s']
				return 9
			case r == 116 : // ['t','t']
				return 65
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S63
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 100 : // ['a','d']
				return 9
			case r == 101 : // ['e','e']
				return 66
			case 102 <= r && r <= 122 : // ['f','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S64
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S65
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S66
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 6
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 21
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case r == 95 : // ['_','_']
				return 6
			case 97 <= r && r <= 122 : // ['a','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
}
