// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

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

// NewSnapmirrorPoliciesGetParams creates a new SnapmirrorPoliciesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorPoliciesGetParams() *SnapmirrorPoliciesGetParams {
	return &SnapmirrorPoliciesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorPoliciesGetParamsWithTimeout creates a new SnapmirrorPoliciesGetParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorPoliciesGetParamsWithTimeout(timeout time.Duration) *SnapmirrorPoliciesGetParams {
	return &SnapmirrorPoliciesGetParams{
		timeout: timeout,
	}
}

// NewSnapmirrorPoliciesGetParamsWithContext creates a new SnapmirrorPoliciesGetParams object
// with the ability to set a context for a request.
func NewSnapmirrorPoliciesGetParamsWithContext(ctx context.Context) *SnapmirrorPoliciesGetParams {
	return &SnapmirrorPoliciesGetParams{
		Context: ctx,
	}
}

// NewSnapmirrorPoliciesGetParamsWithHTTPClient creates a new SnapmirrorPoliciesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorPoliciesGetParamsWithHTTPClient(client *http.Client) *SnapmirrorPoliciesGetParams {
	return &SnapmirrorPoliciesGetParams{
		HTTPClient: client,
	}
}

/* SnapmirrorPoliciesGetParams contains all the parameters to send to the API endpoint
   for the snapmirror policies get operation.

   Typically these are written to a http.Request.
*/
type SnapmirrorPoliciesGetParams struct {

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IdentityPreservation.

	   Filter by identity_preservation
	*/
	IdentityPreservationQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* NetworkCompressionEnabled.

	   Filter by network_compression_enabled
	*/
	NetworkCompressionEnabledQueryParameter *bool

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* RetentionCount.

	   Filter by retention.count
	*/
	RetentionCountQueryParameter *int64

	/* RetentionCreationScheduleName.

	   Filter by retention.creation_schedule.name
	*/
	RetentionCreationScheduleNameQueryParameter *string

	/* RetentionCreationScheduleUUID.

	   Filter by retention.creation_schedule.uuid
	*/
	RetentionCreationScheduleUUIDQueryParameter *string

	/* RetentionLabel.

	   Filter by retention.label
	*/
	RetentionLabelQueryParameter *string

	/* RetentionPrefix.

	   Filter by retention.prefix
	*/
	RetentionPrefixQueryParameter *string

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

	/* SyncCommonSnapshotScheduleName.

	   Filter by sync_common_snapshot_schedule.name
	*/
	SyncCommonSnapshotScheduleNameQueryParameter *string

	/* SyncCommonSnapshotScheduleUUID.

	   Filter by sync_common_snapshot_schedule.uuid
	*/
	SyncCommonSnapshotScheduleUUIDQueryParameter *string

	/* SyncType.

	   Filter by sync_type
	*/
	SyncTypeQueryParameter *string

	/* Throttle.

	   Filter by throttle
	*/
	ThrottleQueryParameter *int64

	/* TransferScheduleName.

	   Filter by transfer_schedule.name
	*/
	TransferScheduleNameQueryParameter *string

	/* TransferScheduleUUID.

	   Filter by transfer_schedule.uuid
	*/
	TransferScheduleUUIDQueryParameter *string

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror policies get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPoliciesGetParams) WithDefaults() *SnapmirrorPoliciesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror policies get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPoliciesGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SnapmirrorPoliciesGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithTimeout(timeout time.Duration) *SnapmirrorPoliciesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithContext(ctx context.Context) *SnapmirrorPoliciesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithHTTPClient(client *http.Client) *SnapmirrorPoliciesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentQueryParameter adds the comment to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithCommentQueryParameter(comment *string) *SnapmirrorPoliciesGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithFields adds the fields to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithFields(fields []string) *SnapmirrorPoliciesGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIdentityPreservationQueryParameter adds the identityPreservation to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithIdentityPreservationQueryParameter(identityPreservation *string) *SnapmirrorPoliciesGetParams {
	o.SetIdentityPreservationQueryParameter(identityPreservation)
	return o
}

