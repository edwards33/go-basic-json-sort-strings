package main

import (
	"fmt"
	"sort"
	"strings"
)

func compareStrings(a, b string) bool{
	if strings.Compare(a, b) == -1 {
		return true
	} 
	return false
}

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return compareStrings(p[i], p[j]) }

func main() {
	studyGroup := people{"William", "Jack", "Liam", "Charlie", "Harry"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

}
