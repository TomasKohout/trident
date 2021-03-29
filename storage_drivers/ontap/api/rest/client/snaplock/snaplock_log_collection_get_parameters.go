// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// NewSnaplockLogCollectionGetParams creates a new SnaplockLogCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLogCollectionGetParams() *SnaplockLogCollectionGetParams {
	return &SnaplockLogCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLogCollectionGetParamsWithTimeout creates a new SnaplockLogCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockLogCollectionGetParamsWithTimeout(timeout time.Duration) *SnaplockLogCollectionGetParams {
	return &SnaplockLogCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnaplockLogCollectionGetParamsWithContext creates a new SnaplockLogCollectionGetParams object
// with the ability to set a context for a request.
func NewSnaplockLogCollectionGetParamsWithContext(ctx context.Context) *SnaplockLogCollectionGetParams {
	return &SnaplockLogCollectionGetParams{
		Context: ctx,
	}
}

// NewSnaplockLogCollectionGetParamsWithHTTPClient creates a new SnaplockLogCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLogCollectionGetParamsWithHTTPClient(client *http.Client) *SnaplockLogCollectionGetParams {
	return &SnaplockLogCollectionGetParams{
		HTTPClient: client,
	}
}

/* SnaplockLogCollectionGetParams contains all the parameters to send to the API endpoint
   for the snaplock log collection get operation.

   Typically these are written to a http.Request.
*/
type SnaplockLogCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LogArchiveBaseName.

	   Filter by log_archive.base_name
	*/
	LogArchiveBaseNameQueryParameter *string

	/* LogArchiveExpiryTime.

	   Filter by log_archive.expiry_time
	*/
	LogArchiveExpiryTimeQueryParameter *string

	/* LogArchivePath.

	   Filter by log_archive.path
	*/
	LogArchivePathQueryParameter *string

	/* LogArchiveSize.

	   Filter by log_archive.size
	*/
	LogArchiveSizeQueryParameter *int64

	/* LogFilesBaseName.

	   Filter by log_files.base_name
	*/
	LogFilesBaseNameQueryParameter *string

	/* LogFilesExpiryTime.

	   Filter by log_files.expiry_time
	*/
	LogFilesExpiryTimeQueryParameter *string

	/* LogFilesPath.

	   Filter by log_files.path
	*/
	LogFilesPathQueryParameter *string

	/* LogFilesSize.

	   Filter by log_files.size
	*/
	LogFilesSizeQueryParameter *int64

	/* LogVolumeMaxLogSize.

	   Filter by log_volume.max_log_size
	*/
	LogVolumeMaxLogSizeQueryParameter *int64

	/* LogVolumeRetentionPeriod.

	   Filter by log_volume.retention_period
	*/
	LogVolumeRetentionPeriodQueryParameter *string

	/* LogVolumeVolumeName.

	   Filter by log_volume.volume.name
	*/
	LogVolumeVolumeNameQueryParameter *string

	/* LogVolumeVolumeUUID.

	   Filter by log_volume.volume.uuid
	*/
	LogVolumeVolumeUUIDQueryParameter *string

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

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock log collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLogCollectionGetParams) WithDefaults() *SnaplockLogCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock log collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLogCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SnaplockLogCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithTimeout(timeout time.Duration) *SnaplockLogCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithContext(ctx context.Context) *SnaplockLogCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithHTTPClient(client *http.Client) *SnaplockLogCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithFields(fields []string) *SnaplockLogCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLogArchiveBaseNameQueryParameter adds the logArchiveBaseName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogArchiveBaseNameQueryParameter(logArchiveBaseName *string) *SnaplockLogCollectionGetParams {
	o.SetLogArchiveBaseNameQueryParameter(logArchiveBaseName)
	return o
}

// SetLogArchiveBaseNameQueryParameter adds the logArchiveBaseName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogArchiveBaseNameQueryParameter(logArchiveBaseName *string) {
	o.LogArchiveBaseNameQueryParameter = logArchiveBaseName
}

// WithLogArchiveExpiryTimeQueryParameter adds the logArchiveExpiryTime to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogArchiveExpiryTimeQueryParameter(logArchiveExpiryTime *string) *SnaplockLogCollectionGetParams {
	o.SetLogArchiveExpiryTimeQueryParameter(logArchiveExpiryTime)
	return o
}

