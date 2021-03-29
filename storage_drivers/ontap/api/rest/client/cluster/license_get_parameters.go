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

// NewLicenseGetParams creates a new LicenseGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLicenseGetParams() *LicenseGetParams {
	return &LicenseGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLicenseGetParamsWithTimeout creates a new LicenseGetParams object
// with the ability to set a timeout on a request.
func NewLicenseGetParamsWithTimeout(timeout time.Duration) *LicenseGetParams {
	return &LicenseGetParams{
		timeout: timeout,
	}
}

// NewLicenseGetParamsWithContext creates a new LicenseGetParams object
// with the ability to set a context for a request.
func NewLicenseGetParamsWithContext(ctx context.Context) *LicenseGetParams {
	return &LicenseGetParams{
		Context: ctx,
	}
}

// NewLicenseGetParamsWithHTTPClient creates a new LicenseGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLicenseGetParamsWithHTTPClient(client *http.Client) *LicenseGetParams {
	return &LicenseGetParams{
		HTTPClient: client,
	}
}

/* LicenseGetParams contains all the parameters to send to the API endpoint
   for the license get operation.

   Typically these are written to a http.Request.
*/
type LicenseGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LicensesActive.

	   Filter by licenses.active
	*/
	LicensesActiveQueryParameter *bool

	/* LicensesCapacityMaximumSize.

	   Filter by licenses.capacity.maximum_size
	*/
	LicensesCapacityMaximumSizeQueryParameter *int64

	/* LicensesCapacityUsedSize.

	   Filter by licenses.capacity.used_size
	*/
	LicensesCapacityUsedSizeQueryParameter *int64

	/* LicensesComplianceState.

	   Filter by licenses.compliance.state
	*/
	LicensesComplianceStateQueryParameter *string

	/* LicensesEvaluation.

	   Filter by licenses.evaluation
	*/
	LicensesEvaluationQueryParameter *bool

	/* LicensesExpiryTime.

	   Filter by licenses.expiry_time
	*/
	LicensesExpiryTimeQueryParameter *string

	/* LicensesOwner.

	   Filter by licenses.owner
	*/
	LicensesOwnerQueryParameter *string

	/* LicensesSerialNumber.

	   Filter by licenses.serial_number
	*/
	LicensesSerialNumberQueryParameter *string

	/* LicensesStartTime.

	   Filter by licenses.start_time
	*/
	LicensesStartTimeQueryParameter *string

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* Name.

	   Name of the license package.
	*/
	NamePathParameter string

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the license get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseGetParams) WithDefaults() *LicenseGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the license get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the license get params
func (o *LicenseGetParams) WithTimeout(timeout time.Duration) *LicenseGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the license get params
func (o *LicenseGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the license get params
func (o *LicenseGetParams) WithContext(ctx context.Context) *LicenseGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the license get params
func (o *LicenseGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the license get params
func (o *LicenseGetParams) WithHTTPClient(client *http.Client) *LicenseGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the license get params
func (o *LicenseGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the license get params
func (o *LicenseGetParams) WithFields(fields []string) *LicenseGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the license get params
func (o *LicenseGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLicensesActiveQueryParameter adds the licensesActive to the license get params
func (o *LicenseGetParams) WithLicensesActiveQueryParameter(licensesActive *bool) *LicenseGetParams {
	o.SetLicensesActiveQueryParameter(licensesActive)
	return o
}

// SetLicensesActiveQueryParameter adds the licensesActive to the license get params
func (o *LicenseGetParams) SetLicensesActiveQueryParameter(licensesActive *bool) {
	o.LicensesActiveQueryParameter = licensesActive
}

// WithLicensesCapacityMaximumSizeQueryParameter adds the licensesCapacityMaximumSize to the license get params
func (o *LicenseGetParams) WithLicensesCapacityMaximumSizeQueryParameter(licensesCapacityMaximumSize *int64) *LicenseGetParams {
	o.SetLicensesCapacityMaximumSizeQueryParameter(licensesCapacityMaximumSize)
	return o
}

// SetLicensesCapacityMaximumSizeQueryParameter adds the licensesCapacityMaximumSize to the license get params
func (o *LicenseGetParams) SetLicensesCapacityMaximumSizeQueryParameter(licensesCapacityMaximumSize *int64) {
	o.LicensesCapacityMaximumSizeQueryParameter = licensesCapacityMaximumSize
}

// WithLicensesCapacityUsedSizeQueryParameter adds the licensesCapacityUsedSize to the license get params
func (o *LicenseGetParams) WithLicensesCapacityUsedSizeQueryParameter(licensesCapacityUsedSize *int64) *LicenseGetParams {
	o.SetLicensesCapacityUsedSizeQueryParameter(licensesCapacityUsedSize)
	return o
}

// SetLicensesCapacityUsedSizeQueryParameter adds the licensesCapacityUsedSize to the license get params
func (o *LicenseGetParams) SetLicensesCapacityUsedSizeQueryParameter(licensesCapacityUsedSize *int64) {
	o.LicensesCapacityUsedSizeQueryParameter = licensesCapacityUsedSize
}

// WithLicensesComplianceStateQueryParameter adds the licensesComplianceState to the license get params
func (o *LicenseGetParams) WithLicensesComplianceStateQueryParameter(licensesComplianceState *string) *LicenseGetParams {
	o.SetLicensesComplianceStateQueryParameter(licensesComplianceState)
	return o
}

// SetLicensesComplianceStateQueryParameter adds the licensesComplianceState to the license get params
func (o *LicenseGetParams) SetLicensesComplianceStateQueryParameter(licensesComplianceState *string) {
	o.LicensesComplianceStateQueryParameter = licensesComplianceState
}

// WithLicensesEvaluationQueryParameter adds the licensesEvaluation to the license get params
func (o *LicenseGetParams) WithLicensesEvaluationQueryParameter(licensesEvaluation *bool) *LicenseGetParams {
	o.SetLicensesEvaluationQueryParameter(licensesEvaluation)
	return o
}

// SetLicensesEvaluationQueryParameter adds the licensesEvaluation to the license get params
func (o *LicenseGetParams) SetLicensesEvaluationQueryParameter(licensesEvaluation *bool) {
	o.LicensesEvaluationQueryParameter = licensesEvaluation
}

// WithLicensesExpiryTimeQueryParameter adds the licensesExpiryTime to the license get params
func (o *LicenseGetParams) WithLicensesExpiryTimeQueryParameter(licensesExpiryTime *string) *LicenseGetParams {
	o.SetLicensesExpiryTimeQueryParameter(licensesExpiryTime)
	return o
}

// SetLicensesExpiryTimeQueryParameter adds the licensesExpiryTime to the license get params
func (o *LicenseGetParams) SetLicensesExpiryTimeQueryParameter(licensesExpiryTime *string) {
	o.LicensesExpiryTimeQueryParameter = licensesExpiryTime
}

// WithLicensesOwnerQueryParameter adds the licensesOwner to the license get params
func (o *LicenseGetParams) WithLicensesOwnerQueryParameter(licensesOwner *string) *LicenseGetParams {
	o.SetLicensesOwnerQueryParameter(licensesOwner)
	return o
}

// SetLicensesOwnerQueryParameter adds the licensesOwner to the license get params
func (o *LicenseGetParams) SetLicensesOwnerQueryParameter(licensesOwner *string) {
	o.LicensesOwnerQueryParameter = licensesOwner
}

// WithLicensesSerialNumberQueryParameter adds the licensesSerialNumber to the license get params
func (o *LicenseGetParams) WithLicensesSerialNumberQueryParameter(licensesSerialNumber *string) *LicenseGetParams {
	o.SetLicensesSerialNumberQueryParameter(licensesSerialNumber)
	return o
}

// SetLicensesSerialNumberQueryParameter adds the licensesSerialNumber to the license get params
func (o *LicenseGetParams) SetLicensesSerialNumberQueryParameter(licensesSerialNumber *string) {
	o.LicensesSerialNumberQueryParameter = licensesSerialNumber
}

// WithLicensesStartTimeQueryParameter adds the licensesStartTime to the license get params
func (o *LicenseGetParams) WithLicensesStartTimeQueryParameter(licensesStartTime *string) *LicenseGetParams {
	o.SetLicensesStartTimeQueryParameter(licensesStartTime)
	return o
}

// SetLicensesStartTimeQueryParameter adds the licensesStartTime to the license get params
func (o *LicenseGetParams) SetLicensesStartTimeQueryParameter(licensesStartTime *string) {
	o.LicensesStartTimeQueryParameter = licensesStartTime
}

// WithNameQueryParameter adds the name to the license get params
func (o *LicenseGetParams) WithNameQueryParameter(name *string) *LicenseGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the license get params
func (o *LicenseGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithNamePathParameter adds the name to the license get params
func (o *LicenseGetParams) WithNamePathParameter(name string) *LicenseGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the license get params
func (o *LicenseGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithScopeQueryParameter adds the scope to the license get params
func (o *LicenseGetParams) WithScopeQueryParameter(scope *string) *LicenseGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the license get params
func (o *LicenseGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithStateQueryParameter adds the state to the license get params
func (o *LicenseGetParams) WithStateQueryParameter(state *string) *LicenseGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the license get params
func (o *LicenseGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WriteToRequest writes these params to a swagger request
func (o *LicenseGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LicensesActiveQueryParameter != nil {

		// query param licenses.active
		var qrLicensesActive bool

		if o.LicensesActiveQueryParameter != nil {
			qrLicensesActive = *o.LicensesActiveQueryParameter
		}
		qLicensesActive := swag.FormatBool(qrLicensesActive)
		if qLicensesActive != "" {

			if err := r.SetQueryParam("licenses.active", qLicensesActive); err != nil {
				return err
			}
		}
	}

	if o.LicensesCapacityMaximumSizeQueryParameter != nil {

		// query param licenses.capacity.maximum_size
		var qrLicensesCapacityMaximumSize int64

		if o.LicensesCapacityMaximumSizeQueryParameter != nil {
			qrLicensesCapacityMaximumSize = *o.LicensesCapacityMaximumSizeQueryParameter
		}
		qLicensesCapacityMaximumSize := swag.FormatInt64(qrLicensesCapacityMaximumSize)
		if qLicensesCapacityMaximumSize != "" {

			if err := r.SetQueryParam("licenses.capacity.maximum_size", qLicensesCapacityMaximumSize); err != nil {
				return err
			}
		}
	}

	if o.LicensesCapacityUsedSizeQueryParameter != nil {

		// query param licenses.capacity.used_size
		var qrLicensesCapacityUsedSize int64

		if o.LicensesCapacityUsedSizeQueryParameter != nil {
			qrLicensesCapacityUsedSize = *o.LicensesCapacityUsedSizeQueryParameter
		}
		qLicensesCapacityUsedSize := swag.FormatInt64(qrLicensesCapacityUsedSize)
		if qLicensesCapacityUsedSize != "" {

			if err := r.SetQueryParam("licenses.capacity.used_size", qLicensesCapacityUsedSize); err != nil {
				return err
			}
		}
	}

	if o.LicensesComplianceStateQueryParameter != nil {

		// query param licenses.compliance.state
		var qrLicensesComplianceState string

		if o.LicensesComplianceStateQueryParameter != nil {
			qrLicensesComplianceState = *o.LicensesComplianceStateQueryParameter
		}
		qLicensesComplianceState := qrLicensesComplianceState
		if qLicensesComplianceState != "" {

			if err := r.SetQueryParam("licenses.compliance.state", qLicensesComplianceState); err != nil {
				return err
			}
		}
	}

	if o.LicensesEvaluationQueryParameter != nil {

		// query param licenses.evaluation
		var qrLicensesEvaluation bool

		if o.LicensesEvaluationQueryParameter != nil {
			qrLicensesEvaluation = *o.LicensesEvaluationQueryParameter
		}
		qLicensesEvaluation := swag.FormatBool(qrLicensesEvaluation)
		if qLicensesEvaluation != "" {

			if err := r.SetQueryParam("licenses.evaluation", qLicensesEvaluation); err != nil {
				return err
			}
		}
	}

	if o.LicensesExpiryTimeQueryParameter != nil {

		// query param licenses.expiry_time
		var qrLicensesExpiryTime string

		if o.LicensesExpiryTimeQueryParameter != nil {
			qrLicensesExpiryTime = *o.LicensesExpiryTimeQueryParameter
		}
		qLicensesExpiryTime := qrLicensesExpiryTime
		if qLicensesExpiryTime != "" {

			if err := r.SetQueryParam("licenses.expiry_time", qLicensesExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.LicensesOwnerQueryParameter != nil {

		// query param licenses.owner
		var qrLicensesOwner string

		if o.LicensesOwnerQueryParameter != nil {
			qrLicensesOwner = *o.LicensesOwnerQueryParameter
		}
		qLicensesOwner := qrLicensesOwner
		if qLicensesOwner != "" {

			if err := r.SetQueryParam("licenses.owner", qLicensesOwner); err != nil {
				return err
			}
		}
	}

	if o.LicensesSerialNumberQueryParameter != nil {

		// query param licenses.serial_number
		var qrLicensesSerialNumber string

		if o.LicensesSerialNumberQueryParameter != nil {
			qrLicensesSerialNumber = *o.LicensesSerialNumberQueryParameter
		}
		qLicensesSerialNumber := qrLicensesSerialNumber
		if qLicensesSerialNumber != "" {

			if err := r.SetQueryParam("licenses.serial_number", qLicensesSerialNumber); err != nil {
				return err
			}
		}
	}

	if o.LicensesStartTimeQueryParameter != nil {

		// query param licenses.start_time
		var qrLicensesStartTime string

		if o.LicensesStartTimeQueryParameter != nil {
			qrLicensesStartTime = *o.LicensesStartTimeQueryParameter
		}
		qLicensesStartTime := qrLicensesStartTime
		if qLicensesStartTime != "" {

			if err := r.SetQueryParam("licenses.start_time", qLicensesStartTime); err != nil {
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

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLicenseGet binds the parameter fields
func (o *LicenseGetParams) bindParamFields(formats strfmt.Registry) []string {
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
