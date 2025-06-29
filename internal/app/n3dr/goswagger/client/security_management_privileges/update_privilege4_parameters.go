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

// NewUpdatePrivilege4Params creates a new UpdatePrivilege4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePrivilege4Params() *UpdatePrivilege4Params {
	return &UpdatePrivilege4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrivilege4ParamsWithTimeout creates a new UpdatePrivilege4Params object
// with the ability to set a timeout on a request.
func NewUpdatePrivilege4ParamsWithTimeout(timeout time.Duration) *UpdatePrivilege4Params {
	return &UpdatePrivilege4Params{
		timeout: timeout,
	}
}

// NewUpdatePrivilege4ParamsWithContext creates a new UpdatePrivilege4Params object
// with the ability to set a context for a request.
func NewUpdatePrivilege4ParamsWithContext(ctx context.Context) *UpdatePrivilege4Params {
	return &UpdatePrivilege4Params{
		Context: ctx,
	}
}

// NewUpdatePrivilege4ParamsWithHTTPClient creates a new UpdatePrivilege4Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePrivilege4ParamsWithHTTPClient(client *http.Client) *UpdatePrivilege4Params {
	return &UpdatePrivilege4Params{
		HTTPClient: client,
	}
}

/*
UpdatePrivilege4Params contains all the parameters to send to the API endpoint

	for the update privilege 4 operation.

	Typically these are written to a http.Request.
*/
type UpdatePrivilege4Params struct {

	/* Body.

	   The privilege to update.
	*/
	Body *models.APIPrivilegeRepositoryAdminRequest

	/* PrivilegeName.

	   The name of the privilege to update.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update privilege 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege4Params) WithDefaults() *UpdatePrivilege4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update privilege 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update privilege 4 params
func (o *UpdatePrivilege4Params) WithTimeout(timeout time.Duration) *UpdatePrivilege4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update privilege 4 params
func (o *UpdatePrivilege4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update privilege 4 params
func (o *UpdatePrivilege4Params) WithContext(ctx context.Context) *UpdatePrivilege4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update privilege 4 params
func (o *UpdatePrivilege4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update privilege 4 params
func (o *UpdatePrivilege4Params) WithHTTPClient(client *http.Client) *UpdatePrivilege4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update privilege 4 params
func (o *UpdatePrivilege4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update privilege 4 params
func (o *UpdatePrivilege4Params) WithBody(body *models.APIPrivilegeRepositoryAdminRequest) *UpdatePrivilege4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update privilege 4 params
func (o *UpdatePrivilege4Params) SetBody(body *models.APIPrivilegeRepositoryAdminRequest) {
	o.Body = body
}

// WithPrivilegeName adds the privilegeName to the update privilege 4 params
func (o *UpdatePrivilege4Params) WithPrivilegeName(privilegeName string) *UpdatePrivilege4Params {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the update privilege 4 params
func (o *UpdatePrivilege4Params) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrivilege4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
