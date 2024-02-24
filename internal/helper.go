package internal

import (
	"html/template"
	"log"
	"net/http"
)

func respondedWithError(w http.ResponseWriter, code int, logMsg string, httmMsg string, err error) {
	log.Println(logMsg, err)
	http.Error(w, httmMsg, code)
}

// Executes the given template path.
func executeTemplate(filePath string, w http.ResponseWriter, ctx map[string]interface{}) {
	t, err := template.ParseFiles(filePath)
	if err != nil {
		respondedWithError(w, http.StatusInternalServerError, "Error parsing temaple:", "Internal server error", err)
		return
	}

	err = t.Execute(w, ctx)
	if err != nil {
		respondedWithError(w, http.StatusInternalServerError, "Error in template execution:", "Internal server error", err)
		return
	}
}
