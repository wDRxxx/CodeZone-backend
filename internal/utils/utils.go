package utils

import (
	"codeZone/internal/models"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"runtime/debug"
)

// WriteJSON writes json by given data
func WriteJSON(data any, w http.ResponseWriter, status int) error {
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		return err
	}

	return nil
}

// WriteJSONError writes json error
func WriteJSONError(err error, w http.ResponseWriter, status int) error {
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(&models.JSONResponse{
		Error:   true,
		Message: err.Error(),
	})

	if err != nil {
		return err
	}

	w.WriteHeader(status)

	_, err = w.Write(res)
	if err != nil {
		return err
	}

	return nil
}

// ReadJSON reads json to given data
func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1024 * 1024 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

// PrintErrWithStack prints error with its stack
func PrintErrWithStack(err error) {
	log.Println(err, string(debug.Stack()))
}
