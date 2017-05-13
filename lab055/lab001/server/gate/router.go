package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	msg.JSONProcessor.SetRouter(&msg.C2S_AddUser{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.C2S_Message{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.C2S_Action{}, game.ChanRPC)
}
