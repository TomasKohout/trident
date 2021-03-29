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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewQtreeModifyParams creates a new QtreeModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQtreeModifyParams() *QtreeModifyParams {
	return &QtreeModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQtreeModifyParamsWithTimeout creates a new QtreeModifyParams object
// with the ability to set a timeout on a request.
func NewQtreeModifyParamsWithTimeout(timeout time.Duration) *QtreeModifyParams {
	return &QtreeModifyParams{
		timeout: timeout,
	}
}

// NewQtreeModifyParamsWithContext creates a new QtreeModifyParams object
// with the ability to set a context for a request.
func NewQtreeModifyParamsWithContext(ctx context.Context) *QtreeModifyParams {
	return &QtreeModifyParams{
		Context: ctx,
	}
}

// NewQtreeModifyParamsWithHTTPClient creates a new QtreeModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewQtreeModifyParamsWithHTTPClient(client *http.Client) *QtreeModifyParams {
	return &QtreeModifyParams{
		HTTPClient: client,
	}
}

/* QtreeModifyParams contains all the parameters to send to the API endpoint
   for the qtree modify operation.

   Typically these are written to a http.Request.
*/
type QtreeModifyParams struct {

	/* ID.

	   Qtree ID
	*/
	IDPathParameter string

	/* Info.

	   The new property values for the qtree.

	*/
	Info *models.Qtree

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the qtree modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QtreeModifyParams) WithDefaults() *QtreeModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the qtree modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QtreeModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := QtreeModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the qtree modify params
func (o *QtreeModifyParams) WithTimeout(timeout time.Duration) *QtreeModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the qtree modify params
func (o *QtreeModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the qtree modify params
func (o *QtreeModifyParams) WithContext(ctx context.Context) *QtreeModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the qtree modify params
func (o *QtreeModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the qtree modify params
func (o *QtreeModifyParams) WithHTTPClient(client *http.Client) *QtreeModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the qtree modify params
func (o *QtreeModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDPathParameter adds the id to the qtree modify params
func (o *QtreeModifyParams) WithIDPathParameter(id string) *QtreeModifyParams {
	o.SetIDPathParameter(id)
	return o
}

// SetIDPathParameter adds the id to the qtree modify params
func (o *QtreeModifyParams) SetIDPathParameter(id string) {
	o.IDPathParameter = id
}

// WithInfo adds the info to the qtree modify params
func (o *QtreeModifyParams) WithInfo(info *models.Qtree) *QtreeModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the qtree modify params
func (o *QtreeModifyParams) SetInfo(info *models.Qtree) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the qtree modify params
func (o *QtreeModifyParams) WithReturnTimeout(returnTimeout *int64) *QtreeModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the qtree modify params
func (o *QtreeModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithVolumeUUIDPathParameter adds the volumeUUID to the qtree modify params
func (o *QtreeModifyParams) WithVolumeUUIDPathParameter(volumeUUID string) *QtreeModifyParams {
	o.SetVolumeUUIDPathParameter(volumeUUID)
	return o
}

// SetVolumeUUIDPathParameter adds the volumeUuid to the qtree modify params
func (o *QtreeModifyParams) SetVolumeUUIDPathParameter(volumeUUID string) {
	o.VolumeUUIDPathParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *QtreeModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.IDPathParameter); err != nil {
		return err
	}
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

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
