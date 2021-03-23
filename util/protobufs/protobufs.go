package protobufs

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/openziti/foundation/channel2"
	"time"
)

func Send(ch channel2.Channel, msg channel2.TypedMessage) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	envelopeMsg := channel2.NewMessage(msg.GetContentType(), body)
	return ch.Send(envelopeMsg)
}

func SendReply(ch channel2.Channel, sourceMsg *channel2.Message, msg channel2.TypedMessage) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	envelopeMsg := channel2.NewMessage(msg.GetContentType(), body)
	envelopeMsg.ReplyTo(sourceMsg)
	return ch.Send(envelopeMsg)
}

func SendWithTimeout(ch channel2.Channel, msg channel2.TypedMessage, timeout time.Duration) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	envelopeMsg := channel2.NewMessage(msg.GetContentType(), body)
	return ch.SendWithTimeout(envelopeMsg, timeout)
}

func SendReplyWithTimeout(ch channel2.Channel, sourceMessage *channel2.Message, msg channel2.TypedMessage, timeout time.Duration) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	envelopeMsg := channel2.NewMessage(msg.GetContentType(), body)
	envelopeMsg.ReplyTo(sourceMessage)
	return ch.SendWithTimeout(envelopeMsg, timeout)
}

func SendForReply(ch channel2.Channel, msg channel2.TypedMessage, timeout time.Duration) (*channel2.Message, error) {
	body, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	envelopeMsg := channel2.NewMessage(msg.GetContentType(), body)
	waitCh, err := ch.SendAndWait(envelopeMsg)
	if err != nil {
		return nil, err
	}

	select {
	case responseMsg := <-waitCh:
		return responseMsg, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("timed out waiting for response to request of type %v", msg.GetContentType())
	}
}

func SendForReplyAndDecode(ch channel2.Channel, msg channel2.TypedMessage, timeout time.Duration, result channel2.TypedMessage) error {
	responseMsg, err := SendForReply(ch, msg, timeout)
	if err != nil {
		return err
	}
	if responseMsg.ContentType != result.GetContentType() {
		return fmt.Errorf("unexpected response type %v to request of type %v. expected %v",
			responseMsg.ContentType, msg.GetContentType(), result.GetContentType())
	}
	if err := proto.Unmarshal(responseMsg.Body, result); err != nil {
		return err
	}
	return nil
}
