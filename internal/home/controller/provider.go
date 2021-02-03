package controller

import "inspiranesia/system/http"

func Provide(adapter *http.NougatHttpHandlerAdapter) {
	controller := newHomeController()

	adapter.GetAndReturnStringF("/home", controller.GetHome)
}
