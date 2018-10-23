package directmail

import (
	"time"
)

type DirectMailApp struct {
	Format           string
	Version          string
	AccessKeyId      string
	SignatureMethod  string
	SignatureVersion string
	RegionId         string
	Timeout          time.Duration
}

func NewDirectMailApp(format, version, accessKeyId, signatureMethod, signatureVersion, regionId string, timeout time.Duration) *DirectMailApp {
	return &DirectMailApp{
		Format:			  format,
		Version:          version,
		AccessKeyId:      accessKeyId,
		SignatureMethod:  signatureMethod,
		SignatureVersion: signatureVersion,
		RegionId:         regionId,
		Timeout:          timeout,
	}
}

func (dm *DirectMailApp) SetHttpClientTimeout() {
	if dm.Timeout != 0 {
		SetClientTimeout(dm.Timeout)
	}
}
