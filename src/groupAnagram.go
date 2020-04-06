package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3288/
import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagrams := groupAnagrams(strs)

	fmt.Println(anagrams)
}

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	var dictionary map[string][]string
	dictionary = make(map[string][]string)

	for i := 0; i < length; i++ {
		str := strs[i]
		sortedStr := sortStr(str)
		dictionary[sortedStr] = append(dictionary[sortedStr], str)
	}
	keys := make([]string, 0, len(dictionary))
	for k := range dictionary {
		keys = append(keys, k)
	}

	result := make([][]string, 0)
	length = len(keys)
	for i := 0; i < length; i++ {
		strings, _ := dictionary[keys[i]]
		sort.Strings(strings)
		result = append(result, strings)
	}

	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortStr(str string) string {
	s := []rune(str)
	sort.Sort(sortRunes(s))
	return string(s)
}
