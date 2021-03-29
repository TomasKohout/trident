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

// NewClusterLdapCreateParams creates a new ClusterLdapCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterLdapCreateParams() *ClusterLdapCreateParams {
	return &ClusterLdapCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterLdapCreateParamsWithTimeout creates a new ClusterLdapCreateParams object
// with the ability to set a timeout on a request.
func NewClusterLdapCreateParamsWithTimeout(timeout time.Duration) *ClusterLdapCreateParams {
	return &ClusterLdapCreateParams{
		timeout: timeout,
	}
}

// NewClusterLdapCreateParamsWithContext creates a new ClusterLdapCreateParams object
// with the ability to set a context for a request.
func NewClusterLdapCreateParamsWithContext(ctx context.Context) *ClusterLdapCreateParams {
	return &ClusterLdapCreateParams{
		Context: ctx,
	}
}

// NewClusterLdapCreateParamsWithHTTPClient creates a new ClusterLdapCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterLdapCreateParamsWithHTTPClient(client *http.Client) *ClusterLdapCreateParams {
	return &ClusterLdapCreateParams{
		HTTPClient: client,
	}
}

/* ClusterLdapCreateParams contains all the parameters to send to the API endpoint
   for the cluster ldap create operation.

   Typically these are written to a http.Request.
*/
type ClusterLdapCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.ClusterLdap

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster ldap create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterLdapCreateParams) WithDefaults() *ClusterLdapCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster ldap create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterLdapCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := ClusterLdapCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cluster ldap create params
func (o *ClusterLdapCreateParams) WithTimeout(timeout time.Duration) *ClusterLdapCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster ldap create params
func (o *ClusterLdapCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster ldap create params
func (o *ClusterLdapCreateParams) WithContext(ctx context.Context) *ClusterLdapCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster ldap create params
func (o *ClusterLdapCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster ldap create params
func (o *ClusterLdapCreateParams) WithHTTPClient(client *http.Client) *ClusterLdapCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster ldap create params
func (o *ClusterLdapCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster ldap create params
func (o *ClusterLdapCreateParams) WithInfo(info *models.ClusterLdap) *ClusterLdapCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster ldap create params
func (o *ClusterLdapCreateParams) SetInfo(info *models.ClusterLdap) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the cluster ldap create params
func (o *ClusterLdapCreateParams) WithReturnRecords(returnRecords *bool) *ClusterLdapCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cluster ldap create params
func (o *ClusterLdapCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterLdapCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
