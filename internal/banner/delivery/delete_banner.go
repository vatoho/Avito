package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
)

func (d *BannerDelivery) DeleteBanner(w http.ResponseWriter, r *http.Request) {
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
	bannerID, ok := vars["id"]
	if !ok {
		d.logger.Errorf("banner id was not passed")
		response.WriteResponse(w, response.Error{Err: "banner id was not passed"},
			http.StatusBadRequest, d.logger)
		return
	}

	bannerIDInt, err := strconv.ParseUint(bannerID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in banner ID conversion: %s", err)
		response.WriteResponse(w, response.Error{Err: "banner ID must be positive integer"},
			http.StatusBadRequest, d.logger)
		return
	}

	err = d.service.DeleteBanner(r.Context(), bannerIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrBannerNotFound) {
			d.logger.Errorf("no banners with id %d", bannerIDInt)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		d.logger.Errorf("internal server error in deleting banner: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
