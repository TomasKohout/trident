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
)

// NewVscanOnAccessGetParams creates a new VscanOnAccessGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanOnAccessGetParams() *VscanOnAccessGetParams {
	return &VscanOnAccessGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanOnAccessGetParamsWithTimeout creates a new VscanOnAccessGetParams object
// with the ability to set a timeout on a request.
func NewVscanOnAccessGetParamsWithTimeout(timeout time.Duration) *VscanOnAccessGetParams {
	return &VscanOnAccessGetParams{
		timeout: timeout,
	}
}

// NewVscanOnAccessGetParamsWithContext creates a new VscanOnAccessGetParams object
// with the ability to set a context for a request.
func NewVscanOnAccessGetParamsWithContext(ctx context.Context) *VscanOnAccessGetParams {
	return &VscanOnAccessGetParams{
		Context: ctx,
	}
}

// NewVscanOnAccessGetParamsWithHTTPClient creates a new VscanOnAccessGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanOnAccessGetParamsWithHTTPClient(client *http.Client) *VscanOnAccessGetParams {
	return &VscanOnAccessGetParams{
		HTTPClient: client,
	}
}

/* VscanOnAccessGetParams contains all the parameters to send to the API endpoint
   for the vscan on access get operation.

   Typically these are written to a http.Request.
*/
type VscanOnAccessGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	// Name.
	NamePathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan on access get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessGetParams) WithDefaults() *VscanOnAccessGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan on access get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vscan on access get params
func (o *VscanOnAccessGetParams) WithTimeout(timeout time.Duration) *VscanOnAccessGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan on access get params
func (o *VscanOnAccessGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan on access get params
func (o *VscanOnAccessGetParams) WithContext(ctx context.Context) *VscanOnAccessGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan on access get params
func (o *VscanOnAccessGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan on access get params
func (o *VscanOnAccessGetParams) WithHTTPClient(client *http.Client) *VscanOnAccessGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan on access get params
func (o *VscanOnAccessGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the vscan on access get params
func (o *VscanOnAccessGetParams) WithFields(fields []string) *VscanOnAccessGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the vscan on access get params
func (o *VscanOnAccessGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithNamePathParameter adds the name to the vscan on access get params
func (o *VscanOnAccessGetParams) WithNamePathParameter(name string) *VscanOnAccessGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the vscan on access get params
func (o *VscanOnAccessGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSvmUUID adds the svmUUID to the vscan on access get params
func (o *VscanOnAccessGetParams) WithSvmUUID(svmUUID string) *VscanOnAccessGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan on access get params
func (o *VscanOnAccessGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanOnAccessGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVscanOnAccessGet binds the parameter fields
func (o *VscanOnAccessGetParams) bindParamFields(formats strfmt.Registry) []string {
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
