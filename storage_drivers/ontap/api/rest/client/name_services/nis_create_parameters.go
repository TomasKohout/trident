// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewNisCreateParams creates a new NisCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNisCreateParams() *NisCreateParams {
	return &NisCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNisCreateParamsWithTimeout creates a new NisCreateParams object
// with the ability to set a timeout on a request.
func NewNisCreateParamsWithTimeout(timeout time.Duration) *NisCreateParams {
	return &NisCreateParams{
		timeout: timeout,
	}
}

// NewNisCreateParamsWithContext creates a new NisCreateParams object
// with the ability to set a context for a request.
func NewNisCreateParamsWithContext(ctx context.Context) *NisCreateParams {
	return &NisCreateParams{
		Context: ctx,
	}
}

// NewNisCreateParamsWithHTTPClient creates a new NisCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNisCreateParamsWithHTTPClient(client *http.Client) *NisCreateParams {
	return &NisCreateParams{
		HTTPClient: client,
	}
}

/* NisCreateParams contains all the parameters to send to the API endpoint
   for the nis create operation.

   Typically these are written to a http.Request.
*/
type NisCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.NisService

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nis create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NisCreateParams) WithDefaults() *NisCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nis create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NisCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NisCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nis create params
func (o *NisCreateParams) WithTimeout(timeout time.Duration) *NisCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nis create params
func (o *NisCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nis create params
func (o *NisCreateParams) WithContext(ctx context.Context) *NisCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nis create params
func (o *NisCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nis create params
func (o *NisCreateParams) WithHTTPClient(client *http.Client) *NisCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nis create params
func (o *NisCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the nis create params
func (o *NisCreateParams) WithInfo(info *models.NisService) *NisCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nis create params
func (o *NisCreateParams) SetInfo(info *models.NisService) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the nis create params
func (o *NisCreateParams) WithReturnRecords(returnRecords *bool) *NisCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nis create params
func (o *NisCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NisCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
