// Code generated by go-swagger; DO NOT EDIT.

package security_management_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"n3dr/internal/app/n3dr/goswagger/models"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*
GetUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUsersOK struct {
	Payload []*models.APIUser
}

// IsSuccess returns true when this get users o k response has a 2xx status code
func (o *GetUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users o k response has a 3xx status code
func (o *GetUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users o k response has a 4xx status code
func (o *GetUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users o k response has a 5xx status code
func (o *GetUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users o k response a status code equal to that given
func (o *GetUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users o k response
func (o *GetUsersOK) Code() int {
	return 200
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) String() string {
	return fmt.Sprintf("[GET /v1/security/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() []*models.APIUser {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*
GetUsersBadRequest describes a response with status code 400, with default header values.

Password was not supplied in the body of the request
*/
type GetUsersBadRequest struct {
}

// IsSuccess returns true when this get users bad request response has a 2xx status code
func (o *GetUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users bad request response has a 3xx status code
func (o *GetUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users bad request response has a 4xx status code
func (o *GetUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users bad request response has a 5xx status code
func (o *GetUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get users bad request response a status code equal to that given
func (o *GetUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get users bad request response
func (o *GetUsersBadRequest) Code() int {
	return 400
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/security/users][%d] getUsersBadRequest ", 400)
}

func (o *GetUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/security/users][%d] getUsersBadRequest ", 400)
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForbidden creates a GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

/*
GetUsersForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type GetUsersForbidden struct {
}

// IsSuccess returns true when this get users forbidden response has a 2xx status code
func (o *GetUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users forbidden response has a 3xx status code
func (o *GetUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users forbidden response has a 4xx status code
func (o *GetUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users forbidden response has a 5xx status code
func (o *GetUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get users forbidden response a status code equal to that given
func (o *GetUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get users forbidden response
func (o *GetUsersForbidden) Code() int {
	return 403
}

func (o *GetUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/users][%d] getUsersForbidden ", 403)
}

func (o *GetUsersForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/users][%d] getUsersForbidden ", 403)
}

func (o *GetUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
