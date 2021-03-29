// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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
	"github.com/go-openapi/swag"
)

// NewCifsSearchPathDeleteParams creates a new CifsSearchPathDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSearchPathDeleteParams() *CifsSearchPathDeleteParams {
	return &CifsSearchPathDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSearchPathDeleteParamsWithTimeout creates a new CifsSearchPathDeleteParams object
// with the ability to set a timeout on a request.
func NewCifsSearchPathDeleteParamsWithTimeout(timeout time.Duration) *CifsSearchPathDeleteParams {
	return &CifsSearchPathDeleteParams{
		timeout: timeout,
	}
}

// NewCifsSearchPathDeleteParamsWithContext creates a new CifsSearchPathDeleteParams object
// with the ability to set a context for a request.
func NewCifsSearchPathDeleteParamsWithContext(ctx context.Context) *CifsSearchPathDeleteParams {
	return &CifsSearchPathDeleteParams{
		Context: ctx,
	}
}

// NewCifsSearchPathDeleteParamsWithHTTPClient creates a new CifsSearchPathDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSearchPathDeleteParamsWithHTTPClient(client *http.Client) *CifsSearchPathDeleteParams {
	return &CifsSearchPathDeleteParams{
		HTTPClient: client,
	}
}

/* CifsSearchPathDeleteParams contains all the parameters to send to the API endpoint
   for the cifs search path delete operation.

   Typically these are written to a http.Request.
*/
type CifsSearchPathDeleteParams struct {

	/* Index.

	   Home directory search path index
	*/
	IndexPathParameter int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs search path delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSearchPathDeleteParams) WithDefaults() *CifsSearchPathDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs search path delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSearchPathDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) WithTimeout(timeout time.Duration) *CifsSearchPathDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) WithContext(ctx context.Context) *CifsSearchPathDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) WithHTTPClient(client *http.Client) *CifsSearchPathDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexPathParameter adds the index to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) WithIndexPathParameter(index int64) *CifsSearchPathDeleteParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) SetIndexPathParameter(index int64) {
	o.IndexPathParameter = index
}

// WithSvmUUID adds the svmUUID to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) WithSvmUUID(svmUUID string) *CifsSearchPathDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs search path delete params
func (o *CifsSearchPathDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSearchPathDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.IndexPathParameter)); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
