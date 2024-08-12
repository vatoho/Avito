package delivery

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilyushkaaa/banner-service/internal/banner/delivery/dto"
	"github.com/ilyushkaaa/banner-service/internal/banner/service"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
)

func (d *BannerDelivery) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r.Context())
	if err != nil {
		d.logger.Errorf("error in getting user from context: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}
	if user.Role != adminRole {
		d.logger.Errorf("user %d has got no access dor updating banner", user.TagID)
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

	bannerToUpdate, err := getBannerToUpdate(rBody, bannerIDInt)
	if err != nil {
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			d.logger.Errorf("invalid json: %s", string(rBody))
			response.WriteResponse(w, response.Error{Err: response.ErrInvalidJSON.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		if errors.Is(err, ErrNoFieldsToUpdate) {
			d.logger.Error(err)
			response.WriteResponse(w, response.Error{Err: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("error in response body unmarshalling: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	err = d.service.UpdateBanner(r.Context(), *bannerToUpdate)
	if err != nil {
		if errors.Is(err, storage.ErrBannerNotFound) {
			d.logger.Errorf("no banners with id %d", bannerIDInt)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if errors.Is(err, storage.ErrDuplicateFeatureTag) {
			d.logger.Error("banner with one of combinations of feature + tag already exists")
			response.WriteResponse(w, response.Error{Err: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("internal server error in updating banner: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getBannerToUpdate(rBody []byte, id uint64) (*service.BannerToUpdate, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(rBody, &raw)
	if err != nil {
		return nil, err
	}

	bannerUpdate := dto.BannerUpdate{}
	bannerToUpdate := &service.BannerToUpdate{}

	hasDiffs := false
	if val, ok := raw["tag_ids"]; ok {
		hasDiffs = true
		err = json.Unmarshal(val, &bannerUpdate.TagIDs)
		if err != nil {
			return nil, err
		}
		bannerToUpdate.TagIDs = &bannerUpdate.TagIDs
	}
	if val, ok := raw["feature_id"]; ok {
		hasDiffs = true
		err = json.Unmarshal(val, &bannerUpdate.FeatureID)
		if err != nil {
			return nil, err
		}
		bannerToUpdate.FeatureID = &bannerUpdate.FeatureID
	}
	if val, ok := raw["content"]; ok {
		hasDiffs = true
		err = json.Unmarshal(val, &bannerUpdate.Content)
		if err != nil {
			return nil, err
		}
		bannerToUpdate.Content = &bannerUpdate.Content
	}
	if val, ok := raw["is_active"]; ok {
		hasDiffs = true
		err = json.Unmarshal(val, &bannerUpdate.IsActive)
		if err != nil {
			return nil, err
		}
		bannerToUpdate.IsActive = &bannerUpdate.IsActive
	}
	if !hasDiffs {
		return nil, ErrNoFieldsToUpdate
	}
	bannerToUpdate.ID = id
	return bannerToUpdate, nil

}
