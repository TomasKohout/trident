// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewQuotaRuleDeleteParams creates a new QuotaRuleDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuotaRuleDeleteParams() *QuotaRuleDeleteParams {
	return &QuotaRuleDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuotaRuleDeleteParamsWithTimeout creates a new QuotaRuleDeleteParams object
// with the ability to set a timeout on a request.
func NewQuotaRuleDeleteParamsWithTimeout(timeout time.Duration) *QuotaRuleDeleteParams {
	return &QuotaRuleDeleteParams{
		timeout: timeout,
	}
}

// NewQuotaRuleDeleteParamsWithContext creates a new QuotaRuleDeleteParams object
// with the ability to set a context for a request.
func NewQuotaRuleDeleteParamsWithContext(ctx context.Context) *QuotaRuleDeleteParams {
	return &QuotaRuleDeleteParams{
		Context: ctx,
	}
}

// NewQuotaRuleDeleteParamsWithHTTPClient creates a new QuotaRuleDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuotaRuleDeleteParamsWithHTTPClient(client *http.Client) *QuotaRuleDeleteParams {
	return &QuotaRuleDeleteParams{
		HTTPClient: client,
	}
}

/* QuotaRuleDeleteParams contains all the parameters to send to the API endpoint
   for the quota rule delete operation.

   Typically these are written to a http.Request.
*/
type QuotaRuleDeleteParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Rule UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quota rule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaRuleDeleteParams) WithDefaults() *QuotaRuleDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quota rule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaRuleDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := QuotaRuleDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the quota rule delete params
func (o *QuotaRuleDeleteParams) WithTimeout(timeout time.Duration) *QuotaRuleDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quota rule delete params
func (o *QuotaRuleDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quota rule delete params
func (o *QuotaRuleDeleteParams) WithContext(ctx context.Context) *QuotaRuleDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quota rule delete params
func (o *QuotaRuleDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quota rule delete params
func (o *QuotaRuleDeleteParams) WithHTTPClient(client *http.Client) *QuotaRuleDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quota rule delete params
func (o *QuotaRuleDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeout adds the returnTimeout to the quota rule delete params
func (o *QuotaRuleDeleteParams) WithReturnTimeout(returnTimeout *int64) *QuotaRuleDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the quota rule delete params
func (o *QuotaRuleDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the quota rule delete params
func (o *QuotaRuleDeleteParams) WithUUIDPathParameter(uuid string) *QuotaRuleDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the quota rule delete params
func (o *QuotaRuleDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *QuotaRuleDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
