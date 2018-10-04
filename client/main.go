package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/taichi-hagiwara/go-chat-demo/service"
)

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if h, ok := err.(*flags.Error); !ok || h.Type != flags.ErrHelp {
			panic(err)
		}
	}

	s, err := json.MarshalIndent(opts, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	client, err := service.Dial(opts.Args.Server, opts.Nickname)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		if stdin.Text() == "/exit" {
			break
		}

		err := client.Post(&service.PostArgs{
			UserID: client.UserID(),
			Text:   stdin.Text(),
		}, &struct{}{})

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}

	if stdin.Err() != nil {
		panic(stdin.Err())
	}
}
