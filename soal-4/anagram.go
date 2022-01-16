package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "soal4",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) > 0 {
			filterToAnagram(args)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func filterToAnagram(data []string) {
	var anagrams [][]string
	var tmpAnagrams []string
	for _, word := range data {
		siblings := validateRangeAnagram(data, word)

		if siblings == nil {
			continue
		}

		if ok := existingAnagrams(tmpAnagrams, siblings); ok {
			continue
		}

		// before append, check the data is already appended
		tmpAnagrams = append(tmpAnagrams, siblings...)
		anagrams = append(anagrams, siblings)
	}

	x, _ := json.MarshalIndent(anagrams, " ", " ")
	log.Printf("data:\n %v", string(x))
}

func existingAnagrams(anagrams []string, newAnagrams []string) bool {
	for _, ags := range anagrams {
		if ok := existingAnagram(ags, newAnagrams); ok {
			return true
		}
	}

	return false
}

func existingAnagram(anagram string, newAnagrams []string) bool {
	for _, nAnagram := range newAnagrams {
		if nAnagram == anagram {
			return true
		}
	}

	return false
}

// validateRangeAnagram
// this find another anagram with same length
// and validation the anagram identic or not
func validateRangeAnagram(data []string, anagram string) []string {
	var siblings []string
	for _, word := range data {
		if len(word) != len(anagram) {
			continue
		}

		if validateAnagramRune(word, anagram) {
			siblings = append(siblings, word)
		}
	}

	return siblings
}

// validateAnagramRune
// this will be check every alphabetic of the words
func validateAnagramRune(data, anagram string) bool {
	for _, d := range data {
		if !strings.ContainsRune(anagram, rune(d)) {
			return false
		}
	}

	return true
}

func main() {
	Execute()
}
