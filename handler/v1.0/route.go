package v1_0

import (
	"github.com/go-chi/chi"
	"google.golang.org/grpc"

	"github.com/chenjun-git/umbrella-common/monitor"
)

func RegisterRouter(r chi.Router) {
	registerRouter("v1.0", r)
}

func registerRouter(version string, r chi.Router) {
	r.Route("/api/v1.0/email", func(r chi.Router) {
		r.Post("/sign_up", monitor.HttpHandlerWrapper("SendSignupVerifyCode", sendSignupVerifyCode))
	})
}

func wrapCode(v interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    0,
		"message": "ok",
		"body":    v,
	}
}

func handlerDial(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	// ui := traceconf.GlobalTracingConfig.GrpcUnaryClientInterceptor()
	//opts = append(opts, grpc.WithUnaryInterceptor(ui))

	return grpc.Dial(target, opts...)
}
