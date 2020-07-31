package fhir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	schema "github.com/MedHubUz/fhirschema"
	"github.com/go-chi/chi"
	"github.com/medhub-hex/pkg/http/rest"
	"github.com/medhub-hex/pkg/rand"
	"github.com/valyala/fastjson"
)

type Handler struct {
	fhirService Service
}

func NewHandler(s Service) rest.RESTful {
	return &Handler{s}
}

func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.List)
	r.Post("/", h.Post)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Put("/", h.Put)
		r.Delete("/", h.Delete)
	})
	return r
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	resource := &Resource{
		ResourceType: chi.URLParam(r, "resourceType"),
		ResourceID:   chi.URLParam(r, "id"),
	}

	output, err := h.fhirService.Read(r.Context(), resource)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(output)
	return
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	resourceType := chi.URLParam(r, "resourceType")
	fhirResource, err := schema.GetFhirResourceInstance(resourceType)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(requestBody, &fhirResource); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resource := &Resource{
		ResourceType:   resourceType,
		Data:           string(requestBody),
		CreatedAt:      time.Now().UTC().Format("2006-01-02 15:04:05"),
		PractitionerID: "1",
	}

	var fjpool fastjson.ParserPool
	fjparser := fjpool.Get()
	jsonResource, err := fjparser.ParseBytes(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if string(jsonResource.GetStringBytes("resourceType")) != resourceType {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if id := jsonResource.GetStringBytes("id"); len(id) != 0 {
		resource.ResourceID = string(id)
		fjpool.Put(fjparser)
	}

	resource.ResourceID = rand.String(16)
	// fmt.Println(newID)

	output, err := h.fhirService.Create(r.Context(), resource)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(output)

	// b, ok := resource.Data.([]byte)
	// fmt.Println(b, ok)
	// dec := json.NewEncoder(w)
	// dec.Encode(resource.Data)
	// render.JSON(w, r, resource.Data)
	return
}

func (h Handler) Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put"))
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
