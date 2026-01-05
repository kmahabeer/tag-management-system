package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/api"
)

func ListContexts(w http.ResponseWriter, r *http.Request) {
	response := api.ContextsGet200Response{
		Contexts: []api.Context{},
	}

	WriteJSON(w, http.StatusOK, response)
}

func CreateContext(w http.ResponseWriter, r *http.Request) {
	var input api.ContextInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	context := api.Context{
		ID:                 uuid.New(),
		Name:               input.Name,
		ClassificationType: input.ClassificationType,
		Description:        input.Description,
		IsActive:           input.IsActive,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	WriteJSON(w, http.StatusCreated, context)
}

func GetContext(w http.ResponseWriter, r *http.Request) {
	context := api.Context{
		ID:                 uuid.New(),
		Name:               "example",
		ClassificationType: "objective",
		Description:        nil,
		IsActive:           true,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	WriteJSON(w, http.StatusOK, context)
}

func UpdateContext(w http.ResponseWriter, r *http.Request) {
	var input api.ContextInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	context := api.Context{
		ID:                 uuid.New(),
		Name:               input.Name,
		ClassificationType: input.ClassificationType,
		Description:        input.Description,
		IsActive:           input.IsActive,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	WriteJSON(w, http.StatusOK, context)
}

func DeleteContext(w http.ResponseWriter, r *http.Request) {
	response := api.TagsIdDelete200Response{
		Status: "deleted",
	}

	WriteJSON(w, http.StatusOK, response)
}

func ListPartsOfSpeech(w http.ResponseWriter, r *http.Request) {
	response := api.PartsOfSpeechGet200Response{
		PartsOfSpeech: []api.PartOfSpeech{},
	}

	WriteJSON(w, http.StatusOK, response)
}

func CreatePartOfSpeech(w http.ResponseWriter, r *http.Request) {
	var input api.PartOfSpeechInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	pos := api.PartOfSpeech{
		ID:          uuid.New(),
		Name:        input.Name,
		Description: input.Description,
		IsActive:    input.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	WriteJSON(w, http.StatusCreated, pos)
}

func GetPartOfSpeech(w http.ResponseWriter, r *http.Request) {
	pos := api.PartOfSpeech{
		ID:          uuid.New(),
		Name:        "noun",
		Description: nil,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	WriteJSON(w, http.StatusOK, pos)
}

func UpdatePartOfSpeech(w http.ResponseWriter, r *http.Request) {
	var input api.PartOfSpeechInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	pos := api.PartOfSpeech{
		ID:          uuid.New(),
		Name:        input.Name,
		Description: input.Description,
		IsActive:    input.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	WriteJSON(w, http.StatusOK, pos)
}

func DeletePartOfSpeech(w http.ResponseWriter, r *http.Request) {
	response := api.TagsIdDelete200Response{
		Status: "deleted",
	}

	WriteJSON(w, http.StatusOK, response)
}

func ListRatings(w http.ResponseWriter, r *http.Request) {
	response := api.RatingsGet200Response{
		Ratings: []api.Rating{},
	}

	WriteJSON(w, http.StatusOK, response)
}

func CreateRating(w http.ResponseWriter, r *http.Request) {
	var input api.RatingInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	rating := api.Rating{
		ID:           uuid.New(),
		Name:         input.Name,
		Score:        input.Score,
		Description:  input.Description,
		RatingTypeID: input.RatingTypeID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	WriteJSON(w, http.StatusCreated, rating)
}

func GetRating(w http.ResponseWriter, r *http.Request) {
	rating := api.Rating{
		ID:           uuid.New(),
		Name:         "excellent",
		Score:        10,
		Description:  nil,
		RatingTypeID: uuid.New(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	WriteJSON(w, http.StatusOK, rating)
}

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	var input api.RatingInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	rating := api.Rating{
		ID:           uuid.New(),
		Name:         input.Name,
		Score:        input.Score,
		Description:  input.Description,
		RatingTypeID: input.RatingTypeID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	WriteJSON(w, http.StatusOK, rating)
}

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	response := api.TagsIdDelete200Response{
		Status: "deleted",
	}

	WriteJSON(w, http.StatusOK, response)
}

func ListRatingTypes(w http.ResponseWriter, r *http.Request) {
	response := api.RatingTypesGet200Response{
		RatingTypes: []api.RatingType{},
	}

	WriteJSON(w, http.StatusOK, response)
}

func CreateRatingType(w http.ResponseWriter, r *http.Request) {
	var input api.RatingTypeInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	rt := api.RatingType{
		ID:           uuid.New(),
		Name:         input.Name,
		IsNormalized: input.IsNormalized,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	WriteJSON(w, http.StatusCreated, rt)
}

func GetRatingType(w http.ResponseWriter, r *http.Request) {
	rt := api.RatingType{
		ID:           uuid.New(),
		Name:         "likeness",
		IsNormalized: true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	WriteJSON(w, http.StatusOK, rt)
}

func UpdateRatingType(w http.ResponseWriter, r *http.Request) {
	var input api.RatingTypeInput
	if err := DecodeAndValidate(r, &input); err != nil {
		return
	}

	rt := api.RatingType{
		ID:           uuid.New(),
		Name:         input.Name,
		IsNormalized: input.IsNormalized,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	WriteJSON(w, http.StatusOK, rt)
}

func DeleteRatingType(w http.ResponseWriter, r *http.Request) {
	response := api.TagsIdDelete200Response{
		Status: "deleted",
	}

	WriteJSON(w, http.StatusOK, response)
}
