/* This file contains definition of DFA */

package fa

type DFA struct {
	Alphabet     []string
	States       []string
	InitialState string
	FinalStates  []string
	Transitions  map[string][]map[string]string // of the form {'q_i': [{'w_i': 'q_j'}, ...], ...}
}

func (dfa *DFA) VerifyDFA() bool {
	return true
}

// VerifyAlphabet verifies if the string s contains alphabet of the language
func (dfa *DFA) VerifyAlphabet(s string) bool {
	return false
}

// VerifyInputAccepted verifies if the input string s is is accepted by the DFA
func (dfa *DFA) VerifyInputAccepted(s string) bool {
	state := dfa.InitialState

	// for every char in string
	for i := 0; i < len(s); i++ {
		// for every transition
		// for every
		var go_to string
		for j := 0; j < len(dfa.Transitions[state]); j++ {

			for c, s := range dfa.Transitions[state][j] {
				if c == string(s[i]) {
					go_to = s
					break
				}
			}
			if go_to != "" {
				break
			}
		}
		state = go_to
	}
	for _, fs := range dfa.FinalStates {
		if fs == state {
			return true
		}
	}
	return false
}

func (dfa *DFA) GenerateGraph() {

}

// newDFA create new DFA
func NewDFA(alphabet, states []string, initial string, final []string, transitions map[string][]map[string]string) *DFA {
	return &DFA{
		Alphabet:     alphabet,
		States:       states,
		InitialState: initial,
		FinalStates:  final,
		Transitions:  transitions,
	}
}
