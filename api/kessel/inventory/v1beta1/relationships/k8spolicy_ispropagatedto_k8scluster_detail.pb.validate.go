// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/relationships/k8spolicy_ispropagatedto_k8scluster_detail.proto

package relationships

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

// Validate checks the field values on K8SPolicyIsPropagatedToK8SClusterDetail
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *K8SPolicyIsPropagatedToK8SClusterDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// K8SPolicyIsPropagatedToK8SClusterDetail with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// K8SPolicyIsPropagatedToK8SClusterDetailMultiError, or nil if none found.
func (m *K8SPolicyIsPropagatedToK8SClusterDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *K8SPolicyIsPropagatedToK8SClusterDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for K8SPolicyId

	// no validation rules for K8SClusterId

	// no validation rules for Status

	if len(errors) > 0 {
		return K8SPolicyIsPropagatedToK8SClusterDetailMultiError(errors)
	}

	return nil
}

// K8SPolicyIsPropagatedToK8SClusterDetailMultiError is an error wrapping
// multiple validation errors returned by
// K8SPolicyIsPropagatedToK8SClusterDetail.ValidateAll() if the designated
// constraints aren't met.
type K8SPolicyIsPropagatedToK8SClusterDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m K8SPolicyIsPropagatedToK8SClusterDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m K8SPolicyIsPropagatedToK8SClusterDetailMultiError) AllErrors() []error { return m }

// K8SPolicyIsPropagatedToK8SClusterDetailValidationError is the validation
// error returned by K8SPolicyIsPropagatedToK8SClusterDetail.Validate if the
// designated constraints aren't met.
type K8SPolicyIsPropagatedToK8SClusterDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e K8SPolicyIsPropagatedToK8SClusterDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e K8SPolicyIsPropagatedToK8SClusterDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e K8SPolicyIsPropagatedToK8SClusterDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e K8SPolicyIsPropagatedToK8SClusterDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e K8SPolicyIsPropagatedToK8SClusterDetailValidationError) ErrorName() string {
	return "K8SPolicyIsPropagatedToK8SClusterDetailValidationError"
}

// Error satisfies the builtin error interface
func (e K8SPolicyIsPropagatedToK8SClusterDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sK8SPolicyIsPropagatedToK8SClusterDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = K8SPolicyIsPropagatedToK8SClusterDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = K8SPolicyIsPropagatedToK8SClusterDetailValidationError{}