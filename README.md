# GraphTheroy

This is my third year project for Graph Theroy 
The project is to build a program developed in Go that can build a non-deterministic finite automaton
 (NFA) from a regular expression and check it against a given string.
 
#Brief

The following project concerns a real-world problem – one that you could
be asked to solve in industry. You are not expected to know how to do it
off the top of your head. Rather, it is expected that you will research and
investigate possible ways to tackle the problem, and then come up with your
own solution based on those. A quick search for solutions online will convince
you that many people have written solutions to the problem already, in many
different programming languages, and many of those are not experienced
software developers. Note that a series of videos will be provided to students
on the course page to help them with the project.
Problem statement
You must write a program in the Go programming language that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “*
” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.
Page 1 of 5
Graph Theory Project 2018
Other special characters you might consider allowing as input are brackets
“()” which can be used for grouping, “+” which means “at least one of”, and
“?” which means “zero or one of”. You might also decide to remove the
concatenation character, so that 1.0 becomes 10, with the concatenation
implicit.
You may initially restrict the non-special characters your program works
with to 0 and 1, if you wish. However, you should at least attempt to expand
these to all of the digits, and the characters a to z, and A to Z.
You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed. You might do that for this project as follows:
1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.
Overall your program might have the following layout.
type nfa struct {
...
}
func regexcompile(r string) nfa {
...
return n
}
func (n nfa) regexmatch(n nfa, r sting) bool {
...
return ismatch
}
func main() {
n := regexcompile("01*0")
t := n.regexmatch("01110")
f := n.regexmatch("1000001")
}

#Research 

#Run Project 
Go compiler can be installed from https://golang.org/doc/install

#Useful links
Google. The go programming language.
https://golang.org/.
Regular Expression Matching Can Be Simple And Fast 
(but is slow in Java, Perl, PHP, Python, Ruby, ...)
https://swtch.com/~rsc/regexp/regexp1.html