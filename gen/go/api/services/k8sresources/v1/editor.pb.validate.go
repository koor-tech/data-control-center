// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/services/k8sresources/v1/editor.proto

package k8sresourcesv1

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

// Validate checks the field values on GetResourcesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetResourcesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourcesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourcesResponseMultiError, or nil if none found.
func (m *GetResourcesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourcesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResources()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetResourcesResponseValidationError{
					field:  "Resources",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetResourcesResponseValidationError{
					field:  "Resources",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResources()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResourcesResponseValidationError{
				field:  "Resources",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetResourcesResponseMultiError(errors)
	}

	return nil
}

// GetResourcesResponseMultiError is an error wrapping multiple validation
// errors returned by GetResourcesResponse.ValidateAll() if the designated
// constraints aren't met.
type GetResourcesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourcesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourcesResponseMultiError) AllErrors() []error { return m }

// GetResourcesResponseValidationError is the validation error returned by
// GetResourcesResponse.Validate if the designated constraints aren't met.
type GetResourcesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourcesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourcesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourcesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourcesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourcesResponseValidationError) ErrorName() string {
	return "GetResourcesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetResourcesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourcesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourcesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourcesResponseValidationError{}

// Validate checks the field values on GetResourcesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetResourcesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourcesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourcesRequestMultiError, or nil if none found.
func (m *GetResourcesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourcesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetResourcesRequestMultiError(errors)
	}

	return nil
}

// GetResourcesRequestMultiError is an error wrapping multiple validation
// errors returned by GetResourcesRequest.ValidateAll() if the designated
// constraints aren't met.
type GetResourcesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourcesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourcesRequestMultiError) AllErrors() []error { return m }

// GetResourcesRequestValidationError is the validation error returned by
// GetResourcesRequest.Validate if the designated constraints aren't met.
type GetResourcesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourcesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourcesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourcesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourcesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourcesRequestValidationError) ErrorName() string {
	return "GetResourcesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetResourcesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourcesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourcesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourcesRequestValidationError{}

// Validate checks the field values on SaveResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveResourceRequestMultiError, or nil if none found.
func (m *SaveResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveResourceRequestValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveResourceRequestValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveResourceRequestValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SaveResourceRequestMultiError(errors)
	}

	return nil
}

// SaveResourceRequestMultiError is an error wrapping multiple validation
// errors returned by SaveResourceRequest.ValidateAll() if the designated
// constraints aren't met.
type SaveResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveResourceRequestMultiError) AllErrors() []error { return m }

// SaveResourceRequestValidationError is the validation error returned by
// SaveResourceRequest.Validate if the designated constraints aren't met.
type SaveResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveResourceRequestValidationError) ErrorName() string {
	return "SaveResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveResourceRequestValidationError{}

// Validate checks the field values on SaveResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveResourceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveResourceResponseMultiError, or nil if none found.
func (m *SaveResourceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveResourceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveResourceResponseValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SaveResourceResponseMultiError(errors)
	}

	return nil
}

// SaveResourceResponseMultiError is an error wrapping multiple validation
// errors returned by SaveResourceResponse.ValidateAll() if the designated
// constraints aren't met.
type SaveResourceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveResourceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveResourceResponseMultiError) AllErrors() []error { return m }

// SaveResourceResponseValidationError is the validation error returned by
// SaveResourceResponse.Validate if the designated constraints aren't met.
type SaveResourceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveResourceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveResourceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveResourceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveResourceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveResourceResponseValidationError) ErrorName() string {
	return "SaveResourceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SaveResourceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveResourceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveResourceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveResourceResponseValidationError{}