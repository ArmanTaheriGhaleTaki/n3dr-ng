// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"n3dr/internal/app/n3dr/goswagger/models"
)

// GetSystemStatusChecksReader is a Reader for the GetSystemStatusChecks structure.
type GetSystemStatusChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemStatusChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemStatusChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystemStatusChecksOK creates a GetSystemStatusChecksOK with default headers values
func NewGetSystemStatusChecksOK() *GetSystemStatusChecksOK {
	return &GetSystemStatusChecksOK{}
}

/*
GetSystemStatusChecksOK describes a response with status code 200, with default header values.

The system status check results
*/
type GetSystemStatusChecksOK struct {
	Payload map[string]models.Result
}

// IsSuccess returns true when this get system status checks o k response has a 2xx status code
func (o *GetSystemStatusChecksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get system status checks o k response has a 3xx status code
func (o *GetSystemStatusChecksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system status checks o k response has a 4xx status code
func (o *GetSystemStatusChecksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system status checks o k response has a 5xx status code
func (o *GetSystemStatusChecksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get system status checks o k response a status code equal to that given
func (o *GetSystemStatusChecksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get system status checks o k response
func (o *GetSystemStatusChecksOK) Code() int {
	return 200
}

func (o *GetSystemStatusChecksOK) Error() string {
	return fmt.Sprintf("[GET /v1/status/check][%d] getSystemStatusChecksOK  %+v", 200, o.Payload)
}

func (o *GetSystemStatusChecksOK) String() string {
	return fmt.Sprintf("[GET /v1/status/check][%d] getSystemStatusChecksOK  %+v", 200, o.Payload)
}

func (o *GetSystemStatusChecksOK) GetPayload() map[string]models.Result {
	return o.Payload
}

func (o *GetSystemStatusChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
