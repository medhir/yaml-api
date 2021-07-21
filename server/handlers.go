package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/medhir/yaml-api/storage"
	"gopkg.in/yaml.v2"
)

func (s *Server) handleGetMetadata() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		searchTerms := map[string]string{}
		for k, v := range values {
			searchTerms[k] = v[0]
		}
		results, err := s.storage.LookupMetadata(searchTerms)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not retreive metadata by provided search terms:\n%v", err), http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(results)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not encode JSON:\n%v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not write JSON to response:\n%v", err), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handlePostMetadata() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		yamlBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not read request body:\n%v", err), http.StatusBadRequest)
			return
		}
		metadata := &storage.Metadata{}
		err = yaml.Unmarshal(yamlBytes, metadata)
		if err != nil {
			http.Error(w, fmt.Sprintf("request does not contain valid YAML:\n%v", err), http.StatusBadRequest)
			return
		}
		err = s.storage.ValidateMetadata(metadata)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.storage.AddMetadata(metadata)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleMetadata() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.handleGetMetadata()(w, r)
		case http.MethodPost:
			s.handlePostMetadata()(w, r)
		default:
			http.Error(w, fmt.Sprintf("unimplemented http handler for method %s", r.Method), http.StatusMethodNotAllowed)
		}
	}
}
