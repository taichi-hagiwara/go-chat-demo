package service

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

// Serve は、チャットのサービスを開始する。
func Serve(listen string) error {
	service := &ChatService{
		nicknames: make(map[xid.ID]string),
	}

	rpc.Register(service)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", listen)
	if err != nil {
		return errors.Wrapf(err, "cannot listen tcp %s", listen)
	}

	return http.Serve(l, nil)
}
