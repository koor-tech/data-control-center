// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/services/ceph/v1/users.proto

package cephv1

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

// Validate checks the field values on ListCephUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCephUsersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCephUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCephUsersRequestMultiError, or nil if none found.
func (m *ListCephUsersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCephUsersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListCephUsersRequestMultiError(errors)
	}

	return nil
}

// ListCephUsersRequestMultiError is an error wrapping multiple validation
// errors returned by ListCephUsersRequest.ValidateAll() if the designated
// constraints aren't met.
type ListCephUsersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCephUsersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCephUsersRequestMultiError) AllErrors() []error { return m }

// ListCephUsersRequestValidationError is the validation error returned by
// ListCephUsersRequest.Validate if the designated constraints aren't met.
type ListCephUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCephUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCephUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCephUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCephUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCephUsersRequestValidationError) ErrorName() string {
	return "ListCephUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCephUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCephUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCephUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCephUsersRequestValidationError{}

// Validate checks the field values on ListCephUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCephUsersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCephUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCephUsersResponseMultiError, or nil if none found.
func (m *ListCephUsersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCephUsersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCephUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCephUsersResponseValidationError{
						field:  fmt.Sprintf("CephUsers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCephUsersResponseValidationError{
						field:  fmt.Sprintf("CephUsers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCephUsersResponseValidationError{
					field:  fmt.Sprintf("CephUsers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListCephUsersResponseMultiError(errors)
	}

	return nil
}

// ListCephUsersResponseMultiError is an error wrapping multiple validation
// errors returned by ListCephUsersResponse.ValidateAll() if the designated
// constraints aren't met.
type ListCephUsersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCephUsersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCephUsersResponseMultiError) AllErrors() []error { return m }

// ListCephUsersResponseValidationError is the validation error returned by
// ListCephUsersResponse.Validate if the designated constraints aren't met.
type ListCephUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCephUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCephUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCephUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCephUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCephUsersResponseValidationError) ErrorName() string {
	return "ListCephUsersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCephUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCephUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCephUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCephUsersResponseValidationError{}

// Validate checks the field values on CreateCephUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCephUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCephUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCephUserRequestMultiError, or nil if none found.
func (m *CreateCephUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCephUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCephUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCephUserRequestValidationError{
					field:  "CephUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCephUserRequestValidationError{
					field:  "CephUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCephUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCephUserRequestValidationError{
				field:  "CephUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCephUserRequestMultiError(errors)
	}

	return nil
}

// CreateCephUserRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCephUserRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCephUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCephUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCephUserRequestMultiError) AllErrors() []error { return m }

// CreateCephUserRequestValidationError is the validation error returned by
// CreateCephUserRequest.Validate if the designated constraints aren't met.
type CreateCephUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCephUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCephUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCephUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCephUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCephUserRequestValidationError) ErrorName() string {
	return "CreateCephUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCephUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCephUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCephUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCephUserRequestValidationError{}

// Validate checks the field values on CreateCephUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCephUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCephUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCephUserResponseMultiError, or nil if none found.
func (m *CreateCephUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCephUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCephUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCephUserResponseValidationError{
					field:  "CephUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCephUserResponseValidationError{
					field:  "CephUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCephUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCephUserResponseValidationError{
				field:  "CephUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCephUserResponseMultiError(errors)
	}

	return nil
}

// CreateCephUserResponseMultiError is an error wrapping multiple validation
// errors returned by CreateCephUserResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateCephUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCephUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCephUserResponseMultiError) AllErrors() []error { return m }

// CreateCephUserResponseValidationError is the validation error returned by
// CreateCephUserResponse.Validate if the designated constraints aren't met.
type CreateCephUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCephUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCephUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCephUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCephUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCephUserResponseValidationError) ErrorName() string {
	return "CreateCephUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCephUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCephUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCephUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCephUserResponseValidationError{}

// Validate checks the field values on DeleteCephUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCephUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCephUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCephUserRequestMultiError, or nil if none found.
func (m *DeleteCephUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCephUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	if len(errors) > 0 {
		return DeleteCephUserRequestMultiError(errors)
	}

	return nil
}

// DeleteCephUserRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteCephUserRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteCephUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCephUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCephUserRequestMultiError) AllErrors() []error { return m }

// DeleteCephUserRequestValidationError is the validation error returned by
// DeleteCephUserRequest.Validate if the designated constraints aren't met.
type DeleteCephUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCephUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCephUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCephUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCephUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCephUserRequestValidationError) ErrorName() string {
	return "DeleteCephUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCephUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCephUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCephUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCephUserRequestValidationError{}

// Validate checks the field values on DeleteCephUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCephUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCephUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCephUserResponseMultiError, or nil if none found.
func (m *DeleteCephUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCephUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return DeleteCephUserResponseMultiError(errors)
	}

	return nil
}

// DeleteCephUserResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteCephUserResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteCephUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCephUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCephUserResponseMultiError) AllErrors() []error { return m }

// DeleteCephUserResponseValidationError is the validation error returned by
// DeleteCephUserResponse.Validate if the designated constraints aren't met.
type DeleteCephUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCephUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCephUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCephUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCephUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCephUserResponseValidationError) ErrorName() string {
	return "DeleteCephUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCephUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCephUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCephUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCephUserResponseValidationError{}
