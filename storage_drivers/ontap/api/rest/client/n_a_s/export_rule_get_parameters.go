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

// NewExportRuleGetParams creates a new ExportRuleGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleGetParams() *ExportRuleGetParams {
	return &ExportRuleGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleGetParamsWithTimeout creates a new ExportRuleGetParams object
// with the ability to set a timeout on a request.
func NewExportRuleGetParamsWithTimeout(timeout time.Duration) *ExportRuleGetParams {
	return &ExportRuleGetParams{
		timeout: timeout,
	}
}

// NewExportRuleGetParamsWithContext creates a new ExportRuleGetParams object
// with the ability to set a context for a request.
func NewExportRuleGetParamsWithContext(ctx context.Context) *ExportRuleGetParams {
	return &ExportRuleGetParams{
		Context: ctx,
	}
}

// NewExportRuleGetParamsWithHTTPClient creates a new ExportRuleGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleGetParamsWithHTTPClient(client *http.Client) *ExportRuleGetParams {
	return &ExportRuleGetParams{
		HTTPClient: client,
	}
}

/* ExportRuleGetParams contains all the parameters to send to the API endpoint
   for the export rule get operation.

   Typically these are written to a http.Request.
*/
type ExportRuleGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Export Rule Index
	*/
	IndexPathParameter int64

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyIDPathParameter int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleGetParams) WithDefaults() *ExportRuleGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export rule get params
func (o *ExportRuleGetParams) WithTimeout(timeout time.Duration) *ExportRuleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule get params
func (o *ExportRuleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule get params
func (o *ExportRuleGetParams) WithContext(ctx context.Context) *ExportRuleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule get params
func (o *ExportRuleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule get params
func (o *ExportRuleGetParams) WithHTTPClient(client *http.Client) *ExportRuleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule get params
func (o *ExportRuleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the export rule get params
func (o *ExportRuleGetParams) WithFields(fields []string) *ExportRuleGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the export rule get params
func (o *ExportRuleGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndexPathParameter adds the index to the export rule get params
func (o *ExportRuleGetParams) WithIndexPathParameter(index int64) *ExportRuleGetParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the export rule get params
func (o *ExportRuleGetParams) SetIndexPathParameter(index int64) {
	o.IndexPathParameter = index
}

// WithPolicyIDPathParameter adds the policyID to the export rule get params
func (o *ExportRuleGetParams) WithPolicyIDPathParameter(policyID int64) *ExportRuleGetParams {
	o.SetPolicyIDPathParameter(policyID)
	return o
}

// SetPolicyIDPathParameter adds the policyId to the export rule get params
func (o *ExportRuleGetParams) SetPolicyIDPathParameter(policyID int64) {
	o.PolicyIDPathParameter = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.IndexPathParameter)); err != nil {
		return err
	}

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyIDPathParameter)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamExportRuleGet binds the parameter fields
func (o *ExportRuleGetParams) bindParamFields(formats strfmt.Registry) []string {
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
