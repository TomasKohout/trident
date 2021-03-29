// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewFcInterfaceGetParams creates a new FcInterfaceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcInterfaceGetParams() *FcInterfaceGetParams {
	return &FcInterfaceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcInterfaceGetParamsWithTimeout creates a new FcInterfaceGetParams object
// with the ability to set a timeout on a request.
func NewFcInterfaceGetParamsWithTimeout(timeout time.Duration) *FcInterfaceGetParams {
	return &FcInterfaceGetParams{
		timeout: timeout,
	}
}

// NewFcInterfaceGetParamsWithContext creates a new FcInterfaceGetParams object
// with the ability to set a context for a request.
func NewFcInterfaceGetParamsWithContext(ctx context.Context) *FcInterfaceGetParams {
	return &FcInterfaceGetParams{
		Context: ctx,
	}
}

// NewFcInterfaceGetParamsWithHTTPClient creates a new FcInterfaceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcInterfaceGetParamsWithHTTPClient(client *http.Client) *FcInterfaceGetParams {
	return &FcInterfaceGetParams{
		HTTPClient: client,
	}
}

/* FcInterfaceGetParams contains all the parameters to send to the API endpoint
   for the fc interface get operation.

   Typically these are written to a http.Request.
*/
type FcInterfaceGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   The unique identifier for the FC interface.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fc interface get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcInterfaceGetParams) WithDefaults() *FcInterfaceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fc interface get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcInterfaceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fc interface get params
func (o *FcInterfaceGetParams) WithTimeout(timeout time.Duration) *FcInterfaceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fc interface get params
func (o *FcInterfaceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fc interface get params
func (o *FcInterfaceGetParams) WithContext(ctx context.Context) *FcInterfaceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fc interface get params
func (o *FcInterfaceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fc interface get params
func (o *FcInterfaceGetParams) WithHTTPClient(client *http.Client) *FcInterfaceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fc interface get params
func (o *FcInterfaceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the fc interface get params
func (o *FcInterfaceGetParams) WithFields(fields []string) *FcInterfaceGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the fc interface get params
func (o *FcInterfaceGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUIDPathParameter adds the uuid to the fc interface get params
func (o *FcInterfaceGetParams) WithUUIDPathParameter(uuid string) *FcInterfaceGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the fc interface get params
func (o *FcInterfaceGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *FcInterfaceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamFcInterfaceGet binds the parameter fields
func (o *FcInterfaceGetParams) bindParamFields(formats strfmt.Registry) []string {
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
