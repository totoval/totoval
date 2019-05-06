package events

import (
	"github.com/golang/protobuf/proto"

	"github.com/totoval/framework/hub"
	pbs "totoval/app/events/protocol_buffers"
)

func init() {
	hub.Make(&UserRegistered{})
}

type UserRegistered struct {
	hub.Event
}

func (ur *UserRegistered) ParamProto() proto.Message {
	return &pbs.UserRegistered{}
}
