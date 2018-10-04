package service

import (
	"net/rpc"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type Client struct {
	c  *rpc.Client
	id xid.ID
}

// Dial は、チャットのクライアントを取得する。
func Dial(server, nickname string) (*Client, error) {
	c, err := rpc.DialHTTP("tcp", server)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial server: %s", server)
	}

	client := &Client{c: c}

	var reply RegisterReply
	if err := client.register(&RegisterArgs{Nickname: nickname}, &reply); err != nil {
		return nil, errors.Wrap(err, "failed to login")
	}
	client.id = reply.UserID

	return client, nil
}

func (c *Client) UserID() xid.ID {
	return c.id
}

func (c *Client) Close() error {
	err := c.leave(&LeaveArgs{UserID: c.id}, &struct{}{})
	if err != nil {
		c.c.Close()
		return errors.Wrap(err, "failed to leave")
	}

	return c.c.Close()
}

func (c *Client) register(args *RegisterArgs, reply *RegisterReply) error {
	if args.Nickname == "SERVER" {
		return errors.Errorf("invalid nickname: %s", args.Nickname)
	}
	return c.c.Call("ChatService.Register", args, reply)
}

func (c *Client) leave(args *LeaveArgs, reply *struct{}) error {
	return c.c.Call("ChatService.Leave", args, reply)
}
func (c *Client) Post(args *PostArgs, reply *struct{}) error {
	return c.c.Call("ChatService.Post", args, reply)
}
