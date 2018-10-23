package v1_0

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	commonErrors "github.com/chenjun-git/umbrella-common/errors"

	"github.com/chenjun-git/umbrella-email/common"
	"github.com/chenjun-git/umbrella-email/pb"
	"github.com/chenjun-git/umbrella-email/server"
	"github.com/chenjun-git/umbrella-email/utils"
)

func sendSignupVerifyCode(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.EmailGatewayRequestIOErr, err.Error()))
		return
	}
	var req pb.EmailSendSingleReq
	if err = json.Unmarshal(body, &req); err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.EmailGatewayJsonUnmarshalErr, err.Error()))
		return
	}

	s := server.NewServer()
	s.PatchDirectMail(*common.Config.DirectMail)

	resp, err := s.SendSignupVerifyCode(context.Background(), &req)
	if err != nil {
		utils.JSON(w, r, common.ConvertError(err))
		return
	}

	utils.JSON(w, r, wrapCode(resp))
}
