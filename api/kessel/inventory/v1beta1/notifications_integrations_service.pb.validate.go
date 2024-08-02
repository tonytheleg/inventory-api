// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/notifications_integrations_service.proto

package v1beta1

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

// Validate checks the field values on CreateNotificationsIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CreateNotificationsIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNotificationsIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreateNotificationsIntegrationRequestMultiError, or nil if none found.
func (m *CreateNotificationsIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNotificationsIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIntegration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateNotificationsIntegrationRequestValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateNotificationsIntegrationRequestValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIntegration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateNotificationsIntegrationRequestValidationError{
				field:  "Integration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateNotificationsIntegrationRequestMultiError(errors)
	}

	return nil
}

// CreateNotificationsIntegrationRequestMultiError is an error wrapping
// multiple validation errors returned by
// CreateNotificationsIntegrationRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateNotificationsIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNotificationsIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNotificationsIntegrationRequestMultiError) AllErrors() []error { return m }

// CreateNotificationsIntegrationRequestValidationError is the validation error
// returned by CreateNotificationsIntegrationRequest.Validate if the
// designated constraints aren't met.
type CreateNotificationsIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNotificationsIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNotificationsIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNotificationsIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNotificationsIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNotificationsIntegrationRequestValidationError) ErrorName() string {
	return "CreateNotificationsIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNotificationsIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNotificationsIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNotificationsIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNotificationsIntegrationRequestValidationError{}

// Validate checks the field values on CreateNotificationsIntegrationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CreateNotificationsIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CreateNotificationsIntegrationResponse with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// CreateNotificationsIntegrationResponseMultiError, or nil if none found.
func (m *CreateNotificationsIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNotificationsIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIntegration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateNotificationsIntegrationResponseValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateNotificationsIntegrationResponseValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIntegration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateNotificationsIntegrationResponseValidationError{
				field:  "Integration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateNotificationsIntegrationResponseMultiError(errors)
	}

	return nil
}

// CreateNotificationsIntegrationResponseMultiError is an error wrapping
// multiple validation errors returned by
// CreateNotificationsIntegrationResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateNotificationsIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNotificationsIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNotificationsIntegrationResponseMultiError) AllErrors() []error { return m }

// CreateNotificationsIntegrationResponseValidationError is the validation
// error returned by CreateNotificationsIntegrationResponse.Validate if the
// designated constraints aren't met.
type CreateNotificationsIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNotificationsIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNotificationsIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNotificationsIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNotificationsIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNotificationsIntegrationResponseValidationError) ErrorName() string {
	return "CreateNotificationsIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNotificationsIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNotificationsIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNotificationsIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNotificationsIntegrationResponseValidationError{}

// Validate checks the field values on UpdateNotificationsIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateNotificationsIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNotificationsIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdateNotificationsIntegrationRequestMultiError, or nil if none found.
func (m *UpdateNotificationsIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNotificationsIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Resource

	if all {
		switch v := interface{}(m.GetIntegration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNotificationsIntegrationRequestValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNotificationsIntegrationRequestValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIntegration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNotificationsIntegrationRequestValidationError{
				field:  "Integration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNotificationsIntegrationRequestMultiError(errors)
	}

	return nil
}

// UpdateNotificationsIntegrationRequestMultiError is an error wrapping
// multiple validation errors returned by
// UpdateNotificationsIntegrationRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateNotificationsIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNotificationsIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNotificationsIntegrationRequestMultiError) AllErrors() []error { return m }

// UpdateNotificationsIntegrationRequestValidationError is the validation error
// returned by UpdateNotificationsIntegrationRequest.Validate if the
// designated constraints aren't met.
type UpdateNotificationsIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNotificationsIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNotificationsIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNotificationsIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNotificationsIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNotificationsIntegrationRequestValidationError) ErrorName() string {
	return "UpdateNotificationsIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNotificationsIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNotificationsIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNotificationsIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNotificationsIntegrationRequestValidationError{}

// Validate checks the field values on UpdateNotificationsIntegrationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateNotificationsIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// UpdateNotificationsIntegrationResponse with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// UpdateNotificationsIntegrationResponseMultiError, or nil if none found.
func (m *UpdateNotificationsIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNotificationsIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateNotificationsIntegrationResponseMultiError(errors)
	}

	return nil
}

// UpdateNotificationsIntegrationResponseMultiError is an error wrapping
// multiple validation errors returned by
// UpdateNotificationsIntegrationResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateNotificationsIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNotificationsIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNotificationsIntegrationResponseMultiError) AllErrors() []error { return m }

// UpdateNotificationsIntegrationResponseValidationError is the validation
// error returned by UpdateNotificationsIntegrationResponse.Validate if the
// designated constraints aren't met.
type UpdateNotificationsIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNotificationsIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNotificationsIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNotificationsIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNotificationsIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNotificationsIntegrationResponseValidationError) ErrorName() string {
	return "UpdateNotificationsIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNotificationsIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNotificationsIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNotificationsIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNotificationsIntegrationResponseValidationError{}

// Validate checks the field values on DeleteNotificationsIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DeleteNotificationsIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNotificationsIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// DeleteNotificationsIntegrationRequestMultiError, or nil if none found.
func (m *DeleteNotificationsIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNotificationsIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Resource

	if len(errors) > 0 {
		return DeleteNotificationsIntegrationRequestMultiError(errors)
	}

	return nil
}

// DeleteNotificationsIntegrationRequestMultiError is an error wrapping
// multiple validation errors returned by
// DeleteNotificationsIntegrationRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteNotificationsIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNotificationsIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNotificationsIntegrationRequestMultiError) AllErrors() []error { return m }

// DeleteNotificationsIntegrationRequestValidationError is the validation error
// returned by DeleteNotificationsIntegrationRequest.Validate if the
// designated constraints aren't met.
type DeleteNotificationsIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNotificationsIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNotificationsIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNotificationsIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNotificationsIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNotificationsIntegrationRequestValidationError) ErrorName() string {
	return "DeleteNotificationsIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNotificationsIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNotificationsIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNotificationsIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNotificationsIntegrationRequestValidationError{}

// Validate checks the field values on DeleteNotificationsIntegrationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DeleteNotificationsIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// DeleteNotificationsIntegrationResponse with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// DeleteNotificationsIntegrationResponseMultiError, or nil if none found.
func (m *DeleteNotificationsIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNotificationsIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteNotificationsIntegrationResponseMultiError(errors)
	}

	return nil
}

// DeleteNotificationsIntegrationResponseMultiError is an error wrapping
// multiple validation errors returned by
// DeleteNotificationsIntegrationResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteNotificationsIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNotificationsIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNotificationsIntegrationResponseMultiError) AllErrors() []error { return m }

// DeleteNotificationsIntegrationResponseValidationError is the validation
// error returned by DeleteNotificationsIntegrationResponse.Validate if the
// designated constraints aren't met.
type DeleteNotificationsIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNotificationsIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNotificationsIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNotificationsIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNotificationsIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNotificationsIntegrationResponseValidationError) ErrorName() string {
	return "DeleteNotificationsIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNotificationsIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNotificationsIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNotificationsIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNotificationsIntegrationResponseValidationError{}
