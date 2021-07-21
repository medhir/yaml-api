package storage

import (
	"errors"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
)

// Storage is the controller used to store, index, and lookup YAML documents
type Storage struct {
	index index
}

// Metadata describes all the properties of the YAML metadata stored & indexed by the API
type Metadata struct {
	Title       string       `yaml:"title" json:"title"`
	Version     string       `yaml:"version" json:"version"`
	Maintainers []Maintainer `yaml:"maintainers" json:"maintainers"`
	Company     string       `yaml:"company" json:"company"`
	Website     string       `yaml:"website" json:"website"`
	Source      string       `yaml:"source" json:"source"`
	License     string       `yaml:"license" json:"license"`
	Description string       `yaml:"description" json:"description"`
}

// Maintainer describes the properties of an application maintainer
type Maintainer struct {
	Name  string `yaml:"name" json:"name"`
	Email string `yaml:"email" json:"email"`
}

// NewStorage initializes a new metadata store
func NewStorage() *Storage {
	return &Storage{
		index: index{
			title:           map[string][]*Metadata{},
			version:         map[string][]*Metadata{},
			maintainerName:  map[string][]*Metadata{},
			maintainerEmail: map[string][]*Metadata{},
			company:         map[string][]*Metadata{},
			website:         map[string][]*Metadata{},
			source:          map[string][]*Metadata{},
			license:         map[string][]*Metadata{},
			description:     map[string][]*Metadata{},
		},
	}
}

// AddMetadata indexes references a metadata object to the in-memory store by the values of every attribute.
func (s *Storage) AddMetadata(metadata *Metadata) error {
	err := indexField(metadata.Title, s.index.title, metadata, true)
	if err != nil {
		return err
	}
	err = indexField(metadata.Version, s.index.version, metadata, false)
	if err != nil {
		return err
	}
	for _, maintainer := range metadata.Maintainers {
		err = indexField(maintainer.Name, s.index.maintainerName, metadata, true)
		if err != nil {
			return err
		}
		err = indexField(maintainer.Email, s.index.maintainerEmail, metadata, true)
		if err != nil {
			return err
		}
	}
	err = indexField(metadata.Company, s.index.company, metadata, true)
	if err != nil {
		return err
	}
	err = indexField(metadata.Website, s.index.website, metadata, true)
	if err != nil {
		return err
	}
	err = indexField(metadata.Source, s.index.source, metadata, true)
	if err != nil {
		return err
	}
	err = indexField(metadata.License, s.index.license, metadata, true)
	if err != nil {
		return err
	}
	err = indexField(metadata.Description, s.index.description, metadata, true)
	if err != nil {
		return err
	}
	return nil
}

// retrieveDocuments returns all metadata that matches a search phrase in a specific attribute
// (such as description, title, etc)
func (s *Storage) retrieveDocuments(attr string, searchInput string) ([]*Metadata, error) {
	resultSet := map[string][]*Metadata{}
	switch attribute(attr) {
	case Title:
		err := getResultsByToken(searchInput, resultSet, s.index.title)
		if err != nil {
			return nil, err
		}
	case Version:
		// do not tokenize version numbers (should include periods)
		resultSet[searchInput] = append(resultSet[searchInput], s.index.version[searchInput]...)
	case MaintainerName:
		err := getResultsByToken(searchInput, resultSet, s.index.maintainerName)
		if err != nil {
			return nil, err
		}
	case MaintainerEmail:
		err := getResultsByToken(searchInput, resultSet, s.index.maintainerEmail)
		if err != nil {
			return nil, err
		}
	case Company:
		err := getResultsByToken(searchInput, resultSet, s.index.company)
		if err != nil {
			return nil, err
		}
	case Website:
		err := getResultsByToken(searchInput, resultSet, s.index.website)
		if err != nil {
			return nil, err
		}
	case Source:
		err := getResultsByToken(searchInput, resultSet, s.index.source)
		if err != nil {
			return nil, err
		}
	case License:
		err := getResultsByToken(searchInput, resultSet, s.index.license)
		if err != nil {
			return nil, err
		}
	case Description:
		err := getResultsByToken(searchInput, resultSet, s.index.description)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("cannot retrieve documents by unknown attribute type")
	}
	// reduce the result set to only those with all the search terms provided
	// this is found by taking the intersection of metadata results for each term
	resultsMatchingAllTerms := matchAllTerms(resultSet)
	return resultsMatchingAllTerms, nil
}

func getResultsByToken(searchInput string, resultSet, field map[string][]*Metadata) error {
	tokens, err := processText(searchInput)
	if err != nil {
		return err
	}
	for _, token := range tokens {
		resultSet[token] = []*Metadata{}
		resultSet[token] = append(resultSet[token], field[token]...)
	}
	return nil
}

// LookupMetadata performs a search of all metadata by the desired attribute(s)
func (s *Storage) LookupMetadata(attrsAndValues map[string]string) ([]*Metadata, error) {
	resultSet := map[string][]*Metadata{}
	for k, v := range attrsAndValues {
		result, err := s.retrieveDocuments(k, v)
		if err != nil {
			return nil, err
		}
		resultSet[k] = result
	}
	return matchAllTerms(resultSet), nil
}

// ValidateMetadata ensures that all metadata fields are formatted properly.
// Assumptions:
// Title, Company, and License can be any string.
// Version is a properly formatted semantic version.
// Maintainers must be a slice of Maintainer structs, each of which as a Name and Email field. Email must be a valid email address.
// Website and Source must use a properly formatted URL.
// Description must be formatted using Markdown.
func (s *Storage) ValidateMetadata(metadata *Metadata) error {
	if metadata.Title == "" {
		return errors.New("metadata must have a title")
	}
	if metadata.Version == "" {
		return errors.New("metadata must have a version")
	}
	_, err := semver.NewVersion(metadata.Version)
	if err != nil {
		return errors.New("version must follow the semantic versioning scheme: https://semver.org")
	}
	if len(metadata.Maintainers) == 0 {
		return errors.New("metadata must have a list of maintainers, each of which has a name and email attribute")
	}
	for _, maintainer := range metadata.Maintainers {
		if maintainer.Name == "" {
			return errors.New("maintainer must have a name")
		}
		if maintainer.Email == "" {
			return errors.New("maintainer must have an email")
		}
		err = validateEmail(maintainer.Email)
		if err != nil {
			return err
		}
	}
	if metadata.Company == "" {
		return errors.New("metadata must have a company")
	}
	if metadata.Website == "" {
		return errors.New("metadata must have a website")
	}
	_, err = url.ParseRequestURI(metadata.Website)
	if err != nil {
		return errors.New("metadata must have a website with a valid URL")
	}
	if metadata.Source == "" {
		return errors.New("metadata must have a source")
	}
	_, err = url.ParseRequestURI(metadata.Source)
	if err != nil {
		return errors.New("metadata must have a source with a valid URL")
	}
	if metadata.License == "" {
		return errors.New("metadata must have a license")
	}
	if metadata.Description == "" {
		return errors.New("metadata must have a description")
	}
	// err = goldmark.Convert([]byte(metadata.Description), &bytes.Buffer{})
	// if err != nil {
	// 	return errors.New("metadata must have a valid description with markdown formatting")
	// }
	return nil
}

// from https://golangcode.com/validate-an-email-address/
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func validateEmail(e string) error {
	if (len(e) < 3 && len(e) > 254) || !emailRegex.MatchString(e) {
		return errors.New("email must be a properly formatted email address")
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return errors.New("email must be a valid email address")
	}
	return nil
}
