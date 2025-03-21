// Code generated by Kitex v0.12.1. DO NOT EDIT.

package memoirservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	memoir "github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GenerateDailyMemoir": kitex.NewMethodInfo(
		generateDailyMemoirHandler,
		newGenerateDailyMemoirArgs,
		newGenerateDailyMemoirResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GenerateWeeklyMemoir": kitex.NewMethodInfo(
		generateWeeklyMemoirHandler,
		newGenerateWeeklyMemoirArgs,
		newGenerateWeeklyMemoirResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GenerateMonthlyMemoir": kitex.NewMethodInfo(
		generateMonthlyMemoirHandler,
		newGenerateMonthlyMemoirArgs,
		newGenerateMonthlyMemoirResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetMemoirList": kitex.NewMethodInfo(
		getMemoirListHandler,
		newGetMemoirListArgs,
		newGetMemoirListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetMemoirDetail": kitex.NewMethodInfo(
		getMemoirDetailHandler,
		newGetMemoirDetailArgs,
		newGetMemoirDetailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	memoirServiceServiceInfo                = NewServiceInfo()
	memoirServiceServiceInfoForClient       = NewServiceInfoForClient()
	memoirServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return memoirServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return memoirServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return memoirServiceServiceInfoForClient
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
	serviceName := "MemoirService"
	handlerType := (*memoir.MemoirService)(nil)
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
		"PackageName": "memoir",
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

func generateDailyMemoirHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(memoir.GenerateDailyMemoirRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(memoir.MemoirService).GenerateDailyMemoir(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GenerateDailyMemoirArgs:
		success, err := handler.(memoir.MemoirService).GenerateDailyMemoir(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GenerateDailyMemoirResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGenerateDailyMemoirArgs() interface{} {
	return &GenerateDailyMemoirArgs{}
}

func newGenerateDailyMemoirResult() interface{} {
	return &GenerateDailyMemoirResult{}
}

type GenerateDailyMemoirArgs struct {
	Req *memoir.GenerateDailyMemoirRequest
}

func (p *GenerateDailyMemoirArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(memoir.GenerateDailyMemoirRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GenerateDailyMemoirArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GenerateDailyMemoirArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GenerateDailyMemoirArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GenerateDailyMemoirArgs) Unmarshal(in []byte) error {
	msg := new(memoir.GenerateDailyMemoirRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GenerateDailyMemoirArgs_Req_DEFAULT *memoir.GenerateDailyMemoirRequest

func (p *GenerateDailyMemoirArgs) GetReq() *memoir.GenerateDailyMemoirRequest {
	if !p.IsSetReq() {
		return GenerateDailyMemoirArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GenerateDailyMemoirArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GenerateDailyMemoirArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GenerateDailyMemoirResult struct {
	Success *memoir.GenerateMemoirResponse
}

var GenerateDailyMemoirResult_Success_DEFAULT *memoir.GenerateMemoirResponse

func (p *GenerateDailyMemoirResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(memoir.GenerateMemoirResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GenerateDailyMemoirResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GenerateDailyMemoirResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GenerateDailyMemoirResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GenerateDailyMemoirResult) Unmarshal(in []byte) error {
	msg := new(memoir.GenerateMemoirResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GenerateDailyMemoirResult) GetSuccess() *memoir.GenerateMemoirResponse {
	if !p.IsSetSuccess() {
		return GenerateDailyMemoirResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GenerateDailyMemoirResult) SetSuccess(x interface{}) {
	p.Success = x.(*memoir.GenerateMemoirResponse)
}

func (p *GenerateDailyMemoirResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GenerateDailyMemoirResult) GetResult() interface{} {
	return p.Success
}

func generateWeeklyMemoirHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(memoir.GenerateWeeklyMemoirRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(memoir.MemoirService).GenerateWeeklyMemoir(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GenerateWeeklyMemoirArgs:
		success, err := handler.(memoir.MemoirService).GenerateWeeklyMemoir(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GenerateWeeklyMemoirResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGenerateWeeklyMemoirArgs() interface{} {
	return &GenerateWeeklyMemoirArgs{}
}

func newGenerateWeeklyMemoirResult() interface{} {
	return &GenerateWeeklyMemoirResult{}
}

type GenerateWeeklyMemoirArgs struct {
	Req *memoir.GenerateWeeklyMemoirRequest
}

func (p *GenerateWeeklyMemoirArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(memoir.GenerateWeeklyMemoirRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GenerateWeeklyMemoirArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GenerateWeeklyMemoirArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GenerateWeeklyMemoirArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GenerateWeeklyMemoirArgs) Unmarshal(in []byte) error {
	msg := new(memoir.GenerateWeeklyMemoirRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GenerateWeeklyMemoirArgs_Req_DEFAULT *memoir.GenerateWeeklyMemoirRequest

func (p *GenerateWeeklyMemoirArgs) GetReq() *memoir.GenerateWeeklyMemoirRequest {
	if !p.IsSetReq() {
		return GenerateWeeklyMemoirArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GenerateWeeklyMemoirArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GenerateWeeklyMemoirArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GenerateWeeklyMemoirResult struct {
	Success *memoir.GenerateMemoirResponse
}

var GenerateWeeklyMemoirResult_Success_DEFAULT *memoir.GenerateMemoirResponse

func (p *GenerateWeeklyMemoirResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(memoir.GenerateMemoirResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GenerateWeeklyMemoirResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GenerateWeeklyMemoirResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GenerateWeeklyMemoirResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GenerateWeeklyMemoirResult) Unmarshal(in []byte) error {
	msg := new(memoir.GenerateMemoirResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GenerateWeeklyMemoirResult) GetSuccess() *memoir.GenerateMemoirResponse {
	if !p.IsSetSuccess() {
		return GenerateWeeklyMemoirResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GenerateWeeklyMemoirResult) SetSuccess(x interface{}) {
	p.Success = x.(*memoir.GenerateMemoirResponse)
}

func (p *GenerateWeeklyMemoirResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GenerateWeeklyMemoirResult) GetResult() interface{} {
	return p.Success
}

func generateMonthlyMemoirHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(memoir.GenerateMonthlyMemoirRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(memoir.MemoirService).GenerateMonthlyMemoir(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GenerateMonthlyMemoirArgs:
		success, err := handler.(memoir.MemoirService).GenerateMonthlyMemoir(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GenerateMonthlyMemoirResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGenerateMonthlyMemoirArgs() interface{} {
	return &GenerateMonthlyMemoirArgs{}
}

func newGenerateMonthlyMemoirResult() interface{} {
	return &GenerateMonthlyMemoirResult{}
}

type GenerateMonthlyMemoirArgs struct {
	Req *memoir.GenerateMonthlyMemoirRequest
}

func (p *GenerateMonthlyMemoirArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(memoir.GenerateMonthlyMemoirRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GenerateMonthlyMemoirArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GenerateMonthlyMemoirArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GenerateMonthlyMemoirArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GenerateMonthlyMemoirArgs) Unmarshal(in []byte) error {
	msg := new(memoir.GenerateMonthlyMemoirRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GenerateMonthlyMemoirArgs_Req_DEFAULT *memoir.GenerateMonthlyMemoirRequest

func (p *GenerateMonthlyMemoirArgs) GetReq() *memoir.GenerateMonthlyMemoirRequest {
	if !p.IsSetReq() {
		return GenerateMonthlyMemoirArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GenerateMonthlyMemoirArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GenerateMonthlyMemoirArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GenerateMonthlyMemoirResult struct {
	Success *memoir.GenerateMemoirResponse
}

var GenerateMonthlyMemoirResult_Success_DEFAULT *memoir.GenerateMemoirResponse

func (p *GenerateMonthlyMemoirResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(memoir.GenerateMemoirResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GenerateMonthlyMemoirResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GenerateMonthlyMemoirResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GenerateMonthlyMemoirResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GenerateMonthlyMemoirResult) Unmarshal(in []byte) error {
	msg := new(memoir.GenerateMemoirResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GenerateMonthlyMemoirResult) GetSuccess() *memoir.GenerateMemoirResponse {
	if !p.IsSetSuccess() {
		return GenerateMonthlyMemoirResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GenerateMonthlyMemoirResult) SetSuccess(x interface{}) {
	p.Success = x.(*memoir.GenerateMemoirResponse)
}

func (p *GenerateMonthlyMemoirResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GenerateMonthlyMemoirResult) GetResult() interface{} {
	return p.Success
}

func getMemoirListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(memoir.GetMemoirListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(memoir.MemoirService).GetMemoirList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetMemoirListArgs:
		success, err := handler.(memoir.MemoirService).GetMemoirList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetMemoirListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetMemoirListArgs() interface{} {
	return &GetMemoirListArgs{}
}

func newGetMemoirListResult() interface{} {
	return &GetMemoirListResult{}
}

type GetMemoirListArgs struct {
	Req *memoir.GetMemoirListRequest
}

func (p *GetMemoirListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(memoir.GetMemoirListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetMemoirListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetMemoirListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetMemoirListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetMemoirListArgs) Unmarshal(in []byte) error {
	msg := new(memoir.GetMemoirListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetMemoirListArgs_Req_DEFAULT *memoir.GetMemoirListRequest

func (p *GetMemoirListArgs) GetReq() *memoir.GetMemoirListRequest {
	if !p.IsSetReq() {
		return GetMemoirListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMemoirListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetMemoirListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetMemoirListResult struct {
	Success *memoir.GetMemoirListResponse
}

var GetMemoirListResult_Success_DEFAULT *memoir.GetMemoirListResponse

func (p *GetMemoirListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(memoir.GetMemoirListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetMemoirListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetMemoirListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetMemoirListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetMemoirListResult) Unmarshal(in []byte) error {
	msg := new(memoir.GetMemoirListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMemoirListResult) GetSuccess() *memoir.GetMemoirListResponse {
	if !p.IsSetSuccess() {
		return GetMemoirListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMemoirListResult) SetSuccess(x interface{}) {
	p.Success = x.(*memoir.GetMemoirListResponse)
}

func (p *GetMemoirListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMemoirListResult) GetResult() interface{} {
	return p.Success
}

func getMemoirDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(memoir.GetMemoirDetailRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(memoir.MemoirService).GetMemoirDetail(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetMemoirDetailArgs:
		success, err := handler.(memoir.MemoirService).GetMemoirDetail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetMemoirDetailResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetMemoirDetailArgs() interface{} {
	return &GetMemoirDetailArgs{}
}

func newGetMemoirDetailResult() interface{} {
	return &GetMemoirDetailResult{}
}

type GetMemoirDetailArgs struct {
	Req *memoir.GetMemoirDetailRequest
}

func (p *GetMemoirDetailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(memoir.GetMemoirDetailRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetMemoirDetailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetMemoirDetailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetMemoirDetailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetMemoirDetailArgs) Unmarshal(in []byte) error {
	msg := new(memoir.GetMemoirDetailRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetMemoirDetailArgs_Req_DEFAULT *memoir.GetMemoirDetailRequest

func (p *GetMemoirDetailArgs) GetReq() *memoir.GetMemoirDetailRequest {
	if !p.IsSetReq() {
		return GetMemoirDetailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMemoirDetailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetMemoirDetailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetMemoirDetailResult struct {
	Success *memoir.GetMemoirDetailResponse
}

var GetMemoirDetailResult_Success_DEFAULT *memoir.GetMemoirDetailResponse

func (p *GetMemoirDetailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(memoir.GetMemoirDetailResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetMemoirDetailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetMemoirDetailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetMemoirDetailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetMemoirDetailResult) Unmarshal(in []byte) error {
	msg := new(memoir.GetMemoirDetailResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMemoirDetailResult) GetSuccess() *memoir.GetMemoirDetailResponse {
	if !p.IsSetSuccess() {
		return GetMemoirDetailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMemoirDetailResult) SetSuccess(x interface{}) {
	p.Success = x.(*memoir.GetMemoirDetailResponse)
}

func (p *GetMemoirDetailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMemoirDetailResult) GetResult() interface{} {
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

func (p *kClient) GenerateDailyMemoir(ctx context.Context, Req *memoir.GenerateDailyMemoirRequest) (r *memoir.GenerateMemoirResponse, err error) {
	var _args GenerateDailyMemoirArgs
	_args.Req = Req
	var _result GenerateDailyMemoirResult
	if err = p.c.Call(ctx, "GenerateDailyMemoir", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GenerateWeeklyMemoir(ctx context.Context, Req *memoir.GenerateWeeklyMemoirRequest) (r *memoir.GenerateMemoirResponse, err error) {
	var _args GenerateWeeklyMemoirArgs
	_args.Req = Req
	var _result GenerateWeeklyMemoirResult
	if err = p.c.Call(ctx, "GenerateWeeklyMemoir", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GenerateMonthlyMemoir(ctx context.Context, Req *memoir.GenerateMonthlyMemoirRequest) (r *memoir.GenerateMemoirResponse, err error) {
	var _args GenerateMonthlyMemoirArgs
	_args.Req = Req
	var _result GenerateMonthlyMemoirResult
	if err = p.c.Call(ctx, "GenerateMonthlyMemoir", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMemoirList(ctx context.Context, Req *memoir.GetMemoirListRequest) (r *memoir.GetMemoirListResponse, err error) {
	var _args GetMemoirListArgs
	_args.Req = Req
	var _result GetMemoirListResult
	if err = p.c.Call(ctx, "GetMemoirList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMemoirDetail(ctx context.Context, Req *memoir.GetMemoirDetailRequest) (r *memoir.GetMemoirDetailResponse, err error) {
	var _args GetMemoirDetailArgs
	_args.Req = Req
	var _result GetMemoirDetailResult
	if err = p.c.Call(ctx, "GetMemoirDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
