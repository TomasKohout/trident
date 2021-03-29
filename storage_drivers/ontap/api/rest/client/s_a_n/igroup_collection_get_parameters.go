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
	"github.com/go-openapi/swag"
)

// NewIgroupCollectionGetParams creates a new IgroupCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIgroupCollectionGetParams() *IgroupCollectionGetParams {
	return &IgroupCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIgroupCollectionGetParamsWithTimeout creates a new IgroupCollectionGetParams object
// with the ability to set a timeout on a request.
func NewIgroupCollectionGetParamsWithTimeout(timeout time.Duration) *IgroupCollectionGetParams {
	return &IgroupCollectionGetParams{
		timeout: timeout,
	}
}

// NewIgroupCollectionGetParamsWithContext creates a new IgroupCollectionGetParams object
// with the ability to set a context for a request.
func NewIgroupCollectionGetParamsWithContext(ctx context.Context) *IgroupCollectionGetParams {
	return &IgroupCollectionGetParams{
		Context: ctx,
	}
}

// NewIgroupCollectionGetParamsWithHTTPClient creates a new IgroupCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIgroupCollectionGetParamsWithHTTPClient(client *http.Client) *IgroupCollectionGetParams {
	return &IgroupCollectionGetParams{
		HTTPClient: client,
	}
}

