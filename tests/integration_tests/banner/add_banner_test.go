//go:build integration
// +build integration

package banner

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ilyushkaaa/banner-service/tests/test_json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	tableName1 = "banner_tags"
	tableName2 = "previous_banners"
	tableName3 = "banners"
	tableName4 = "users"
)

var tableNames = []string{tableName1, tableName2, tableName3, tableName4}

func TestAddBanner(t *testing.T) {
	tf := newBannerTestFixtures(t)
	defer tf.Close(t)

	t.Run("error duplicate feature + tag", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodPost, "/banner", strings.NewReader(test_json.BannerAddOld))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.AddBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"such pair of feature and tag already exists"}`, string(body))
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodPost, "/banner", strings.NewReader(test_json.BannerAddNew))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.AddBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})
}