// SetIdentityPreservationQueryParameter adds the identityPreservation to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetIdentityPreservationQueryParameter(identityPreservation *string) {
	o.IdentityPreservationQueryParameter = identityPreservation
}

// WithMaxRecords adds the maxRecords to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithMaxRecords(maxRecords *int64) *SnapmirrorPoliciesGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithNameQueryParameter(name *string) *SnapmirrorPoliciesGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithNetworkCompressionEnabledQueryParameter adds the networkCompressionEnabled to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithNetworkCompressionEnabledQueryParameter(networkCompressionEnabled *bool) *SnapmirrorPoliciesGetParams {
	o.SetNetworkCompressionEnabledQueryParameter(networkCompressionEnabled)
	return o
}

// SetNetworkCompressionEnabledQueryParameter adds the networkCompressionEnabled to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetNetworkCompressionEnabledQueryParameter(networkCompressionEnabled *bool) {
	o.NetworkCompressionEnabledQueryParameter = networkCompressionEnabled
}

// WithOrderBy adds the orderBy to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithOrderBy(orderBy []string) *SnapmirrorPoliciesGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithRetentionCountQueryParameter adds the retentionCount to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithRetentionCountQueryParameter(retentionCount *int64) *SnapmirrorPoliciesGetParams {
	o.SetRetentionCountQueryParameter(retentionCount)
	return o
}

// SetRetentionCountQueryParameter adds the retentionCount to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetRetentionCountQueryParameter(retentionCount *int64) {
	o.RetentionCountQueryParameter = retentionCount
}

// WithRetentionCreationScheduleNameQueryParameter adds the retentionCreationScheduleName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithRetentionCreationScheduleNameQueryParameter(retentionCreationScheduleName *string) *SnapmirrorPoliciesGetParams {
	o.SetRetentionCreationScheduleNameQueryParameter(retentionCreationScheduleName)
	return o
}

// SetRetentionCreationScheduleNameQueryParameter adds the retentionCreationScheduleName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetRetentionCreationScheduleNameQueryParameter(retentionCreationScheduleName *string) {
	o.RetentionCreationScheduleNameQueryParameter = retentionCreationScheduleName
}

// WithRetentionCreationScheduleUUIDQueryParameter adds the retentionCreationScheduleUUID to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithRetentionCreationScheduleUUIDQueryParameter(retentionCreationScheduleUUID *string) *SnapmirrorPoliciesGetParams {
	o.SetRetentionCreationScheduleUUIDQueryParameter(retentionCreationScheduleUUID)
	return o
}

// SetRetentionCreationScheduleUUIDQueryParameter adds the retentionCreationScheduleUuid to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetRetentionCreationScheduleUUIDQueryParameter(retentionCreationScheduleUUID *string) {
	o.RetentionCreationScheduleUUIDQueryParameter = retentionCreationScheduleUUID
}

// WithRetentionLabelQueryParameter adds the retentionLabel to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithRetentionLabelQueryParameter(retentionLabel *string) *SnapmirrorPoliciesGetParams {
	o.SetRetentionLabelQueryParameter(retentionLabel)
	return o
}

// SetRetentionLabelQueryParameter adds the retentionLabel to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetRetentionLabelQueryParameter(retentionLabel *string) {
	o.RetentionLabelQueryParameter = retentionLabel
}

// WithRetentionPrefixQueryParameter adds the retentionPrefix to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithRetentionPrefixQueryParameter(retentionPrefix *string) *SnapmirrorPoliciesGetParams {
	o.SetRetentionPrefixQueryParameter(retentionPrefix)
	return o
}

// SetRetentionPrefixQueryParameter adds the retentionPrefix to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetRetentionPrefixQueryParameter(retentionPrefix *string) {
	o.RetentionPrefixQueryParameter = retentionPrefix
}

