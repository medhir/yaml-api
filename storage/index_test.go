package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processText(t *testing.T) {
	text := "The quick brown fox jumped on things, and rocks, and emailed a letter at Mail@Gmail.com"
	tokens := tokenize(text)
	t.Run("Splits on any character that is not a letter or number", func(t *testing.T) {
		assert.Equal(t, tokens, []string{"The", "quick", "brown", "fox", "jumped", "on", "things", "and", "rocks", "and", "emailed", "a", "letter", "at", "Mail", "Gmail", "com"})
	})

	tokens = toLowercase(tokens)
	t.Run("Removes capitalization from tokens", func(t *testing.T) {
		assert.Equal(t, tokens, []string{"the", "quick", "brown", "fox", "jumped", "on", "things", "and", "rocks", "and", "emailed", "a", "letter", "at", "mail", "gmail", "com"})
	})

	tokens = removeCommonWords(tokens)
	t.Run("Removes tokens representing common words", func(t *testing.T) {
		assert.Equal(t, tokens, []string{"quick", "brown", "fox", "jumped", "things", "rocks", "emailed", "letter", "at", "mail", "gmail", "com"})
	})

	tokens, err := stem(tokens)
	t.Run("stems tokens to transform words to their base form", func(t *testing.T) {
		assert.Equal(t, tokens, []string{"quick", "brown", "fox", "jump", "thing", "rock", "email", "letter", "at", "mail", "gmail", "com"})
		assert.NoError(t, err)
	})
}

func Test_indexField(t *testing.T) {
	descriptionIndex := map[string][]*Metadata{}
	versionIndex := map[string][]*Metadata{}
	md := &Metadata{
		Title:   "App title 1",
		Version: "1.0.0",
		Maintainers: []Maintainer{
			{
				Name:  "Bill Bob",
				Email: "bill@gmail.com",
			},
			{
				Name:  "Medhir Bhargava",
				Email: "email@gmail.com",
			},
		},
		Company:     "BigCorp",
		Website:     "https://www.wikipedia.com",
		Source:      "https://github.com",
		License:     "All rights reserved",
		Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
	}

	t.Run("indexes using tokens", func(t *testing.T) {
		err := indexField(md.Description, descriptionIndex, md, true)
		assert.NoError(t, err)
		tokens, err := processText(md.Description)
		assert.NoError(t, err)
		for _, token := range tokens {
			assert.Equal(t, 1, len(descriptionIndex[token]))
			assert.Equal(t, md, descriptionIndex[token][0])
		}
	})

	t.Run("indexes without using tokens", func(t *testing.T) {
		err := indexField(md.Version, versionIndex, md, false)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(versionIndex[md.Version]))
		assert.Equal(t, md, versionIndex[md.Version][0])
	})
}
