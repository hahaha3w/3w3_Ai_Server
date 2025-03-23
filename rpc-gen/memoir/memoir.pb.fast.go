// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package memoir

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Memoir) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Memoir[number], err)
}

func (x *Memoir) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Style, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.StartDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.EndDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Memoir) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.CreatedAt, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenerateMemoirRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 7:
		offset, err = x.fastReadField7(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GenerateMemoirRequest[number], err)
}

func (x *GenerateMemoirRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GenerateMemoirRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenerateMemoirRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenerateMemoirRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenerateMemoirRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Style, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenerateMemoirRequest) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.StartDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenerateMemoirRequest) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.EndDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenerateMemoirResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GenerateMemoirResponse[number], err)
}

func (x *GenerateMemoirResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GenerateMemoirResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Memoir
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Memoir = &v
	return offset, nil
}

func (x *GenerateMemoirResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ErrorMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMemoirListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 7:
		offset, err = x.fastReadField7(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMemoirListRequest[number], err)
}

func (x *GetMemoirListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetMemoirListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMemoirListRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Style, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMemoirListRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.StartDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMemoirListRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.EndDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMemoirListRequest) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Page, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetMemoirListRequest) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetMemoirListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMemoirListResponse[number], err)
}

func (x *GetMemoirListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Memoir
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Memoirs = append(x.Memoirs, &v)
	return offset, nil
}

func (x *GetMemoirListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetMemoirDetailRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMemoirDetailRequest[number], err)
}

func (x *GetMemoirDetailRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MemoirId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetMemoirDetailRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetMemoirDetailResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMemoirDetailResponse[number], err)
}

func (x *GetMemoirDetailResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Memoir
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Memoir = &v
	return offset, nil
}

func (x *DeleteMemoirRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteMemoirRequest[number], err)
}

func (x *DeleteMemoirRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MemoirId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DeleteMemoirRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DeleteMemoirResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteMemoirResponse[number], err)
}

func (x *DeleteMemoirResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *DeleteMemoirResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ErrorMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Memoir) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	return offset
}

func (x *Memoir) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Memoir) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *Memoir) fastWriteField3(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetTitle())
	return offset
}

func (x *Memoir) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetContent())
	return offset
}

func (x *Memoir) fastWriteField5(buf []byte) (offset int) {
	if x.Type == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetType())
	return offset
}

func (x *Memoir) fastWriteField6(buf []byte) (offset int) {
	if x.Style == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetStyle())
	return offset
}

func (x *Memoir) fastWriteField7(buf []byte) (offset int) {
	if x.StartDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetStartDate())
	return offset
}

func (x *Memoir) fastWriteField8(buf []byte) (offset int) {
	if x.EndDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetEndDate())
	return offset
}

func (x *Memoir) fastWriteField9(buf []byte) (offset int) {
	if x.CreatedAt == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 9, x.GetCreatedAt())
	return offset
}

func (x *GenerateMemoirRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *GenerateMemoirRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GenerateMemoirRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTitle())
	return offset
}

func (x *GenerateMemoirRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetContent())
	return offset
}

func (x *GenerateMemoirRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Type == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetType())
	return offset
}

func (x *GenerateMemoirRequest) fastWriteField5(buf []byte) (offset int) {
	if x.Style == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetStyle())
	return offset
}

func (x *GenerateMemoirRequest) fastWriteField6(buf []byte) (offset int) {
	if x.StartDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetStartDate())
	return offset
}

func (x *GenerateMemoirRequest) fastWriteField7(buf []byte) (offset int) {
	if x.EndDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetEndDate())
	return offset
}

func (x *GenerateMemoirResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GenerateMemoirResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *GenerateMemoirResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Memoir == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetMemoir())
	return offset
}

func (x *GenerateMemoirResponse) fastWriteField3(buf []byte) (offset int) {
	if x.ErrorMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetErrorMsg())
	return offset
}

func (x *GetMemoirListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *GetMemoirListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetMemoirListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Type == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetType())
	return offset
}

func (x *GetMemoirListRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Style == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetStyle())
	return offset
}

func (x *GetMemoirListRequest) fastWriteField4(buf []byte) (offset int) {
	if x.StartDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetStartDate())
	return offset
}

func (x *GetMemoirListRequest) fastWriteField5(buf []byte) (offset int) {
	if x.EndDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetEndDate())
	return offset
}

func (x *GetMemoirListRequest) fastWriteField6(buf []byte) (offset int) {
	if x.Page == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.GetPage())
	return offset
}

func (x *GetMemoirListRequest) fastWriteField7(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 7, x.GetPageSize())
	return offset
}

func (x *GetMemoirListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetMemoirListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Memoirs == nil {
		return offset
	}
	for i := range x.GetMemoirs() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetMemoirs()[i])
	}
	return offset
}

func (x *GetMemoirListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetMemoirDetailRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetMemoirDetailRequest) fastWriteField1(buf []byte) (offset int) {
	if x.MemoirId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetMemoirId())
	return offset
}

func (x *GetMemoirDetailRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *GetMemoirDetailResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetMemoirDetailResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Memoir == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetMemoir())
	return offset
}

func (x *DeleteMemoirRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteMemoirRequest) fastWriteField1(buf []byte) (offset int) {
	if x.MemoirId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetMemoirId())
	return offset
}

