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

// NewUpdateRepository30Params creates a new UpdateRepository30Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepository30Params() *UpdateRepository30Params {
	return &UpdateRepository30Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepository30ParamsWithTimeout creates a new UpdateRepository30Params object
// with the ability to set a timeout on a request.
func NewUpdateRepository30ParamsWithTimeout(timeout time.Duration) *UpdateRepository30Params {
	return &UpdateRepository30Params{
		timeout: timeout,
	}
}

// NewUpdateRepository30ParamsWithContext creates a new UpdateRepository30Params object
// with the ability to set a context for a request.
func NewUpdateRepository30ParamsWithContext(ctx context.Context) *UpdateRepository30Params {
	return &UpdateRepository30Params{
		Context: ctx,
	}
}

// NewUpdateRepository30ParamsWithHTTPClient creates a new UpdateRepository30Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepository30ParamsWithHTTPClient(client *http.Client) *UpdateRepository30Params {
	return &UpdateRepository30Params{
		HTTPClient: client,
	}
}

/*
UpdateRepository30Params contains all the parameters to send to the API endpoint

	for the update repository 30 operation.

	Typically these are written to a http.Request.
*/
type UpdateRepository30Params struct {

	// Body.
	Body *models.ConanProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 30 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository30Params) WithDefaults() *UpdateRepository30Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 30 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository30Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 30 params
func (o *UpdateRepository30Params) WithTimeout(timeout time.Duration) *UpdateRepository30Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 30 params
func (o *UpdateRepository30Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 30 params
func (o *UpdateRepository30Params) WithContext(ctx context.Context) *UpdateRepository30Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 30 params
func (o *UpdateRepository30Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 30 params
func (o *UpdateRepository30Params) WithHTTPClient(client *http.Client) *UpdateRepository30Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 30 params
func (o *UpdateRepository30Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 30 params
func (o *UpdateRepository30Params) WithBody(body *models.ConanProxyRepositoryAPIRequest) *UpdateRepository30Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 30 params
func (o *UpdateRepository30Params) SetBody(body *models.ConanProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 30 params
func (o *UpdateRepository30Params) WithRepositoryName(repositoryName string) *UpdateRepository30Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 30 params
func (o *UpdateRepository30Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepository30Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
