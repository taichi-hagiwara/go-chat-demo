package main

import (
	"encoding/json"
	"fmt"

	flags "github.com/jessevdk/go-flags"
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

	if err := service.Serve(opts.Args.Listen); err != nil {
		panic(err)
	}
}
