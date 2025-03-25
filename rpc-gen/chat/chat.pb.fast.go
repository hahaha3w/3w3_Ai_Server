// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package chat

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Conversation) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Conversation[number], err)
}

func (x *Conversation) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ConversationId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Conversation) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Conversation) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SessionTitle, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Conversation) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Mode, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Conversation) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CreateTime, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Message[number], err)
}

func (x *Message) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MessageId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Message) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ConversationId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Message) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Message) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SenderType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.SendTime, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateConversationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateConversationRequest[number], err)
}

func (x *CreateConversationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CreateConversationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SessionTitle, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateConversationRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Mode, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateConversationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateConversationResponse[number], err)
}

func (x *CreateConversationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Conversation
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Conversation = &v
	return offset, nil
}

func (x *ListConversationsRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListConversationsRequest[number], err)
}

func (x *ListConversationsRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListConversationsRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListConversationsRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PageNumber, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListConversationsResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListConversationsResponse[number], err)
}

func (x *ListConversationsResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Conversation
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Conversations = append(x.Conversations, &v)
	return offset, nil
}

func (x *ListConversationsResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetConversationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetConversationRequest[number], err)
}

func (x *GetConversationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ConversationId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetConversationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetConversationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetConversationResponse[number], err)
}

func (x *GetConversationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Conversation
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Conversation = &v
	return offset, nil
}

func (x *SendMessageRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMessageRequest[number], err)
}

func (x *SendMessageRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ConversationId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *SendMessageRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *SendMessageRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendMessageResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMessageResponse[number], err)
}

func (x *SendMessageResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Message
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Message = &v
	return offset, nil
}

func (x *ListMessagesRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListMessagesRequest[number], err)
}

func (x *ListMessagesRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ConversationId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListMessagesRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListMessagesRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListMessagesRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PageNumber, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListMessagesResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListMessagesResponse[number], err)
}

func (x *ListMessagesResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Message
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Messages = append(x.Messages, &v)
	return offset, nil
}

func (x *ListMessagesResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DeleteConversationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteConversationRequest[number], err)
}

func (x *DeleteConversationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ConversationId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DeleteConversationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteConversationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteConversationResponse[number], err)
}

func (x *DeleteConversationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Conversation) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *Conversation) fastWriteField1(buf []byte) (offset int) {
	if x.ConversationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetConversationId())
	return offset
}

func (x *Conversation) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *Conversation) fastWriteField3(buf []byte) (offset int) {
	if x.SessionTitle == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSessionTitle())
	return offset
}

func (x *Conversation) fastWriteField4(buf []byte) (offset int) {
	if x.Mode == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetMode())
	return offset
}

func (x *Conversation) fastWriteField5(buf []byte) (offset int) {
	if x.CreateTime == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetCreateTime())
	return offset
}

func (x *Message) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *Message) fastWriteField1(buf []byte) (offset int) {
	if x.MessageId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetMessageId())
	return offset
}

func (x *Message) fastWriteField2(buf []byte) (offset int) {
	if x.ConversationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetConversationId())
	return offset
}

func (x *Message) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *Message) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetContent())
	return offset
}

func (x *Message) fastWriteField5(buf []byte) (offset int) {
	if x.SenderType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSenderType())
	return offset
}

func (x *Message) fastWriteField6(buf []byte) (offset int) {
	if x.SendTime == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetSendTime())
	return offset
}

func (x *CreateConversationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CreateConversationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CreateConversationRequest) fastWriteField2(buf []byte) (offset int) {
	if x.SessionTitle == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSessionTitle())
	return offset
}

func (x *CreateConversationRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Mode == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetMode())
	return offset
}

func (x *CreateConversationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateConversationResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Conversation == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetConversation())
	return offset
}

func (x *ListConversationsRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ListConversationsRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *ListConversationsRequest) fastWriteField2(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetPageSize())
	return offset
}

func (x *ListConversationsRequest) fastWriteField3(buf []byte) (offset int) {
	if x.PageNumber == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetPageNumber())
	return offset
}

func (x *ListConversationsResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListConversationsResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Conversations == nil {
		return offset
	}
	for i := range x.GetConversations() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetConversations()[i])
	}
	return offset
}

func (x *ListConversationsResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetConversationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetConversationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ConversationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetConversationId())
	return offset
}

func (x *GetConversationRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *GetConversationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetConversationResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Conversation == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetConversation())
	return offset
}

func (x *SendMessageRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *SendMessageRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ConversationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetConversationId())
	return offset
}

func (x *SendMessageRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *SendMessageRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetContent())
	return offset
}

func (x *SendMessageResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SendMessageResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Message == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetMessage())
	return offset
}

func (x *ListMessagesRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ListMessagesRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ConversationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetConversationId())
	return offset
}

func (x *ListMessagesRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *ListMessagesRequest) fastWriteField3(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetPageSize())
	return offset
}

func (x *ListMessagesRequest) fastWriteField4(buf []byte) (offset int) {
	if x.PageNumber == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.GetPageNumber())
	return offset
}

func (x *ListMessagesResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListMessagesResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Messages == nil {
		return offset
	}
	for i := range x.GetMessages() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetMessages()[i])
	}
	return offset
}

func (x *ListMessagesResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *DeleteConversationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteConversationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ConversationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetConversationId())
	return offset
}

func (x *DeleteConversationRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *DeleteConversationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteConversationResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *Conversation) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *Conversation) sizeField1() (n int) {
	if x.ConversationId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetConversationId())
	return n
}

