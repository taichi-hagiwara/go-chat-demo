package main

var opts struct {
	Args struct {
		Listen string `description:"Listen option"`
	} `positional-args:"yes" required:"true"`
}
