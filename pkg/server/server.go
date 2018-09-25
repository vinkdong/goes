package server

import (
	"fmt"
	"github.com/vinkdong/gox/log"
	"net/http"
	"io/ioutil"
)

type Server struct {
	Mux  *http.ServeMux
	Addr string
	Next string
	Version string
}

func NewServer() *Server {
	s := &Server{}
	mux := http.NewServeMux()
	s.Mux = mux
	return s
}

func (s *Server) Run() error {
	s.Mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		requestIp := request.RemoteAddr
		log.Infof("get request from %s", requestIp)
		writer.Write([]byte(fmt.Sprintf("Your request address is %s\nServer version is %s\n", requestIp, s.Version)))
	})

	// just sample request another
	s.Mux.HandleFunc("/redirect", func(writer http.ResponseWriter, request *http.Request) {
		requestIp := request.RemoteAddr
		log.Infof("get request from %s", requestIp)
		log.Infof("headers %s", request.Header)
		log.Infof("send request to %s", s.Next)
		resp, err := http.Get(s.Next)
		if err != nil {
			log.Error(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		writer.Write(body)
	})
	return http.ListenAndServe(s.Addr, s.Mux)
}