func (x *Conversation) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *Conversation) sizeField3() (n int) {
	if x.SessionTitle == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSessionTitle())
	return n
}

func (x *Conversation) sizeField4() (n int) {
	if x.Mode == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetMode())
	return n
}

func (x *Conversation) sizeField5() (n int) {
	if x.CreateTime == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetCreateTime())
	return n
}

func (x *Message) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *Message) sizeField1() (n int) {
	if x.MessageId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetMessageId())
	return n
}

func (x *Message) sizeField2() (n int) {
	if x.ConversationId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetConversationId())
	return n
}

func (x *Message) sizeField3() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetUserId())
	return n
}

func (x *Message) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetContent())
	return n
}

func (x *Message) sizeField5() (n int) {
	if x.SenderType == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetSenderType())
	return n
}

func (x *Message) sizeField6() (n int) {
	if x.SendTime == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetSendTime())
	return n
}

func (x *CreateConversationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CreateConversationRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *CreateConversationRequest) sizeField2() (n int) {
	if x.SessionTitle == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSessionTitle())
	return n
}

func (x *CreateConversationRequest) sizeField3() (n int) {
	if x.Mode == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetMode())
	return n
}

func (x *CreateConversationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateConversationResponse) sizeField1() (n int) {
	if x.Conversation == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetConversation())
	return n
}

func (x *ListConversationsRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ListConversationsRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *ListConversationsRequest) sizeField2() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetPageSize())
	return n
}

func (x *ListConversationsRequest) sizeField3() (n int) {
	if x.PageNumber == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetPageNumber())
	return n
}

func (x *ListConversationsResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListConversationsResponse) sizeField1() (n int) {
	if x.Conversations == nil {
		return n
	}
	for i := range x.GetConversations() {
		n += fastpb.SizeMessage(1, x.GetConversations()[i])
	}
	return n
}

func (x *ListConversationsResponse) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetTotal())
	return n
}

func (x *GetConversationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetConversationRequest) sizeField1() (n int) {
	if x.ConversationId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetConversationId())
	return n
}

func (x *GetConversationRequest) sizeField2() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUserId())
	return n
}

func (x *GetConversationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetConversationResponse) sizeField1() (n int) {
	if x.Conversation == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetConversation())
	return n
}

func (x *SendMessageRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *SendMessageRequest) sizeField1() (n int) {
	if x.ConversationId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetConversationId())
	return n
}

func (x *SendMessageRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *SendMessageRequest) sizeField3() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetContent())
	return n
}

func (x *SendMessageResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SendMessageResponse) sizeField1() (n int) {
	if x.Message == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetMessage())
	return n
}

func (x *ListMessagesRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ListMessagesRequest) sizeField1() (n int) {
	if x.ConversationId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetConversationId())
	return n
}

func (x *ListMessagesRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *ListMessagesRequest) sizeField3() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetPageSize())
	return n
}

func (x *ListMessagesRequest) sizeField4() (n int) {
	if x.PageNumber == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.GetPageNumber())
	return n
}

func (x *ListMessagesResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListMessagesResponse) sizeField1() (n int) {
	if x.Messages == nil {
		return n
	}
	for i := range x.GetMessages() {
		n += fastpb.SizeMessage(1, x.GetMessages()[i])
	}
	return n
}

func (x *ListMessagesResponse) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetTotal())
	return n
}

func (x *DeleteConversationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteConversationRequest) sizeField1() (n int) {
	if x.ConversationId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetConversationId())
	return n
}

func (x *DeleteConversationRequest) sizeField2() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUserId())
	return n
}

func (x *DeleteConversationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteConversationResponse) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

var fieldIDToName_Conversation = map[int32]string{
	1: "ConversationId",
	2: "UserId",
	3: "SessionTitle",
	4: "Mode",
	5: "CreateTime",
}

var fieldIDToName_Message = map[int32]string{
	1: "MessageId",
	2: "ConversationId",
	3: "UserId",
	4: "Content",
	5: "SenderType",
	6: "SendTime",
}

var fieldIDToName_CreateConversationRequest = map[int32]string{
	1: "UserId",
	2: "SessionTitle",
	3: "Mode",
}

var fieldIDToName_CreateConversationResponse = map[int32]string{
	1: "Conversation",
}

var fieldIDToName_ListConversationsRequest = map[int32]string{
	1: "UserId",
	2: "PageSize",
	3: "PageNumber",
}

var fieldIDToName_ListConversationsResponse = map[int32]string{
	1: "Conversations",
	2: "Total",
}

var fieldIDToName_GetConversationRequest = map[int32]string{
	1: "ConversationId",
	2: "UserId",
}

var fieldIDToName_GetConversationResponse = map[int32]string{
	1: "Conversation",
}

var fieldIDToName_SendMessageRequest = map[int32]string{
	1: "ConversationId",
	2: "UserId",
	3: "Content",
}

var fieldIDToName_SendMessageResponse = map[int32]string{
	1: "Message",
}

var fieldIDToName_ListMessagesRequest = map[int32]string{
	1: "ConversationId",
	2: "UserId",
	3: "PageSize",
	4: "PageNumber",
}

var fieldIDToName_ListMessagesResponse = map[int32]string{
	1: "Messages",
	2: "Total",
}

var fieldIDToName_DeleteConversationRequest = map[int32]string{
	1: "ConversationId",
	2: "UserId",
}

var fieldIDToName_DeleteConversationResponse = map[int32]string{
	1: "Success",
}
