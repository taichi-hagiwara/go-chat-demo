package service

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

// ChatService は、チャットの管理を表す。
type ChatService struct {
	nicknames map[xid.ID]string
}

type RegisterArgs struct {
	Nickname string
}

type RegisterReply struct {
	UserID xid.ID
}

type LeaveArgs struct {
	UserID xid.ID
}

type PostArgs struct {
	UserID xid.ID
	Text   string
}

func (c *ChatService) Register(args *RegisterArgs, reply *RegisterReply) error {
	id := xid.New()
	c.nicknames[id] = args.Nickname
	reply.UserID = id
	fmt.Printf("<SERVER> %s has joined as %s.\n", id, args.Nickname)
	return nil
}

func (c *ChatService) Leave(args *LeaveArgs, reply *struct{}) error {
	if nn, ok := c.nicknames[args.UserID]; ok {
		fmt.Printf("<SERVER> %s(%s) has been left.\n", args.UserID, nn)
		delete(c.nicknames, args.UserID)
		return nil
	}
	return errors.Errorf("unknown id: %s", args.UserID)
}

func (c *ChatService) Post(args *PostArgs, reply *struct{}) error {
	if nn, ok := c.nicknames[args.UserID]; ok {
		fmt.Printf("<%s> %s\n", nn, args.Text)
		return nil
	}
	return errors.Errorf("unknown id: %s", args.UserID)
}
