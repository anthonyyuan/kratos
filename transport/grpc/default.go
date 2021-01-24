package grpc

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DefaultErrorEncoder is default error encoder.
func DefaultErrorEncoder(err error) error {
	se, ok := err.(*errors.StatusError)
	if !ok {
		se = &errors.StatusError{
			Code:    2,
			Message: "Unknown: " + err.Error(),
		}
	}
	gs := status.Newf(codes.Code(se.Code), "%s: %s", se.Reason, se.Message)
	details := []proto.Message{
		&errdetails.ErrorInfo{
			Reason:   se.Reason,
			Metadata: map[string]string{"message": se.Message},
		},
	}
	gs, err = gs.WithDetails(details...)
	if err != nil {
		return err
	}
	return gs.Err()
}

// DefaultErrorDecoder is default error decoder.
func DefaultErrorDecoder(err error) error {
	gs := status.Convert(err)
	var (
		reason  string
		message string
	)
	for _, detail := range gs.Details() {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			reason = d.Reason
			message = d.Metadata["message"]
		}
	}
	return &errors.StatusError{
		Code:    int32(gs.Code()),
		Reason:  reason,
		Message: message,
	}
}
