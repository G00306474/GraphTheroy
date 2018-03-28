package main

 import (
	"fmt"
 )

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

 func main(){
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

	
 }