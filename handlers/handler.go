package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mholt/binding"
	"github.com/unrolled/render"
)

type Errors struct {
	binding.Errors
}

func (e Errors) GetErrors() string {
	messages := []string{}
	for _, err := range e.Errors {
		messages = append(messages, strings.ToLower(fmt.Sprintf("%s %s", strings.Join(err.FieldNames, " "), err.Error())))
	}

	return strings.Join(messages, ", ")
}

func RenderJSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	r := render.New()
	r.JSON(rw, statusCode, map[string]interface{}{"data": data})
}

func RenderError(rw http.ResponseWriter, statusCode int, errMsg string) {
	r := render.New()
	r.JSON(rw, statusCode, map[string]string{"error": errMsg})
}
