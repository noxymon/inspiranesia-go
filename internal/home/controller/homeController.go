package controller

import (
	"inspiranesia/system/http"
	nh "net/http"
)

type HomeController struct {
	name string
}

func newHomeController() *HomeController {
	return &HomeController{"Home Controller"}
}

func (h HomeController) GetHome(param ...interface{}) *http.NougatHttpResponse {
	if param[0] == "nok" {
		return &http.NougatHttpResponse{StatusCode: nh.StatusBadRequest, Body: "NOK !"}
	}
	return &http.NougatHttpResponse{StatusCode: nh.StatusOK, Body: "Ok !"}
}