// SetLogArchiveExpiryTimeQueryParameter adds the logArchiveExpiryTime to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogArchiveExpiryTimeQueryParameter(logArchiveExpiryTime *string) {
	o.LogArchiveExpiryTimeQueryParameter = logArchiveExpiryTime
}

// WithLogArchivePathQueryParameter adds the logArchivePath to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogArchivePathQueryParameter(logArchivePath *string) *SnaplockLogCollectionGetParams {
	o.SetLogArchivePathQueryParameter(logArchivePath)
	return o
}

// SetLogArchivePathQueryParameter adds the logArchivePath to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogArchivePathQueryParameter(logArchivePath *string) {
	o.LogArchivePathQueryParameter = logArchivePath
}

// WithLogArchiveSizeQueryParameter adds the logArchiveSize to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogArchiveSizeQueryParameter(logArchiveSize *int64) *SnaplockLogCollectionGetParams {
	o.SetLogArchiveSizeQueryParameter(logArchiveSize)
	return o
}

// SetLogArchiveSizeQueryParameter adds the logArchiveSize to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogArchiveSizeQueryParameter(logArchiveSize *int64) {
	o.LogArchiveSizeQueryParameter = logArchiveSize
}

// WithLogFilesBaseNameQueryParameter adds the logFilesBaseName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogFilesBaseNameQueryParameter(logFilesBaseName *string) *SnaplockLogCollectionGetParams {
	o.SetLogFilesBaseNameQueryParameter(logFilesBaseName)
	return o
}

// SetLogFilesBaseNameQueryParameter adds the logFilesBaseName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogFilesBaseNameQueryParameter(logFilesBaseName *string) {
	o.LogFilesBaseNameQueryParameter = logFilesBaseName
}

// WithLogFilesExpiryTimeQueryParameter adds the logFilesExpiryTime to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogFilesExpiryTimeQueryParameter(logFilesExpiryTime *string) *SnaplockLogCollectionGetParams {
	o.SetLogFilesExpiryTimeQueryParameter(logFilesExpiryTime)
	return o
}

// SetLogFilesExpiryTimeQueryParameter adds the logFilesExpiryTime to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogFilesExpiryTimeQueryParameter(logFilesExpiryTime *string) {
	o.LogFilesExpiryTimeQueryParameter = logFilesExpiryTime
}

// WithLogFilesPathQueryParameter adds the logFilesPath to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogFilesPathQueryParameter(logFilesPath *string) *SnaplockLogCollectionGetParams {
	o.SetLogFilesPathQueryParameter(logFilesPath)
	return o
}

// SetLogFilesPathQueryParameter adds the logFilesPath to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogFilesPathQueryParameter(logFilesPath *string) {
	o.LogFilesPathQueryParameter = logFilesPath
}

// WithLogFilesSizeQueryParameter adds the logFilesSize to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogFilesSizeQueryParameter(logFilesSize *int64) *SnaplockLogCollectionGetParams {
	o.SetLogFilesSizeQueryParameter(logFilesSize)
	return o
}

// SetLogFilesSizeQueryParameter adds the logFilesSize to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogFilesSizeQueryParameter(logFilesSize *int64) {
	o.LogFilesSizeQueryParameter = logFilesSize
}

// WithLogVolumeMaxLogSizeQueryParameter adds the logVolumeMaxLogSize to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogVolumeMaxLogSizeQueryParameter(logVolumeMaxLogSize *int64) *SnaplockLogCollectionGetParams {
	o.SetLogVolumeMaxLogSizeQueryParameter(logVolumeMaxLogSize)
	return o
}

// SetLogVolumeMaxLogSizeQueryParameter adds the logVolumeMaxLogSize to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogVolumeMaxLogSizeQueryParameter(logVolumeMaxLogSize *int64) {
	o.LogVolumeMaxLogSizeQueryParameter = logVolumeMaxLogSize
}

// WithLogVolumeRetentionPeriodQueryParameter adds the logVolumeRetentionPeriod to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogVolumeRetentionPeriodQueryParameter(logVolumeRetentionPeriod *string) *SnaplockLogCollectionGetParams {
	o.SetLogVolumeRetentionPeriodQueryParameter(logVolumeRetentionPeriod)
	return o
}

