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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSecurityAuditModifyParams creates a new SecurityAuditModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityAuditModifyParams() *SecurityAuditModifyParams {
	return &SecurityAuditModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityAuditModifyParamsWithTimeout creates a new SecurityAuditModifyParams object
// with the ability to set a timeout on a request.
func NewSecurityAuditModifyParamsWithTimeout(timeout time.Duration) *SecurityAuditModifyParams {
	return &SecurityAuditModifyParams{
		timeout: timeout,
	}
}

// NewSecurityAuditModifyParamsWithContext creates a new SecurityAuditModifyParams object
// with the ability to set a context for a request.
func NewSecurityAuditModifyParamsWithContext(ctx context.Context) *SecurityAuditModifyParams {
	return &SecurityAuditModifyParams{
		Context: ctx,
	}
}

// NewSecurityAuditModifyParamsWithHTTPClient creates a new SecurityAuditModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityAuditModifyParamsWithHTTPClient(client *http.Client) *SecurityAuditModifyParams {
	return &SecurityAuditModifyParams{
		HTTPClient: client,
	}
}

/* SecurityAuditModifyParams contains all the parameters to send to the API endpoint
   for the security audit modify operation.

   Typically these are written to a http.Request.
*/
type SecurityAuditModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SecurityAudit

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security audit modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityAuditModifyParams) WithDefaults() *SecurityAuditModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security audit modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityAuditModifyParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := SecurityAuditModifyParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security audit modify params
func (o *SecurityAuditModifyParams) WithTimeout(timeout time.Duration) *SecurityAuditModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security audit modify params
func (o *SecurityAuditModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security audit modify params
func (o *SecurityAuditModifyParams) WithContext(ctx context.Context) *SecurityAuditModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security audit modify params
func (o *SecurityAuditModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security audit modify params
func (o *SecurityAuditModifyParams) WithHTTPClient(client *http.Client) *SecurityAuditModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security audit modify params
func (o *SecurityAuditModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security audit modify params
func (o *SecurityAuditModifyParams) WithInfo(info *models.SecurityAudit) *SecurityAuditModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security audit modify params
func (o *SecurityAuditModifyParams) SetInfo(info *models.SecurityAudit) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the security audit modify params
func (o *SecurityAuditModifyParams) WithReturnRecords(returnRecords *bool) *SecurityAuditModifyParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security audit modify params
func (o *SecurityAuditModifyParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security audit modify params
func (o *SecurityAuditModifyParams) WithReturnTimeout(returnTimeout *int64) *SecurityAuditModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security audit modify params
func (o *SecurityAuditModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityAuditModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
