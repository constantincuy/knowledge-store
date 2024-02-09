package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status   int
	bytes    []byte
	jsonData any
	header   http.Header
	cookie   *http.Cookie
}

func (r Response) Header(key, value string) Response {
	r.header.Add(key, value)
	return r
}

func (r Response) Status(s int) Response {
	r.status = s
	return r
}

func (r Response) Json(d any) Response {
	r.jsonData = d
	return r
}

func (r Response) Cookie(cookie *http.Cookie) Response {
	r.cookie = cookie
	return r
}

func (r Response) Bytes(b []byte) Response {
	r.bytes = b
	return r
}

func (r Response) Send(w http.ResponseWriter) error {
	if r.cookie != nil {
		http.SetCookie(w, r.cookie)
	}

	for key, value := range r.header {
		w.Header().Set(key, value[0])
	}

	if r.jsonData != nil {
		w.Header().Add("Content-Type", "application/json")
		b, err := json.Marshal(r.jsonData)

		if err != nil {
			return err
		}

		_, err = w.Write(b)

		return err
	}

	w.WriteHeader(r.status)
	_, err := w.Write(r.bytes)
	return err
}

func New() Response {
	return Response{
		status: http.StatusOK,
		header: make(http.Header),
	}
}
