package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirsean/adamwestum/service"
	"net/http"
	"strconv"
)

func Lines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, _ := strconv.ParseInt(vars["number"], 10, 32)

	lines := service.Lines(int(number))

	type LinesResponse struct {
		Number int64
		Lines []string
	}

	response, _ := json.Marshal(LinesResponse{
		Number: number,
		Lines: lines,
	})
	w.Write(response)
}
