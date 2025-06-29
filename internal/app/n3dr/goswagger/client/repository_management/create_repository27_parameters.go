// Code generated by go-swagger; DO NOT EDIT.

package repository_management

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

// NewCreateRepository27Params creates a new CreateRepository27Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository27Params() *CreateRepository27Params {
	return &CreateRepository27Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository27ParamsWithTimeout creates a new CreateRepository27Params object
// with the ability to set a timeout on a request.
func NewCreateRepository27ParamsWithTimeout(timeout time.Duration) *CreateRepository27Params {
	return &CreateRepository27Params{
		timeout: timeout,
	}
}

// NewCreateRepository27ParamsWithContext creates a new CreateRepository27Params object
// with the ability to set a context for a request.
func NewCreateRepository27ParamsWithContext(ctx context.Context) *CreateRepository27Params {
	return &CreateRepository27Params{
		Context: ctx,
	}
}

// NewCreateRepository27ParamsWithHTTPClient creates a new CreateRepository27Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository27ParamsWithHTTPClient(client *http.Client) *CreateRepository27Params {
	return &CreateRepository27Params{
		HTTPClient: client,
	}
}

/*
CreateRepository27Params contains all the parameters to send to the API endpoint

	for the create repository 27 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository27Params struct {

	// Body.
	Body *models.PypiHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 27 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository27Params) WithDefaults() *CreateRepository27Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 27 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository27Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 27 params
func (o *CreateRepository27Params) WithTimeout(timeout time.Duration) *CreateRepository27Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 27 params
func (o *CreateRepository27Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 27 params
func (o *CreateRepository27Params) WithContext(ctx context.Context) *CreateRepository27Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 27 params
func (o *CreateRepository27Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 27 params
func (o *CreateRepository27Params) WithHTTPClient(client *http.Client) *CreateRepository27Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 27 params
func (o *CreateRepository27Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 27 params
func (o *CreateRepository27Params) WithBody(body *models.PypiHostedRepositoryAPIRequest) *CreateRepository27Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 27 params
func (o *CreateRepository27Params) SetBody(body *models.PypiHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository27Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