// SetLogVolumeRetentionPeriodQueryParameter adds the logVolumeRetentionPeriod to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogVolumeRetentionPeriodQueryParameter(logVolumeRetentionPeriod *string) {
	o.LogVolumeRetentionPeriodQueryParameter = logVolumeRetentionPeriod
}

// WithLogVolumeVolumeNameQueryParameter adds the logVolumeVolumeName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogVolumeVolumeNameQueryParameter(logVolumeVolumeName *string) *SnaplockLogCollectionGetParams {
	o.SetLogVolumeVolumeNameQueryParameter(logVolumeVolumeName)
	return o
}

// SetLogVolumeVolumeNameQueryParameter adds the logVolumeVolumeName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogVolumeVolumeNameQueryParameter(logVolumeVolumeName *string) {
	o.LogVolumeVolumeNameQueryParameter = logVolumeVolumeName
}

// WithLogVolumeVolumeUUIDQueryParameter adds the logVolumeVolumeUUID to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithLogVolumeVolumeUUIDQueryParameter(logVolumeVolumeUUID *string) *SnaplockLogCollectionGetParams {
	o.SetLogVolumeVolumeUUIDQueryParameter(logVolumeVolumeUUID)
	return o
}

// SetLogVolumeVolumeUUIDQueryParameter adds the logVolumeVolumeUuid to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetLogVolumeVolumeUUIDQueryParameter(logVolumeVolumeUUID *string) {
	o.LogVolumeVolumeUUIDQueryParameter = logVolumeVolumeUUID
}

