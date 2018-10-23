package common

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	commonErrors "github.com/chenjun-git/umbrella-common/errors"
)

func RPCError(e commonErrors.Error) error {
	return status.Error(codes.Code(e.GetCode()), e.GetMessage())
}

func NewRPCError(code int, message string) error {
	return status.Error(codes.Code(code), message)
}

func ConvertError(err error) commonErrors.Error {
	if err == nil {
		return commonErrors.NewError(OK, "")
	}

	if e, ok := err.(commonErrors.Error); ok {
		return e
	}

	return commonErrors.NewError(EmailConvertErr, fmt.Sprintf("cannot convert error(%v) to Error", err.Error()))
}

func ConvertRPCError(err error) commonErrors.Error {
	if err == nil {
		return commonErrors.NewError(OK, "")
	}

	stat, ok := status.FromError(err)
	if !ok {
		return commonErrors.NewError(EmailConvertErr, fmt.Sprintf("cannot convert rpc error(%v) to Error", err.Error()))
	}

	return commonErrors.NewError(int(stat.Code()), stat.Message())
}
