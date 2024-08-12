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

func TestGetBanners(t *testing.T) {
	tf := newBannerTestFixtures(t)
	defer tf.Close(t)

	t.Run("ok, get all banners", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodGet, "/banner", nil)
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.GetBanners(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ExpectedAllBanners, string(body))
	})

	t.Run("ok, get banners by feature", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodGet, "/banner?feature_id=1", nil)
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.GetBanners(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ExpectedBannersByFeature1, string(body))
	})
}
