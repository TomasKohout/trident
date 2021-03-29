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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewDNSModifyParams creates a new DNSModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDNSModifyParams() *DNSModifyParams {
	return &DNSModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDNSModifyParamsWithTimeout creates a new DNSModifyParams object
// with the ability to set a timeout on a request.
func NewDNSModifyParamsWithTimeout(timeout time.Duration) *DNSModifyParams {
	return &DNSModifyParams{
		timeout: timeout,
	}
}

// NewDNSModifyParamsWithContext creates a new DNSModifyParams object
// with the ability to set a context for a request.
func NewDNSModifyParamsWithContext(ctx context.Context) *DNSModifyParams {
	return &DNSModifyParams{
		Context: ctx,
	}
}

// NewDNSModifyParamsWithHTTPClient creates a new DNSModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDNSModifyParamsWithHTTPClient(client *http.Client) *DNSModifyParams {
	return &DNSModifyParams{
		HTTPClient: client,
	}
}

/* DNSModifyParams contains all the parameters to send to the API endpoint
   for the dns modify operation.

   Typically these are written to a http.Request.
*/
type DNSModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.DNS

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dns modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSModifyParams) WithDefaults() *DNSModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dns modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dns modify params
func (o *DNSModifyParams) WithTimeout(timeout time.Duration) *DNSModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dns modify params
func (o *DNSModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dns modify params
func (o *DNSModifyParams) WithContext(ctx context.Context) *DNSModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dns modify params
func (o *DNSModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dns modify params
func (o *DNSModifyParams) WithHTTPClient(client *http.Client) *DNSModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dns modify params
func (o *DNSModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the dns modify params
func (o *DNSModifyParams) WithInfo(info *models.DNS) *DNSModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the dns modify params
func (o *DNSModifyParams) SetInfo(info *models.DNS) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the dns modify params
func (o *DNSModifyParams) WithSvmUUID(svmUUID string) *DNSModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the dns modify params
func (o *DNSModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DNSModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
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
