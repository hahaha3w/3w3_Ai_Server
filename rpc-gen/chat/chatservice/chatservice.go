// Code generated by Kitex v0.12.1. DO NOT EDIT.

package chatservice

import (
	"context"
	"errors"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	chat "github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateConversation": kitex.NewMethodInfo(
		createConversationHandler,
		newCreateConversationArgs,
		newCreateConversationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListConversations": kitex.NewMethodInfo(
		listConversationsHandler,
		newListConversationsArgs,
		newListConversationsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SendMessage": kitex.NewMethodInfo(
		sendMessageHandler,
		newSendMessageArgs,
		newSendMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingServer),
	),
	"ListMessages": kitex.NewMethodInfo(
		listMessagesHandler,
		newListMessagesArgs,
		newListMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteConversation": kitex.NewMethodInfo(
		deleteConversationHandler,
		newDeleteConversationArgs,
		newDeleteConversationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	chatServiceServiceInfo                = NewServiceInfo()
	chatServiceServiceInfoForClient       = NewServiceInfoForClient()
	chatServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return chatServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return chatServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return chatServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(true, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ChatService"
	handlerType := (*chat.ChatService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "chat",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func createConversationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chat.CreateConversationRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chat.ChatService).CreateConversation(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateConversationArgs:
		success, err := handler.(chat.ChatService).CreateConversation(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateConversationResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateConversationArgs() interface{} {
	return &CreateConversationArgs{}
}

func newCreateConversationResult() interface{} {
	return &CreateConversationResult{}
}

type CreateConversationArgs struct {
	Req *chat.CreateConversationRequest
}

func (p *CreateConversationArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chat.CreateConversationRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateConversationArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateConversationArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateConversationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateConversationArgs) Unmarshal(in []byte) error {
	msg := new(chat.CreateConversationRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateConversationArgs_Req_DEFAULT *chat.CreateConversationRequest

func (p *CreateConversationArgs) GetReq() *chat.CreateConversationRequest {
	if !p.IsSetReq() {
		return CreateConversationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateConversationArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateConversationArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateConversationResult struct {
	Success *chat.CreateConversationResponse
}

var CreateConversationResult_Success_DEFAULT *chat.CreateConversationResponse

func (p *CreateConversationResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chat.CreateConversationResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateConversationResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateConversationResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateConversationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateConversationResult) Unmarshal(in []byte) error {
	msg := new(chat.CreateConversationResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateConversationResult) GetSuccess() *chat.CreateConversationResponse {
	if !p.IsSetSuccess() {
		return CreateConversationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateConversationResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.CreateConversationResponse)
}

func (p *CreateConversationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateConversationResult) GetResult() interface{} {
	return p.Success
}

func listConversationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chat.ListConversationsRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chat.ChatService).ListConversations(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListConversationsArgs:
		success, err := handler.(chat.ChatService).ListConversations(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListConversationsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListConversationsArgs() interface{} {
	return &ListConversationsArgs{}
}

func newListConversationsResult() interface{} {
	return &ListConversationsResult{}
}

type ListConversationsArgs struct {
	Req *chat.ListConversationsRequest
}

func (p *ListConversationsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chat.ListConversationsRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListConversationsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListConversationsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListConversationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListConversationsArgs) Unmarshal(in []byte) error {
	msg := new(chat.ListConversationsRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListConversationsArgs_Req_DEFAULT *chat.ListConversationsRequest

func (p *ListConversationsArgs) GetReq() *chat.ListConversationsRequest {
	if !p.IsSetReq() {
		return ListConversationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListConversationsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListConversationsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListConversationsResult struct {
	Success *chat.ListConversationsResponse
}

var ListConversationsResult_Success_DEFAULT *chat.ListConversationsResponse

func (p *ListConversationsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chat.ListConversationsResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListConversationsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListConversationsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListConversationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListConversationsResult) Unmarshal(in []byte) error {
	msg := new(chat.ListConversationsResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListConversationsResult) GetSuccess() *chat.ListConversationsResponse {
	if !p.IsSetSuccess() {
		return ListConversationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListConversationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.ListConversationsResponse)
}

func (p *ListConversationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListConversationsResult) GetResult() interface{} {
	return p.Success
}

func sendMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	streamingArgs, ok := arg.(*streaming.Args)
	if !ok {
		return errInvalidMessageType
	}
	st := streamingArgs.Stream
	stream := &chatServiceSendMessageServer{st}
	req := new(chat.SendMessageRequest)
	if err := st.RecvMsg(req); err != nil {
		return err
	}
	return handler.(chat.ChatService).SendMessage(req, stream)
}

type chatServiceSendMessageClient struct {
	streaming.Stream
}

func (x *chatServiceSendMessageClient) DoFinish(err error) {
	if finisher, ok := x.Stream.(streaming.WithDoFinish); ok {
		finisher.DoFinish(err)
	} else {
		panic(fmt.Sprintf("streaming.WithDoFinish is not implemented by %T", x.Stream))
	}
}
func (x *chatServiceSendMessageClient) Recv() (*chat.SendMessageResponse, error) {
	m := new(chat.SendMessageResponse)
	return m, x.Stream.RecvMsg(m)
}

type chatServiceSendMessageServer struct {
	streaming.Stream
}

func (x *chatServiceSendMessageServer) Send(m *chat.SendMessageResponse) error {
	return x.Stream.SendMsg(m)
}

func newSendMessageArgs() interface{} {
	return &SendMessageArgs{}
}

func newSendMessageResult() interface{} {
	return &SendMessageResult{}
}

type SendMessageArgs struct {
	Req *chat.SendMessageRequest
}

func (p *SendMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chat.SendMessageRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SendMessageArgs) Unmarshal(in []byte) error {
	msg := new(chat.SendMessageRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendMessageArgs_Req_DEFAULT *chat.SendMessageRequest

func (p *SendMessageArgs) GetReq() *chat.SendMessageRequest {
	if !p.IsSetReq() {
		return SendMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendMessageArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendMessageResult struct {
	Success *chat.SendMessageResponse
}

var SendMessageResult_Success_DEFAULT *chat.SendMessageResponse

func (p *SendMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chat.SendMessageResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SendMessageResult) Unmarshal(in []byte) error {
	msg := new(chat.SendMessageResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendMessageResult) GetSuccess() *chat.SendMessageResponse {
	if !p.IsSetSuccess() {
		return SendMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.SendMessageResponse)
}

func (p *SendMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendMessageResult) GetResult() interface{} {
	return p.Success
}

func listMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chat.ListMessagesRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chat.ChatService).ListMessages(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListMessagesArgs:
		success, err := handler.(chat.ChatService).ListMessages(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListMessagesResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListMessagesArgs() interface{} {
	return &ListMessagesArgs{}
}

func newListMessagesResult() interface{} {
	return &ListMessagesResult{}
}

type ListMessagesArgs struct {
	Req *chat.ListMessagesRequest
}

func (p *ListMessagesArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chat.ListMessagesRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListMessagesArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListMessagesArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListMessagesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListMessagesArgs) Unmarshal(in []byte) error {
	msg := new(chat.ListMessagesRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListMessagesArgs_Req_DEFAULT *chat.ListMessagesRequest

func (p *ListMessagesArgs) GetReq() *chat.ListMessagesRequest {
	if !p.IsSetReq() {
		return ListMessagesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListMessagesArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListMessagesArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListMessagesResult struct {
	Success *chat.ListMessagesResponse
}

var ListMessagesResult_Success_DEFAULT *chat.ListMessagesResponse

func (p *ListMessagesResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chat.ListMessagesResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListMessagesResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListMessagesResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListMessagesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListMessagesResult) Unmarshal(in []byte) error {
	msg := new(chat.ListMessagesResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListMessagesResult) GetSuccess() *chat.ListMessagesResponse {
	if !p.IsSetSuccess() {
		return ListMessagesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListMessagesResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.ListMessagesResponse)
}

func (p *ListMessagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListMessagesResult) GetResult() interface{} {
	return p.Success
}

func deleteConversationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chat.DeleteConversationRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chat.ChatService).DeleteConversation(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteConversationArgs:
		success, err := handler.(chat.ChatService).DeleteConversation(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteConversationResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteConversationArgs() interface{} {
	return &DeleteConversationArgs{}
}

func newDeleteConversationResult() interface{} {
	return &DeleteConversationResult{}
}

type DeleteConversationArgs struct {
	Req *chat.DeleteConversationRequest
}

func (p *DeleteConversationArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chat.DeleteConversationRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteConversationArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteConversationArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteConversationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteConversationArgs) Unmarshal(in []byte) error {
	msg := new(chat.DeleteConversationRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteConversationArgs_Req_DEFAULT *chat.DeleteConversationRequest

func (p *DeleteConversationArgs) GetReq() *chat.DeleteConversationRequest {
	if !p.IsSetReq() {
		return DeleteConversationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteConversationArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteConversationArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteConversationResult struct {
	Success *chat.DeleteConversationResponse
}

var DeleteConversationResult_Success_DEFAULT *chat.DeleteConversationResponse

func (p *DeleteConversationResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chat.DeleteConversationResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteConversationResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteConversationResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteConversationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteConversationResult) Unmarshal(in []byte) error {
	msg := new(chat.DeleteConversationResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteConversationResult) GetSuccess() *chat.DeleteConversationResponse {
	if !p.IsSetSuccess() {
		return DeleteConversationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteConversationResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.DeleteConversationResponse)
}

func (p *DeleteConversationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteConversationResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateConversation(ctx context.Context, Req *chat.CreateConversationRequest) (r *chat.CreateConversationResponse, err error) {
	var _args CreateConversationArgs
	_args.Req = Req
	var _result CreateConversationResult
	if err = p.c.Call(ctx, "CreateConversation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListConversations(ctx context.Context, Req *chat.ListConversationsRequest) (r *chat.ListConversationsResponse, err error) {
	var _args ListConversationsArgs
	_args.Req = Req
	var _result ListConversationsResult
	if err = p.c.Call(ctx, "ListConversations", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (ChatService_SendMessageClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "SendMessage", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &chatServiceSendMessageClient{res.Stream}

	if err := stream.Stream.SendMsg(req); err != nil {
		return nil, err
	}
	if err := stream.Stream.Close(); err != nil {
		return nil, err
	}
	return stream, nil
}

func (p *kClient) ListMessages(ctx context.Context, Req *chat.ListMessagesRequest) (r *chat.ListMessagesResponse, err error) {
	var _args ListMessagesArgs
	_args.Req = Req
	var _result ListMessagesResult
	if err = p.c.Call(ctx, "ListMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteConversation(ctx context.Context, Req *chat.DeleteConversationRequest) (r *chat.DeleteConversationResponse, err error) {
	var _args DeleteConversationArgs
	_args.Req = Req
	var _result DeleteConversationResult
	if err = p.c.Call(ctx, "DeleteConversation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
