// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIgroupModifyParams creates a new IgroupModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIgroupModifyParams() *IgroupModifyParams {
	return &IgroupModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIgroupModifyParamsWithTimeout creates a new IgroupModifyParams object
// with the ability to set a timeout on a request.
func NewIgroupModifyParamsWithTimeout(timeout time.Duration) *IgroupModifyParams {
	return &IgroupModifyParams{
		timeout: timeout,
	}
}

// NewIgroupModifyParamsWithContext creates a new IgroupModifyParams object
// with the ability to set a context for a request.
func NewIgroupModifyParamsWithContext(ctx context.Context) *IgroupModifyParams {
	return &IgroupModifyParams{
		Context: ctx,
	}
}

// NewIgroupModifyParamsWithHTTPClient creates a new IgroupModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewIgroupModifyParamsWithHTTPClient(client *http.Client) *IgroupModifyParams {
	return &IgroupModifyParams{
		HTTPClient: client,
	}
}

/* IgroupModifyParams contains all the parameters to send to the API endpoint
   for the igroup modify operation.

   Typically these are written to a http.Request.
*/
type IgroupModifyParams struct {

	/* Info.

	   The new property values for the initiator group.

	*/
	Info *models.Igroup

	/* UUID.

	   The unique identifier of the initiator group.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the igroup modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupModifyParams) WithDefaults() *IgroupModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the igroup modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the igroup modify params
func (o *IgroupModifyParams) WithTimeout(timeout time.Duration) *IgroupModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the igroup modify params
func (o *IgroupModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the igroup modify params
func (o *IgroupModifyParams) WithContext(ctx context.Context) *IgroupModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the igroup modify params
func (o *IgroupModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the igroup modify params
func (o *IgroupModifyParams) WithHTTPClient(client *http.Client) *IgroupModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the igroup modify params
func (o *IgroupModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the igroup modify params
func (o *IgroupModifyParams) WithInfo(info *models.Igroup) *IgroupModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the igroup modify params
func (o *IgroupModifyParams) SetInfo(info *models.Igroup) {
	o.Info = info
}

// WithUUIDPathParameter adds the uuid to the igroup modify params
func (o *IgroupModifyParams) WithUUIDPathParameter(uuid string) *IgroupModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the igroup modify params
func (o *IgroupModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IgroupModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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
