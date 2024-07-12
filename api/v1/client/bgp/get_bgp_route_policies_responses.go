// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetBgpRoutePoliciesReader is a Reader for the GetBgpRoutePolicies structure.
type GetBgpRoutePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBgpRoutePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBgpRoutePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetBgpRoutePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetBgpRoutePoliciesDisabled()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /bgp/route-policies] GetBgpRoutePolicies", response, response.Code())
	}
}

// NewGetBgpRoutePoliciesOK creates a GetBgpRoutePoliciesOK with default headers values
func NewGetBgpRoutePoliciesOK() *GetBgpRoutePoliciesOK {
	return &GetBgpRoutePoliciesOK{}
}

/*
GetBgpRoutePoliciesOK describes a response with status code 200, with default header values.

Success
*/
type GetBgpRoutePoliciesOK struct {
	Payload []*models.BgpRoutePolicy
}

// IsSuccess returns true when this get bgp route policies o k response has a 2xx status code
func (o *GetBgpRoutePoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bgp route policies o k response has a 3xx status code
func (o *GetBgpRoutePoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bgp route policies o k response has a 4xx status code
func (o *GetBgpRoutePoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bgp route policies o k response has a 5xx status code
func (o *GetBgpRoutePoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bgp route policies o k response a status code equal to that given
func (o *GetBgpRoutePoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bgp route policies o k response
func (o *GetBgpRoutePoliciesOK) Code() int {
	return 200
}

func (o *GetBgpRoutePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /bgp/route-policies][%d] getBgpRoutePoliciesOK  %+v", 200, o.Payload)
}

func (o *GetBgpRoutePoliciesOK) String() string {
	return fmt.Sprintf("[GET /bgp/route-policies][%d] getBgpRoutePoliciesOK  %+v", 200, o.Payload)
}

func (o *GetBgpRoutePoliciesOK) GetPayload() []*models.BgpRoutePolicy {
	return o.Payload
}

func (o *GetBgpRoutePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBgpRoutePoliciesInternalServerError creates a GetBgpRoutePoliciesInternalServerError with default headers values
func NewGetBgpRoutePoliciesInternalServerError() *GetBgpRoutePoliciesInternalServerError {
	return &GetBgpRoutePoliciesInternalServerError{}
}

/*
GetBgpRoutePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBgpRoutePoliciesInternalServerError struct {
	Payload models.Error
}

// IsSuccess returns true when this get bgp route policies internal server error response has a 2xx status code
func (o *GetBgpRoutePoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bgp route policies internal server error response has a 3xx status code
func (o *GetBgpRoutePoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bgp route policies internal server error response has a 4xx status code
func (o *GetBgpRoutePoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bgp route policies internal server error response has a 5xx status code
func (o *GetBgpRoutePoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get bgp route policies internal server error response a status code equal to that given
func (o *GetBgpRoutePoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get bgp route policies internal server error response
func (o *GetBgpRoutePoliciesInternalServerError) Code() int {
	return 500
}

func (o *GetBgpRoutePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bgp/route-policies][%d] getBgpRoutePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBgpRoutePoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /bgp/route-policies][%d] getBgpRoutePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBgpRoutePoliciesInternalServerError) GetPayload() models.Error {
	return o.Payload
}

func (o *GetBgpRoutePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBgpRoutePoliciesDisabled creates a GetBgpRoutePoliciesDisabled with default headers values
func NewGetBgpRoutePoliciesDisabled() *GetBgpRoutePoliciesDisabled {
	return &GetBgpRoutePoliciesDisabled{}
}

/*
GetBgpRoutePoliciesDisabled describes a response with status code 501, with default header values.

BGP Control Plane disabled
*/
type GetBgpRoutePoliciesDisabled struct {
	Payload models.Error
}

// IsSuccess returns true when this get bgp route policies disabled response has a 2xx status code
func (o *GetBgpRoutePoliciesDisabled) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bgp route policies disabled response has a 3xx status code
func (o *GetBgpRoutePoliciesDisabled) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bgp route policies disabled response has a 4xx status code
func (o *GetBgpRoutePoliciesDisabled) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bgp route policies disabled response has a 5xx status code
func (o *GetBgpRoutePoliciesDisabled) IsServerError() bool {
	return true
}

// IsCode returns true when this get bgp route policies disabled response a status code equal to that given
func (o *GetBgpRoutePoliciesDisabled) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the get bgp route policies disabled response
func (o *GetBgpRoutePoliciesDisabled) Code() int {
	return 501
}

func (o *GetBgpRoutePoliciesDisabled) Error() string {
	return fmt.Sprintf("[GET /bgp/route-policies][%d] getBgpRoutePoliciesDisabled  %+v", 501, o.Payload)
}

func (o *GetBgpRoutePoliciesDisabled) String() string {
	return fmt.Sprintf("[GET /bgp/route-policies][%d] getBgpRoutePoliciesDisabled  %+v", 501, o.Payload)
}

func (o *GetBgpRoutePoliciesDisabled) GetPayload() models.Error {
	return o.Payload
}

func (o *GetBgpRoutePoliciesDisabled) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
