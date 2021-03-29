// Code generated by go-swagger; DO NOT EDIT.

package security

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
)

// NewClusterAccountAdProxyGetParams creates a new ClusterAccountAdProxyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterAccountAdProxyGetParams() *ClusterAccountAdProxyGetParams {
	return &ClusterAccountAdProxyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterAccountAdProxyGetParamsWithTimeout creates a new ClusterAccountAdProxyGetParams object
// with the ability to set a timeout on a request.
func NewClusterAccountAdProxyGetParamsWithTimeout(timeout time.Duration) *ClusterAccountAdProxyGetParams {
	return &ClusterAccountAdProxyGetParams{
		timeout: timeout,
	}
}

// NewClusterAccountAdProxyGetParamsWithContext creates a new ClusterAccountAdProxyGetParams object
// with the ability to set a context for a request.
func NewClusterAccountAdProxyGetParamsWithContext(ctx context.Context) *ClusterAccountAdProxyGetParams {
	return &ClusterAccountAdProxyGetParams{
		Context: ctx,
	}
}

// NewClusterAccountAdProxyGetParamsWithHTTPClient creates a new ClusterAccountAdProxyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterAccountAdProxyGetParamsWithHTTPClient(client *http.Client) *ClusterAccountAdProxyGetParams {
	return &ClusterAccountAdProxyGetParams{
		HTTPClient: client,
	}
}

/* ClusterAccountAdProxyGetParams contains all the parameters to send to the API endpoint
   for the cluster account ad proxy get operation.

   Typically these are written to a http.Request.
*/
type ClusterAccountAdProxyGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster account ad proxy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyGetParams) WithDefaults() *ClusterAccountAdProxyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster account ad proxy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster account ad proxy get params
func (o *ClusterAccountAdProxyGetParams) WithTimeout(timeout time.Duration) *ClusterAccountAdProxyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster account ad proxy get params
func (o *ClusterAccountAdProxyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster account ad proxy get params
func (o *ClusterAccountAdProxyGetParams) WithContext(ctx context.Context) *ClusterAccountAdProxyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster account ad proxy get params
func (o *ClusterAccountAdProxyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster account ad proxy get params
func (o *ClusterAccountAdProxyGetParams) WithHTTPClient(client *http.Client) *ClusterAccountAdProxyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster account ad proxy get params
func (o *ClusterAccountAdProxyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterAccountAdProxyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
