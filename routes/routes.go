package routes

import (
	"net/http"
	"text/template"

	"github.com/go_server/config"
	"github.com/go_server/logger"
	"github.com/go_server/models"
	"github.com/go_server/response"
	"github.com/go_server/store"
	"github.com/go_server/utils"
	"github.com/go_server/validation"
	"github.com/gorilla/mux"
)

var httpMethods = struct {
	GET    string
	POST   string
	DELETE string
	PUT    string
	PATCH  string
}{
	"GET",
	"POST",
	"DELETE",
	"PUT",
	"PATCH",
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl, _ := template.ParseFiles("templates/index.html")
	tpl.Execute(w, nil)

}

func upload(w http.ResponseWriter, r *http.Request) {
	auth := r.PostFormValue("auth")
	if auth != config.EnvironmentValues.AuthToken {
		logger.Debug("Auth token not matched")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	file, fileHeader, err := r.FormFile("upload")
	if err != nil {
		logger.Debug("Invalid Upload")
		logger.Debug(err.Error())
	}
	defer file.Close()

	if !validation.MimeType((validation.FileValidationParams{Seeker: file})) {
		logger.Debug("Invalid Mime Type")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !validation.FileSize(validation.FileValidationParams{Size: int(fileHeader.Size)}) {
		logger.Debug("Invalid File Size")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = utils.WriteDataToFile(fileHeader, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mimeType, _ := utils.GetContentType(file)

	fileInfo := models.FileData{
		ContentType: mimeType,
		Size:        int(fileHeader.Size),
		FileName:    fileHeader.Filename,
	}

	id, err := store.Store.Write(fileInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(w, http.StatusCreated, id)

}

//Register All Routes
func RegisterRoutes() http.Handler {
	r := mux.NewRouter().PathPrefix(config.AppParams.ApiBase).Subrouter()
	r.Handle("/", http.HandlerFunc(index))
	r.Handle("/upload", http.HandlerFunc(upload)).Methods("POST")
	return r
}
