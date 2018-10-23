package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chenjun-git/umbrella-email/common"
	"github.com/chenjun-git/umbrella-email/directmail"
	"github.com/chenjun-git/umbrella-email/pb"
)

func (s *Server) SendSignupVerifyCode(ctx context.Context, req *pb.EmailSendSingleReq) (*pb.EmailSendSingleResp, error) {

	sendReq := directmail.EmailSendSingleReq{
		Action:         directmail.SingleSendMail,
		AccountName:    common.Config.DirectMail.AccountName,
		ReplyToAddress: common.Config.DirectMail.ReplyToAddress,
		AddressType:    common.Config.DirectMail.AddressType,
		ToAddress:      req.Email,
		FromAlias:      common.Config.DirectMail.FromAlias,
		Subject:        "",
		HtmlBody:       "",
		TextBody:       req.VerifyCode,
		ClickTrace:     common.Config.DirectMail.ClickTrace,
	}
	resp := s.DirectMail.SendSingleVerifyCode(sendReq)
	fmt.Printf("%v\n", resp)
	code, err := strconv.Atoi(resp.Code)
	if err != nil {
		return nil, common.NewRPCError(common.EmailServiceInternalErr, fmt.Sprintf("code: %v, msg: %v", resp.Code, resp.Message))
	}

	if code != directmail.DIRECT_MAIL_RESULT_CODE_OK {
		return nil, common.NewRPCError(common.EmailServiceInternalErr, fmt.Sprintf("code: %v, msg: %v", resp.Code, resp.Message))
	}

	return nil, nil
}
