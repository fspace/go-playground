package api

import "net/http"

func RenderError(err error) {
	/**
	if err != nil {
		api.RenderError(err)
		return
	}
	*/
}

func RenderJSON(w http.ResponseWriter, data interface{}) {

}
