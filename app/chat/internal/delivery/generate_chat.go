package delivery

import (
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
	"log"
)

// StreamResponseEcho implements the Echo interface.
func (d *ChatDelivery) Chat(req *chat.ChatReq, stream chat.Echo_ChatServer) (err error) {
	err = d.usecase.Chat(req.Message, stream)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
