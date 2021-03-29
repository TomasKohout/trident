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

// NewNetworkIPRoutesGetParams creates a new NetworkIPRoutesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPRoutesGetParams() *NetworkIPRoutesGetParams {
	return &NetworkIPRoutesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPRoutesGetParamsWithTimeout creates a new NetworkIPRoutesGetParams object
// with the ability to set a timeout on a request.
func NewNetworkIPRoutesGetParamsWithTimeout(timeout time.Duration) *NetworkIPRoutesGetParams {
	return &NetworkIPRoutesGetParams{
		timeout: timeout,
	}
}

// NewNetworkIPRoutesGetParamsWithContext creates a new NetworkIPRoutesGetParams object
// with the ability to set a context for a request.
func NewNetworkIPRoutesGetParamsWithContext(ctx context.Context) *NetworkIPRoutesGetParams {
	return &NetworkIPRoutesGetParams{
		Context: ctx,
	}
}

// NewNetworkIPRoutesGetParamsWithHTTPClient creates a new NetworkIPRoutesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPRoutesGetParamsWithHTTPClient(client *http.Client) *NetworkIPRoutesGetParams {
	return &NetworkIPRoutesGetParams{
		HTTPClient: client,
	}
}

/* NetworkIPRoutesGetParams contains all the parameters to send to the API endpoint
   for the network ip routes get operation.

   Typically these are written to a http.Request.
*/
type NetworkIPRoutesGetParams struct {

	/* DestinationAddress.

	   Filter by destination.address
	*/
	DestinationAddressQueryParameter *string

	/* DestinationFamily.

	   Filter by destination.family
	*/
	DestinationFamilyQueryParameter *string

	/* DestinationNetmask.

	   Filter by destination.netmask
	*/
	DestinationNetmaskQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Gateway.

	   Filter by gateway
	*/
	GatewayQueryParameter *string

	/* IpspaceName.

	   Filter by ipspace.name
	*/
	IpspaceNameQueryParameter *string

	/* IpspaceUUID.

	   Filter by ipspace.uuid
	*/
	IpspaceUUIDQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

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

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip routes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRoutesGetParams) WithDefaults() *NetworkIPRoutesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip routes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRoutesGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NetworkIPRoutesGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithTimeout(timeout time.Duration) *NetworkIPRoutesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithContext(ctx context.Context) *NetworkIPRoutesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithHTTPClient(client *http.Client) *NetworkIPRoutesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationAddressQueryParameter adds the destinationAddress to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithDestinationAddressQueryParameter(destinationAddress *string) *NetworkIPRoutesGetParams {
	o.SetDestinationAddressQueryParameter(destinationAddress)
	return o
}

// SetDestinationAddressQueryParameter adds the destinationAddress to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetDestinationAddressQueryParameter(destinationAddress *string) {
	o.DestinationAddressQueryParameter = destinationAddress
}

// WithDestinationFamilyQueryParameter adds the destinationFamily to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithDestinationFamilyQueryParameter(destinationFamily *string) *NetworkIPRoutesGetParams {
	o.SetDestinationFamilyQueryParameter(destinationFamily)
	return o
}

// SetDestinationFamilyQueryParameter adds the destinationFamily to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetDestinationFamilyQueryParameter(destinationFamily *string) {
	o.DestinationFamilyQueryParameter = destinationFamily
}

// WithDestinationNetmaskQueryParameter adds the destinationNetmask to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithDestinationNetmaskQueryParameter(destinationNetmask *string) *NetworkIPRoutesGetParams {
	o.SetDestinationNetmaskQueryParameter(destinationNetmask)
	return o
}

// SetDestinationNetmaskQueryParameter adds the destinationNetmask to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetDestinationNetmaskQueryParameter(destinationNetmask *string) {
	o.DestinationNetmaskQueryParameter = destinationNetmask
}

// WithFields adds the fields to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithFields(fields []string) *NetworkIPRoutesGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithGatewayQueryParameter adds the gateway to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithGatewayQueryParameter(gateway *string) *NetworkIPRoutesGetParams {
	o.SetGatewayQueryParameter(gateway)
	return o
}

// SetGatewayQueryParameter adds the gateway to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetGatewayQueryParameter(gateway *string) {
	o.GatewayQueryParameter = gateway
}

// WithIpspaceNameQueryParameter adds the ipspaceName to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithIpspaceNameQueryParameter(ipspaceName *string) *NetworkIPRoutesGetParams {
	o.SetIpspaceNameQueryParameter(ipspaceName)
	return o
}

