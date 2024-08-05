package main

import (
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// normalize function converts text to lowercase and replaces accented characters with their ASCII equivalents
func normalize(text string) string {
	// Convert accented characters to their ASCII equivalents
	t := norm.NFD.String(text)
	re := regexp.MustCompile(`[\p{Mn}]`)
	t = re.ReplaceAllString(t, "")

	// Convert to lowercase
	return strings.ToLower(t)
}

// jaro calculates the Jaro similarity score between two strings
func jaro(s1, s2 string) float64 {
	s1Len := len(s1)
	s2Len := len(s2)

	if s1Len == 0 && s2Len == 0 {
		return 1
	}
	if s1Len == 0 || s2Len == 0 {
		return 0
	}

	matchDistance := (max(s1Len, s2Len) / 2) - 1

	s1Matches := make([]bool, s1Len)
	s2Matches := make([]bool, s2Len)

	matches := 0
	transpositions := 0

	for i := 0; i < s1Len; i++ {
		start := max(0, i-matchDistance)
		end := min(i+matchDistance+1, s2Len)

		for j := start; j < end; j++ {
			if s2Matches[j] {
				continue
			}
			if s1[i] != s2[j] {
				continue
			}
			s1Matches[i] = true
			s2Matches[j] = true
			matches++
			break
		}
	}

	if matches == 0 {
		return 0
	}

	k := 0
	for i := 0; i < s1Len; i++ {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if s1[i] != s2[k] {
			transpositions++
		}
		k++
	}

	transpositions /= 2

	return (float64(matches)/float64(s1Len) +
		float64(matches)/float64(s2Len) +
		float64(matches-transpositions)/float64(matches)) / 3.0
}

// jaroWinkler calculates the Jaro-Winkler similarity score between two strings
func jaroWinkler(s1, s2 string) float64 {
	jaroScore := jaro(s1, s2)
	prefixLength := 0
	maxPrefixLength := 4

	for i := 0; i < min(min(len(s1), len(s2)), maxPrefixLength); i++ {
		if s1[i] == s2[i] {
			prefixLength++
		} else {
			break
		}
	}

	return jaroScore + float64(prefixLength)*0.1*(1.0-jaroScore)
}

// Utility functions for min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
