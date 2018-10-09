package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/taichi-hagiwara/ezrpc"
	"github.com/taichi-hagiwara/go-chat-demo/service"
)

func main() {
	client, err := ezrpc.NewClient(service.ChatService(), "localhost:8080", "Server", &ezrpc.CertInfo{
		CACert:  "ca.pem",
		Cert:    "cert.pem",
		Private: "private.pem",
	})

	if err != nil {
		log.Fatal(err)
	}

	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		if stdin.Text() == "/exit" {
			break
		}

		r, err := client.Invoke("post", &service.PostArgs{Text: stdin.Text()})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		result := r.(*service.PostResult)
		for i := range result.Log {
			fmt.Printf("%s <%s> %s\n", result.Log[i].Time, result.Log[i].Name, result.Log[i].Text)
		}

	}

	if stdin.Err() != nil {
		panic(stdin.Err())
	}
}