// SetIpspaceNameQueryParameter adds the ipspaceName to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetIpspaceNameQueryParameter(ipspaceName *string) {
	o.IpspaceNameQueryParameter = ipspaceName
}

// WithIpspaceUUIDQueryParameter adds the ipspaceUUID to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithIpspaceUUIDQueryParameter(ipspaceUUID *string) *NetworkIPRoutesGetParams {
	o.SetIpspaceUUIDQueryParameter(ipspaceUUID)
	return o
}

// SetIpspaceUUIDQueryParameter adds the ipspaceUuid to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetIpspaceUUIDQueryParameter(ipspaceUUID *string) {
	o.IpspaceUUIDQueryParameter = ipspaceUUID
}

// WithMaxRecords adds the maxRecords to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithMaxRecords(maxRecords *int64) *NetworkIPRoutesGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithOrderBy(orderBy []string) *NetworkIPRoutesGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithReturnRecords(returnRecords *bool) *NetworkIPRoutesGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithReturnTimeout(returnTimeout *int64) *NetworkIPRoutesGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScopeQueryParameter adds the scope to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithScopeQueryParameter(scope *string) *NetworkIPRoutesGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithSVMNameQueryParameter adds the svmName to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithSVMNameQueryParameter(svmName *string) *NetworkIPRoutesGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NetworkIPRoutesGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the network ip routes get params
func (o *NetworkIPRoutesGetParams) WithUUIDQueryParameter(uuid *string) *NetworkIPRoutesGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the network ip routes get params
func (o *NetworkIPRoutesGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPRoutesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DestinationAddressQueryParameter != nil {

		// query param destination.address
		var qrDestinationAddress string

		if o.DestinationAddressQueryParameter != nil {
			qrDestinationAddress = *o.DestinationAddressQueryParameter
		}
		qDestinationAddress := qrDestinationAddress
		if qDestinationAddress != "" {

			if err := r.SetQueryParam("destination.address", qDestinationAddress); err != nil {
				return err
			}
		}
	}

	if o.DestinationFamilyQueryParameter != nil {

		// query param destination.family
		var qrDestinationFamily string

		if o.DestinationFamilyQueryParameter != nil {
			qrDestinationFamily = *o.DestinationFamilyQueryParameter
		}
		qDestinationFamily := qrDestinationFamily
		if qDestinationFamily != "" {

			if err := r.SetQueryParam("destination.family", qDestinationFamily); err != nil {
				return err
			}
		}
	}

	if o.DestinationNetmaskQueryParameter != nil {

		// query param destination.netmask
		var qrDestinationNetmask string

		if o.DestinationNetmaskQueryParameter != nil {
			qrDestinationNetmask = *o.DestinationNetmaskQueryParameter
		}
		qDestinationNetmask := qrDestinationNetmask
		if qDestinationNetmask != "" {

			if err := r.SetQueryParam("destination.netmask", qDestinationNetmask); err != nil {
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

	if o.GatewayQueryParameter != nil {

		// query param gateway
		var qrGateway string

		if o.GatewayQueryParameter != nil {
			qrGateway = *o.GatewayQueryParameter
		}
		qGateway := qrGateway
		if qGateway != "" {

			if err := r.SetQueryParam("gateway", qGateway); err != nil {
				return err
			}
		}
	}

	if o.IpspaceNameQueryParameter != nil {

		// query param ipspace.name
		var qrIpspaceName string

		if o.IpspaceNameQueryParameter != nil {
			qrIpspaceName = *o.IpspaceNameQueryParameter
		}
		qIpspaceName := qrIpspaceName
		if qIpspaceName != "" {

			if err := r.SetQueryParam("ipspace.name", qIpspaceName); err != nil {
				return err
			}
		}
	}

	if o.IpspaceUUIDQueryParameter != nil {

		// query param ipspace.uuid
		var qrIpspaceUUID string

		if o.IpspaceUUIDQueryParameter != nil {
			qrIpspaceUUID = *o.IpspaceUUIDQueryParameter
		}
		qIpspaceUUID := qrIpspaceUUID
		if qIpspaceUUID != "" {

			if err := r.SetQueryParam("ipspace.uuid", qIpspaceUUID); err != nil {
				return err
			}
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

	if o.ScopeQueryParameter != nil {

		// query param scope
		var qrScope string

		if o.ScopeQueryParameter != nil {
			qrScope = *o.ScopeQueryParameter
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNetworkIPRoutesGet binds the parameter fields
func (o *NetworkIPRoutesGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNetworkIPRoutesGet binds the parameter order_by
func (o *NetworkIPRoutesGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
