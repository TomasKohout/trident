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

// NewKerberosRealmCreateParams creates a new KerberosRealmCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKerberosRealmCreateParams() *KerberosRealmCreateParams {
	return &KerberosRealmCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKerberosRealmCreateParamsWithTimeout creates a new KerberosRealmCreateParams object
// with the ability to set a timeout on a request.
func NewKerberosRealmCreateParamsWithTimeout(timeout time.Duration) *KerberosRealmCreateParams {
	return &KerberosRealmCreateParams{
		timeout: timeout,
	}
}

// NewKerberosRealmCreateParamsWithContext creates a new KerberosRealmCreateParams object
// with the ability to set a context for a request.
func NewKerberosRealmCreateParamsWithContext(ctx context.Context) *KerberosRealmCreateParams {
	return &KerberosRealmCreateParams{
		Context: ctx,
	}
}

// NewKerberosRealmCreateParamsWithHTTPClient creates a new KerberosRealmCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewKerberosRealmCreateParamsWithHTTPClient(client *http.Client) *KerberosRealmCreateParams {
	return &KerberosRealmCreateParams{
		HTTPClient: client,
	}
}

/* KerberosRealmCreateParams contains all the parameters to send to the API endpoint
   for the kerberos realm create operation.

   Typically these are written to a http.Request.
*/
type KerberosRealmCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.KerberosRealm

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kerberos realm create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosRealmCreateParams) WithDefaults() *KerberosRealmCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kerberos realm create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosRealmCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := KerberosRealmCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the kerberos realm create params
func (o *KerberosRealmCreateParams) WithTimeout(timeout time.Duration) *KerberosRealmCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kerberos realm create params
func (o *KerberosRealmCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kerberos realm create params
func (o *KerberosRealmCreateParams) WithContext(ctx context.Context) *KerberosRealmCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kerberos realm create params
func (o *KerberosRealmCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kerberos realm create params
func (o *KerberosRealmCreateParams) WithHTTPClient(client *http.Client) *KerberosRealmCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kerberos realm create params
func (o *KerberosRealmCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the kerberos realm create params
func (o *KerberosRealmCreateParams) WithInfo(info *models.KerberosRealm) *KerberosRealmCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the kerberos realm create params
func (o *KerberosRealmCreateParams) SetInfo(info *models.KerberosRealm) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the kerberos realm create params
func (o *KerberosRealmCreateParams) WithReturnRecords(returnRecords *bool) *KerberosRealmCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the kerberos realm create params
func (o *KerberosRealmCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *KerberosRealmCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
