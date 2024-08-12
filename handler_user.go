package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shetty456/rssgg/internal/auth"
	"github.com/shetty456/rssgg/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramertes struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramertes{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json:", err))
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
