package main

import (
	"fmt"

	"github.com/MohamedGouaouri/automata-go/automata/fa"
)

func main() {
	alphabet := []string{"0", "1"}
	states := []string{"q0", "q1"}
	initial := "q0"
	final := []string{"q1"}
	transitions := make(map[string][]map[string]string)

	trans := make(map[string]string)
	trans["0"] = "q0"
	trans["1"] = "q1"
	transitions["q0"] = append(transitions["q0"], trans)

	trans = make(map[string]string)
	trans["0"] = "q1"
	trans["1"] = "q0"

	transitions["q1"] = append(transitions["q1"], trans)

	dfa := fa.NewDFA(alphabet, states, initial, final, transitions)
	fmt.Println(dfa.VerifyInputAccepted("0100"))
}
