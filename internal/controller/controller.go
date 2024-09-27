package controller

import (
	"LinkShortener/internal/db/postgres"
	"LinkShortener/internal/hashing"
	"LinkShortener/internal/jwtToken"
	"LinkShortener/internal/shortener"
	"LinkShortener/internal/util"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	client := postgres.Clients{}
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &client)
	if err != nil {
		log.Println("error unmarshalling client. Error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	client.Password = hashing.Make(client.Password)

	insert := postgres.Db.Create(&client)
	if insert.Error != nil {
		log.Println("error inserting to table client. Error:", insert.Error)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	client := postgres.Clients{}
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &client)
	if err != nil {
		log.Println("error unmarshalling client. Error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	client.Password = hashing.Make(client.Password)

	find := postgres.Clients{}
	postgres.Db.Find(&find, "login = ? AND password = ?", client.Login, client.Password)
	//if find is nil
	if (postgres.Clients{}) == find {
		w.WriteHeader(http.StatusForbidden)
	} else {
		token, err := jwtToken.Create(client.Login)
		if err != nil {
			log.Println("error creating jwt: ", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func (c *Controller) GetLink(w http.ResponseWriter, r *http.Request) {
	link := postgres.Links{}
	short := strings.Trim(r.URL.Path, "/")

	postgres.Db.Find(&link, "short = ?", short)
	if (postgres.Links{}) == link {
		w.WriteHeader(http.StatusNotFound)
	}
	http.Redirect(w, r, link.Original, http.StatusFound)
	link.SearchCount++
	postgres.Db.Save(&link)
}

func (c *Controller) AllLinks(w http.ResponseWriter, r *http.Request) {
	clientId := util.GetClientIdByToken(r.Header["Authorization"][0])
	links := []postgres.Links{}
	postgres.Db.Find(&links, "clients_id = ?", clientId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

func (c *Controller) CreateLink(w http.ResponseWriter, r *http.Request) {
	//TODO: send multiple links in a row
	var link postgres.Links
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &link)
	if err != nil {
		log.Println("error unmarshalling link. Error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	link.Short = shortener.Short(link.Original)

	link.ClientsID = util.GetClientIdByToken(r.Header["Authorization"][0])

	insert := postgres.Db.Create(&link)
	if insert.Error != nil {
		log.Println("error inserting to table links. Error:", insert.Error)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c *Controller) DeleteLink(w http.ResponseWriter, r *http.Request) {
	link := postgres.Links{}
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &link)
	if err != nil {
		log.Println("error unmarshalling link. Error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	postgres.Db.Find(&link, "short = ?", link.Short)
	if id := util.GetClientIdByToken(r.Header["Authorization"][0]); id == link.ClientsID {
		postgres.Db.Delete(&link)
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusForbidden)
}
