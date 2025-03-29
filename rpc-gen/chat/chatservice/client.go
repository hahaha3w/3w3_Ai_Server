// Code generated by Kitex v0.12.1. DO NOT EDIT.

package chatservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/transport"
	chat "github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
	"github.com/cloudwego/kitex/client/streamclient"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateConversation(ctx context.Context, Req *chat.CreateConversationRequest, callOptions ...callopt.Option) (r *chat.CreateConversationResponse, err error)
	ListConversations(ctx context.Context, Req *chat.ListConversationsRequest, callOptions ...callopt.Option) (r *chat.ListConversationsResponse, err error)
	SendMessage(ctx context.Context, Req *chat.SendMessageRequest, callOptions ...callopt.Option) (stream ChatService_SendMessageClient, err error)
	ListMessages(ctx context.Context, Req *chat.ListMessagesRequest, callOptions ...callopt.Option) (r *chat.ListMessagesResponse, err error)
	DeleteConversation(ctx context.Context, Req *chat.DeleteConversationRequest, callOptions ...callopt.Option) (r *chat.DeleteConversationResponse, err error)
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	SendMessage(ctx context.Context, Req *chat.SendMessageRequest, callOptions ...streamcall.Option) (stream ChatService_SendMessageClient, err error)
}

type ChatService_SendMessageClient interface {
	streaming.Stream
	Recv() (*chat.SendMessageResponse, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kChatServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kChatServiceClient struct {
	*kClient
}

func (p *kChatServiceClient) CreateConversation(ctx context.Context, Req *chat.CreateConversationRequest, callOptions ...callopt.Option) (r *chat.CreateConversationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateConversation(ctx, Req)
}

func (p *kChatServiceClient) ListConversations(ctx context.Context, Req *chat.ListConversationsRequest, callOptions ...callopt.Option) (r *chat.ListConversationsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListConversations(ctx, Req)
}

func (p *kChatServiceClient) SendMessage(ctx context.Context, Req *chat.SendMessageRequest, callOptions ...callopt.Option) (stream ChatService_SendMessageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendMessage(ctx, Req)
}

func (p *kChatServiceClient) ListMessages(ctx context.Context, Req *chat.ListMessagesRequest, callOptions ...callopt.Option) (r *chat.ListMessagesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMessages(ctx, Req)
}

func (p *kChatServiceClient) DeleteConversation(ctx context.Context, Req *chat.DeleteConversationRequest, callOptions ...callopt.Option) (r *chat.DeleteConversationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteConversation(ctx, Req)
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kChatServiceStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kChatServiceStreamClient struct {
	*kClient
}

func (p *kChatServiceStreamClient) SendMessage(ctx context.Context, Req *chat.SendMessageRequest, callOptions ...streamcall.Option) (stream ChatService_SendMessageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.SendMessage(ctx, Req)
}
