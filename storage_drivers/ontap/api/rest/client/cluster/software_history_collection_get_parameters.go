// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewSoftwareHistoryCollectionGetParams creates a new SoftwareHistoryCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSoftwareHistoryCollectionGetParams() *SoftwareHistoryCollectionGetParams {
	return &SoftwareHistoryCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareHistoryCollectionGetParamsWithTimeout creates a new SoftwareHistoryCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSoftwareHistoryCollectionGetParamsWithTimeout(timeout time.Duration) *SoftwareHistoryCollectionGetParams {
	return &SoftwareHistoryCollectionGetParams{
		timeout: timeout,
	}
}

// NewSoftwareHistoryCollectionGetParamsWithContext creates a new SoftwareHistoryCollectionGetParams object
// with the ability to set a context for a request.
func NewSoftwareHistoryCollectionGetParamsWithContext(ctx context.Context) *SoftwareHistoryCollectionGetParams {
	return &SoftwareHistoryCollectionGetParams{
		Context: ctx,
	}
}

// NewSoftwareHistoryCollectionGetParamsWithHTTPClient creates a new SoftwareHistoryCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSoftwareHistoryCollectionGetParamsWithHTTPClient(client *http.Client) *SoftwareHistoryCollectionGetParams {
	return &SoftwareHistoryCollectionGetParams{
		HTTPClient: client,
	}
}

/* SoftwareHistoryCollectionGetParams contains all the parameters to send to the API endpoint
   for the software history collection get operation.

   Typically these are written to a http.Request.
*/
type SoftwareHistoryCollectionGetParams struct {

	/* EndTime.

	   Filter by end_time
	*/
	EndTimeQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* FromVersion.

	   Filter by from_version
	*/
	FromVersionQueryParameter *string

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

	/* StartTime.

	   Filter by start_time
	*/
	StartTimeQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* ToVersion.

	   Filter by to_version
	*/
	ToVersionQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the software history collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwareHistoryCollectionGetParams) WithDefaults() *SoftwareHistoryCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the software history collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwareHistoryCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SoftwareHistoryCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithTimeout(timeout time.Duration) *SoftwareHistoryCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithContext(ctx context.Context) *SoftwareHistoryCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithHTTPClient(client *http.Client) *SoftwareHistoryCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTimeQueryParameter adds the endTime to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithEndTimeQueryParameter(endTime *string) *SoftwareHistoryCollectionGetParams {
	o.SetEndTimeQueryParameter(endTime)
	return o
}

// SetEndTimeQueryParameter adds the endTime to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetEndTimeQueryParameter(endTime *string) {
	o.EndTimeQueryParameter = endTime
}

// WithFields adds the fields to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithFields(fields []string) *SoftwareHistoryCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFromVersionQueryParameter adds the fromVersion to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithFromVersionQueryParameter(fromVersion *string) *SoftwareHistoryCollectionGetParams {
	o.SetFromVersionQueryParameter(fromVersion)
	return o
}

// SetFromVersionQueryParameter adds the fromVersion to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetFromVersionQueryParameter(fromVersion *string) {
	o.FromVersionQueryParameter = fromVersion
}

// WithMaxRecords adds the maxRecords to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithMaxRecords(maxRecords *int64) *SoftwareHistoryCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *SoftwareHistoryCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *SoftwareHistoryCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderBy adds the orderBy to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithOrderBy(orderBy []string) *SoftwareHistoryCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithReturnRecords(returnRecords *bool) *SoftwareHistoryCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SoftwareHistoryCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStartTimeQueryParameter adds the startTime to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithStartTimeQueryParameter(startTime *string) *SoftwareHistoryCollectionGetParams {
	o.SetStartTimeQueryParameter(startTime)
	return o
}

// SetStartTimeQueryParameter adds the startTime to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetStartTimeQueryParameter(startTime *string) {
	o.StartTimeQueryParameter = startTime
}

// WithStateQueryParameter adds the state to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithStateQueryParameter(state *string) *SoftwareHistoryCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithToVersionQueryParameter adds the toVersion to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) WithToVersionQueryParameter(toVersion *string) *SoftwareHistoryCollectionGetParams {
	o.SetToVersionQueryParameter(toVersion)
	return o
}

// SetToVersionQueryParameter adds the toVersion to the software history collection get params
func (o *SoftwareHistoryCollectionGetParams) SetToVersionQueryParameter(toVersion *string) {
	o.ToVersionQueryParameter = toVersion
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareHistoryCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndTimeQueryParameter != nil {

		// query param end_time
		var qrEndTime string

		if o.EndTimeQueryParameter != nil {
			qrEndTime = *o.EndTimeQueryParameter
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
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

	if o.FromVersionQueryParameter != nil {

		// query param from_version
		var qrFromVersion string

		if o.FromVersionQueryParameter != nil {
			qrFromVersion = *o.FromVersionQueryParameter
		}
		qFromVersion := qrFromVersion
		if qFromVersion != "" {

			if err := r.SetQueryParam("from_version", qFromVersion); err != nil {
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

	if o.StartTimeQueryParameter != nil {

		// query param start_time
		var qrStartTime string

		if o.StartTimeQueryParameter != nil {
			qrStartTime = *o.StartTimeQueryParameter
		}
		qStartTime := qrStartTime
		if qStartTime != "" {

			if err := r.SetQueryParam("start_time", qStartTime); err != nil {
				return err
			}
		}
	}

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.ToVersionQueryParameter != nil {

		// query param to_version
		var qrToVersion string

		if o.ToVersionQueryParameter != nil {
			qrToVersion = *o.ToVersionQueryParameter
		}
		qToVersion := qrToVersion
		if qToVersion != "" {

			if err := r.SetQueryParam("to_version", qToVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSoftwareHistoryCollectionGet binds the parameter fields
func (o *SoftwareHistoryCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSoftwareHistoryCollectionGet binds the parameter order_by
func (o *SoftwareHistoryCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