// WithReturnRecords adds the returnRecords to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithReturnRecords(returnRecords *bool) *SnapmirrorPoliciesGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithReturnTimeout(returnTimeout *int64) *SnapmirrorPoliciesGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScopeQueryParameter adds the scope to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithScopeQueryParameter(scope *string) *SnapmirrorPoliciesGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithSVMNameQueryParameter adds the svmName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithSVMNameQueryParameter(svmName *string) *SnapmirrorPoliciesGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SnapmirrorPoliciesGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithSyncCommonSnapshotScheduleNameQueryParameter adds the syncCommonSnapshotScheduleName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithSyncCommonSnapshotScheduleNameQueryParameter(syncCommonSnapshotScheduleName *string) *SnapmirrorPoliciesGetParams {
	o.SetSyncCommonSnapshotScheduleNameQueryParameter(syncCommonSnapshotScheduleName)
	return o
}

// SetSyncCommonSnapshotScheduleNameQueryParameter adds the syncCommonSnapshotScheduleName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetSyncCommonSnapshotScheduleNameQueryParameter(syncCommonSnapshotScheduleName *string) {
	o.SyncCommonSnapshotScheduleNameQueryParameter = syncCommonSnapshotScheduleName
}

// WithSyncCommonSnapshotScheduleUUIDQueryParameter adds the syncCommonSnapshotScheduleUUID to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithSyncCommonSnapshotScheduleUUIDQueryParameter(syncCommonSnapshotScheduleUUID *string) *SnapmirrorPoliciesGetParams {
	o.SetSyncCommonSnapshotScheduleUUIDQueryParameter(syncCommonSnapshotScheduleUUID)
	return o
}

// SetSyncCommonSnapshotScheduleUUIDQueryParameter adds the syncCommonSnapshotScheduleUuid to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetSyncCommonSnapshotScheduleUUIDQueryParameter(syncCommonSnapshotScheduleUUID *string) {
	o.SyncCommonSnapshotScheduleUUIDQueryParameter = syncCommonSnapshotScheduleUUID
}

// WithSyncTypeQueryParameter adds the syncType to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithSyncTypeQueryParameter(syncType *string) *SnapmirrorPoliciesGetParams {
	o.SetSyncTypeQueryParameter(syncType)
	return o
}

// SetSyncTypeQueryParameter adds the syncType to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetSyncTypeQueryParameter(syncType *string) {
	o.SyncTypeQueryParameter = syncType
}

// WithThrottleQueryParameter adds the throttle to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithThrottleQueryParameter(throttle *int64) *SnapmirrorPoliciesGetParams {
	o.SetThrottleQueryParameter(throttle)
	return o
}

// SetThrottleQueryParameter adds the throttle to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetThrottleQueryParameter(throttle *int64) {
	o.ThrottleQueryParameter = throttle
}

// WithTransferScheduleNameQueryParameter adds the transferScheduleName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithTransferScheduleNameQueryParameter(transferScheduleName *string) *SnapmirrorPoliciesGetParams {
	o.SetTransferScheduleNameQueryParameter(transferScheduleName)
	return o
}

// SetTransferScheduleNameQueryParameter adds the transferScheduleName to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetTransferScheduleNameQueryParameter(transferScheduleName *string) {
	o.TransferScheduleNameQueryParameter = transferScheduleName
}

// WithTransferScheduleUUIDQueryParameter adds the transferScheduleUUID to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithTransferScheduleUUIDQueryParameter(transferScheduleUUID *string) *SnapmirrorPoliciesGetParams {
	o.SetTransferScheduleUUIDQueryParameter(transferScheduleUUID)
	return o
}

