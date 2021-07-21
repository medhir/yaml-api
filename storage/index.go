package storage

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/kljensen/snowball"
)

// Indexing & search implementation inspired by https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/

type attribute string

const (
	Title           = attribute("title")
	Version         = attribute("version")
	MaintainerName  = attribute("maintainer_name")
	MaintainerEmail = attribute("maintainer_email")
	Company         = attribute("company")
	Website         = attribute("website")
	Source          = attribute("source")
	License         = attribute("license")
	Description     = attribute("description")
)

type index struct {
	title           map[string][]*Metadata
	version         map[string][]*Metadata
	maintainerName  map[string][]*Metadata
	maintainerEmail map[string][]*Metadata
	company         map[string][]*Metadata
	website         map[string][]*Metadata
	source          map[string][]*Metadata
	license         map[string][]*Metadata
	description     map[string][]*Metadata
}

func indexField(text string, field map[string][]*Metadata, metadata *Metadata, tokenize bool) error {
	if tokenize {
		tokens, err := processText(text)
		if err != nil {
			return err
		}
		for _, token := range tokens {
			_, ok := field[token]
			if !ok {
				field[token] = []*Metadata{}
			}
			// the particular token being assessed has already been indexed for this document
			// so do not add a duplicate reference to the document
			if len(field[token]) > 0 && field[token][len(field[token])-1] == metadata {
				continue
			}
			field[token] = append(field[token], metadata)
		}
	} else {
		_, ok := field[text]
		if !ok {
			field[text] = []*Metadata{}
		}
		field[text] = append(field[text], metadata)
	}
	return nil
}

func matchAllTerms(resultSet map[string][]*Metadata) []*Metadata {
	keys := []string{}
	for k := range resultSet {
		keys = append(keys, k)
	}
	if len(keys) <= 0 {
		return []*Metadata{}
	} else if len(keys) == 1 {
		return resultSet[keys[0]]
	}

	result := intersection(resultSet[keys[0]], resultSet[keys[1]])
	if len(keys) > 2 {
		for i := 2; i < len(keys); i++ {
			newResult := intersection(result, resultSet[keys[i]])
			result = newResult
		}
	}
	return result
}

func intersection(slice1 []*Metadata, slice2 []*Metadata) []*Metadata {
	intersection := []*Metadata{}
	exists := map[*Metadata]bool{}
	for _, md := range slice1 {
		exists[md] = true
	}
	for _, md := range slice2 {
		if _, ok := exists[md]; ok {
			intersection = append(intersection, md)
		}
	}
	return intersection
}

func processText(text string) ([]string, error) {
	tokens := tokenize(text)
	tokens = toLowercase(tokens)
	tokens = removeCommonWords(tokens)
	tokens, err := stem(tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func toLowercase(tokens []string) []string {
	lowercaseTokens := []string{}
	for _, token := range tokens {
		lowercaseTokens = append(lowercaseTokens, strings.ToLower(token))
	}
	return lowercaseTokens
}

func removeCommonWords(tokens []string) []string {
	// Top 15 words (OEC rank)
	commonWords := map[string]bool{
		"the":  true,
		"be":   true,
		"to":   true,
		"of":   true,
		"and":  true,
		"a":    true,
		"in":   true,
		"that": true,
		"have": true,
		"I":    true,
		"it":   true,
		"for":  true,
		"not":  true,
		"on":   true,
		"with": true,
	}
	tokensWithoutCommonWords := []string{}
	for _, token := range tokens {
		if _, ok := commonWords[token]; !ok {
			tokensWithoutCommonWords = append(tokensWithoutCommonWords, token)
		}
	}
	return tokensWithoutCommonWords
}

func stem(tokens []string) ([]string, error) {
	stemmedTokens := []string{}
	for _, token := range tokens {
		stemmedToken, err := snowball.Stem(token, "english", true)
		if err != nil {
			return nil, fmt.Errorf("unable to stem token: %s", err.Error())
		}
		stemmedTokens = append(stemmedTokens, stemmedToken)
	}
	return stemmedTokens, nil
}
