package handler

import (
	"encoding/json"
	"net/http"
)

type ResultResponse struct {
	StarCount int      `json:"star_count"`
	Followers []string `json:"followers"`
}

func dataHandler(user_name string, repo_name string) (*ResultResponse, error) {
	var followers []string

	organizations, err := fetchFollowersList(user_name)
	if err != nil {
		return nil, err
	}
	for _, organization := range organizations {
		followers = append(followers, organization.GetLogin())
	}

	repositry, err := fetchStarCountOfRepository(user_name, repo_name)
	if err != nil {
		return nil, err
	}

	resultResponse := ResultResponse{repositry.GetStargazersCount(), followers}

	return &resultResponse, nil
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	user_name := r.URL.Query().Get("user")
	repo_name := r.URL.Query().Get("repo")

	result, err := dataHandler(user_name, repo_name)

	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	} else {
		respondJSON(w, http.StatusOK, result)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, http.StatusNotFound, map[string]string{"error": message})
}
