// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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
)

// NewNvmeSubsystemHostDeleteParams creates a new NvmeSubsystemHostDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemHostDeleteParams() *NvmeSubsystemHostDeleteParams {
	return &NvmeSubsystemHostDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemHostDeleteParamsWithTimeout creates a new NvmeSubsystemHostDeleteParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemHostDeleteParamsWithTimeout(timeout time.Duration) *NvmeSubsystemHostDeleteParams {
	return &NvmeSubsystemHostDeleteParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemHostDeleteParamsWithContext creates a new NvmeSubsystemHostDeleteParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemHostDeleteParamsWithContext(ctx context.Context) *NvmeSubsystemHostDeleteParams {
	return &NvmeSubsystemHostDeleteParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemHostDeleteParamsWithHTTPClient creates a new NvmeSubsystemHostDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemHostDeleteParamsWithHTTPClient(client *http.Client) *NvmeSubsystemHostDeleteParams {
	return &NvmeSubsystemHostDeleteParams{
		HTTPClient: client,
	}
}

/* NvmeSubsystemHostDeleteParams contains all the parameters to send to the API endpoint
   for the nvme subsystem host delete operation.

   Typically these are written to a http.Request.
*/
type NvmeSubsystemHostDeleteParams struct {

	/* Nqn.

	   The NVMe qualified name (NQN) used to identify the NVMe subsystem host.

	*/
	NqnPathParameter string

	/* SubsystemUUID.

	   The unique identifier of the NVMe subsystem.

	*/
	SubsystemUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem host delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostDeleteParams) WithDefaults() *NvmeSubsystemHostDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem host delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) WithTimeout(timeout time.Duration) *NvmeSubsystemHostDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) WithContext(ctx context.Context) *NvmeSubsystemHostDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) WithHTTPClient(client *http.Client) *NvmeSubsystemHostDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNqnPathParameter adds the nqn to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) WithNqnPathParameter(nqn string) *NvmeSubsystemHostDeleteParams {
	o.SetNqnPathParameter(nqn)
	return o
}

// SetNqnPathParameter adds the nqn to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) SetNqnPathParameter(nqn string) {
	o.NqnPathParameter = nqn
}

// WithSubsystemUUIDPathParameter adds the subsystemUUID to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) WithSubsystemUUIDPathParameter(subsystemUUID string) *NvmeSubsystemHostDeleteParams {
	o.SetSubsystemUUIDPathParameter(subsystemUUID)
	return o
}

// SetSubsystemUUIDPathParameter adds the subsystemUuid to the nvme subsystem host delete params
func (o *NvmeSubsystemHostDeleteParams) SetSubsystemUUIDPathParameter(subsystemUUID string) {
	o.SubsystemUUIDPathParameter = subsystemUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemHostDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nqn
	if err := r.SetPathParam("nqn", o.NqnPathParameter); err != nil {
		return err
	}

	// path param subsystem.uuid
	if err := r.SetPathParam("subsystem.uuid", o.SubsystemUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
