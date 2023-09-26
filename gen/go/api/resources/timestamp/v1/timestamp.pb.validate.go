// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/resources/timestamp/v1/timestamp.proto

package timestampv1

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

// Validate checks the field values on Timestamp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Timestamp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Timestamp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimestampMultiError, or nil
// if none found.
func (m *Timestamp) ValidateAll() error {
	return m.validate(true)
}

func (m *Timestamp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TimestampValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TimestampValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimestampValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TimestampMultiError(errors)
	}

	return nil
}

// TimestampMultiError is an error wrapping multiple validation errors returned
// by Timestamp.ValidateAll() if the designated constraints aren't met.
type TimestampMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimestampMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimestampMultiError) AllErrors() []error { return m }

// TimestampValidationError is the validation error returned by
// Timestamp.Validate if the designated constraints aren't met.
type TimestampValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimestampValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimestampValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimestampValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimestampValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimestampValidationError) ErrorName() string { return "TimestampValidationError" }

// Error satisfies the builtin error interface
func (e TimestampValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimestamp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimestampValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimestampValidationError{}