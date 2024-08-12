//go:build integration
// +build integration

package banner

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestApplyBannerVersion(t *testing.T) {
	tf := newBannerTestFixtures(t)
	defer tf.Close(t)

	t.Run("banner version not found", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodPost, "/banner/versions/10", nil)
		request = mux.SetURLVars(request, map[string]string{"id": "10"})
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.ApplyBannerVersion(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodPost, "/banner/versions/1", nil)
		request = mux.SetURLVars(request, map[string]string{"id": "1"})
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.ApplyBannerVersion(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
