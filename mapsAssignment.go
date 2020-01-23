package main

import (
	"fmt"
	"strings"
)
func WordCount(str string) map[string]int {
    wordList := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range wordList {
        _, stringExists := counts[word]
        if stringExists {
            counts[word] += 1
        } else {
            counts[word] = 1
        }
    }
    return counts
}
func main(){
    strLine := "Pooja Suresh Biswa Pooja"
	fmt.Println(wordCount(strLine))	
}