package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilyushkaaa/banner-service/internal/banner/delivery"
	"github.com/ilyushkaaa/banner-service/internal/middleware"
)

func GetRouter(handlers *delivery.BannerDelivery, mw *middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	assignRoutes(router, handlers)
	assignMiddleware(router, mw)
	return router
}

func assignRoutes(router *mux.Router, handlers *delivery.BannerDelivery) {
	router.HandleFunc("/user_banner", handlers.GetUserBanner).Methods(http.MethodGet)
	router.HandleFunc("/banner", handlers.GetBanners).Methods(http.MethodGet)
	router.HandleFunc("/banner", handlers.AddBanner).Methods(http.MethodPost)
	router.HandleFunc("/banner/{id}", handlers.UpdateBanner).Methods(http.MethodPatch)
	router.HandleFunc("/banner/{id}", handlers.DeleteBanner).Methods(http.MethodDelete)
	router.HandleFunc("/banner/{id}/versions", handlers.GetBannerVersions).Methods(http.MethodGet)
	router.HandleFunc("/banner/versions/{id}", handlers.ApplyBannerVersion).Methods(http.MethodPost)
	router.HandleFunc("/banners", handlers.DeleteBannersByFeatureTag).Methods(http.MethodDelete)

}

func assignMiddleware(router *mux.Router, mw *middleware.Middleware) {
	router.Use(mw.AccessLog)
	router.Use(mw.Auth)
}
