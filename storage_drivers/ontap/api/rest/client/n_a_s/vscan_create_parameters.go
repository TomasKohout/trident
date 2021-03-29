// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewVscanCreateParams creates a new VscanCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanCreateParams() *VscanCreateParams {
	return &VscanCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanCreateParamsWithTimeout creates a new VscanCreateParams object
// with the ability to set a timeout on a request.
func NewVscanCreateParamsWithTimeout(timeout time.Duration) *VscanCreateParams {
	return &VscanCreateParams{
		timeout: timeout,
	}
}

// NewVscanCreateParamsWithContext creates a new VscanCreateParams object
// with the ability to set a context for a request.
func NewVscanCreateParamsWithContext(ctx context.Context) *VscanCreateParams {
	return &VscanCreateParams{
		Context: ctx,
	}
}

// NewVscanCreateParamsWithHTTPClient creates a new VscanCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanCreateParamsWithHTTPClient(client *http.Client) *VscanCreateParams {
	return &VscanCreateParams{
		HTTPClient: client,
	}
}

/* VscanCreateParams contains all the parameters to send to the API endpoint
   for the vscan create operation.

   Typically these are written to a http.Request.
*/
type VscanCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Vscan

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanCreateParams) WithDefaults() *VscanCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := VscanCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vscan create params
func (o *VscanCreateParams) WithTimeout(timeout time.Duration) *VscanCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan create params
func (o *VscanCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan create params
func (o *VscanCreateParams) WithContext(ctx context.Context) *VscanCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan create params
func (o *VscanCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan create params
func (o *VscanCreateParams) WithHTTPClient(client *http.Client) *VscanCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan create params
func (o *VscanCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the vscan create params
func (o *VscanCreateParams) WithInfo(info *models.Vscan) *VscanCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the vscan create params
func (o *VscanCreateParams) SetInfo(info *models.Vscan) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the vscan create params
func (o *VscanCreateParams) WithReturnRecords(returnRecords *bool) *VscanCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the vscan create params
func (o *VscanCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *VscanCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
