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
)

// NewClusterPeerGetParams creates a new ClusterPeerGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterPeerGetParams() *ClusterPeerGetParams {
	return &ClusterPeerGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterPeerGetParamsWithTimeout creates a new ClusterPeerGetParams object
// with the ability to set a timeout on a request.
func NewClusterPeerGetParamsWithTimeout(timeout time.Duration) *ClusterPeerGetParams {
	return &ClusterPeerGetParams{
		timeout: timeout,
	}
}

// NewClusterPeerGetParamsWithContext creates a new ClusterPeerGetParams object
// with the ability to set a context for a request.
func NewClusterPeerGetParamsWithContext(ctx context.Context) *ClusterPeerGetParams {
	return &ClusterPeerGetParams{
		Context: ctx,
	}
}

// NewClusterPeerGetParamsWithHTTPClient creates a new ClusterPeerGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterPeerGetParamsWithHTTPClient(client *http.Client) *ClusterPeerGetParams {
	return &ClusterPeerGetParams{
		HTTPClient: client,
	}
}

/* ClusterPeerGetParams contains all the parameters to send to the API endpoint
   for the cluster peer get operation.

   Typically these are written to a http.Request.
*/
type ClusterPeerGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Cluster peer relationship UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster peer get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterPeerGetParams) WithDefaults() *ClusterPeerGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster peer get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterPeerGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster peer get params
func (o *ClusterPeerGetParams) WithTimeout(timeout time.Duration) *ClusterPeerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster peer get params
func (o *ClusterPeerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster peer get params
func (o *ClusterPeerGetParams) WithContext(ctx context.Context) *ClusterPeerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster peer get params
func (o *ClusterPeerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster peer get params
func (o *ClusterPeerGetParams) WithHTTPClient(client *http.Client) *ClusterPeerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster peer get params
func (o *ClusterPeerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cluster peer get params
func (o *ClusterPeerGetParams) WithFields(fields []string) *ClusterPeerGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cluster peer get params
func (o *ClusterPeerGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUIDPathParameter adds the uuid to the cluster peer get params
func (o *ClusterPeerGetParams) WithUUIDPathParameter(uuid string) *ClusterPeerGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the cluster peer get params
func (o *ClusterPeerGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterPeerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
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

// bindParamClusterPeerGet binds the parameter fields
func (o *ClusterPeerGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
