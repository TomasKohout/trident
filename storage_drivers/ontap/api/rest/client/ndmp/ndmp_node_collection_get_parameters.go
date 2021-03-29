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

// NewNdmpNodeCollectionGetParams creates a new NdmpNodeCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpNodeCollectionGetParams() *NdmpNodeCollectionGetParams {
	return &NdmpNodeCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpNodeCollectionGetParamsWithTimeout creates a new NdmpNodeCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNdmpNodeCollectionGetParamsWithTimeout(timeout time.Duration) *NdmpNodeCollectionGetParams {
	return &NdmpNodeCollectionGetParams{
		timeout: timeout,
	}
}

// NewNdmpNodeCollectionGetParamsWithContext creates a new NdmpNodeCollectionGetParams object
// with the ability to set a context for a request.
func NewNdmpNodeCollectionGetParamsWithContext(ctx context.Context) *NdmpNodeCollectionGetParams {
	return &NdmpNodeCollectionGetParams{
		Context: ctx,
	}
}

// NewNdmpNodeCollectionGetParamsWithHTTPClient creates a new NdmpNodeCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpNodeCollectionGetParamsWithHTTPClient(client *http.Client) *NdmpNodeCollectionGetParams {
	return &NdmpNodeCollectionGetParams{
		HTTPClient: client,
	}
}

/* NdmpNodeCollectionGetParams contains all the parameters to send to the API endpoint
   for the ndmp node collection get operation.

   Typically these are written to a http.Request.
*/
type NdmpNodeCollectionGetParams struct {

	/* AuthenticationTypes.

	   Filter by authentication_types
	*/
	AuthenticationTypesQueryParameter *string

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* User.

	   Filter by user
	*/
	UserQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp node collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeCollectionGetParams) WithDefaults() *NdmpNodeCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp node collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NdmpNodeCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithTimeout(timeout time.Duration) *NdmpNodeCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithContext(ctx context.Context) *NdmpNodeCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithHTTPClient(client *http.Client) *NdmpNodeCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationTypesQueryParameter adds the authenticationTypes to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithAuthenticationTypesQueryParameter(authenticationTypes *string) *NdmpNodeCollectionGetParams {
	o.SetAuthenticationTypesQueryParameter(authenticationTypes)
	return o
}

// SetAuthenticationTypesQueryParameter adds the authenticationTypes to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetAuthenticationTypesQueryParameter(authenticationTypes *string) {
	o.AuthenticationTypesQueryParameter = authenticationTypes
}

// WithEnabledQueryParameter adds the enabled to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithEnabledQueryParameter(enabled *bool) *NdmpNodeCollectionGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithFields adds the fields to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithFields(fields []string) *NdmpNodeCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithMaxRecords(maxRecords *int64) *NdmpNodeCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *NdmpNodeCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *NdmpNodeCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderBy adds the orderBy to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithOrderBy(orderBy []string) *NdmpNodeCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithReturnRecords(returnRecords *bool) *NdmpNodeCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *NdmpNodeCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUserQueryParameter adds the user to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) WithUserQueryParameter(user *string) *NdmpNodeCollectionGetParams {
	o.SetUserQueryParameter(user)
	return o
}

// SetUserQueryParameter adds the user to the ndmp node collection get params
func (o *NdmpNodeCollectionGetParams) SetUserQueryParameter(user *string) {
	o.UserQueryParameter = user
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpNodeCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthenticationTypesQueryParameter != nil {

		// query param authentication_types
		var qrAuthenticationTypes string

		if o.AuthenticationTypesQueryParameter != nil {
			qrAuthenticationTypes = *o.AuthenticationTypesQueryParameter
		}
		qAuthenticationTypes := qrAuthenticationTypes
		if qAuthenticationTypes != "" {

			if err := r.SetQueryParam("authentication_types", qAuthenticationTypes); err != nil {
				return err
			}
		}
	}

	if o.EnabledQueryParameter != nil {

		// query param enabled
		var qrEnabled bool

		if o.EnabledQueryParameter != nil {
			qrEnabled = *o.EnabledQueryParameter
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.UserQueryParameter != nil {

		// query param user
		var qrUser string

		if o.UserQueryParameter != nil {
			qrUser = *o.UserQueryParameter
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNdmpNodeCollectionGet binds the parameter fields
func (o *NdmpNodeCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNdmpNodeCollectionGet binds the parameter order_by
func (o *NdmpNodeCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
