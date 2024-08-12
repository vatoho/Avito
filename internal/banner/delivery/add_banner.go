package delivery

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ilyushkaaa/banner-service/internal/banner/delivery/dto"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
)

const adminRole = "admin"

func (d *BannerDelivery) AddBanner(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r.Context())
	if err != nil {
		d.logger.Errorf("error in getting user from context: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}
	if user.Role != adminRole {
		d.logger.Errorf("user %d has got no access dor adding banner", user.TagID)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		d.logger.Errorf("error in reading request body: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	defer func() {
		err = r.Body.Close()
		if err != nil {
			d.logger.Errorf("error in closing request body")
		}
	}()

	var bannerDTO dto.BannerToAdd
	err = json.Unmarshal(rBody, &bannerDTO)
	if err != nil {
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			d.logger.Errorf("invalid json: %s", string(rBody))
			response.WriteResponse(w, response.Error{Err: response.ErrInvalidJSON.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("error in response body unmarshalling: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	err = bannerDTO.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in adding banner: %v", err)
		response.WriteResponse(w, response.Error{Err: err.Error()}, http.StatusBadRequest, d.logger)
		return
	}

	bannerToAdd := dto.ConvertToBanner(bannerDTO)
	addedBanner, err := d.service.AddBanner(r.Context(), bannerToAdd)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicateFeatureTag) {
			d.logger.Errorf("banner with one of combibnations of featiure + tag already exists: %d, %v", bannerToAdd.FeatureID, bannerToAdd.TagIDs)
			response.WriteResponse(w, response.Error{Err: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("internal server error in adding banner: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.WriteResponse(w, dto.BannerResponse{ID: addedBanner.ID}, http.StatusCreated, d.logger)
}
