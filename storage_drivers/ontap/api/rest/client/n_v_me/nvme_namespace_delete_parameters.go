// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewNvmeNamespaceDeleteParams creates a new NvmeNamespaceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeNamespaceDeleteParams() *NvmeNamespaceDeleteParams {
	return &NvmeNamespaceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeNamespaceDeleteParamsWithTimeout creates a new NvmeNamespaceDeleteParams object
// with the ability to set a timeout on a request.
func NewNvmeNamespaceDeleteParamsWithTimeout(timeout time.Duration) *NvmeNamespaceDeleteParams {
	return &NvmeNamespaceDeleteParams{
		timeout: timeout,
	}
}

// NewNvmeNamespaceDeleteParamsWithContext creates a new NvmeNamespaceDeleteParams object
// with the ability to set a context for a request.
func NewNvmeNamespaceDeleteParamsWithContext(ctx context.Context) *NvmeNamespaceDeleteParams {
	return &NvmeNamespaceDeleteParams{
		Context: ctx,
	}
}

// NewNvmeNamespaceDeleteParamsWithHTTPClient creates a new NvmeNamespaceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeNamespaceDeleteParamsWithHTTPClient(client *http.Client) *NvmeNamespaceDeleteParams {
	return &NvmeNamespaceDeleteParams{
		HTTPClient: client,
	}
}

/* NvmeNamespaceDeleteParams contains all the parameters to send to the API endpoint
   for the nvme namespace delete operation.

   Typically these are written to a http.Request.
*/
type NvmeNamespaceDeleteParams struct {

	/* AllowDeleteWhileMapped.

	     Allows deletion of a mapped NVMe namespace.
	A mapped NVMe namespace might be in use. Deleting a mapped namespace also deletes the namespace map and makes the data no longer available, possibly causing a disruption in the availability of data.
	**This parameter should be used with caution.**

	*/
	AllowDeleteWhileMappedQueryParameter *bool

	/* UUID.

	   The unique identifier of the NVMe namespace to delete.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme namespace delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeNamespaceDeleteParams) WithDefaults() *NvmeNamespaceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme namespace delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeNamespaceDeleteParams) SetDefaults() {
	var (
		allowDeleteWhileMappedQueryParameterDefault = bool(false)
	)

	val := NvmeNamespaceDeleteParams{
		AllowDeleteWhileMappedQueryParameter: &allowDeleteWhileMappedQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) WithTimeout(timeout time.Duration) *NvmeNamespaceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) WithContext(ctx context.Context) *NvmeNamespaceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) WithHTTPClient(client *http.Client) *NvmeNamespaceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeleteWhileMappedQueryParameter adds the allowDeleteWhileMapped to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) WithAllowDeleteWhileMappedQueryParameter(allowDeleteWhileMapped *bool) *NvmeNamespaceDeleteParams {
	o.SetAllowDeleteWhileMappedQueryParameter(allowDeleteWhileMapped)
	return o
}

// SetAllowDeleteWhileMappedQueryParameter adds the allowDeleteWhileMapped to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) SetAllowDeleteWhileMappedQueryParameter(allowDeleteWhileMapped *bool) {
	o.AllowDeleteWhileMappedQueryParameter = allowDeleteWhileMapped
}

// WithUUIDPathParameter adds the uuid to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) WithUUIDPathParameter(uuid string) *NvmeNamespaceDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the nvme namespace delete params
func (o *NvmeNamespaceDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeNamespaceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeleteWhileMappedQueryParameter != nil {

		// query param allow_delete_while_mapped
		var qrAllowDeleteWhileMapped bool

		if o.AllowDeleteWhileMappedQueryParameter != nil {
			qrAllowDeleteWhileMapped = *o.AllowDeleteWhileMappedQueryParameter
		}
		qAllowDeleteWhileMapped := swag.FormatBool(qrAllowDeleteWhileMapped)
		if qAllowDeleteWhileMapped != "" {

			if err := r.SetQueryParam("allow_delete_while_mapped", qAllowDeleteWhileMapped); err != nil {
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