// SetTransferScheduleUUIDQueryParameter adds the transferScheduleUuid to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetTransferScheduleUUIDQueryParameter(transferScheduleUUID *string) {
	o.TransferScheduleUUIDQueryParameter = transferScheduleUUID
}

// WithTypeQueryParameter adds the typeVar to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithTypeQueryParameter(typeVar *string) *SnapmirrorPoliciesGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithUUIDQueryParameter adds the uuid to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) WithUUIDQueryParameter(uuid *string) *SnapmirrorPoliciesGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the snapmirror policies get params
func (o *SnapmirrorPoliciesGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorPoliciesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
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

	if o.IdentityPreservationQueryParameter != nil {

		// query param identity_preservation
		var qrIdentityPreservation string

		if o.IdentityPreservationQueryParameter != nil {
			qrIdentityPreservation = *o.IdentityPreservationQueryParameter
		}
		qIdentityPreservation := qrIdentityPreservation
		if qIdentityPreservation != "" {

			if err := r.SetQueryParam("identity_preservation", qIdentityPreservation); err != nil {
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

	if o.NetworkCompressionEnabledQueryParameter != nil {

		// query param network_compression_enabled
		var qrNetworkCompressionEnabled bool

		if o.NetworkCompressionEnabledQueryParameter != nil {
			qrNetworkCompressionEnabled = *o.NetworkCompressionEnabledQueryParameter
		}
		qNetworkCompressionEnabled := swag.FormatBool(qrNetworkCompressionEnabled)
		if qNetworkCompressionEnabled != "" {

			if err := r.SetQueryParam("network_compression_enabled", qNetworkCompressionEnabled); err != nil {
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

	if o.RetentionCountQueryParameter != nil {

		// query param retention.count
		var qrRetentionCount int64

		if o.RetentionCountQueryParameter != nil {
			qrRetentionCount = *o.RetentionCountQueryParameter
		}
		qRetentionCount := swag.FormatInt64(qrRetentionCount)
		if qRetentionCount != "" {

			if err := r.SetQueryParam("retention.count", qRetentionCount); err != nil {
				return err
			}
		}
	}

	if o.RetentionCreationScheduleNameQueryParameter != nil {

		// query param retention.creation_schedule.name
		var qrRetentionCreationScheduleName string

		if o.RetentionCreationScheduleNameQueryParameter != nil {
			qrRetentionCreationScheduleName = *o.RetentionCreationScheduleNameQueryParameter
		}
		qRetentionCreationScheduleName := qrRetentionCreationScheduleName
		if qRetentionCreationScheduleName != "" {

			if err := r.SetQueryParam("retention.creation_schedule.name", qRetentionCreationScheduleName); err != nil {
				return err
			}
		}
	}

	if o.RetentionCreationScheduleUUIDQueryParameter != nil {

		// query param retention.creation_schedule.uuid
		var qrRetentionCreationScheduleUUID string

		if o.RetentionCreationScheduleUUIDQueryParameter != nil {
			qrRetentionCreationScheduleUUID = *o.RetentionCreationScheduleUUIDQueryParameter
		}
		qRetentionCreationScheduleUUID := qrRetentionCreationScheduleUUID
		if qRetentionCreationScheduleUUID != "" {

			if err := r.SetQueryParam("retention.creation_schedule.uuid", qRetentionCreationScheduleUUID); err != nil {
				return err
			}
		}
	}

	if o.RetentionLabelQueryParameter != nil {

		// query param retention.label
		var qrRetentionLabel string

		if o.RetentionLabelQueryParameter != nil {
			qrRetentionLabel = *o.RetentionLabelQueryParameter
		}
		qRetentionLabel := qrRetentionLabel
		if qRetentionLabel != "" {

			if err := r.SetQueryParam("retention.label", qRetentionLabel); err != nil {
				return err
			}
		}
	}

	if o.RetentionPrefixQueryParameter != nil {

		// query param retention.prefix
		var qrRetentionPrefix string

		if o.RetentionPrefixQueryParameter != nil {
			qrRetentionPrefix = *o.RetentionPrefixQueryParameter
		}
		qRetentionPrefix := qrRetentionPrefix
		if qRetentionPrefix != "" {

			if err := r.SetQueryParam("retention.prefix", qRetentionPrefix); err != nil {
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

	if o.SyncCommonSnapshotScheduleNameQueryParameter != nil {

		// query param sync_common_snapshot_schedule.name
		var qrSyncCommonSnapshotScheduleName string

		if o.SyncCommonSnapshotScheduleNameQueryParameter != nil {
			qrSyncCommonSnapshotScheduleName = *o.SyncCommonSnapshotScheduleNameQueryParameter
		}
		qSyncCommonSnapshotScheduleName := qrSyncCommonSnapshotScheduleName
		if qSyncCommonSnapshotScheduleName != "" {

			if err := r.SetQueryParam("sync_common_snapshot_schedule.name", qSyncCommonSnapshotScheduleName); err != nil {
				return err
			}
		}
	}

	if o.SyncCommonSnapshotScheduleUUIDQueryParameter != nil {

		// query param sync_common_snapshot_schedule.uuid
		var qrSyncCommonSnapshotScheduleUUID string

		if o.SyncCommonSnapshotScheduleUUIDQueryParameter != nil {
			qrSyncCommonSnapshotScheduleUUID = *o.SyncCommonSnapshotScheduleUUIDQueryParameter
		}
		qSyncCommonSnapshotScheduleUUID := qrSyncCommonSnapshotScheduleUUID
		if qSyncCommonSnapshotScheduleUUID != "" {

			if err := r.SetQueryParam("sync_common_snapshot_schedule.uuid", qSyncCommonSnapshotScheduleUUID); err != nil {
				return err
			}
		}
	}

	if o.SyncTypeQueryParameter != nil {

		// query param sync_type
		var qrSyncType string

		if o.SyncTypeQueryParameter != nil {
			qrSyncType = *o.SyncTypeQueryParameter
		}
		qSyncType := qrSyncType
		if qSyncType != "" {

			if err := r.SetQueryParam("sync_type", qSyncType); err != nil {
				return err
			}
		}
	}

	if o.ThrottleQueryParameter != nil {

		// query param throttle
		var qrThrottle int64

		if o.ThrottleQueryParameter != nil {
			qrThrottle = *o.ThrottleQueryParameter
		}
		qThrottle := swag.FormatInt64(qrThrottle)
		if qThrottle != "" {

			if err := r.SetQueryParam("throttle", qThrottle); err != nil {
				return err
			}
		}
	}

	if o.TransferScheduleNameQueryParameter != nil {

		// query param transfer_schedule.name
		var qrTransferScheduleName string

		if o.TransferScheduleNameQueryParameter != nil {
			qrTransferScheduleName = *o.TransferScheduleNameQueryParameter
		}
		qTransferScheduleName := qrTransferScheduleName
		if qTransferScheduleName != "" {

			if err := r.SetQueryParam("transfer_schedule.name", qTransferScheduleName); err != nil {
				return err
			}
		}
	}

	if o.TransferScheduleUUIDQueryParameter != nil {

		// query param transfer_schedule.uuid
		var qrTransferScheduleUUID string

		if o.TransferScheduleUUIDQueryParameter != nil {
			qrTransferScheduleUUID = *o.TransferScheduleUUIDQueryParameter
		}
		qTransferScheduleUUID := qrTransferScheduleUUID
		if qTransferScheduleUUID != "" {

			if err := r.SetQueryParam("transfer_schedule.uuid", qTransferScheduleUUID); err != nil {
				return err
			}
		}
	}

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
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

// bindParamSnapmirrorPoliciesGet binds the parameter fields
func (o *SnapmirrorPoliciesGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnapmirrorPoliciesGet binds the parameter order_by
func (o *SnapmirrorPoliciesGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
