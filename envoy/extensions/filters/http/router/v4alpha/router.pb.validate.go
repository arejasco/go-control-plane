// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/router/v4alpha/router.proto

package envoy_extensions_filters_http_router_v4alpha

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

// Validate checks the field values on Router with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Router) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDynamicStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouterValidationError{
				field:  "DynamicStats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for StartChildSpan

	for idx, item := range m.GetUpstreamLog() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouterValidationError{
					field:  fmt.Sprintf("UpstreamLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for SuppressEnvoyHeaders

	for idx, item := range m.GetStrictCheckHeaders() {
		_, _ = idx, item

		if _, ok := _Router_StrictCheckHeaders_InLookup[item]; !ok {
			return RouterValidationError{
				field:  fmt.Sprintf("StrictCheckHeaders[%v]", idx),
				reason: "value must be in list [x-envoy-upstream-rq-timeout-ms x-envoy-upstream-rq-per-try-timeout-ms x-envoy-max-retries x-envoy-retry-grpc-on x-envoy-retry-on]",
			}
		}

	}

	// no validation rules for RespectExpectedRqTimeout

	return nil
}

// RouterValidationError is the validation error returned by Router.Validate if
// the designated constraints aren't met.
type RouterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouterValidationError) ErrorName() string { return "RouterValidationError" }

// Error satisfies the builtin error interface
func (e RouterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouterValidationError{}

var _Router_StrictCheckHeaders_InLookup = map[string]struct{}{
	"x-envoy-upstream-rq-timeout-ms":         {},
	"x-envoy-upstream-rq-per-try-timeout-ms": {},
	"x-envoy-max-retries":                    {},
	"x-envoy-retry-grpc-on":                  {},
	"x-envoy-retry-on":                       {},
}
