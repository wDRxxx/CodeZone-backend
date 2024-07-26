package api

import (
	"codeZone/internal/models"
	"codeZone/internal/repository/docker"
	"codeZone/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"net/http"
)

var errInternal = errors.New("Internal error, try later")

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	resp, err := s.apiService.Home()
	if err != nil {
		utils.PrintErrWithStack(err)
		utils.WriteJSONError(errInternal, w, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(resp, w, http.StatusOK)
	if err != nil {
		utils.PrintErrWithStack(err)
		utils.WriteJSONError(errInternal, w, http.StatusInternalServerError)
		return
	}
}

func (s *server) run(w http.ResponseWriter, r *http.Request) {
	var req models.RunV1Request
	err := utils.ReadJSON(w, r, &req)
	if err != nil {
		utils.PrintErrWithStack(err)
		utils.WriteJSONError(errInternal, w, http.StatusInternalServerError)
		return
	}

	resp, err := s.apiService.Run(r.Context(), &req)
	if err != nil {
		utils.PrintErrWithStack(err)
		if errors.Is(err, docker.ErrUnsupported) {
			utils.WriteJSONError(err, w, http.StatusNotFound)
			return
		}

		utils.WriteJSONError(errInternal, w, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(resp, w, http.StatusAccepted)
	if err != nil {
		utils.PrintErrWithStack(err)
		utils.WriteJSONError(errInternal, w, http.StatusInternalServerError)
		return
	}
}

func (s *server) check(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := s.apiService.Check(r.Context(), id)
	if err != nil {
		utils.PrintErrWithStack(err)
		utils.WriteJSONError(errInternal, w, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(resp, w, http.StatusOK)
	if err != nil {
		utils.PrintErrWithStack(err)
		utils.WriteJSONError(errInternal, w, http.StatusInternalServerError)
		return
	}
}
