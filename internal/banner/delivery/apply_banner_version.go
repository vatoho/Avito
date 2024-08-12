package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
)

func (d *BannerDelivery) ApplyBannerVersion(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r.Context())
	if err != nil {
		d.logger.Errorf("error in getting user from context: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}
	if user.Role != adminRole {
		d.logger.Errorf("user %d has got no access for deleting banner", user.TagID)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	versionID, ok := vars["id"]
	if !ok {
		d.logger.Errorf("banner id was not passed")
		response.WriteResponse(w, response.Error{Err: "banner id was not passed"},
			http.StatusBadRequest, d.logger)
		return
	}

	versionIDInt, err := strconv.ParseUint(versionID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in banner ID conversion: %s", err)
		response.WriteResponse(w, response.Error{Err: "banner ID must be positive integer"},
			http.StatusBadRequest, d.logger)
		return
	}

	err = d.service.ApplyBannerVersion(r.Context(), versionIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrNoBannerVersions) {
			d.logger.Errorf("no banner versions with id %d", versionIDInt)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		d.logger.Errorf("internal server error in applying banner version: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	w.WriteHeader(http.StatusOK)
}
