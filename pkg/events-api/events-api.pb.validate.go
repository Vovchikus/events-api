// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/events-api/v1/events-api.proto

package events_api_v1

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

// Validate checks the field values on CreateEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateEventRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateEventRequestMultiError, or nil if none found.
func (m *CreateEventRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateEventRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateEventRequestMultiError(errors)
	}

	return nil
}

// CreateEventRequestMultiError is an error wrapping multiple validation errors
// returned by CreateEventRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateEventRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateEventRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateEventRequestMultiError) AllErrors() []error { return m }

// CreateEventRequestValidationError is the validation error returned by
// CreateEventRequest.Validate if the designated constraints aren't met.
type CreateEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEventRequestValidationError) ErrorName() string {
	return "CreateEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEventRequestValidationError{}

// Validate checks the field values on CreateEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateEventResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateEventResponseMultiError, or nil if none found.
func (m *CreateEventResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateEventResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return CreateEventResponseMultiError(errors)
	}

	return nil
}

// CreateEventResponseMultiError is an error wrapping multiple validation
// errors returned by CreateEventResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateEventResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateEventResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateEventResponseMultiError) AllErrors() []error { return m }

// CreateEventResponseValidationError is the validation error returned by
// CreateEventResponse.Validate if the designated constraints aren't met.
type CreateEventResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEventResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEventResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEventResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEventResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEventResponseValidationError) ErrorName() string {
	return "CreateEventResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEventResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEventResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEventResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEventResponseValidationError{}

// Validate checks the field values on Health with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Health) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Health with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HealthMultiError, or nil if none found.
func (m *Health) ValidateAll() error {
	return m.validate(true)
}

func (m *Health) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uptime

	// no validation rules for AllocatedMemory

	// no validation rules for TotalAllocatedMemory

	// no validation rules for Goroutines

	// no validation rules for GCCycles

	// no validation rules for NumberOfCPUs

	// no validation rules for HeapSys

	// no validation rules for HeapAllocated

	// no validation rules for ObjectsInUse

	// no validation rules for OSMemoryObtained

	if len(errors) > 0 {
		return HealthMultiError(errors)
	}

	return nil
}

// HealthMultiError is an error wrapping multiple validation errors returned by
// Health.ValidateAll() if the designated constraints aren't met.
type HealthMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthMultiError) AllErrors() []error { return m }

// HealthValidationError is the validation error returned by Health.Validate if
// the designated constraints aren't met.
type HealthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthValidationError) ErrorName() string { return "HealthValidationError" }

// Error satisfies the builtin error interface
func (e HealthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthValidationError{}