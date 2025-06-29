// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"n3dr/internal/app/n3dr/goswagger/models"
)

// NewUpdatePrivilege1Params creates a new UpdatePrivilege1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePrivilege1Params() *UpdatePrivilege1Params {
	return &UpdatePrivilege1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrivilege1ParamsWithTimeout creates a new UpdatePrivilege1Params object
// with the ability to set a timeout on a request.
func NewUpdatePrivilege1ParamsWithTimeout(timeout time.Duration) *UpdatePrivilege1Params {
	return &UpdatePrivilege1Params{
		timeout: timeout,
	}
}

// NewUpdatePrivilege1ParamsWithContext creates a new UpdatePrivilege1Params object
// with the ability to set a context for a request.
func NewUpdatePrivilege1ParamsWithContext(ctx context.Context) *UpdatePrivilege1Params {
	return &UpdatePrivilege1Params{
		Context: ctx,
	}
}

// NewUpdatePrivilege1ParamsWithHTTPClient creates a new UpdatePrivilege1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePrivilege1ParamsWithHTTPClient(client *http.Client) *UpdatePrivilege1Params {
	return &UpdatePrivilege1Params{
		HTTPClient: client,
	}
}

/*
UpdatePrivilege1Params contains all the parameters to send to the API endpoint

	for the update privilege 1 operation.

	Typically these are written to a http.Request.
*/
type UpdatePrivilege1Params struct {

	/* Body.

	   The privilege to update.
	*/
	Body *models.APIPrivilegeApplicationRequest

	/* PrivilegeName.

	   The name of the privilege to update.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update privilege 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege1Params) WithDefaults() *UpdatePrivilege1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update privilege 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update privilege 1 params
func (o *UpdatePrivilege1Params) WithTimeout(timeout time.Duration) *UpdatePrivilege1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update privilege 1 params
func (o *UpdatePrivilege1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update privilege 1 params
func (o *UpdatePrivilege1Params) WithContext(ctx context.Context) *UpdatePrivilege1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update privilege 1 params
func (o *UpdatePrivilege1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update privilege 1 params
func (o *UpdatePrivilege1Params) WithHTTPClient(client *http.Client) *UpdatePrivilege1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update privilege 1 params
func (o *UpdatePrivilege1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update privilege 1 params
func (o *UpdatePrivilege1Params) WithBody(body *models.APIPrivilegeApplicationRequest) *UpdatePrivilege1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update privilege 1 params
func (o *UpdatePrivilege1Params) SetBody(body *models.APIPrivilegeApplicationRequest) {
	o.Body = body
}

// WithPrivilegeName adds the privilegeName to the update privilege 1 params
func (o *UpdatePrivilege1Params) WithPrivilegeName(privilegeName string) *UpdatePrivilege1Params {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the update privilege 1 params
func (o *UpdatePrivilege1Params) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrivilege1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param privilegeName
	if err := r.SetPathParam("privilegeName", o.PrivilegeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
