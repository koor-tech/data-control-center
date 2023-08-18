// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/services/stats/stats.proto

package stats

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on StreamRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamRequestMultiError, or
// nil if none found.
func (m *StreamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StreamRequestMultiError(errors)
	}

	return nil
}

// StreamRequestMultiError is an error wrapping multiple validation errors
// returned by StreamRequest.ValidateAll() if the designated constraints
// aren't met.
type StreamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamRequestMultiError) AllErrors() []error { return m }

// StreamRequestValidationError is the validation error returned by
// StreamRequest.Validate if the designated constraints aren't met.
type StreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamRequestValidationError) ErrorName() string { return "StreamRequestValidationError" }

// Error satisfies the builtin error interface
func (e StreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamRequestValidationError{}

// Validate checks the field values on StreamResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamResponseMultiError,
// or nil if none found.
func (m *StreamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Message.(type) {
	case *StreamResponse_Todo:
		if v == nil {
			err := StreamResponseValidationError{
				field:  "Message",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Todo
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return StreamResponseMultiError(errors)
	}

	return nil
}

// StreamResponseMultiError is an error wrapping multiple validation errors
// returned by StreamResponse.ValidateAll() if the designated constraints
// aren't met.
type StreamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamResponseMultiError) AllErrors() []error { return m }

// StreamResponseValidationError is the validation error returned by
// StreamResponse.Validate if the designated constraints aren't met.
type StreamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamResponseValidationError) ErrorName() string { return "StreamResponseValidationError" }

// Error satisfies the builtin error interface
func (e StreamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamResponseValidationError{}
