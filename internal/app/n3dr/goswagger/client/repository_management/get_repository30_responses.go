// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"n3dr/internal/app/n3dr/goswagger/models"
)

// GetRepository30Reader is a Reader for the GetRepository30 structure.
type GetRepository30Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository30Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository30OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository30OK creates a GetRepository30OK with default headers values
func NewGetRepository30OK() *GetRepository30OK {
	return &GetRepository30OK{}
}

/*
GetRepository30OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository30OK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repository30 o k response has a 2xx status code
func (o *GetRepository30OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository30 o k response has a 3xx status code
func (o *GetRepository30OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository30 o k response has a 4xx status code
func (o *GetRepository30OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository30 o k response has a 5xx status code
func (o *GetRepository30OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository30 o k response a status code equal to that given
func (o *GetRepository30OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository30 o k response
func (o *GetRepository30OK) Code() int {
	return 200
}

func (o *GetRepository30OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/conda/proxy/{repositoryName}][%d] getRepository30OK  %+v", 200, o.Payload)
}

func (o *GetRepository30OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/conda/proxy/{repositoryName}][%d] getRepository30OK  %+v", 200, o.Payload)
}

func (o *GetRepository30OK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepository30OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
