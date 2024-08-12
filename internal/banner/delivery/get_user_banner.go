package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ilyushkaaa/banner-service/internal/banner/delivery/dto"
	"github.com/ilyushkaaa/banner-service/internal/banner/service"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
)

func (d *BannerDelivery) GetUserBanner(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	tagID := params.Get("tag_id")
	if tagID == "" {
		d.logger.Error("tag id was not passed")
		response.WriteResponse(w, response.Error{Err: "tag id is required"},
			http.StatusBadRequest, d.logger)
		return
	}
	tagIDInt, err := d.parseParam("tag id", tagID, w)
	if err != nil {
		return
	}
	user, err := getUserFromContext(r.Context())
	if err != nil {
		d.logger.Errorf("error in getting user from context: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}
	if user.TagID != tagIDInt && user.Role != "admin" {
		d.logger.Errorf("user %d has got no access dor getting banner with tag %d", user.TagID, tagIDInt)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	featureID := params.Get("feature_id")
	if featureID == "" {
		d.logger.Error("feature id was not passed")
		response.WriteResponse(w, response.Error{Err: "feature id is required"},
			http.StatusBadRequest, d.logger)
		return
	}
	featureIDInt, err := d.parseParam("feature_id", featureID, w)
	if err != nil {
		return
	}
	useLastVersion := params.Get("use_last_revision")
	useLastVersionBool := false
	if useLastVersion != "" {
		useLastVersionBool, err = strconv.ParseBool(useLastVersion)
		if err != nil {
			d.logger.Errorf("error in use last version conversion: %s", err)
			response.WriteResponse(w, response.Error{Err: "use_last_version must be false or true"},
				http.StatusBadRequest, d.logger)
			return
		}
	}
	bannerContent, err := d.service.GetUserBanner(r.Context(), featureIDInt, tagIDInt, useLastVersionBool)
	if err != nil {
		if errors.Is(err, storage.ErrBannerNotFound) {
			d.logger.Errorf("no banners with feature id = %d and tag id = %d", featureIDInt, tagIDInt)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrBannerIsInactive) {
			d.logger.Errorf("banner with feature %d and tag %d is inactive", featureIDInt, tagIDInt)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		d.logger.Errorf("internal server error in getting banners: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	content := dto.BannerForUser{Content: bannerContent}
	response.WriteResponse(w, content, http.StatusOK, d.logger)
}
