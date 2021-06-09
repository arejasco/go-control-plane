// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v2/service.proto

package envoy_config_trace_v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on TraceServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TraceServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGrpcService() == nil {
		return TraceServiceConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TraceServiceConfigValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TraceServiceConfigValidationError is the validation error returned by
// TraceServiceConfig.Validate if the designated constraints aren't met.
type TraceServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraceServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraceServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraceServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraceServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraceServiceConfigValidationError) ErrorName() string {
	return "TraceServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e TraceServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTraceServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraceServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraceServiceConfigValidationError{}
