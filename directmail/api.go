package directmail

import (
	"encoding/json"
	"net/url"
	"fmt"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

type DirectMailReq struct {
	Format           string
	Version          string
	AccessKeyId      string
	Signature        string
	SignatureMethod  string
	Timestamp        string
	SignatureVersion string
	SignatureNonce   string
	RegionId         string
}

func (r *DirectMailReq) setAppInfo(format, version, accessKeyId, signatureMethod, signatureVersion, regionId string) error {
	r.Format = format
	r.Version = version
	r.AccessKeyId = accessKeyId
	r.SignatureMethod = signatureMethod
	r.SignatureVersion = signatureVersion
	r.Timestamp = url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	sn, err := uuid.NewV4()
	if err != nil {
		return err
	}
	r.SignatureNonce = sn.String()
	r.RegionId = regionId

	return nil
}

type DirectMailResp struct {
	RequestId string
	Code 	  string
	HostId    string
	Message   string
}

type EmailSendSingleReq struct {
	DirectMailReq
	Action         string
	AccountName    string
	ReplyToAddress string
	AddressType    string
	ToAddress      string
	FromAlias      string
	Subject        string
	HtmlBody       string
	TextBody       string
	ClickTrace     string
}

type EmailSendSingleResp struct {
	DirectMailResp
}

type EmailSendGroupReq struct {
	DirectMailReq
	Action        string
	AccountName   string
	AddressType   string
	TemplateName  string
	ReceiversName string
	TagName       string
	ClickTrace    string
}

type EmailSendGroupResp struct {
	DirectMailResp
}

func (dm *DirectMailApp) SendSingleVerifyCode(req EmailSendSingleReq) *EmailSendSingleResp {
	req.setAppInfo(dm.Format, dm.Version, dm.AccessKeyId, dm.SignatureMethod, dm.SignatureVersion, dm.RegionId)
	req.Signature = signEmailSendSingleReq(req)

	emailSingleSendParamMap := map[string]string{
		FormatKey:           req.Format,
		VersionKey:          req.Version,
		AccessKeyIdKey:      req.AccessKeyId,
		SignatureMethodKey:  req.SignatureMethod,
		TimestampKey:        req.Timestamp,
		SignatureVersionKey: req.SignatureVersion,
		SignatureNonceKey:   req.SignatureNonce,
		RegionIdKey:         req.RegionId,
		ActionKey:           req.Action,
		AccountNameKey:		 req.AccountName,
		ReplyToAddressKey:   req.ReplyToAddress,
		AddressTypeKey:      req.AddressType,
		ToAddressKey:        req.ToAddress,
		FromAliasKey:        req.FromAlias,
		SubjectKey:          req.Subject,
		HtmlBodyKey:         req.HtmlBody,
		TextBodyKey:         req.TextBody,
		ClickTraceKey:       req.ClickTrace,
		SignatureKey:        req.Signature,
	}

	allData := make([]string, 0)
	for _, v := range emailSingleSendParamKey {
		allData = append(allData, fmt.Sprintf("\"%s\":\"%s\"", v, emailSingleSendParamMap[v]))
	}
	body := "{"
	body += strings.Join(allData, ",")
	body += "}"
	fmt.Printf("%v\n", body)
	resp, err := httpReqWithParams("POST", DIRECT_MAIL_API_HOST, body)
	if err != nil {
		fmt.Printf("httpReqWithParams host: %v, err: %v, req: %v, resp: %v\n", DIRECT_MAIL_API_HOST, err, req, resp)
		return &EmailSendSingleResp{
			DirectMailResp: DirectMailResp{},
		}
	}

	var result EmailSendSingleResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		fmt.Printf("Unmarshal err: %v\n", err)
		return &EmailSendSingleResp{
			DirectMailResp: DirectMailResp{},
		}
	}

	return &result
}
