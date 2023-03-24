// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/common/key_value/v3/config.proto

package key_valuev3

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

// Validate checks the field values on KeyValueStoreConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *KeyValueStoreConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KeyValueStoreConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// KeyValueStoreConfigMultiError, or nil if none found.
func (m *KeyValueStoreConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *KeyValueStoreConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetConfig() == nil {
		err := KeyValueStoreConfigValidationError{
			field:  "Config",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, KeyValueStoreConfigValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, KeyValueStoreConfigValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KeyValueStoreConfigValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return KeyValueStoreConfigMultiError(errors)
	}
	return nil
}

// KeyValueStoreConfigMultiError is an error wrapping multiple validation
// errors returned by KeyValueStoreConfig.ValidateAll() if the designated
// constraints aren't met.
type KeyValueStoreConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KeyValueStoreConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KeyValueStoreConfigMultiError) AllErrors() []error { return m }

// KeyValueStoreConfigValidationError is the validation error returned by
// KeyValueStoreConfig.Validate if the designated constraints aren't met.
type KeyValueStoreConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyValueStoreConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyValueStoreConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyValueStoreConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyValueStoreConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyValueStoreConfigValidationError) ErrorName() string {
	return "KeyValueStoreConfigValidationError"
}

// Error satisfies the builtin error interface
func (e KeyValueStoreConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyValueStoreConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyValueStoreConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyValueStoreConfigValidationError{}
