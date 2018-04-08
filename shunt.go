/*Code Adapted from 
https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b 
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
By Ian Mcloughlan
*//*
Student Name: Kevin Moran Student Number G00306474
*/
package main

 import (
	"fmt"
	"os"
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
			//only pop one frag off the nfa stack
			frag :=nfastack[len(nfastack)-1]
			nfastack= nfastack[:len(nfastack)-1]

			accept:= state{}
			initial := state {edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa {initial: &initial, accept: &accept})
		case '+':
			frag :=nfastack[len(nfastack)-1]
			nfastack= nfastack[:len(nfastack)-1]

			accept:= state{}
			initial := state {edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = &initial
			//frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default:
			//not a special charcter 
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa {initial: &initial, accept: &accept})
		}//closes switch
	}//closes for
	/*
	if len(nfastack) != 1 {
		fmt.Println("Uh oh!", len(nfastack), nfastack)
	}//closes if
	*/
	return nfastack[0]
}
//Thompson  End




//Shunt Start
 func intopost(infix string) string {

	specials := map[rune]int{'*':10,'+':9,'.':8,'|':7}

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
func addstate(l []*state, s *state, a *state) []*state{
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addstate(l, s.edge1, a)
		if s.edge2 != nil{
			l = addstate(l, s.edge2, a)
		}
	}// if
	return l
}
//Pomatch checks if string is a match
func pomatch(po string, s string) bool {
	ismatch := false
	ponfa := poregtonfa(po)
	
	current := []*state{}
	next :=[]*state{}

	current = addstate(current[:], ponfa.initial, ponfa.accept)


	for _, r := range s {
		for _, c := range current{
			if c.symbol ==r{
				//adds c state and any other state i can get to 
				next = addstate(next[:], c.edge1, ponfa.accept)
			}
		}
		//gets next current state adds it to current 
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept{
			ismatch = true
			break
		}
	}
	return ismatch

}
//
 func main(){

	run();
 }
 func run (){
	var userChoice string
	for {
		
		fmt.Println("1. User enters Test Condition againt String ")
		fmt.Println("2. Test shunt, thompson & Rega with hardcoded tests ")
		fmt.Println("3. Exit  ")
		fmt.Scanln(&userChoice)
		switch userChoice {
		case "1":
			//Run the test
			userInput()
		case "2":
			noUserInput()
		case "3":
			//Close the program
			os.Exit(3)
		default:
			fmt.Println("Error please choose 1,2 or 3")
		}
	}
 }
 func userInput(){

	 var condition string
	 var testString string
 
	 fmt.Println("Please enter the Condition")
	 fmt.Scanln(&condition)
	 fmt.Println("Please enter the test String")
	 fmt.Scanln(&testString)

	 fmt.Println("Infix: ", condition)
	 fmt.Println("PoFix: ",intopost(condition))
	 fmt.Println("Match: ", pomatch(condition, testString))
	 fmt.Println()
 }
 func noUserInput(){	
	//Shunt test begin
	 fmt.Println("Shunt Test")
	//Expencted Answer: ab.c*.
	fmt.Println("Infix:   ", "a.b.c*")
	fmt.Println("Postfix:   ",intopost( "a.b.c*"))

	//Expencted Answer: abd|.*
	fmt.Println("Infix:   ", "(a.(b|d))*")
	fmt.Println("Postfix:   ",intopost( "(a.(b|d))*"))

	//Expencted Answer: abd|.c*
	fmt.Println("Infix:   ", "a.(b|d)c*")
	fmt.Println("Postfix:   ",intopost( "a.(b|d)c*"))

	//Expencted Answer: abb.+.c.
	fmt.Println("Infix:   ", "a.(b.b)+.c")
	fmt.Println("Postfix:   ",intopost( "a.(b.b)+.c"))
 //Shunt test end
 
	fmt.Println("Thompson algo test being given ab.c*|")
	//NFA  test begin
	//a followed by b and any amount of C 
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
	//NFA  test end
	
	fmt.Println("Match test case ab.c*| cccc")
	fmt.Println(pomatch("ab.c*|","cccc"))
}