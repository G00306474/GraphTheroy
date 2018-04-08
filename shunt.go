package main

 import (
	"fmt"
 )


//Thompson  Start

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct{
	initial *state
	accept *state

}
//poregtonfa = post fix regular expression to non dertiminenistic finite atonamon
//takes string as input 
func poregtonfa(postfix string) *nfa {
	nfastack := []*nfa{}
	for _, r := range postfix {
	
		switch r {
		case '.':
			//pops 2 fragments off the stack of nfa fragments 
			frag2 := nfastack[len(nfastack)-1]
			nfastack= nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack= nfastack[:len(nfastack)-1]

			//joins the accept state of the first one to the inital state of thw other fragment 
			//thats the conncatanation
			frag1.accept.edge1 = frag2.initial

			//pushes new fragment onto the nfa stack (frag1 initial state with frag2 accepts state)
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		
		case '|':
				//pops 2 fragments off the stack of nfa fragments 
				frag2 :=nfastack[len(nfastack)-1]
				nfastack= nfastack[:len(nfastack)-1]
				frag1 :=nfastack[len(nfastack)-1]
				nfastack= nfastack[:len(nfastack)-1]

				//new accept state
				//join fraag1 and frag2 accet states to that
				//create new inital state
				accept:= state{}
				initial := state{edge1: frag1.initial, edge2: frag2.initial}
				frag1.accept.edge1 = &accept
				frag2.accept.edge1 = &accept
	
				 //push to stack 
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*':
			frag :=nfastack[len(nfastack)-1]
			nfastack= nfastack[:len(nfastack)-1]

			accept:= state{}
			initial := state {edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa {initial: &initial, accept: &accept})

		default:
			//not a special charcter 
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa {initial: &initial, accept: &accept})
		}//closes switch
	}//closes for 
	return nfastack[0]
}




//Thompson  End




//Shunt Start
 func intopost(infix string) string {

	specials := map[rune]int{'*':10,'.':9,'|':8}

	pofix := []rune{}
	s := []rune{}

	for _, r:= range infix{
		switch{

		case r == '(':// if we read an open bracket we will add it to the stack 
			s = append(s,r)
		case r == ')':// if we read an closing bracket we will remove everything till we hit an open bracket on the stack
			for s[len(s)-1] != '('{
				pofix = append(pofix, s[len(s)-1])
				s = s[:len(s)-1]//takes everything but the last charcter and makes it equal to s 
			}
			s = s[:len(s)-1]//discards open bracket
		case specials [r] > 0: //gets the value 0 back if not in specials 
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix = append(pofix, s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = append(s, r)

		default:

			pofix = append(pofix, r)
		}

	}
	for len(s) > 0 {
		
			pofix = append(pofix, s[len(s)-1])
			s = s[:len(s)-1]// puts top element of stack and puts it in pofix
	

	}
	return string(pofix)//changes rune back  to string 
 }
//Shunt End
//Rega checks if string is a match
func pomatch(po string, s string) bool {
	ismatch := false
	ponfa := poregtonfa(po)
	
	current := []*state{}
	next :=[]*staate{}



	for _, r := range s {
		for _, c := current{
			if c.symbol ==r{
				//adds c state and any other state i can get to 
			}
		}
		//gets next current state adds it to current 
		current, next = next, []*state{}
	}

	for _, c := range current {
		
	}
	return ismatch

}
//
 func main(){
	 //Shunt test begin
	//Answer: ab.c*.
	fmt.Println("Infix:   ", "a.b.c*")
	fmt.Println("Postfix:   ",intopost( "a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix:   ", "(a.(b|d))*")
	fmt.Println("Postfix:   ",intopost( "(a.(b|d))*"))

	//Answer: abd|.c*
	fmt.Println("Infix:   ", "a.(b|d)c*")
	fmt.Println("Postfix:   ",intopost( "a.(b|d)c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix:   ", "a.(b.b)+.c")
	fmt.Println("Postfix:   ",intopost( "a.(b.b)+.c"))
 //Shunt test end

	//NFA  test begin
		//a followed by b and any amount of C 
		nfa := poregtonfa("ab.c*|")
		fmt.Println(nfa)
		//NFA  test end

		fmt.Println(pomatch("ab.c*|","cccc"))
 }
 