func (x *DeleteMemoirRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *DeleteMemoirResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteMemoirResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *DeleteMemoirResponse) fastWriteField2(buf []byte) (offset int) {
	if x.ErrorMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetErrorMsg())
	return offset
}

func (x *Memoir) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	return n
}

func (x *Memoir) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetId())
	return n
}

func (x *Memoir) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *Memoir) sizeField3() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetTitle())
	return n
}

func (x *Memoir) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetContent())
	return n
}

func (x *Memoir) sizeField5() (n int) {
	if x.Type == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetType())
	return n
}

func (x *Memoir) sizeField6() (n int) {
	if x.Style == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetStyle())
	return n
}

func (x *Memoir) sizeField7() (n int) {
	if x.StartDate == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetStartDate())
	return n
}

func (x *Memoir) sizeField8() (n int) {
	if x.EndDate == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetEndDate())
	return n
}

func (x *Memoir) sizeField9() (n int) {
	if x.CreatedAt == "" {
		return n
	}
	n += fastpb.SizeString(9, x.GetCreatedAt())
	return n
}

func (x *GenerateMemoirRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *GenerateMemoirRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *GenerateMemoirRequest) sizeField2() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTitle())
	return n
}

func (x *GenerateMemoirRequest) sizeField3() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetContent())
	return n
}

func (x *GenerateMemoirRequest) sizeField4() (n int) {
	if x.Type == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetType())
	return n
}

func (x *GenerateMemoirRequest) sizeField5() (n int) {
	if x.Style == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetStyle())
	return n
}

func (x *GenerateMemoirRequest) sizeField6() (n int) {
	if x.StartDate == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetStartDate())
	return n
}

func (x *GenerateMemoirRequest) sizeField7() (n int) {
	if x.EndDate == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetEndDate())
	return n
}

func (x *GenerateMemoirResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GenerateMemoirResponse) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *GenerateMemoirResponse) sizeField2() (n int) {
	if x.Memoir == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetMemoir())
	return n
}

func (x *GenerateMemoirResponse) sizeField3() (n int) {
	if x.ErrorMsg == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetErrorMsg())
	return n
}

func (x *GetMemoirListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *GetMemoirListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *GetMemoirListRequest) sizeField2() (n int) {
	if x.Type == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetType())
	return n
}

func (x *GetMemoirListRequest) sizeField3() (n int) {
	if x.Style == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetStyle())
	return n
}

func (x *GetMemoirListRequest) sizeField4() (n int) {
	if x.StartDate == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetStartDate())
	return n
}

func (x *GetMemoirListRequest) sizeField5() (n int) {
	if x.EndDate == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetEndDate())
	return n
}

func (x *GetMemoirListRequest) sizeField6() (n int) {
	if x.Page == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.GetPage())
	return n
}

func (x *GetMemoirListRequest) sizeField7() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt32(7, x.GetPageSize())
	return n
}

func (x *GetMemoirListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetMemoirListResponse) sizeField1() (n int) {
	if x.Memoirs == nil {
		return n
	}
	for i := range x.GetMemoirs() {
		n += fastpb.SizeMessage(1, x.GetMemoirs()[i])
	}
	return n
}

func (x *GetMemoirListResponse) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetTotal())
	return n
}

func (x *GetMemoirDetailRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetMemoirDetailRequest) sizeField1() (n int) {
	if x.MemoirId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetMemoirId())
	return n
}

func (x *GetMemoirDetailRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *GetMemoirDetailResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetMemoirDetailResponse) sizeField1() (n int) {
	if x.Memoir == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetMemoir())
	return n
}

func (x *DeleteMemoirRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteMemoirRequest) sizeField1() (n int) {
	if x.MemoirId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetMemoirId())
	return n
}

func (x *DeleteMemoirRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *DeleteMemoirResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteMemoirResponse) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *DeleteMemoirResponse) sizeField2() (n int) {
	if x.ErrorMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetErrorMsg())
	return n
}

var fieldIDToName_Memoir = map[int32]string{
	1: "Id",
	2: "UserId",
	3: "Title",
	4: "Content",
	5: "Type",
	6: "Style",
	7: "StartDate",
	8: "EndDate",
	9: "CreatedAt",
}

var fieldIDToName_GenerateMemoirRequest = map[int32]string{
	1: "UserId",
	2: "Title",
	3: "Content",
	4: "Type",
	5: "Style",
	6: "StartDate",
	7: "EndDate",
}

var fieldIDToName_GenerateMemoirResponse = map[int32]string{
	1: "Success",
	2: "Memoir",
	3: "ErrorMsg",
}

var fieldIDToName_GetMemoirListRequest = map[int32]string{
	1: "UserId",
	2: "Type",
	3: "Style",
	4: "StartDate",
	5: "EndDate",
	6: "Page",
	7: "PageSize",
}

var fieldIDToName_GetMemoirListResponse = map[int32]string{
	1: "Memoirs",
	2: "Total",
}

var fieldIDToName_GetMemoirDetailRequest = map[int32]string{
	1: "MemoirId",
	2: "UserId",
}

var fieldIDToName_GetMemoirDetailResponse = map[int32]string{
	1: "Memoir",
}

var fieldIDToName_DeleteMemoirRequest = map[int32]string{
	1: "MemoirId",
	2: "UserId",
}

var fieldIDToName_DeleteMemoirResponse = map[int32]string{
	1: "Success",
	2: "ErrorMsg",
}
