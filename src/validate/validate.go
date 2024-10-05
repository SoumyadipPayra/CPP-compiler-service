package validate

import (
	cpb "github.com/SoumyadipPayra/protobufs/go-protos/cpp_compiler"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func PingPong(req *cpb.PingRequest) error {
	return validation.ValidateStruct(
		validation.Field(&req.Msg, validation.Required),
	)
}

func CompileAndRun(req *cpb.CompileAndRunRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.UserName, validation.Required, validation.Length(1, 0)),
		validation.Field(&req.FilePath, validation.Required, validation.Length(1, 0)),
		validation.Field(&req.FileData, validation.Required),
	)
}
