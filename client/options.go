package main

var opts struct {
	Nickname string `short:"n" long:"nickname" default:"Guest" description:"Nickname"`
	Args     struct {
		Server string `description:"Server address"`
	} `positional-args:"yes" required:"true"`
}
