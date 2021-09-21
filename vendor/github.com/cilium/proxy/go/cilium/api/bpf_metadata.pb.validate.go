// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/bpf_metadata.proto

package cilium

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on BpfMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BpfMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BpfRoot

	// no validation rules for IsIngress

	// no validation rules for MayUseOriginalSourceAddress

	// no validation rules for EgressMarkSourceEndpointId

	return nil
}

// BpfMetadataValidationError is the validation error returned by
// BpfMetadata.Validate if the designated constraints aren't met.
type BpfMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BpfMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BpfMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BpfMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BpfMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BpfMetadataValidationError) ErrorName() string { return "BpfMetadataValidationError" }

// Error satisfies the builtin error interface
func (e BpfMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBpfMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BpfMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BpfMetadataValidationError{}
