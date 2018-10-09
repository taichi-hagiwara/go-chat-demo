package main

import (
	"fmt"
	"log"
	"time"

	"github.com/taichi-hagiwara/go-chat-demo/service"

	"github.com/taichi-hagiwara/ezrpc"
)

var logs = make(map[string][]*service.Post)

func main() {
	server := ezrpc.NewServer(service.ChatService())
	server.RegisterHandler("post", postHandler)
	if err := server.Listen("localhost:8080", &ezrpc.CertInfo{
		CACert:  "ca.pem",
		Cert:    "cert.pem",
		Private: "private.pem",
	}); err != nil {
		log.Fatal(err)
	}
}

func postHandler(client *ezrpc.ClientInfo, args interface{}) (result interface{}, err error) {
	a := args.(*service.PostArgs)
	post := &service.Post{Text: a.Text, Name: client.TLSSubject.CommonName, Time: time.Now()}
	fmt.Printf("%s %s <%s> %s\n", client.Remote, post.Time, post.Name, post.Text)

	if _, ok := logs[client.TLSSubject.CommonName]; !ok {
		logs[client.TLSSubject.CommonName] = []*service.Post{}
	}
	for i := range logs {
		logs[i] = append(logs[i], post)
	}

	log := logs[client.TLSSubject.CommonName]
	logs[client.TLSSubject.CommonName] = []*service.Post{}
	return &service.PostResult{Log: log}, nil
}
