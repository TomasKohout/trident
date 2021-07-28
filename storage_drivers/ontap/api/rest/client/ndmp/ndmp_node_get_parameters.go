// Code generated by go-swagger; DO NOT EDIT.

package ndmp

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

// NewNdmpNodeGetParams creates a new NdmpNodeGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpNodeGetParams() *NdmpNodeGetParams {
	return &NdmpNodeGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpNodeGetParamsWithTimeout creates a new NdmpNodeGetParams object
// with the ability to set a timeout on a request.
func NewNdmpNodeGetParamsWithTimeout(timeout time.Duration) *NdmpNodeGetParams {
	return &NdmpNodeGetParams{
		timeout: timeout,
	}
}

// NewNdmpNodeGetParamsWithContext creates a new NdmpNodeGetParams object
// with the ability to set a context for a request.
func NewNdmpNodeGetParamsWithContext(ctx context.Context) *NdmpNodeGetParams {
	return &NdmpNodeGetParams{
		Context: ctx,
	}
}

// NewNdmpNodeGetParamsWithHTTPClient creates a new NdmpNodeGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpNodeGetParamsWithHTTPClient(client *http.Client) *NdmpNodeGetParams {
	return &NdmpNodeGetParams{
		HTTPClient: client,
	}
}

/* NdmpNodeGetParams contains all the parameters to send to the API endpoint
   for the ndmp node get operation.

   Typically these are written to a http.Request.
*/
type NdmpNodeGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp node get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeGetParams) WithDefaults() *NdmpNodeGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp node get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ndmp node get params
func (o *NdmpNodeGetParams) WithTimeout(timeout time.Duration) *NdmpNodeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp node get params
func (o *NdmpNodeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp node get params
func (o *NdmpNodeGetParams) WithContext(ctx context.Context) *NdmpNodeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp node get params
func (o *NdmpNodeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp node get params
func (o *NdmpNodeGetParams) WithHTTPClient(client *http.Client) *NdmpNodeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp node get params
func (o *NdmpNodeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the ndmp node get params
func (o *NdmpNodeGetParams) WithFields(fields []string) *NdmpNodeGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ndmp node get params
func (o *NdmpNodeGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithNodeUUIDPathParameter adds the nodeUUID to the ndmp node get params
func (o *NdmpNodeGetParams) WithNodeUUIDPathParameter(nodeUUID string) *NdmpNodeGetParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the ndmp node get params
func (o *NdmpNodeGetParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpNodeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNdmpNodeGet binds the parameter fields
func (o *NdmpNodeGetParams) bindParamFields(formats strfmt.Registry) []string {
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