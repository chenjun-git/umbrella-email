package directmail

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/chenjun-git/umbrella-email/common"
)

var (
	testParamKey []string
	emailSingleSendParamKey []string
)

const (
	FormatKey           = "Format"
	VersionKey          = "Version"
	AccessKeyIdKey      = "AccessKeyId"
	SignatureKey        = "Signature"
	SignatureMethodKey  = "SignatureMethod"
	TimestampKey        = "Timestamp"
	SignatureVersionKey = "SignatureVersion"
	SignatureNonceKey   = "SignatureNonce"
	RegionIdKey         = "RegionId"
	ActionKey           = "Action"
	AccountNameKey      = "AccountName"
	ReplyToAddressKey   = "ReplyToAddress"
	AddressTypeKey      = "AddressType"
	ToAddressKey        = "ToAddress"
	FromAliasKey        = "FromAlias"
	SubjectKey          = "Subject"
	HtmlBodyKey         = "HtmlBody"
	TextBodyKey         = "TextBody"
	ClickTraceKey       = "ClickTrace"
	TagNameKey          = "TagName"

	//queryStringEmailSingleSend = "AccessKeyId=%s&Action=%s&AddressType=%s&ClickTrace=%s&Format=%s&FromAlias=%s&HtmlBody=%s&RegionId=%s&SignatureMethod=%s&SignatureNonce=%s&SignatureVersion=%s&Subject=%s&TextBody=%s&Timestamp=%s&ToAddress=%s&Version=%s"
)

func init() {
	testParamKey = []string{FormatKey, VersionKey, AccessKeyIdKey, SignatureKey, SignatureMethodKey, TimestampKey, SignatureVersionKey, SignatureNonceKey, RegionIdKey, ActionKey, AccountNameKey, ReplyToAddressKey, AddressTypeKey, ToAddressKey, SubjectKey, HtmlBodyKey, TagNameKey}
	emailSingleSendParamKey = []string{FormatKey, VersionKey, AccessKeyIdKey, SignatureKey, SignatureMethodKey, TimestampKey, SignatureVersionKey, SignatureNonceKey, RegionIdKey, ActionKey, AccountNameKey, ReplyToAddressKey, AddressTypeKey, ToAddressKey, FromAliasKey, SubjectKey, HtmlBodyKey, TextBodyKey, ClickTraceKey}
}

func encodeLocal(encodeStr string) string {
	urlEncode := url.QueryEscape(encodeStr)
	urlEncode = strings.Replace(urlEncode, "+", "%%20", -1)
	urlEncode = strings.Replace(urlEncode, "*", "%2A", -1)
	urlEncode = strings.Replace(urlEncode, "%%7E", "~", -1)
	urlEncode = strings.Replace(urlEncode, "/", "%%2F", -1)
	return urlEncode
}

func signSample() string {
	testParamMap := map[string]string{
		FormatKey:           "XML",
		VersionKey:          "2015-11-23",
		AccessKeyIdKey:      "testid",
		SignatureMethodKey:  "HMAC-SHA1",
		TimestampKey:        "2016-10-20T06:27:56Z",
		SignatureVersionKey: "1.0",
		SignatureNonceKey:   "c1b2c332-4cfb-4a0f-b8cc-ebe622aa0a5c",
		RegionIdKey:         "cn-hangzhou",
		ActionKey:           "SingleSendMail",
		AccountNameKey:		 "<a%b'",
		ReplyToAddressKey:   "true",
		AddressTypeKey:      "1",
		ToAddressKey:        "1@test.com",
		SubjectKey:          "3",
		HtmlBodyKey:         "4",
		TagNameKey:			 "2",
	}

	sort.Strings(testParamKey)

	queryStr := strings.Join(testParamKey, "=%s&")
	queryStr += "=%s"

	sortedData := make([]interface{}, 0)
	for _, d := range testParamKey {
		sortedData = append(sortedData, encodeLocal(testParamMap[d]))
	}

	sortQueryString := fmt.Sprintf(queryStr, sortedData...)
	sortQueryString=strings.Replace(sortQueryString, "=", "%3D", -1)

	//urlEncode := encodeLocal(sortQueryString)
	signStr := fmt.Sprintf("POST&%%2F&%s", sortQueryString)
	//POST&%2F&AccessKeyId%3Dtestid&AccountName%3D%253Ca%2525b%2527%253E&Action%3DSingleSendMail&AddressType%3D1&Format%3DXML&HtmlBody%3D4&RegionId%3Dcn-hangzhou&ReplyToAddress%3Dtrue&SignatureMethod%3DHMAC-SHA1&SignatureNonce%3Dc1b2c332-4cfb-4a0f-b8cc-ebe622aa0a5c&SignatureVersion%3D1.0&Subject%3D3&TagName%3D2&Timestamp%3D2016-10-20T06%253A27%253A56Z&ToAddress%3D1%2540test.com&Version%3D2015-11-23

	//signStr = "POST%26%2F%26AccessKeyId%3Dtestid%26AccountName%3D%253Ca%2525b%2527%253E%26Action%3DSingleSendMail%26AddressType%3D1%26Format%3DXML%26HtmlBody%3D4%26RegionId%3Dcn-hangzhou%26ReplyToAddress%3Dtrue%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3Dc1b2c332-4cfb-4a0f-b8cc-ebe622aa0a5c%26SignatureVersion%3D1.0%26Subject%3D3%26TagName%3D2%26Timestamp%3D2016-10-20T06%253A27%253A56Z%26ToAddress%3D1%2540test.com%26Version%3D2015-11-23"
	key := []byte("testsecret&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(signStr))
	signture := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	//signture = encodeLocal(signture)
	return signture
}

func signEmailSendSingleReq(req EmailSendSingleReq) string {
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
	}

	sort.Strings(emailSingleSendParamKey)

	queryStringEmailSingleSend := strings.Join(emailSingleSendParamKey, "=?&")
	queryStringEmailSingleSend += "=%s"

	sortedData := make([]interface{}, 0)
	for _, d := range emailSingleSendParamKey {
		sortedData = append(sortedData, emailSingleSendParamMap[d])
	}

	sortQueryString := fmt.Sprintf(queryStringEmailSingleSend, sortedData...)
	sortQueryString=strings.Replace(sortQueryString, "=", "%3D", -1)

	urlEncode := encodeLocal(sortQueryString)
	signStr := fmt.Sprintf("POST&%%2F&%s", urlEncode)

	key := []byte(common.Config.DirectMail.AccessKeySecret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(signStr))
	signture := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signture = encodeLocal(signture)
	return signture
}
