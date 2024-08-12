package delivery

import (
	"net/http"

	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
)

func (d *BannerDelivery) DeleteBannersByFeatureTag(w http.ResponseWriter, r *http.Request) {
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
	params := r.URL.Query()
	tagID := params.Get("tag_id")
	var tagIDInt uint64
	if tagID != "" {
		tagIDInt, err = d.parseParam("tag id", tagID, w)
		if err != nil {
			return
		}
	}
	featureID := params.Get("feature_id")
	var featureIDInt uint64
	if featureID != "" {
		featureIDInt, err = d.parseParam("feature id", featureID, w)
		if err != nil {
			return
		}
	}

	if featureIDInt == 0 && tagIDInt == 0 {
		errText := "no feature_id and tag_id in query params"
		d.logger.Error(errText)
		response.WriteResponse(w, response.Error{Err: errText}, http.StatusBadRequest, d.logger)
		return
	}

	sendResult := d.service.DeleteBannersByFeatureTag(featureIDInt, tagIDInt)
	if sendResult.Error != nil {
		d.logger.Errorf("internal server error in deleting banner: %v", err)
		response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	d.logger.Infof("message was put in kafka: partition: %d, offset: %d", sendResult.Partition, sendResult.Offset)
	w.WriteHeader(http.StatusNoContent)
}
