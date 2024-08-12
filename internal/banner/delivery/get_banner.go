package delivery

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ilyushkaaa/banner-service/internal/banner/delivery/dto"
	"github.com/ilyushkaaa/banner-service/internal/banner/filter"
	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
)

func (d *BannerDelivery) GetBanners(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r.Context())
	if err != nil {
		d.logger.Errorf("error in getting user from context: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}
	if user.Role != adminRole {
		d.logger.Errorf("user %d has got no access for getting banners", user.TagID)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	params := r.URL.Query()
	filters, err := d.getFilters(params, w)
	if err != nil {
		return
	}
	banners, err := d.service.GetBanners(r.Context(), *filters)
	if err != nil {
		d.logger.Errorf("internal server error in getting banners: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.WriteResponse(w, dto.GetBannerForAdminSlice(banners), http.StatusOK, d.logger)
}

func (d *BannerDelivery) getFilters(params url.Values, w http.ResponseWriter) (*filter.Filter, error) {
	featureID := params.Get("feature_id")
	var featureIDInt uint64
	var err error
	if featureID != "" {
		featureIDInt, err = d.parseParam("feature id", featureID, w)
		if err != nil {
			return nil, err
		}
	}
	tagID := params.Get("tag_id")
	var tagIDInt uint64
	if tagID != "" {
		tagIDInt, err = d.parseParam("tag id", tagID, w)
		if err != nil {
			return nil, err
		}
	}
	offset := params.Get("offset")
	var offsetInt uint64
	if offset != "" {
		offsetInt, err = d.parseParam("offset", offset, w)
		if err != nil {
			return nil, err
		}
	}
	limit := params.Get("limit")
	var limitInt uint64
	if limit != "" {
		limitInt, err = d.parseParam("limit", limit, w)
		if err != nil {
			return nil, err
		}
	}
	return &filter.Filter{
		FeatureID: featureIDInt,
		TagID:     tagIDInt,
		Offset:    offsetInt,
		Limit:     limitInt,
	}, nil

}

func (d *BannerDelivery) parseParam(paramName, paramValue string, w http.ResponseWriter) (uint64, error) {
	value, err := strconv.ParseUint(paramValue, 10, 64)
	if err != nil || value < 1 {
		d.logger.Errorf("error in %s conversion: %s", paramName, err)
		errText := fmt.Sprintf("%s must positive integer", paramName)
		response.WriteResponse(w, response.Error{Err: errText},
			http.StatusBadRequest, d.logger)
		return 0, fmt.Errorf("parse error")
	}
	return value, nil
}
