package home

import "inspiranesia/system/http"

type HomeController struct {

}

func (home HomeController) TestEndpoint(ctx http.NougatHttpHandler) string {
	return "ini test routing !!"
}