// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/rhel_host.proto

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

// Validate checks the field values on RHELHost with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RHELHost) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RHELHost with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RHELHostMultiError, or nil
// if none found.
func (m *RHELHost) ValidateAll() error {
	return m.validate(true)
}

func (m *RHELHost) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMetadata() == nil {
		err := RHELHostValidationError{
			field:  "Metadata",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetMetadata(); a != nil {

	}

	if m.GetReporterData() == nil {
		err := RHELHostValidationError{
			field:  "ReporterData",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetReporterData(); a != nil {

	}

	for idx, item := range m.GetReporters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RHELHostValidationError{
						field:  fmt.Sprintf("Reporters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RHELHostValidationError{
						field:  fmt.Sprintf("Reporters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RHELHostValidationError{
					field:  fmt.Sprintf("Reporters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RHELHostMultiError(errors)
	}

	return nil
}

// RHELHostMultiError is an error wrapping multiple validation errors returned
// by RHELHost.ValidateAll() if the designated constraints aren't met.
type RHELHostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RHELHostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RHELHostMultiError) AllErrors() []error { return m }

// RHELHostValidationError is the validation error returned by
// RHELHost.Validate if the designated constraints aren't met.
type RHELHostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RHELHostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RHELHostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RHELHostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RHELHostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RHELHostValidationError) ErrorName() string { return "RHELHostValidationError" }

// Error satisfies the builtin error interface
func (e RHELHostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRHELHost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RHELHostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RHELHostValidationError{}