package handler

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"errors"
	"encoding/json"
	"fmt"
	"./responder"
)

var fundamentalParamiters = []string{"imgLink", "orgName", "name", "role", "area", "address", "phone1", "phone2"}

func CreateNewSignature(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		responder.Respond(w, responder.JSONError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		responder.Respond(w, responder.JSONError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	err = validateBody(m)
	if err != nil {
		responder.Respond(w, responder.JSONError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	tel2 := m["phone2"]
	if tel2 != nil {
		tpl, _ := template.ParseFiles("./templates/model2.html")
		w.WriteHeader(http.StatusOK)
		tpl.Execute(w, m)
		return
	}

	tpl, _ := template.ParseFiles("./templates/model.html")
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, m)
}

func validateBody(body map[string]interface{}) error {
	errorMsg := "Missing paramiter: \"%s\""
	for _, val := range fundamentalParamiters {
		if isMissingParam(body, val) {
			return errors.New(fmt.Sprintf(errorMsg, val))
		}
	}

	return nil
}

func isMissingParam(body map[string]interface{}, fieldName string) bool {
	_, ok := body[fieldName]
	return !ok
}
