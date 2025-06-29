// Code generated by go-swagger; DO NOT EDIT.

package script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"n3dr/internal/app/n3dr/goswagger/models"
)

// BrowseReader is a Reader for the Browse structure.
type BrowseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BrowseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBrowseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBrowseOK creates a BrowseOK with default headers values
func NewBrowseOK() *BrowseOK {
	return &BrowseOK{}
}

/*
BrowseOK describes a response with status code 200, with default header values.

successful operation
*/
type BrowseOK struct {
	Payload []*models.ScriptXO
}

// IsSuccess returns true when this browse o k response has a 2xx status code
func (o *BrowseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this browse o k response has a 3xx status code
func (o *BrowseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this browse o k response has a 4xx status code
func (o *BrowseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this browse o k response has a 5xx status code
func (o *BrowseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this browse o k response a status code equal to that given
func (o *BrowseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the browse o k response
func (o *BrowseOK) Code() int {
	return 200
}

func (o *BrowseOK) Error() string {
	return fmt.Sprintf("[GET /v1/script][%d] browseOK  %+v", 200, o.Payload)
}

func (o *BrowseOK) String() string {
	return fmt.Sprintf("[GET /v1/script][%d] browseOK  %+v", 200, o.Payload)
}

func (o *BrowseOK) GetPayload() []*models.ScriptXO {
	return o.Payload
}

func (o *BrowseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
