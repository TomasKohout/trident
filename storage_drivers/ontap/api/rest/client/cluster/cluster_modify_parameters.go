// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewClusterModifyParams creates a new ClusterModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterModifyParams() *ClusterModifyParams {
	return &ClusterModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterModifyParamsWithTimeout creates a new ClusterModifyParams object
// with the ability to set a timeout on a request.
func NewClusterModifyParamsWithTimeout(timeout time.Duration) *ClusterModifyParams {
	return &ClusterModifyParams{
		timeout: timeout,
	}
}

// NewClusterModifyParamsWithContext creates a new ClusterModifyParams object
// with the ability to set a context for a request.
func NewClusterModifyParamsWithContext(ctx context.Context) *ClusterModifyParams {
	return &ClusterModifyParams{
		Context: ctx,
	}
}

// NewClusterModifyParamsWithHTTPClient creates a new ClusterModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterModifyParamsWithHTTPClient(client *http.Client) *ClusterModifyParams {
	return &ClusterModifyParams{
		HTTPClient: client,
	}
}

/* ClusterModifyParams contains all the parameters to send to the API endpoint
   for the cluster modify operation.

   Typically these are written to a http.Request.
*/
type ClusterModifyParams struct {

	/* Info.

	   Cluster information
	*/
	Info *models.Cluster

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

// WithDefaults hydrates default values in the cluster modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterModifyParams) WithDefaults() *ClusterModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterModifyParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := ClusterModifyParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cluster modify params
func (o *ClusterModifyParams) WithTimeout(timeout time.Duration) *ClusterModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster modify params
func (o *ClusterModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster modify params
func (o *ClusterModifyParams) WithContext(ctx context.Context) *ClusterModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster modify params
func (o *ClusterModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster modify params
func (o *ClusterModifyParams) WithHTTPClient(client *http.Client) *ClusterModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster modify params
func (o *ClusterModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster modify params
func (o *ClusterModifyParams) WithInfo(info *models.Cluster) *ClusterModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster modify params
func (o *ClusterModifyParams) SetInfo(info *models.Cluster) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the cluster modify params
func (o *ClusterModifyParams) WithReturnRecords(returnRecords *bool) *ClusterModifyParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cluster modify params
func (o *ClusterModifyParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cluster modify params
func (o *ClusterModifyParams) WithReturnTimeout(returnTimeout *int64) *ClusterModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cluster modify params
func (o *ClusterModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
