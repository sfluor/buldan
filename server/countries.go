package main

import (
	"fmt"
	"maps"
	"sort"

	mrand "math/rand"
)

const minScore = 0.85
const minScoreExact = 0.99

type rawCountry struct {
	name string
	flag string
}

type Country struct {
	Name           string
	NormalizedName string
	Flag           string
}

type CountryStatus struct {
	Country
	GuessedBy *string
}

type Countries struct {
	byName    map[string]Country
	guessedBy map[string]string
	all       []Country
}

func keys[T comparable, V any](m map[T]V) []T {
	res := make([]T, 0, len(m))
	for k := range m {
		res = append(res, k)
	}

	return res
}

func getLang(lang Language) (map[byte]map[string]Country, error) {
	countries, ok := byLang[lang]
	if !ok {
		return nil, fmt.Errorf("Unknown language: %s only has: %s", lang, keys(byLang))
	}
	return countries, nil
}

func countriesStartingWith(lang Language, char byte) (Countries, error) {
	countries, err := getLang(lang)
	if err != nil {
		return Countries{}, err
	}

	byName := countries[char]
	all := make([]Country, 0, len(byName))

	for _, c := range byName {
		all = append(all, c)
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].NormalizedName < all[j].NormalizedName
	})

	return Countries{byName: maps.Clone(byName), all: all, guessedBy: map[string]string{}}, nil
}

func (c *Countries) remaining() int {
	return len(c.byName)
}

func (c *Countries) status() []CountryStatus {
	res := make([]CountryStatus, 0, len(c.all))

	for _, country := range c.all {
		status := CountryStatus{Country: country}
		guesser, ok := c.guessedBy[country.Name]
		if ok {
			status.GuessedBy = &guesser
		}

		res = append(res, status)
	}

	return res
}

func (c *Countries) guess(guess string, from string) (Country, string, bool) {
	normalizedQuery := normalize(guess)

	// countryScore holds a country name and its similarity score
	type countryScore struct {
		key     string
		country Country
		score   float64
	}

	match := countryScore{}

	for name, country := range c.byName {
		// Technically no need to renormalize
		normalizedCountry := normalize(name)
		score := jaroWinkler(normalizedQuery, normalizedCountry)
		if match.score < score {
			match = countryScore{key: name, country: country, score: score}
		}
	}

	valid := match.score >= minScore
	guessStr := guess
	if valid {
		delete(c.byName, match.key)

		c.guessedBy[match.country.Name] = from

		if match.score < minScoreExact {
			guessStr = fmt.Sprintf("%s (%s)", guess, match.country.Name)
		}
	}

	return match.country, guessStr, valid
}

type Language string

const (
	LanguageFrench  Language = "french"
	LanguageEnglish Language = "english"
)

// Letter -> Normalized Name -> Country
var byLang = map[Language]map[byte]map[string]Country{}

func newLetters(lang Language) ([]byte, error) {
	countries, err := getLang(lang)
	if err != nil {
		return nil, err
	}

	letters := []byte{}
	for l := range countries {
		letters = append(letters, l)
	}

	mrand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})

	return letters, nil
}

func init() {

	for lang, rawCountries := range map[Language][]rawCountry{
		LanguageFrench:  frenchCountries,
		LanguageEnglish: englishCountries,
	} {
		countries := map[byte]map[string]Country{}
		for _, country := range rawCountries {
			normalized := normalize(country.name)
			firstLetter := normalized[0]

			if _, ok := countries[firstLetter]; !ok {
				countries[firstLetter] = map[string]Country{}
			}

			if prev, ok := countries[firstLetter][normalized]; ok {
				panic(fmt.Sprintf("Found duplicate countries: %+v and %+v", prev, country))
			}

			countries[firstLetter][normalized] = Country{
				Name:           country.name,
				Flag:           country.flag,
				NormalizedName: normalized,
			}
		}

		byLang[lang] = countries
	}
}
