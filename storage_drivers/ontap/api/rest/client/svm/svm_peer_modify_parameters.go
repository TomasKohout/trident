// Code generated by go-swagger; DO NOT EDIT.

package svm

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

// NewSvmPeerModifyParams creates a new SvmPeerModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmPeerModifyParams() *SvmPeerModifyParams {
	return &SvmPeerModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmPeerModifyParamsWithTimeout creates a new SvmPeerModifyParams object
// with the ability to set a timeout on a request.
func NewSvmPeerModifyParamsWithTimeout(timeout time.Duration) *SvmPeerModifyParams {
	return &SvmPeerModifyParams{
		timeout: timeout,
	}
}

// NewSvmPeerModifyParamsWithContext creates a new SvmPeerModifyParams object
// with the ability to set a context for a request.
func NewSvmPeerModifyParamsWithContext(ctx context.Context) *SvmPeerModifyParams {
	return &SvmPeerModifyParams{
		Context: ctx,
	}
}

// NewSvmPeerModifyParamsWithHTTPClient creates a new SvmPeerModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmPeerModifyParamsWithHTTPClient(client *http.Client) *SvmPeerModifyParams {
	return &SvmPeerModifyParams{
		HTTPClient: client,
	}
}

/* SvmPeerModifyParams contains all the parameters to send to the API endpoint
   for the svm peer modify operation.

   Typically these are written to a http.Request.
*/
type SvmPeerModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SvmPeer

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   SVM peer relationship UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm peer modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerModifyParams) WithDefaults() *SvmPeerModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm peer modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SvmPeerModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm peer modify params
func (o *SvmPeerModifyParams) WithTimeout(timeout time.Duration) *SvmPeerModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm peer modify params
func (o *SvmPeerModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm peer modify params
func (o *SvmPeerModifyParams) WithContext(ctx context.Context) *SvmPeerModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm peer modify params
func (o *SvmPeerModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm peer modify params
func (o *SvmPeerModifyParams) WithHTTPClient(client *http.Client) *SvmPeerModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm peer modify params
func (o *SvmPeerModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the svm peer modify params
func (o *SvmPeerModifyParams) WithInfo(info *models.SvmPeer) *SvmPeerModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm peer modify params
func (o *SvmPeerModifyParams) SetInfo(info *models.SvmPeer) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the svm peer modify params
func (o *SvmPeerModifyParams) WithReturnTimeout(returnTimeout *int64) *SvmPeerModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm peer modify params
func (o *SvmPeerModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the svm peer modify params
func (o *SvmPeerModifyParams) WithUUIDPathParameter(uuid string) *SvmPeerModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the svm peer modify params
func (o *SvmPeerModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmPeerModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
