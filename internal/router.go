package internal

import (
	"inspiranesia/internal/home"
	"inspiranesia/system/http"
)

func RegisterRoute(handler http.NougatHttpHandler){
	homeController := home.HomeController{}
	handler.Get("/test", handler.GetContext())
}
