// Code generated by go-swagger; DO NOT EDIT.

package security_certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"n3dr/internal/app/n3dr/goswagger/models"
)

// RetrieveCertificateReader is a Reader for the RetrieveCertificate structure.
type RetrieveCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetrieveCertificateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrieveCertificateOK creates a RetrieveCertificateOK with default headers values
func NewRetrieveCertificateOK() *RetrieveCertificateOK {
	return &RetrieveCertificateOK{}
}

/*
RetrieveCertificateOK describes a response with status code 200, with default header values.

successful operation
*/
type RetrieveCertificateOK struct {
	Payload *models.APICertificate
}

// IsSuccess returns true when this retrieve certificate o k response has a 2xx status code
func (o *RetrieveCertificateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retrieve certificate o k response has a 3xx status code
func (o *RetrieveCertificateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve certificate o k response has a 4xx status code
func (o *RetrieveCertificateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve certificate o k response has a 5xx status code
func (o *RetrieveCertificateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve certificate o k response a status code equal to that given
func (o *RetrieveCertificateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the retrieve certificate o k response
func (o *RetrieveCertificateOK) Code() int {
	return 200
}

func (o *RetrieveCertificateOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/ssl][%d] retrieveCertificateOK  %+v", 200, o.Payload)
}

func (o *RetrieveCertificateOK) String() string {
	return fmt.Sprintf("[GET /v1/security/ssl][%d] retrieveCertificateOK  %+v", 200, o.Payload)
}

func (o *RetrieveCertificateOK) GetPayload() *models.APICertificate {
	return o.Payload
}

func (o *RetrieveCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveCertificateBadRequest creates a RetrieveCertificateBadRequest with default headers values
func NewRetrieveCertificateBadRequest() *RetrieveCertificateBadRequest {
	return &RetrieveCertificateBadRequest{}
}

/*
RetrieveCertificateBadRequest describes a response with status code 400, with default header values.

A certificate could not be retrieved, see the message for details.
*/
type RetrieveCertificateBadRequest struct {
}

// IsSuccess returns true when this retrieve certificate bad request response has a 2xx status code
func (o *RetrieveCertificateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve certificate bad request response has a 3xx status code
func (o *RetrieveCertificateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve certificate bad request response has a 4xx status code
func (o *RetrieveCertificateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve certificate bad request response has a 5xx status code
func (o *RetrieveCertificateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve certificate bad request response a status code equal to that given
func (o *RetrieveCertificateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the retrieve certificate bad request response
func (o *RetrieveCertificateBadRequest) Code() int {
	return 400
}

func (o *RetrieveCertificateBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/security/ssl][%d] retrieveCertificateBadRequest ", 400)
}

func (o *RetrieveCertificateBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/security/ssl][%d] retrieveCertificateBadRequest ", 400)
}

func (o *RetrieveCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrieveCertificateForbidden creates a RetrieveCertificateForbidden with default headers values
func NewRetrieveCertificateForbidden() *RetrieveCertificateForbidden {
	return &RetrieveCertificateForbidden{}
}

/*
RetrieveCertificateForbidden describes a response with status code 403, with default header values.

Insufficient permissions to retrieve remote certificate.
*/
type RetrieveCertificateForbidden struct {
}

// IsSuccess returns true when this retrieve certificate forbidden response has a 2xx status code
func (o *RetrieveCertificateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve certificate forbidden response has a 3xx status code
func (o *RetrieveCertificateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve certificate forbidden response has a 4xx status code
func (o *RetrieveCertificateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve certificate forbidden response has a 5xx status code
func (o *RetrieveCertificateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve certificate forbidden response a status code equal to that given
func (o *RetrieveCertificateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the retrieve certificate forbidden response
func (o *RetrieveCertificateForbidden) Code() int {
	return 403
}

func (o *RetrieveCertificateForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/ssl][%d] retrieveCertificateForbidden ", 403)
}

func (o *RetrieveCertificateForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/ssl][%d] retrieveCertificateForbidden ", 403)
}

func (o *RetrieveCertificateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
