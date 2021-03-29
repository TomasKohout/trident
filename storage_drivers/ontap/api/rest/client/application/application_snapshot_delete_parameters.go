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
)

// NewApplicationSnapshotDeleteParams creates a new ApplicationSnapshotDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationSnapshotDeleteParams() *ApplicationSnapshotDeleteParams {
	return &ApplicationSnapshotDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationSnapshotDeleteParamsWithTimeout creates a new ApplicationSnapshotDeleteParams object
// with the ability to set a timeout on a request.
func NewApplicationSnapshotDeleteParamsWithTimeout(timeout time.Duration) *ApplicationSnapshotDeleteParams {
	return &ApplicationSnapshotDeleteParams{
		timeout: timeout,
	}
}

// NewApplicationSnapshotDeleteParamsWithContext creates a new ApplicationSnapshotDeleteParams object
// with the ability to set a context for a request.
func NewApplicationSnapshotDeleteParamsWithContext(ctx context.Context) *ApplicationSnapshotDeleteParams {
	return &ApplicationSnapshotDeleteParams{
		Context: ctx,
	}
}

// NewApplicationSnapshotDeleteParamsWithHTTPClient creates a new ApplicationSnapshotDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationSnapshotDeleteParamsWithHTTPClient(client *http.Client) *ApplicationSnapshotDeleteParams {
	return &ApplicationSnapshotDeleteParams{
		HTTPClient: client,
	}
}

/* ApplicationSnapshotDeleteParams contains all the parameters to send to the API endpoint
   for the application snapshot delete operation.

   Typically these are written to a http.Request.
*/
type ApplicationSnapshotDeleteParams struct {

	/* ApplicationUUID.

	   Application UUID
	*/
	ApplicationUUIDPathParameter string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Snapshot copy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application snapshot delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSnapshotDeleteParams) WithDefaults() *ApplicationSnapshotDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application snapshot delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSnapshotDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := ApplicationSnapshotDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) WithTimeout(timeout time.Duration) *ApplicationSnapshotDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) WithContext(ctx context.Context) *ApplicationSnapshotDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) WithHTTPClient(client *http.Client) *ApplicationSnapshotDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationUUIDPathParameter adds the applicationUUID to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) WithApplicationUUIDPathParameter(applicationUUID string) *ApplicationSnapshotDeleteParams {
	o.SetApplicationUUIDPathParameter(applicationUUID)
	return o
}

// SetApplicationUUIDPathParameter adds the applicationUuid to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) SetApplicationUUIDPathParameter(applicationUUID string) {
	o.ApplicationUUIDPathParameter = applicationUUID
}

// WithReturnTimeout adds the returnTimeout to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) WithReturnTimeout(returnTimeout *int64) *ApplicationSnapshotDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) WithUUIDPathParameter(uuid string) *ApplicationSnapshotDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the application snapshot delete params
func (o *ApplicationSnapshotDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationSnapshotDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.uuid
	if err := r.SetPathParam("application.uuid", o.ApplicationUUIDPathParameter); err != nil {
		return err
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
