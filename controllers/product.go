package controllers

import (
	"github.com/gorilla/mux"
	"gowebapplication/controllers/util"
	"gowebapplication/viewmodels"
	"net/http"
	"strconv"
	"text/template"
)

type productController struct {
	template *template.Template
}

func (this *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.GetProduct(id)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()

		this.template.Execute(responseWriter, vm)
	} else {
		w.WriteHeader(404)
	}
}
