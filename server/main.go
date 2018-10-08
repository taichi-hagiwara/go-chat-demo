package main

import (
	"fmt"
	"log"
	"time"

	"github.com/taichi-hagiwara/go-chat-demo-v2/service"

	"github.com/taichi-hagiwara/ezrpc"
)

func main() {
	if err := ezrpc.Serve(&chatServer{
		Service: service.ChatService(),
		Log:     make(map[string][]*service.Post),
	}, "localhost:8080", &ezrpc.CertInfo{
		CACert:  "ca.pem",
		Cert:    "cert.pem",
		Private: "private.pem",
	}); err != nil {
		log.Fatal(err)
	}
}

type chatServer struct {
	ezrpc.Service

	Log map[string][]*service.Post
}

func (s *chatServer) Invoke(name string, client *ezrpc.ClientInfo, args interface{}) (result interface{}, err error) {
	if name == "post" {
		args := args.(*service.PostArgs)
		post := &service.Post{Text: args.Text, Name: client.TLSSubject.CommonName, Time: time.Now()}
		fmt.Printf("%s <%s> %s\n", post.Time, post.Name, post.Text)

		if _, ok := s.Log[client.TLSSubject.CommonName]; !ok {
			s.Log[client.TLSSubject.CommonName] = []*service.Post{}
		}
		for i := range s.Log {
			s.Log[i] = append(s.Log[i], post)
		}

		log := s.Log[client.TLSSubject.CommonName]
		s.Log[client.TLSSubject.CommonName] = []*service.Post{}
		return &service.PostResult{Log: log}, nil

	}
	return nil, nil
}
