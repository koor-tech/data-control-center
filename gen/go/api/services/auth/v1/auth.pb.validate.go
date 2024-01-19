// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/services/auth/v1/auth.proto

package authv1

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

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUsername()); l < 3 || l > 24 {
		err := LoginRequestValidationError{
			field:  "Username",
			reason: "value length must be between 3 and 24 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_LoginRequest_Username_Pattern.MatchString(m.GetUsername()) {
		err := LoginRequestValidationError{
			field:  "Username",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9-_]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 6 {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetPassword()) > 70 {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be at most 70 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

var _LoginRequest_Username_Pattern = regexp.MustCompile("^[a-zA-Z0-9-_]+$")

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResponseMultiError, or
// nil if none found.
func (m *LoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if all {
		switch v := interface{}(m.GetExpires()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginResponseValidationError{
					field:  "Expires",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginResponseValidationError{
					field:  "Expires",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpires()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginResponseValidationError{
				field:  "Expires",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AccountId

	if len(errors) > 0 {
		return LoginResponseMultiError(errors)
	}

	return nil
}

// LoginResponseMultiError is an error wrapping multiple validation errors
// returned by LoginResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponseMultiError) AllErrors() []error { return m }

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on LogoutRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutRequestMultiError, or
// nil if none found.
func (m *LogoutRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LogoutRequestMultiError(errors)
	}

	return nil
}

// LogoutRequestMultiError is an error wrapping multiple validation errors
// returned by LogoutRequest.ValidateAll() if the designated constraints
// aren't met.
type LogoutRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutRequestMultiError) AllErrors() []error { return m }

// LogoutRequestValidationError is the validation error returned by
// LogoutRequest.Validate if the designated constraints aren't met.
type LogoutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutRequestValidationError) ErrorName() string { return "LogoutRequestValidationError" }

// Error satisfies the builtin error interface
func (e LogoutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutRequestValidationError{}

// Validate checks the field values on LogoutResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutResponseMultiError,
// or nil if none found.
func (m *LogoutResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if m.RedirectTo != nil {
		// no validation rules for RedirectTo
	}

	if len(errors) > 0 {
		return LogoutResponseMultiError(errors)
	}

	return nil
}

// LogoutResponseMultiError is an error wrapping multiple validation errors
// returned by LogoutResponse.ValidateAll() if the designated constraints
// aren't met.
type LogoutResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutResponseMultiError) AllErrors() []error { return m }

// LogoutResponseValidationError is the validation error returned by
// LogoutResponse.Validate if the designated constraints aren't met.
type LogoutResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutResponseValidationError) ErrorName() string { return "LogoutResponseValidationError" }

// Error satisfies the builtin error interface
func (e LogoutResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutResponseValidationError{}

// Validate checks the field values on CheckTokenRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CheckTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckTokenRequestMultiError, or nil if none found.
func (m *CheckTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return CheckTokenRequestMultiError(errors)
	}

	return nil
}

// CheckTokenRequestMultiError is an error wrapping multiple validation errors
// returned by CheckTokenRequest.ValidateAll() if the designated constraints
// aren't met.
type CheckTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckTokenRequestMultiError) AllErrors() []error { return m }

// CheckTokenRequestValidationError is the validation error returned by
// CheckTokenRequest.Validate if the designated constraints aren't met.
type CheckTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckTokenRequestValidationError) ErrorName() string {
	return "CheckTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckTokenRequestValidationError{}

// Validate checks the field values on CheckTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckTokenResponseMultiError, or nil if none found.
func (m *CheckTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return CheckTokenResponseMultiError(errors)
	}

	return nil
}

// CheckTokenResponseMultiError is an error wrapping multiple validation errors
// returned by CheckTokenResponse.ValidateAll() if the designated constraints
// aren't met.
type CheckTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckTokenResponseMultiError) AllErrors() []error { return m }

// CheckTokenResponseValidationError is the validation error returned by
// CheckTokenResponse.Validate if the designated constraints aren't met.
type CheckTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckTokenResponseValidationError) ErrorName() string {
	return "CheckTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckTokenResponseValidationError{}
