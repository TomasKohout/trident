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

// NewNodesCreateParams creates a new NodesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodesCreateParams() *NodesCreateParams {
	return &NodesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodesCreateParamsWithTimeout creates a new NodesCreateParams object
// with the ability to set a timeout on a request.
func NewNodesCreateParamsWithTimeout(timeout time.Duration) *NodesCreateParams {
	return &NodesCreateParams{
		timeout: timeout,
	}
}

// NewNodesCreateParamsWithContext creates a new NodesCreateParams object
// with the ability to set a context for a request.
func NewNodesCreateParamsWithContext(ctx context.Context) *NodesCreateParams {
	return &NodesCreateParams{
		Context: ctx,
	}
}

// NewNodesCreateParamsWithHTTPClient creates a new NodesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodesCreateParamsWithHTTPClient(client *http.Client) *NodesCreateParams {
	return &NodesCreateParams{
		HTTPClient: client,
	}
}

/* NodesCreateParams contains all the parameters to send to the API endpoint
   for the nodes create operation.

   Typically these are written to a http.Request.
*/
type NodesCreateParams struct {

	/* CreateRecommendedAggregates.

	   Creates aggregates based on an optimal layout recommended by the system.
	*/
	CreateRecommendedAggregatesQueryParameter *bool

	/* Info.

	   An object containing an array of nodes.
	*/
	Info *models.Node

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

// WithDefaults hydrates default values in the nodes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodesCreateParams) WithDefaults() *NodesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nodes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodesCreateParams) SetDefaults() {
	var (
		createRecommendedAggregatesQueryParameterDefault = bool(false)

		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := NodesCreateParams{
		CreateRecommendedAggregatesQueryParameter: &createRecommendedAggregatesQueryParameterDefault,
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nodes create params
func (o *NodesCreateParams) WithTimeout(timeout time.Duration) *NodesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes create params
func (o *NodesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes create params
func (o *NodesCreateParams) WithContext(ctx context.Context) *NodesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes create params
func (o *NodesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes create params
func (o *NodesCreateParams) WithHTTPClient(client *http.Client) *NodesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes create params
func (o *NodesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateRecommendedAggregatesQueryParameter adds the createRecommendedAggregates to the nodes create params
func (o *NodesCreateParams) WithCreateRecommendedAggregatesQueryParameter(createRecommendedAggregates *bool) *NodesCreateParams {
	o.SetCreateRecommendedAggregatesQueryParameter(createRecommendedAggregates)
	return o
}

// SetCreateRecommendedAggregatesQueryParameter adds the createRecommendedAggregates to the nodes create params
func (o *NodesCreateParams) SetCreateRecommendedAggregatesQueryParameter(createRecommendedAggregates *bool) {
	o.CreateRecommendedAggregatesQueryParameter = createRecommendedAggregates
}

// WithInfo adds the info to the nodes create params
func (o *NodesCreateParams) WithInfo(info *models.Node) *NodesCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nodes create params
func (o *NodesCreateParams) SetInfo(info *models.Node) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the nodes create params
func (o *NodesCreateParams) WithReturnRecords(returnRecords *bool) *NodesCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nodes create params
func (o *NodesCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the nodes create params
func (o *NodesCreateParams) WithReturnTimeout(returnTimeout *int64) *NodesCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the nodes create params
func (o *NodesCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *NodesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateRecommendedAggregatesQueryParameter != nil {

		// query param create_recommended_aggregates
		var qrCreateRecommendedAggregates bool

		if o.CreateRecommendedAggregatesQueryParameter != nil {
			qrCreateRecommendedAggregates = *o.CreateRecommendedAggregatesQueryParameter
		}
		qCreateRecommendedAggregates := swag.FormatBool(qrCreateRecommendedAggregates)
		if qCreateRecommendedAggregates != "" {

			if err := r.SetQueryParam("create_recommended_aggregates", qCreateRecommendedAggregates); err != nil {
				return err
			}
		}
	}
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
