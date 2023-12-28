package handlers

import (
	"bytes"
	"embed"
	"encoding/json"
	"github.com/aadi-1024/quickshare/pkg/auth"
	"github.com/aadi-1024/quickshare/pkg/config"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*.gohtml
var templates embed.FS

type Repository struct {
	config *config.Config
}

var Repo *Repository

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.New("base")
	var err error
	t, err = t.ParseFS(templates, "templates/*.gohtml")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	buf := new(bytes.Buffer)

	//if err = t.Execute(buf, nil); err != nil {
	//	log.Fatalln("error while writing into buffer", err)
	//}
	if err = t.ExecuteTemplate(buf, "base", nil); err != nil {
		log.Fatalln("error", err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalln(err)
	}
}

type jsonResponse struct {
	Ok   bool   `json:"ok"`
	Link string `json:"link"`
}

type jsonRequest struct {
	Code string `json:"code"`
}

func Verify(w http.ResponseWriter, r *http.Request) {
	req := &jsonRequest{}
	resp := &jsonResponse{}

	err := json.NewDecoder(r.Body).Decode(req)

	pass := req.Code

	if auth.MatchAndVerify(pass) {
		resp.Ok = true
		resp.Link = "/file" //link will actually be a randomly generated hash, placeholder for now
	} else {
		resp.Ok = false
		resp.Link = "/"
	}

	js, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Fatalln(err)
	}
}

func File(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, Repo.config.Filename)
}

func InitRepo(app *config.Config) {
	Repo = &Repository{
		config: app,
	}
}