/* IgroupCollectionGetParams contains all the parameters to send to the API endpoint
   for the igroup collection get operation.

   Typically these are written to a http.Request.
*/
type IgroupCollectionGetParams struct {

	/* DeleteOnUnmap.

	   Filter by delete_on_unmap
	*/
	DeleteOnUnmapQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* InitiatorsIgroupUUID.

	   Filter by initiators.igroup.uuid
	*/
	InitiatorsIgroupUUIDQueryParameter *string

	/* InitiatorsName.

	   Filter by initiators.name
	*/
	InitiatorsNameQueryParameter *string

	/* LunMapsLogicalUnitNumber.

	   Filter by lun_maps.logical_unit_number
	*/
	LunMapsLogicalUnitNumberQueryParameter *int64

	/* LunMapsLunName.

	   Filter by lun_maps.lun.name
	*/
	LunMapsLunNameQueryParameter *string

	/* LunMapsLunNodeName.

	   Filter by lun_maps.lun.node.name
	*/
	LunMapsLunNodeNameQueryParameter *string

	/* LunMapsLunNodeUUID.

	   Filter by lun_maps.lun.node.uuid
	*/
	LunMapsLunNodeUUIDQueryParameter *string

	/* LunMapsLunUUID.

	   Filter by lun_maps.lun.uuid
	*/
	LunMapsLunUUIDQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* OsType.

	   Filter by os_type
	*/
	OsTypeQueryParameter *string

	/* Protocol.

	   Filter by protocol
	*/
	ProtocolQueryParameter *string

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

// WithDefaults hydrates default values in the igroup collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupCollectionGetParams) WithDefaults() *IgroupCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the igroup collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := IgroupCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the igroup collection get params
func (o *IgroupCollectionGetParams) WithTimeout(timeout time.Duration) *IgroupCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the igroup collection get params
func (o *IgroupCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the igroup collection get params
func (o *IgroupCollectionGetParams) WithContext(ctx context.Context) *IgroupCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the igroup collection get params
func (o *IgroupCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the igroup collection get params
func (o *IgroupCollectionGetParams) WithHTTPClient(client *http.Client) *IgroupCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the igroup collection get params
func (o *IgroupCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteOnUnmapQueryParameter adds the deleteOnUnmap to the igroup collection get params
func (o *IgroupCollectionGetParams) WithDeleteOnUnmapQueryParameter(deleteOnUnmap *bool) *IgroupCollectionGetParams {
	o.SetDeleteOnUnmapQueryParameter(deleteOnUnmap)
	return o
}

// SetDeleteOnUnmapQueryParameter adds the deleteOnUnmap to the igroup collection get params
func (o *IgroupCollectionGetParams) SetDeleteOnUnmapQueryParameter(deleteOnUnmap *bool) {
	o.DeleteOnUnmapQueryParameter = deleteOnUnmap
}

// WithFields adds the fields to the igroup collection get params
func (o *IgroupCollectionGetParams) WithFields(fields []string) *IgroupCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the igroup collection get params
func (o *IgroupCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInitiatorsIgroupUUIDQueryParameter adds the initiatorsIgroupUUID to the igroup collection get params
func (o *IgroupCollectionGetParams) WithInitiatorsIgroupUUIDQueryParameter(initiatorsIgroupUUID *string) *IgroupCollectionGetParams {
	o.SetInitiatorsIgroupUUIDQueryParameter(initiatorsIgroupUUID)
	return o
}

// SetInitiatorsIgroupUUIDQueryParameter adds the initiatorsIgroupUuid to the igroup collection get params
func (o *IgroupCollectionGetParams) SetInitiatorsIgroupUUIDQueryParameter(initiatorsIgroupUUID *string) {
	o.InitiatorsIgroupUUIDQueryParameter = initiatorsIgroupUUID
}

// WithInitiatorsNameQueryParameter adds the initiatorsName to the igroup collection get params
func (o *IgroupCollectionGetParams) WithInitiatorsNameQueryParameter(initiatorsName *string) *IgroupCollectionGetParams {
	o.SetInitiatorsNameQueryParameter(initiatorsName)
	return o
}

// SetInitiatorsNameQueryParameter adds the initiatorsName to the igroup collection get params
func (o *IgroupCollectionGetParams) SetInitiatorsNameQueryParameter(initiatorsName *string) {
	o.InitiatorsNameQueryParameter = initiatorsName
}

// WithLunMapsLogicalUnitNumberQueryParameter adds the lunMapsLogicalUnitNumber to the igroup collection get params
func (o *IgroupCollectionGetParams) WithLunMapsLogicalUnitNumberQueryParameter(lunMapsLogicalUnitNumber *int64) *IgroupCollectionGetParams {
	o.SetLunMapsLogicalUnitNumberQueryParameter(lunMapsLogicalUnitNumber)
	return o
}

// SetLunMapsLogicalUnitNumberQueryParameter adds the lunMapsLogicalUnitNumber to the igroup collection get params
func (o *IgroupCollectionGetParams) SetLunMapsLogicalUnitNumberQueryParameter(lunMapsLogicalUnitNumber *int64) {
	o.LunMapsLogicalUnitNumberQueryParameter = lunMapsLogicalUnitNumber
}

// WithLunMapsLunNameQueryParameter adds the lunMapsLunName to the igroup collection get params
func (o *IgroupCollectionGetParams) WithLunMapsLunNameQueryParameter(lunMapsLunName *string) *IgroupCollectionGetParams {
	o.SetLunMapsLunNameQueryParameter(lunMapsLunName)
	return o
}

// SetLunMapsLunNameQueryParameter adds the lunMapsLunName to the igroup collection get params
func (o *IgroupCollectionGetParams) SetLunMapsLunNameQueryParameter(lunMapsLunName *string) {
	o.LunMapsLunNameQueryParameter = lunMapsLunName
}

// WithLunMapsLunNodeNameQueryParameter adds the lunMapsLunNodeName to the igroup collection get params
func (o *IgroupCollectionGetParams) WithLunMapsLunNodeNameQueryParameter(lunMapsLunNodeName *string) *IgroupCollectionGetParams {
	o.SetLunMapsLunNodeNameQueryParameter(lunMapsLunNodeName)
	return o
}

// SetLunMapsLunNodeNameQueryParameter adds the lunMapsLunNodeName to the igroup collection get params
func (o *IgroupCollectionGetParams) SetLunMapsLunNodeNameQueryParameter(lunMapsLunNodeName *string) {
	o.LunMapsLunNodeNameQueryParameter = lunMapsLunNodeName
}

// WithLunMapsLunNodeUUIDQueryParameter adds the lunMapsLunNodeUUID to the igroup collection get params
func (o *IgroupCollectionGetParams) WithLunMapsLunNodeUUIDQueryParameter(lunMapsLunNodeUUID *string) *IgroupCollectionGetParams {
	o.SetLunMapsLunNodeUUIDQueryParameter(lunMapsLunNodeUUID)
	return o
}

// SetLunMapsLunNodeUUIDQueryParameter adds the lunMapsLunNodeUuid to the igroup collection get params
func (o *IgroupCollectionGetParams) SetLunMapsLunNodeUUIDQueryParameter(lunMapsLunNodeUUID *string) {
	o.LunMapsLunNodeUUIDQueryParameter = lunMapsLunNodeUUID
}

// WithLunMapsLunUUIDQueryParameter adds the lunMapsLunUUID to the igroup collection get params
func (o *IgroupCollectionGetParams) WithLunMapsLunUUIDQueryParameter(lunMapsLunUUID *string) *IgroupCollectionGetParams {
	o.SetLunMapsLunUUIDQueryParameter(lunMapsLunUUID)
	return o
}

// SetLunMapsLunUUIDQueryParameter adds the lunMapsLunUuid to the igroup collection get params
func (o *IgroupCollectionGetParams) SetLunMapsLunUUIDQueryParameter(lunMapsLunUUID *string) {
	o.LunMapsLunUUIDQueryParameter = lunMapsLunUUID
}

// WithMaxRecords adds the maxRecords to the igroup collection get params
func (o *IgroupCollectionGetParams) WithMaxRecords(maxRecords *int64) *IgroupCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the igroup collection get params
func (o *IgroupCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the igroup collection get params
func (o *IgroupCollectionGetParams) WithNameQueryParameter(name *string) *IgroupCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the igroup collection get params
func (o *IgroupCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderBy adds the orderBy to the igroup collection get params
func (o *IgroupCollectionGetParams) WithOrderBy(orderBy []string) *IgroupCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the igroup collection get params
func (o *IgroupCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithOsTypeQueryParameter adds the osType to the igroup collection get params
func (o *IgroupCollectionGetParams) WithOsTypeQueryParameter(osType *string) *IgroupCollectionGetParams {
	o.SetOsTypeQueryParameter(osType)
	return o
}

// SetOsTypeQueryParameter adds the osType to the igroup collection get params
func (o *IgroupCollectionGetParams) SetOsTypeQueryParameter(osType *string) {
	o.OsTypeQueryParameter = osType
}

// WithProtocolQueryParameter adds the protocol to the igroup collection get params
func (o *IgroupCollectionGetParams) WithProtocolQueryParameter(protocol *string) *IgroupCollectionGetParams {
	o.SetProtocolQueryParameter(protocol)
	return o
}

// SetProtocolQueryParameter adds the protocol to the igroup collection get params
func (o *IgroupCollectionGetParams) SetProtocolQueryParameter(protocol *string) {
	o.ProtocolQueryParameter = protocol
}

// WithReturnRecords adds the returnRecords to the igroup collection get params
func (o *IgroupCollectionGetParams) WithReturnRecords(returnRecords *bool) *IgroupCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the igroup collection get params
func (o *IgroupCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the igroup collection get params
func (o *IgroupCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *IgroupCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the igroup collection get params
func (o *IgroupCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the igroup collection get params
func (o *IgroupCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *IgroupCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the igroup collection get params
func (o *IgroupCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the igroup collection get params
func (o *IgroupCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *IgroupCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the igroup collection get params
func (o *IgroupCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the igroup collection get params
func (o *IgroupCollectionGetParams) WithUUIDQueryParameter(uuid *string) *IgroupCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the igroup collection get params
func (o *IgroupCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IgroupCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteOnUnmapQueryParameter != nil {

		// query param delete_on_unmap
		var qrDeleteOnUnmap bool

		if o.DeleteOnUnmapQueryParameter != nil {
			qrDeleteOnUnmap = *o.DeleteOnUnmapQueryParameter
		}
		qDeleteOnUnmap := swag.FormatBool(qrDeleteOnUnmap)
		if qDeleteOnUnmap != "" {

			if err := r.SetQueryParam("delete_on_unmap", qDeleteOnUnmap); err != nil {
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

	if o.InitiatorsIgroupUUIDQueryParameter != nil {

		// query param initiators.igroup.uuid
		var qrInitiatorsIgroupUUID string

		if o.InitiatorsIgroupUUIDQueryParameter != nil {
			qrInitiatorsIgroupUUID = *o.InitiatorsIgroupUUIDQueryParameter
		}
		qInitiatorsIgroupUUID := qrInitiatorsIgroupUUID
		if qInitiatorsIgroupUUID != "" {

			if err := r.SetQueryParam("initiators.igroup.uuid", qInitiatorsIgroupUUID); err != nil {
				return err
			}
		}
	}

	if o.InitiatorsNameQueryParameter != nil {

		// query param initiators.name
		var qrInitiatorsName string

		if o.InitiatorsNameQueryParameter != nil {
			qrInitiatorsName = *o.InitiatorsNameQueryParameter
		}
		qInitiatorsName := qrInitiatorsName
		if qInitiatorsName != "" {

			if err := r.SetQueryParam("initiators.name", qInitiatorsName); err != nil {
				return err
			}
		}
	}

	if o.LunMapsLogicalUnitNumberQueryParameter != nil {

		// query param lun_maps.logical_unit_number
		var qrLunMapsLogicalUnitNumber int64

		if o.LunMapsLogicalUnitNumberQueryParameter != nil {
			qrLunMapsLogicalUnitNumber = *o.LunMapsLogicalUnitNumberQueryParameter
		}
		qLunMapsLogicalUnitNumber := swag.FormatInt64(qrLunMapsLogicalUnitNumber)
		if qLunMapsLogicalUnitNumber != "" {

			if err := r.SetQueryParam("lun_maps.logical_unit_number", qLunMapsLogicalUnitNumber); err != nil {
				return err
			}
		}
	}

	if o.LunMapsLunNameQueryParameter != nil {

		// query param lun_maps.lun.name
		var qrLunMapsLunName string

		if o.LunMapsLunNameQueryParameter != nil {
			qrLunMapsLunName = *o.LunMapsLunNameQueryParameter
		}
		qLunMapsLunName := qrLunMapsLunName
		if qLunMapsLunName != "" {

			if err := r.SetQueryParam("lun_maps.lun.name", qLunMapsLunName); err != nil {
				return err
			}
		}
	}

	if o.LunMapsLunNodeNameQueryParameter != nil {

		// query param lun_maps.lun.node.name
		var qrLunMapsLunNodeName string

		if o.LunMapsLunNodeNameQueryParameter != nil {
			qrLunMapsLunNodeName = *o.LunMapsLunNodeNameQueryParameter
		}
		qLunMapsLunNodeName := qrLunMapsLunNodeName
		if qLunMapsLunNodeName != "" {

			if err := r.SetQueryParam("lun_maps.lun.node.name", qLunMapsLunNodeName); err != nil {
				return err
			}
		}
	}

	if o.LunMapsLunNodeUUIDQueryParameter != nil {

		// query param lun_maps.lun.node.uuid
		var qrLunMapsLunNodeUUID string

		if o.LunMapsLunNodeUUIDQueryParameter != nil {
			qrLunMapsLunNodeUUID = *o.LunMapsLunNodeUUIDQueryParameter
		}
		qLunMapsLunNodeUUID := qrLunMapsLunNodeUUID
		if qLunMapsLunNodeUUID != "" {

			if err := r.SetQueryParam("lun_maps.lun.node.uuid", qLunMapsLunNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.LunMapsLunUUIDQueryParameter != nil {

		// query param lun_maps.lun.uuid
		var qrLunMapsLunUUID string

		if o.LunMapsLunUUIDQueryParameter != nil {
			qrLunMapsLunUUID = *o.LunMapsLunUUIDQueryParameter
		}
		qLunMapsLunUUID := qrLunMapsLunUUID
		if qLunMapsLunUUID != "" {

			if err := r.SetQueryParam("lun_maps.lun.uuid", qLunMapsLunUUID); err != nil {
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

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.OsTypeQueryParameter != nil {

		// query param os_type
		var qrOsType string

		if o.OsTypeQueryParameter != nil {
			qrOsType = *o.OsTypeQueryParameter
		}
		qOsType := qrOsType
		if qOsType != "" {

			if err := r.SetQueryParam("os_type", qOsType); err != nil {
				return err
			}
		}
	}

	if o.ProtocolQueryParameter != nil {

		// query param protocol
		var qrProtocol string

		if o.ProtocolQueryParameter != nil {
			qrProtocol = *o.ProtocolQueryParameter
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
				return err
			}
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

// bindParamIgroupCollectionGet binds the parameter fields
func (o *IgroupCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamIgroupCollectionGet binds the parameter order_by
func (o *IgroupCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
