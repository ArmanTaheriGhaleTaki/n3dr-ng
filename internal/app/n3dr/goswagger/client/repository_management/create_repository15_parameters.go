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

// NewCreateRepository15Params creates a new CreateRepository15Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository15Params() *CreateRepository15Params {
	return &CreateRepository15Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository15ParamsWithTimeout creates a new CreateRepository15Params object
// with the ability to set a timeout on a request.
func NewCreateRepository15ParamsWithTimeout(timeout time.Duration) *CreateRepository15Params {
	return &CreateRepository15Params{
		timeout: timeout,
	}
}

// NewCreateRepository15ParamsWithContext creates a new CreateRepository15Params object
// with the ability to set a context for a request.
func NewCreateRepository15ParamsWithContext(ctx context.Context) *CreateRepository15Params {
	return &CreateRepository15Params{
		Context: ctx,
	}
}

// NewCreateRepository15ParamsWithHTTPClient creates a new CreateRepository15Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository15ParamsWithHTTPClient(client *http.Client) *CreateRepository15Params {
	return &CreateRepository15Params{
		HTTPClient: client,
	}
}

/*
CreateRepository15Params contains all the parameters to send to the API endpoint

	for the create repository 15 operation.

	Typically these are written to a http.Request.
*/
type CreateRepository15Params struct {

	// Body.
	Body *models.RubyGemsHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 15 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository15Params) WithDefaults() *CreateRepository15Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 15 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository15Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 15 params
func (o *CreateRepository15Params) WithTimeout(timeout time.Duration) *CreateRepository15Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 15 params
func (o *CreateRepository15Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 15 params
func (o *CreateRepository15Params) WithContext(ctx context.Context) *CreateRepository15Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 15 params
func (o *CreateRepository15Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 15 params
func (o *CreateRepository15Params) WithHTTPClient(client *http.Client) *CreateRepository15Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 15 params
func (o *CreateRepository15Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 15 params
func (o *CreateRepository15Params) WithBody(body *models.RubyGemsHostedRepositoryAPIRequest) *CreateRepository15Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 15 params
func (o *CreateRepository15Params) SetBody(body *models.RubyGemsHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository15Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
