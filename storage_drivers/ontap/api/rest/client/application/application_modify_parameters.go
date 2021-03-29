// Code generated by go-swagger; DO NOT EDIT.

package application

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewApplicationModifyParams creates a new ApplicationModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationModifyParams() *ApplicationModifyParams {
	return &ApplicationModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationModifyParamsWithTimeout creates a new ApplicationModifyParams object
// with the ability to set a timeout on a request.
func NewApplicationModifyParamsWithTimeout(timeout time.Duration) *ApplicationModifyParams {
	return &ApplicationModifyParams{
		timeout: timeout,
	}
}

// NewApplicationModifyParamsWithContext creates a new ApplicationModifyParams object
// with the ability to set a context for a request.
func NewApplicationModifyParamsWithContext(ctx context.Context) *ApplicationModifyParams {
	return &ApplicationModifyParams{
		Context: ctx,
	}
}

// NewApplicationModifyParamsWithHTTPClient creates a new ApplicationModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationModifyParamsWithHTTPClient(client *http.Client) *ApplicationModifyParams {
	return &ApplicationModifyParams{
		HTTPClient: client,
	}
}

/* ApplicationModifyParams contains all the parameters to send to the API endpoint
   for the application modify operation.

   Typically these are written to a http.Request.
*/
type ApplicationModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Application

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Application UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationModifyParams) WithDefaults() *ApplicationModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := ApplicationModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the application modify params
func (o *ApplicationModifyParams) WithTimeout(timeout time.Duration) *ApplicationModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application modify params
func (o *ApplicationModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application modify params
func (o *ApplicationModifyParams) WithContext(ctx context.Context) *ApplicationModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application modify params
func (o *ApplicationModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application modify params
func (o *ApplicationModifyParams) WithHTTPClient(client *http.Client) *ApplicationModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application modify params
func (o *ApplicationModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the application modify params
func (o *ApplicationModifyParams) WithInfo(info *models.Application) *ApplicationModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the application modify params
func (o *ApplicationModifyParams) SetInfo(info *models.Application) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the application modify params
func (o *ApplicationModifyParams) WithReturnTimeout(returnTimeout *int64) *ApplicationModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the application modify params
func (o *ApplicationModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the application modify params
func (o *ApplicationModifyParams) WithUUIDPathParameter(uuid string) *ApplicationModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the application modify params
func (o *ApplicationModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
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
