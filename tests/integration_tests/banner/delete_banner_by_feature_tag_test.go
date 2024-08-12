//go:build integration
// +build integration

package banner

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBannerByFeatureTag(t *testing.T) {
	tf := newBannerTestFixtures(t)
	defer tf.Close(t)

	t.Run("ok", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodDelete, "/banners?feature_id=1", nil)
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.DeleteBannersByFeatureTag(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()

		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}
