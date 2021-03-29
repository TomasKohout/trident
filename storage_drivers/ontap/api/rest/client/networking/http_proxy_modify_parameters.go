// Code generated by go-swagger; DO NOT EDIT.

package networking

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewHTTPProxyModifyParams creates a new HTTPProxyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHTTPProxyModifyParams() *HTTPProxyModifyParams {
	return &HTTPProxyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPProxyModifyParamsWithTimeout creates a new HTTPProxyModifyParams object
// with the ability to set a timeout on a request.
func NewHTTPProxyModifyParamsWithTimeout(timeout time.Duration) *HTTPProxyModifyParams {
	return &HTTPProxyModifyParams{
		timeout: timeout,
	}
}

// NewHTTPProxyModifyParamsWithContext creates a new HTTPProxyModifyParams object
// with the ability to set a context for a request.
func NewHTTPProxyModifyParamsWithContext(ctx context.Context) *HTTPProxyModifyParams {
	return &HTTPProxyModifyParams{
		Context: ctx,
	}
}

// NewHTTPProxyModifyParamsWithHTTPClient creates a new HTTPProxyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewHTTPProxyModifyParamsWithHTTPClient(client *http.Client) *HTTPProxyModifyParams {
	return &HTTPProxyModifyParams{
		HTTPClient: client,
	}
}

/* HTTPProxyModifyParams contains all the parameters to send to the API endpoint
   for the http proxy modify operation.

   Typically these are written to a http.Request.
*/
type HTTPProxyModifyParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.NetworkHTTPProxy

	/* UUID.

	   HTTP proxy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the http proxy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyModifyParams) WithDefaults() *HTTPProxyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the http proxy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the http proxy modify params
func (o *HTTPProxyModifyParams) WithTimeout(timeout time.Duration) *HTTPProxyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the http proxy modify params
func (o *HTTPProxyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the http proxy modify params
func (o *HTTPProxyModifyParams) WithContext(ctx context.Context) *HTTPProxyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the http proxy modify params
func (o *HTTPProxyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the http proxy modify params
func (o *HTTPProxyModifyParams) WithHTTPClient(client *http.Client) *HTTPProxyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the http proxy modify params
func (o *HTTPProxyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the http proxy modify params
func (o *HTTPProxyModifyParams) WithInfo(info *models.NetworkHTTPProxy) *HTTPProxyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the http proxy modify params
func (o *HTTPProxyModifyParams) SetInfo(info *models.NetworkHTTPProxy) {
	o.Info = info
}

// WithUUIDPathParameter adds the uuid to the http proxy modify params
func (o *HTTPProxyModifyParams) WithUUIDPathParameter(uuid string) *HTTPProxyModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the http proxy modify params
func (o *HTTPProxyModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPProxyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
