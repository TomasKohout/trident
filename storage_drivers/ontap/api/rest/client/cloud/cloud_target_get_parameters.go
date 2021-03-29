// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewCloudTargetGetParams creates a new CloudTargetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudTargetGetParams() *CloudTargetGetParams {
	return &CloudTargetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudTargetGetParamsWithTimeout creates a new CloudTargetGetParams object
// with the ability to set a timeout on a request.
func NewCloudTargetGetParamsWithTimeout(timeout time.Duration) *CloudTargetGetParams {
	return &CloudTargetGetParams{
		timeout: timeout,
	}
}

// NewCloudTargetGetParamsWithContext creates a new CloudTargetGetParams object
// with the ability to set a context for a request.
func NewCloudTargetGetParamsWithContext(ctx context.Context) *CloudTargetGetParams {
	return &CloudTargetGetParams{
		Context: ctx,
	}
}

// NewCloudTargetGetParamsWithHTTPClient creates a new CloudTargetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudTargetGetParamsWithHTTPClient(client *http.Client) *CloudTargetGetParams {
	return &CloudTargetGetParams{
		HTTPClient: client,
	}
}

/* CloudTargetGetParams contains all the parameters to send to the API endpoint
   for the cloud target get operation.

   Typically these are written to a http.Request.
*/
type CloudTargetGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Cloud target UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud target get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetGetParams) WithDefaults() *CloudTargetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud target get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud target get params
func (o *CloudTargetGetParams) WithTimeout(timeout time.Duration) *CloudTargetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud target get params
func (o *CloudTargetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud target get params
func (o *CloudTargetGetParams) WithContext(ctx context.Context) *CloudTargetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud target get params
func (o *CloudTargetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud target get params
func (o *CloudTargetGetParams) WithHTTPClient(client *http.Client) *CloudTargetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud target get params
func (o *CloudTargetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cloud target get params
func (o *CloudTargetGetParams) WithFields(fields []string) *CloudTargetGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cloud target get params
func (o *CloudTargetGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUIDPathParameter adds the uuid to the cloud target get params
func (o *CloudTargetGetParams) WithUUIDPathParameter(uuid string) *CloudTargetGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the cloud target get params
func (o *CloudTargetGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *CloudTargetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamCloudTargetGet binds the parameter fields
func (o *CloudTargetGetParams) bindParamFields(formats strfmt.Registry) []string {
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
