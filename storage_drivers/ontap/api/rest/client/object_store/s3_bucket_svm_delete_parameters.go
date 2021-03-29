// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// NewS3BucketSvmDeleteParams creates a new S3BucketSvmDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3BucketSvmDeleteParams() *S3BucketSvmDeleteParams {
	return &S3BucketSvmDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3BucketSvmDeleteParamsWithTimeout creates a new S3BucketSvmDeleteParams object
// with the ability to set a timeout on a request.
func NewS3BucketSvmDeleteParamsWithTimeout(timeout time.Duration) *S3BucketSvmDeleteParams {
	return &S3BucketSvmDeleteParams{
		timeout: timeout,
	}
}

// NewS3BucketSvmDeleteParamsWithContext creates a new S3BucketSvmDeleteParams object
// with the ability to set a context for a request.
func NewS3BucketSvmDeleteParamsWithContext(ctx context.Context) *S3BucketSvmDeleteParams {
	return &S3BucketSvmDeleteParams{
		Context: ctx,
	}
}

// NewS3BucketSvmDeleteParamsWithHTTPClient creates a new S3BucketSvmDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3BucketSvmDeleteParamsWithHTTPClient(client *http.Client) *S3BucketSvmDeleteParams {
	return &S3BucketSvmDeleteParams{
		HTTPClient: client,
	}
}

/* S3BucketSvmDeleteParams contains all the parameters to send to the API endpoint
   for the s3 bucket svm delete operation.

   Typically these are written to a http.Request.
*/
type S3BucketSvmDeleteParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* UUID.

	   The unique identifier of the bucket.
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 bucket svm delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketSvmDeleteParams) WithDefaults() *S3BucketSvmDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 bucket svm delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketSvmDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := S3BucketSvmDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) WithTimeout(timeout time.Duration) *S3BucketSvmDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) WithContext(ctx context.Context) *S3BucketSvmDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) WithHTTPClient(client *http.Client) *S3BucketSvmDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeout adds the returnTimeout to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) WithReturnTimeout(returnTimeout *int64) *S3BucketSvmDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) WithSvmUUID(svmUUID string) *S3BucketSvmDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithUUIDPathParameter adds the uuid to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) WithUUIDPathParameter(uuid string) *S3BucketSvmDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the s3 bucket svm delete params
func (o *S3BucketSvmDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *S3BucketSvmDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
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
