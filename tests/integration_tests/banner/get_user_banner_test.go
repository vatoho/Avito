//go:build integration
// +build integration

package banner

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilyushkaaa/banner-service/tests/test_json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserBanner(t *testing.T) {
	tf := newBannerTestFixtures(t)
	defer tf.Close(t)

	t.Run("banner not found", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodGet, "/user_banner?tag_id=1&feature_id=18", nil)
		request.Header.Set("Token", "user_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.GetUserBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodGet, "/user_banner?tag_id=1&feature_id=1", nil)
		request.Header.Set("Token", "user_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.GetUserBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ExpectedBanner1, string(body))
	})
}