// WithMaxRecords adds the maxRecords to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithMaxRecords(maxRecords *int64) *SnaplockLogCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithOrderBy(orderBy []string) *SnaplockLogCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithReturnRecords(returnRecords *bool) *SnaplockLogCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SnaplockLogCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *SnaplockLogCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SnaplockLogCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the snaplock log collection get params
func (o *SnaplockLogCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLogCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LogArchiveBaseNameQueryParameter != nil {

		// query param log_archive.base_name
		var qrLogArchiveBaseName string

		if o.LogArchiveBaseNameQueryParameter != nil {
			qrLogArchiveBaseName = *o.LogArchiveBaseNameQueryParameter
		}
		qLogArchiveBaseName := qrLogArchiveBaseName
		if qLogArchiveBaseName != "" {

			if err := r.SetQueryParam("log_archive.base_name", qLogArchiveBaseName); err != nil {
				return err
			}
		}
	}

	if o.LogArchiveExpiryTimeQueryParameter != nil {

		// query param log_archive.expiry_time
		var qrLogArchiveExpiryTime string

		if o.LogArchiveExpiryTimeQueryParameter != nil {
			qrLogArchiveExpiryTime = *o.LogArchiveExpiryTimeQueryParameter
		}
		qLogArchiveExpiryTime := qrLogArchiveExpiryTime
		if qLogArchiveExpiryTime != "" {

			if err := r.SetQueryParam("log_archive.expiry_time", qLogArchiveExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.LogArchivePathQueryParameter != nil {

		// query param log_archive.path
		var qrLogArchivePath string

		if o.LogArchivePathQueryParameter != nil {
			qrLogArchivePath = *o.LogArchivePathQueryParameter
		}
		qLogArchivePath := qrLogArchivePath
		if qLogArchivePath != "" {

			if err := r.SetQueryParam("log_archive.path", qLogArchivePath); err != nil {
				return err
			}
		}
	}

	if o.LogArchiveSizeQueryParameter != nil {

		// query param log_archive.size
		var qrLogArchiveSize int64

		if o.LogArchiveSizeQueryParameter != nil {
			qrLogArchiveSize = *o.LogArchiveSizeQueryParameter
		}
		qLogArchiveSize := swag.FormatInt64(qrLogArchiveSize)
		if qLogArchiveSize != "" {

			if err := r.SetQueryParam("log_archive.size", qLogArchiveSize); err != nil {
				return err
			}
		}
	}

	if o.LogFilesBaseNameQueryParameter != nil {

		// query param log_files.base_name
		var qrLogFilesBaseName string

		if o.LogFilesBaseNameQueryParameter != nil {
			qrLogFilesBaseName = *o.LogFilesBaseNameQueryParameter
		}
		qLogFilesBaseName := qrLogFilesBaseName
		if qLogFilesBaseName != "" {

			if err := r.SetQueryParam("log_files.base_name", qLogFilesBaseName); err != nil {
				return err
			}
		}
	}

	if o.LogFilesExpiryTimeQueryParameter != nil {

		// query param log_files.expiry_time
		var qrLogFilesExpiryTime string

		if o.LogFilesExpiryTimeQueryParameter != nil {
			qrLogFilesExpiryTime = *o.LogFilesExpiryTimeQueryParameter
		}
		qLogFilesExpiryTime := qrLogFilesExpiryTime
		if qLogFilesExpiryTime != "" {

			if err := r.SetQueryParam("log_files.expiry_time", qLogFilesExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.LogFilesPathQueryParameter != nil {

		// query param log_files.path
		var qrLogFilesPath string

		if o.LogFilesPathQueryParameter != nil {
			qrLogFilesPath = *o.LogFilesPathQueryParameter
		}
		qLogFilesPath := qrLogFilesPath
		if qLogFilesPath != "" {

			if err := r.SetQueryParam("log_files.path", qLogFilesPath); err != nil {
				return err
			}
		}
	}

	if o.LogFilesSizeQueryParameter != nil {

		// query param log_files.size
		var qrLogFilesSize int64

		if o.LogFilesSizeQueryParameter != nil {
			qrLogFilesSize = *o.LogFilesSizeQueryParameter
		}
		qLogFilesSize := swag.FormatInt64(qrLogFilesSize)
		if qLogFilesSize != "" {

			if err := r.SetQueryParam("log_files.size", qLogFilesSize); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeMaxLogSizeQueryParameter != nil {

		// query param log_volume.max_log_size
		var qrLogVolumeMaxLogSize int64

		if o.LogVolumeMaxLogSizeQueryParameter != nil {
			qrLogVolumeMaxLogSize = *o.LogVolumeMaxLogSizeQueryParameter
		}
		qLogVolumeMaxLogSize := swag.FormatInt64(qrLogVolumeMaxLogSize)
		if qLogVolumeMaxLogSize != "" {

			if err := r.SetQueryParam("log_volume.max_log_size", qLogVolumeMaxLogSize); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeRetentionPeriodQueryParameter != nil {

		// query param log_volume.retention_period
		var qrLogVolumeRetentionPeriod string

		if o.LogVolumeRetentionPeriodQueryParameter != nil {
			qrLogVolumeRetentionPeriod = *o.LogVolumeRetentionPeriodQueryParameter
		}
		qLogVolumeRetentionPeriod := qrLogVolumeRetentionPeriod
		if qLogVolumeRetentionPeriod != "" {

			if err := r.SetQueryParam("log_volume.retention_period", qLogVolumeRetentionPeriod); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeVolumeNameQueryParameter != nil {

		// query param log_volume.volume.name
		var qrLogVolumeVolumeName string

		if o.LogVolumeVolumeNameQueryParameter != nil {
			qrLogVolumeVolumeName = *o.LogVolumeVolumeNameQueryParameter
		}
		qLogVolumeVolumeName := qrLogVolumeVolumeName
		if qLogVolumeVolumeName != "" {

			if err := r.SetQueryParam("log_volume.volume.name", qLogVolumeVolumeName); err != nil {
				return err
			}
		}
	}

	if o.LogVolumeVolumeUUIDQueryParameter != nil {

		// query param log_volume.volume.uuid
		var qrLogVolumeVolumeUUID string

		if o.LogVolumeVolumeUUIDQueryParameter != nil {
			qrLogVolumeVolumeUUID = *o.LogVolumeVolumeUUIDQueryParameter
		}
		qLogVolumeVolumeUUID := qrLogVolumeVolumeUUID
		if qLogVolumeVolumeUUID != "" {

			if err := r.SetQueryParam("log_volume.volume.uuid", qLogVolumeVolumeUUID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockLogCollectionGet binds the parameter fields
func (o *SnaplockLogCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnaplockLogCollectionGet binds the parameter order_by
func (o *SnaplockLogCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
