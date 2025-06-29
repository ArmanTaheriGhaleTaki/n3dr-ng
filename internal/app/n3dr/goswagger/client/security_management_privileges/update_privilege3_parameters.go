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

// NewUpdatePrivilege3Params creates a new UpdatePrivilege3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePrivilege3Params() *UpdatePrivilege3Params {
	return &UpdatePrivilege3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrivilege3ParamsWithTimeout creates a new UpdatePrivilege3Params object
// with the ability to set a timeout on a request.
func NewUpdatePrivilege3ParamsWithTimeout(timeout time.Duration) *UpdatePrivilege3Params {
	return &UpdatePrivilege3Params{
		timeout: timeout,
	}
}

// NewUpdatePrivilege3ParamsWithContext creates a new UpdatePrivilege3Params object
// with the ability to set a context for a request.
func NewUpdatePrivilege3ParamsWithContext(ctx context.Context) *UpdatePrivilege3Params {
	return &UpdatePrivilege3Params{
		Context: ctx,
	}
}

// NewUpdatePrivilege3ParamsWithHTTPClient creates a new UpdatePrivilege3Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePrivilege3ParamsWithHTTPClient(client *http.Client) *UpdatePrivilege3Params {
	return &UpdatePrivilege3Params{
		HTTPClient: client,
	}
}

/*
UpdatePrivilege3Params contains all the parameters to send to the API endpoint

	for the update privilege 3 operation.

	Typically these are written to a http.Request.
*/
type UpdatePrivilege3Params struct {

	/* Body.

	   The privilege to update.
	*/
	Body *models.APIPrivilegeRepositoryContentSelectorRequest

	/* PrivilegeName.

	   The name of the privilege to update.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update privilege 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege3Params) WithDefaults() *UpdatePrivilege3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update privilege 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update privilege 3 params
func (o *UpdatePrivilege3Params) WithTimeout(timeout time.Duration) *UpdatePrivilege3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update privilege 3 params
func (o *UpdatePrivilege3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update privilege 3 params
func (o *UpdatePrivilege3Params) WithContext(ctx context.Context) *UpdatePrivilege3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update privilege 3 params
func (o *UpdatePrivilege3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update privilege 3 params
func (o *UpdatePrivilege3Params) WithHTTPClient(client *http.Client) *UpdatePrivilege3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update privilege 3 params
func (o *UpdatePrivilege3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update privilege 3 params
func (o *UpdatePrivilege3Params) WithBody(body *models.APIPrivilegeRepositoryContentSelectorRequest) *UpdatePrivilege3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update privilege 3 params
func (o *UpdatePrivilege3Params) SetBody(body *models.APIPrivilegeRepositoryContentSelectorRequest) {
	o.Body = body
}

// WithPrivilegeName adds the privilegeName to the update privilege 3 params
func (o *UpdatePrivilege3Params) WithPrivilegeName(privilegeName string) *UpdatePrivilege3Params {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the update privilege 3 params
func (o *UpdatePrivilege3Params) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrivilege3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
