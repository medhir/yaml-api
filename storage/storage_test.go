package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateMetadata(t *testing.T) {
	tests := []struct {
		name       string
		metadata   Metadata
		err        bool
		errMessage string
	}{
		{
			name: "Validation fails when metadata has no title",
			metadata: Metadata{
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
			},
			err:        true,
			errMessage: "metadata must have a title",
		},
		{
			name: "Validation fails when metadata has no version",
			metadata: Metadata{
				Title: "App title 1",
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
			},
			err:        true,
			errMessage: "metadata must have a version",
		},
		{
			name: "Validation fails when metadata has an improperly formatted version",
			metadata: Metadata{
				Title:   "App title 1",
				Version: "version 2.2",
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
			},
			err:        true,
			errMessage: "version must follow the semantic versioning scheme: https://semver.org",
		},
		{
			name: "Validation fails when metadata has no maintainer list",
			metadata: Metadata{
				Title:       "App title 1",
				Version:     "1.0.0",
				Company:     "BigCorp",
				Website:     "https://www.wikipedia.com",
				Source:      "https://github.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "metadata must have a list of maintainers, each of which has a name and email attribute",
		},
		{
			name: "Validation fails when maintainer is missing a name",
			metadata: Metadata{
				Title:   "App title 1",
				Version: "1.0.0",
				Maintainers: []Maintainer{
					{
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
			},
			err:        true,
			errMessage: "maintainer must have a name",
		},
		{
			name: "Validation fails when maintainer is missing an email",
			metadata: Metadata{
				Title:   "App title 1",
				Version: "1.0.0",
				Maintainers: []Maintainer{
					{
						Name:  "Bill Bob",
						Email: "bill@gmail.com",
					},
					{
						Name: "Medhir Bhargava",
					},
				},
				Company:     "BigCorp",
				Website:     "https://www.wikipedia.com",
				Source:      "https://github.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "maintainer must have an email",
		},
		{
			name: "Validation fails when maintainer has an improperly formatted email",
			metadata: Metadata{
				Title:   "App title 1",
				Version: "1.0.0",
				Maintainers: []Maintainer{
					{
						Name:  "Bill Bob",
						Email: "bill@gmail.com",
					},
					{
						Name:  "Medhir Bhargava",
						Email: "mailmedhir.com",
					},
				},
				Company:     "BigCorp",
				Website:     "https://www.wikipedia.com",
				Source:      "https://github.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "email must be a properly formatted email address",
		},
		{
			name: "Validation fails when maintainer has an email that does not exist",
			metadata: Metadata{
				Title:   "App title 1",
				Version: "1.0.0",
				Maintainers: []Maintainer{
					{
						Name:  "Bill Bob",
						Email: "bill@gmail.com",
					},
					{
						Name:  "Medhir Bhargava",
						Email: "snailmail@flipflapjack.com",
					},
				},
				Company:     "BigCorp",
				Website:     "https://www.wikipedia.com",
				Source:      "https://github.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "email must be a valid email address",
		},
		{
			name: "Validation fails when metadata has no company",
			metadata: Metadata{
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
				Website:     "https://www.wikipedia.com",
				Source:      "https://github.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "metadata must have a company",
		},
		{
			name: "Validation fails when metadata has no website",
			metadata: Metadata{
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
				Source:      "https://github.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "metadata must have a website",
		},
		{
			name: "Validation fails when website is an invalid URL",
			metadata: Metadata{
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
				Website:     "https//wwwwikipedia.com",
				Source:      "https://github.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "metadata must have a website with a valid URL",
		},
		{
			name: "Validation fails when metadata has no source",
			metadata: Metadata{
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
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "metadata must have a source",
		},
		{
			name: "Validation fails when source has an invalid URL",
			metadata: Metadata{
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
				Source:      "httpsgithub.com",
				License:     "All rights reserved",
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "metadata must have a source with a valid URL",
		},
		{
			name: "Validation fails when metadata has no license",
			metadata: Metadata{
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
				Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)",
			},
			err:        true,
			errMessage: "metadata must have a license",
		},
		{
			name: "Validation fails when metadata has no description",
			metadata: Metadata{
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
				Company: "BigCorp",
				Website: "https://www.wikipedia.com",
				Source:  "https://github.com",
				License: "All rights reserved",
			},
			err:        true,
			errMessage: "metadata must have a description",
		},
		{
			name: "Validation succeeds when all metadata fields are present and properly formatted",
			metadata: Metadata{
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
			},
			err: false,
		},
	}

	s := &Storage{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.ValidateMetadata(&tt.metadata)
			if tt.err {
				assert.Error(t, err)
				if err != nil {
					assert.Equal(t, tt.errMessage, err.Error())
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_AddDocument(t *testing.T) {
	s := NewStorage()
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

	err := s.AddMetadata(md)
	assert.NoError(t, err)
	// check title
	checkIndexForTokens(t, md.Title, md, s.index.title)
	// check version
	assert.NotEmpty(t, s.index.version[md.Version])
	assert.Equal(t, md, s.index.version[md.Version][0])
	// check maintainers
	for _, maintainer := range md.Maintainers {
		checkIndexForTokens(t, maintainer.Name, md, s.index.maintainerName)
		checkIndexForTokens(t, maintainer.Email, md, s.index.maintainerEmail)
	}
	// check company
	checkIndexForTokens(t, md.Company, md, s.index.company)
	// check website
	checkIndexForTokens(t, md.Website, md, s.index.website)
	// check source
	checkIndexForTokens(t, md.Source, md, s.index.source)
	// check license
	checkIndexForTokens(t, md.License, md, s.index.license)
	// check description
	checkIndexForTokens(t, md.Description, md, s.index.description)
}

func checkIndexForTokens(t *testing.T, text string, md *Metadata, field map[string][]*Metadata) {
	tokens, err := processText(text)
	assert.NoError(t, err)
	for _, token := range tokens {
		assert.NotEmpty(t, field[token])
		assert.Equal(t, md, field[token][0])
	}
}

func Test_retrieveDocuments(t *testing.T) {
	s := NewStorage()
	documents := []*Metadata{
		{
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
			Source:      "https://github.com/app1",
			License:     "MIT",
			Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)\nThe quick blue fox jumped on the hen.",
		},
		{
			Title:   "App title 2",
			Version: "2.0.1",
			Maintainers: []Maintainer{
				{
					Name:  "Billy Bob",
					Email: "billy@gmail.com",
				},
				{
					Name:  "Medhir Bhargava",
					Email: "email@gmail.com",
				},
			},
			Company:     "SmallCorp",
			Website:     "https://www.smallcorp.com",
			Source:      "https://github.com/app2",
			License:     "Apache-2.0",
			Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)\nThe quick brown fox jumped on the hen.",
		},
	}
	for _, document := range documents {
		err := s.AddMetadata(document)
		assert.NoError(t, err)
	}
	// exercise search against titles
	results, err := s.retrieveDocuments("title", "app title 1")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[0], results[0])
	results, err = s.retrieveDocuments("title", "app title 2")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.retrieveDocuments("title", "app title")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, documents[0])
	assert.Contains(t, results, documents[1])
	results, err = s.retrieveDocuments("title", "application name")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// exercise search against versions
	results, err = s.retrieveDocuments("version", "1.0.0")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[0], results[0])
	results, err = s.retrieveDocuments("version", "1.0.2")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// exercise search against maintainer name
	results, err = s.retrieveDocuments("maintainer_name", "Billy")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.retrieveDocuments("maintainer_name", "Joanne")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// exercise search against maintainer email
	results, err = s.retrieveDocuments("maintainer_email", "email@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, documents[0])
	assert.Contains(t, results, documents[1])
	results, err = s.retrieveDocuments("maintainer_email", "snail@gmail.com")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// exercise search against company
	results, err = s.retrieveDocuments("company", "SmallCorp")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.retrieveDocuments("company", "MediumCorp")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// exercise search against website
	results, err = s.retrieveDocuments("website", "smallcorp.com")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.retrieveDocuments("website", "medhir.com")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// search against source
	results, err = s.retrieveDocuments("source", "github.com")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, documents[0])
	assert.Contains(t, results, documents[1])
	results, err = s.retrieveDocuments("source", "gitlab.com")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// exercise search against license
	results, err = s.retrieveDocuments("license", "MIT")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[0], results[0])
	results, err = s.retrieveDocuments("license", "apache")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.retrieveDocuments("license", "GPL")
	assert.NoError(t, err)
	assert.Empty(t, results)
	// exercise search against description
	results, err = s.retrieveDocuments("description", "quick blue fox")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[0], results[0])
	results, err = s.retrieveDocuments("description", "quick brown fox")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.retrieveDocuments("description", "quick fox")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, documents[0])
	assert.Contains(t, results, documents[1])
	results, err = s.retrieveDocuments("description", "dog jumped over the moon")
	assert.NoError(t, err)
	assert.Empty(t, results)
}

func Test_LookupMetadata(t *testing.T) {
	s := NewStorage()
	documents := []*Metadata{
		{
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
			Source:      "https://github.com/app1",
			License:     "MIT",
			Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)\nThe quick blue fox jumped on the hen.",
		},
		{
			Title:   "App title 2",
			Version: "2.0.1",
			Maintainers: []Maintainer{
				{
					Name:  "Billy Bob",
					Email: "billy@gmail.com",
				},
				{
					Name:  "Medhir Bhargava",
					Email: "email@gmail.com",
				},
			},
			Company:     "SmallCorp",
			Website:     "https://www.smallcorp.com",
			Source:      "https://github.com/app2",
			License:     "Apache-2.0",
			Description: "# A main heading\n## A secondary heading\nA paragraph\n![an image](https://image.com/png)\nThe quick brown fox jumped on the hen.",
		},
	}
	for _, document := range documents {
		err := s.AddMetadata(document)
		assert.NoError(t, err)
	}
	results, err := s.LookupMetadata(map[string]string{
		"title":   "app title",
		"company": "smallcorp",
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.LookupMetadata(map[string]string{
		"license": "apache",
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, documents[1], results[0])
	results, err = s.LookupMetadata(map[string]string{
		"title":       "app title",
		"license":     "apache",
		"description": "no match",
	})
	assert.NoError(t, err)
	assert.Empty(t, results)
}
