// Code generated by Kitex v0.12.1. DO NOT EDIT.

package activityservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserActivities(ctx context.Context, Req *activity.GetUserActivityReq, callOptions ...callopt.Option) (r *activity.GetUserActivityResp, err error)
	CreateUserActivity(ctx context.Context, Req *activity.CreateUserActivityReq, callOptions ...callopt.Option) (r *activity.CreateUserActivityResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kActivityServiceClient{
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

type kActivityServiceClient struct {
	*kClient
}

func (p *kActivityServiceClient) GetUserActivities(ctx context.Context, Req *activity.GetUserActivityReq, callOptions ...callopt.Option) (r *activity.GetUserActivityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserActivities(ctx, Req)
}

func (p *kActivityServiceClient) CreateUserActivity(ctx context.Context, Req *activity.CreateUserActivityReq, callOptions ...callopt.Option) (r *activity.CreateUserActivityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUserActivity(ctx, Req)
}
