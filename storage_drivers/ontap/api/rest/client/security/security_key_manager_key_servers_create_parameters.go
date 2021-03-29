// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityKeyManagerKeyServersCreateParams creates a new SecurityKeyManagerKeyServersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerKeyServersCreateParams() *SecurityKeyManagerKeyServersCreateParams {
	return &SecurityKeyManagerKeyServersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerKeyServersCreateParamsWithTimeout creates a new SecurityKeyManagerKeyServersCreateParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerKeyServersCreateParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersCreateParams {
	return &SecurityKeyManagerKeyServersCreateParams{
		timeout: timeout,
	}
}

// NewSecurityKeyManagerKeyServersCreateParamsWithContext creates a new SecurityKeyManagerKeyServersCreateParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerKeyServersCreateParamsWithContext(ctx context.Context) *SecurityKeyManagerKeyServersCreateParams {
	return &SecurityKeyManagerKeyServersCreateParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerKeyServersCreateParamsWithHTTPClient creates a new SecurityKeyManagerKeyServersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerKeyServersCreateParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersCreateParams {
	return &SecurityKeyManagerKeyServersCreateParams{
		HTTPClient: client,
	}
}

/* SecurityKeyManagerKeyServersCreateParams contains all the parameters to send to the API endpoint
   for the security key manager key servers create operation.

   Typically these are written to a http.Request.
*/
type SecurityKeyManagerKeyServersCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.KeyServer

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* UUID.

	   External key manager UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security key manager key servers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersCreateParams) WithDefaults() *SecurityKeyManagerKeyServersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager key servers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SecurityKeyManagerKeyServersCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) WithTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) WithContext(ctx context.Context) *SecurityKeyManagerKeyServersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) WithInfo(info *models.KeyServer) *SecurityKeyManagerKeyServersCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) SetInfo(info *models.KeyServer) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) WithReturnRecords(returnRecords *bool) *SecurityKeyManagerKeyServersCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithUUIDPathParameter adds the uuid to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) WithUUIDPathParameter(uuid string) *SecurityKeyManagerKeyServersCreateParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the security key manager key servers create params
func (o *SecurityKeyManagerKeyServersCreateParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerKeyServersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
