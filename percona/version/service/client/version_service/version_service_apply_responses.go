// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ssk199441/agora-postgresql-operator/percona/version/service/client/models"
)

// VersionServiceApplyReader is a Reader for the VersionServiceApply structure.
type VersionServiceApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionServiceApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionServiceApplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionServiceApplyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionServiceApplyOK creates a VersionServiceApplyOK with default headers values
func NewVersionServiceApplyOK() *VersionServiceApplyOK {
	return &VersionServiceApplyOK{}
}

/*
VersionServiceApplyOK describes a response with status code 200, with default header values.

A successful response.
*/
type VersionServiceApplyOK struct {
	Payload *models.VersionVersionResponse
}

// IsSuccess returns true when this version service apply o k response has a 2xx status code
func (o *VersionServiceApplyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this version service apply o k response has a 3xx status code
func (o *VersionServiceApplyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this version service apply o k response has a 4xx status code
func (o *VersionServiceApplyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this version service apply o k response has a 5xx status code
func (o *VersionServiceApplyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this version service apply o k response a status code equal to that given
func (o *VersionServiceApplyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the version service apply o k response
func (o *VersionServiceApplyOK) Code() int {
	return 200
}

func (o *VersionServiceApplyOK) Error() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}/{apply}][%d] versionServiceApplyOK  %+v", 200, o.Payload)
}

func (o *VersionServiceApplyOK) String() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}/{apply}][%d] versionServiceApplyOK  %+v", 200, o.Payload)
}

func (o *VersionServiceApplyOK) GetPayload() *models.VersionVersionResponse {
	return o.Payload
}

func (o *VersionServiceApplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionServiceApplyDefault creates a VersionServiceApplyDefault with default headers values
func NewVersionServiceApplyDefault(code int) *VersionServiceApplyDefault {
	return &VersionServiceApplyDefault{
		_statusCode: code,
	}
}

/*
VersionServiceApplyDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type VersionServiceApplyDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this version service apply default response has a 2xx status code
func (o *VersionServiceApplyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this version service apply default response has a 3xx status code
func (o *VersionServiceApplyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this version service apply default response has a 4xx status code
func (o *VersionServiceApplyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this version service apply default response has a 5xx status code
func (o *VersionServiceApplyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this version service apply default response a status code equal to that given
func (o *VersionServiceApplyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the version service apply default response
func (o *VersionServiceApplyDefault) Code() int {
	return o._statusCode
}

func (o *VersionServiceApplyDefault) Error() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}/{apply}][%d] VersionService_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *VersionServiceApplyDefault) String() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}/{apply}][%d] VersionService_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *VersionServiceApplyDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *VersionServiceApplyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
