package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	msgArr:= strings.Split(msg, " ")
	fmt.Println("msgArr",msgArr)
	for _,v := range msgArr{
		currWord := v
		wordFormedWithDelay := ""
		// fmt.Println("current word", currWord, len(currWord))
		for subIndex:=0;subIndex<len(currWord);subIndex++{
				// fmt.Println("somethig", string(currWord[subIndex]))
				for subSubIndex:=0;subSubIndex<=subIndex;subSubIndex++{
						wordFormedWithDelay = wordFormedWithDelay + string(currWord[subIndex])
				}
				
				// print(wordFormedWithDelay)
		}
		print(wordFormedWithDelay)
	}

	
}

func main() {
	msg := "Time to learn about Go strings!"

	slowDown(msg)
}