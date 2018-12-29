package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type httpHandler struct {
	g *Gang
}

func ServeHttp(g *Gang, serv string) error {
	h := &httpHandler{g: g}
	return http.ListenAndServe(serv, h)
}

func (h *httpHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	p, err := UnmarshalProfile(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	if err := h.g.Post(*p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "OK")
}

func (h *httpHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	if err := h.g.Delete(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	io.WriteString(w, "OK")
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("http: %+v", *r)
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodDelete:
		h.handleDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}