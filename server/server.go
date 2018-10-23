package server

import (
	"github.com/chenjun-git/umbrella-email/common"
	"github.com/chenjun-git/umbrella-email/directmail"
)

type Server struct {
	DirectMail *directmail.DirectMailApp
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) PatchDirectMail(config common.DirectMailConfig) {
	directmail := directmail.NewDirectMailApp(
		config.Format,
		config.Version,
		config.AccessKeyId,
		config.SignatureMethod,
		config.SignatureVersion,
		config.RegionId,
		config.Timeout.D(),
	)
	directmail.SetHttpClientTimeout()
	s.DirectMail = directmail
}
