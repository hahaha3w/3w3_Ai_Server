// Code generated by Kitex v0.12.3. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"SendCode": kitex.NewMethodInfo(
		sendCodeHandler,
		newSendCodeArgs,
		newSendCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newRegisterArgs,
		newRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newLoginArgs,
		newLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ChangePassword": kitex.NewMethodInfo(
		changePasswordHandler,
		newChangePasswordArgs,
		newChangePasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUserInfo": kitex.NewMethodInfo(
		updateUserInfoHandler,
		newUpdateUserInfoArgs,
		newUpdateUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newGetUserInfoArgs,
		newGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteUser": kitex.NewMethodInfo(
		deleteUserHandler,
		newDeleteUserArgs,
		newDeleteUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.3",
		Extra:           extra,
	}
	return svcInfo
}

func sendCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.SendCodeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).SendCode(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SendCodeArgs:
		success, err := handler.(user.UserService).SendCode(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendCodeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSendCodeArgs() interface{} {
	return &SendCodeArgs{}
}

func newSendCodeResult() interface{} {
	return &SendCodeResult{}
}

type SendCodeArgs struct {
	Req *user.SendCodeReq
}

func (p *SendCodeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.SendCodeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendCodeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendCodeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SendCodeArgs) Unmarshal(in []byte) error {
	msg := new(user.SendCodeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendCodeArgs_Req_DEFAULT *user.SendCodeReq

func (p *SendCodeArgs) GetReq() *user.SendCodeReq {
	if !p.IsSetReq() {
		return SendCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendCodeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendCodeResult struct {
	Success *user.SendCodeResp
}

var SendCodeResult_Success_DEFAULT *user.SendCodeResp

func (p *SendCodeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.SendCodeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendCodeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendCodeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SendCodeResult) Unmarshal(in []byte) error {
	msg := new(user.SendCodeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendCodeResult) GetSuccess() *user.SendCodeResp {
	if !p.IsSetSuccess() {
		return SendCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.SendCodeResp)
}

func (p *SendCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendCodeResult) GetResult() interface{} {
	return p.Success
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.RegisterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RegisterArgs:
		success, err := handler.(user.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.RegisterReq
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.RegisterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.RegisterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.RegisterReq

func (p *RegisterArgs) GetReq() *user.RegisterReq {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *user.RegisterResp
}

var RegisterResult_Success_DEFAULT *user.RegisterResp

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.RegisterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.RegisterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.RegisterResp {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.RegisterResp)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.LoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LoginArgs:
		success, err := handler.(user.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.LoginReq
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.LoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.LoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.LoginReq

func (p *LoginArgs) GetReq() *user.LoginReq {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *user.LoginResp
}

var LoginResult_Success_DEFAULT *user.LoginResp

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.LoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.LoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.LoginResp {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.LoginResp)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func changePasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.ChangePasswordReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).ChangePassword(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ChangePasswordArgs:
		success, err := handler.(user.UserService).ChangePassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChangePasswordResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newChangePasswordArgs() interface{} {
	return &ChangePasswordArgs{}
}

func newChangePasswordResult() interface{} {
	return &ChangePasswordResult{}
}

type ChangePasswordArgs struct {
	Req *user.ChangePasswordReq
}

func (p *ChangePasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.ChangePasswordReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChangePasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChangePasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChangePasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ChangePasswordArgs) Unmarshal(in []byte) error {
	msg := new(user.ChangePasswordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChangePasswordArgs_Req_DEFAULT *user.ChangePasswordReq

func (p *ChangePasswordArgs) GetReq() *user.ChangePasswordReq {
	if !p.IsSetReq() {
		return ChangePasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChangePasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChangePasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChangePasswordResult struct {
	Success *user.ChangePasswordResp
}

var ChangePasswordResult_Success_DEFAULT *user.ChangePasswordResp

func (p *ChangePasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.ChangePasswordResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChangePasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChangePasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChangePasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ChangePasswordResult) Unmarshal(in []byte) error {
	msg := new(user.ChangePasswordResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangePasswordResult) GetSuccess() *user.ChangePasswordResp {
	if !p.IsSetSuccess() {
		return ChangePasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChangePasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.ChangePasswordResp)
}

func (p *ChangePasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChangePasswordResult) GetResult() interface{} {
	return p.Success
}

func updateUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserInfoArgs:
		success, err := handler.(user.UserService).UpdateUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserInfoArgs() interface{} {
	return &UpdateUserInfoArgs{}
}

func newUpdateUserInfoResult() interface{} {
	return &UpdateUserInfoResult{}
}

type UpdateUserInfoArgs struct {
	Req *user.UpdateUserInfoReq
}

func (p *UpdateUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserInfoArgs_Req_DEFAULT *user.UpdateUserInfoReq

func (p *UpdateUserInfoArgs) GetReq() *user.UpdateUserInfoReq {
	if !p.IsSetReq() {
		return UpdateUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserInfoResult struct {
	Success *user.UpdateUserInfoResp
}

var UpdateUserInfoResult_Success_DEFAULT *user.UpdateUserInfoResp

func (p *UpdateUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserInfoResult) GetSuccess() *user.UpdateUserInfoResp {
	if !p.IsSetSuccess() {
		return UpdateUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserInfoResp)
}

func (p *UpdateUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserInfoResult) GetResult() interface{} {
	return p.Success
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserInfoArgs:
		success, err := handler.(user.UserService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *user.GetUserInfoReq
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *user.GetUserInfoReq

func (p *GetUserInfoArgs) GetReq() *user.GetUserInfoReq {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoResult struct {
	Success *user.GetUserInfoResp
}

var GetUserInfoResult_Success_DEFAULT *user.GetUserInfoResp

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *user.GetUserInfoResp {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserInfoResp)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoResult) GetResult() interface{} {
	return p.Success
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DeleteUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).DeleteUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteUserArgs:
		success, err := handler.(user.UserService).DeleteUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteUserArgs() interface{} {
	return &DeleteUserArgs{}
}

func newDeleteUserResult() interface{} {
	return &DeleteUserResult{}
}

type DeleteUserArgs struct {
	Req *user.DeleteUserReq
}

func (p *DeleteUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DeleteUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteUserArgs) Unmarshal(in []byte) error {
	msg := new(user.DeleteUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteUserArgs_Req_DEFAULT *user.DeleteUserReq

func (p *DeleteUserArgs) GetReq() *user.DeleteUserReq {
	if !p.IsSetReq() {
		return DeleteUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteUserResult struct {
	Success *user.DeleteUserResp
}

var DeleteUserResult_Success_DEFAULT *user.DeleteUserResp

func (p *DeleteUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DeleteUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteUserResult) Unmarshal(in []byte) error {
	msg := new(user.DeleteUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserResult) GetSuccess() *user.DeleteUserResp {
	if !p.IsSetSuccess() {
		return DeleteUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DeleteUserResp)
}

func (p *DeleteUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUserResult) GetResult() interface{} {
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

func (p *kClient) SendCode(ctx context.Context, Req *user.SendCodeReq) (r *user.SendCodeResp, err error) {
	var _args SendCodeArgs
	_args.Req = Req
	var _result SendCodeResult
	if err = p.c.Call(ctx, "SendCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, Req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangePassword(ctx context.Context, Req *user.ChangePasswordReq) (r *user.ChangePasswordResp, err error) {
	var _args ChangePasswordArgs
	_args.Req = Req
	var _result ChangePasswordResult
	if err = p.c.Call(ctx, "ChangePassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq) (r *user.UpdateUserInfoResp, err error) {
	var _args UpdateUserInfoArgs
	_args.Req = Req
	var _result UpdateUserInfoResult
	if err = p.c.Call(ctx, "UpdateUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq) (r *user.GetUserInfoResp, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, Req *user.DeleteUserReq) (r *user.DeleteUserResp, err error) {
	var _args DeleteUserArgs
	_args.Req = Req
	var _result DeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
