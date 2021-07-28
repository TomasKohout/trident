// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmsEventResponse ems event response
//
// swagger:model ems_event_response
type EmsEventResponse struct {

	// links
	Links *EmsEventResponseLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 3
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*EmsEventResponseRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this ems event response
func (m *EmsEventResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponse) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems event response based on the context it is used
func (m *EmsEventResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponse) UnmarshalBinary(b []byte) error {
	var res EmsEventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseLinks ems event response links
//
// swagger:model EmsEventResponseLinks
type EmsEventResponseLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems event response links
func (m *EmsEventResponseLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems event response links based on the context it is used
func (m *EmsEventResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0 ems event response records items0
//
// swagger:model EmsEventResponseRecordsItems0
type EmsEventResponseRecordsItems0 struct {

	// links
	Links *EmsEventResponseRecordsItems0Links `json:"_links,omitempty"`

	// Index of the event. Returned by default.
	// Example: 1
	// Read Only: true
	Index int64 `json:"index,omitempty"`

	// A formatted text string populated with parameter details. Returned by default.
	// Read Only: true
	LogMessage string `json:"log_message,omitempty"`

	// message
	Message *EmsEventResponseRecordsItems0Message `json:"message,omitempty"`

	// node
	Node *EmsEventResponseRecordsItems0Node `json:"node,omitempty"`

	// A list of parameters provided with the EMS event.
	Parameters []*EmsEventResponseRecordsItems0ParametersItems0 `json:"parameters,omitempty"`

	// Source
	// Read Only: true
	Source string `json:"source,omitempty"`

	// Timestamp of the event. Returned by default.
	// Read Only: true
	Time string `json:"time,omitempty"`
}

// Validate validates this ems event response records items0
func (m *EmsEventResponseRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems event response records items0 based on the context it is used
func (m *EmsEventResponseRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", int64(m.Index)); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateLogMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "log_message", "body", string(m.LogMessage)); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.Message != nil {
		if err := m.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source", "body", string(m.Source)); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "time", "body", string(m.Time)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0Links ems event response records items0 links
//
// swagger:model EmsEventResponseRecordsItems0Links
type EmsEventResponseRecordsItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems event response records items0 links
func (m *EmsEventResponseRecordsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems event response records items0 links based on the context it is used
func (m *EmsEventResponseRecordsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0Links) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0Message ems event response records items0 message
//
// swagger:model EmsEventResponseRecordsItems0Message
type EmsEventResponseRecordsItems0Message struct {

	// links
	Links *EmsEventResponseRecordsItems0MessageLinks `json:"_links,omitempty"`

	// Message name of the event. Returned by default.
	// Example: callhome.spares.low
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Severity of the event. Returned by default.
	// Example: emergency
	// Read Only: true
	// Enum: [emergency alert error notice informational debug]
	Severity string `json:"severity,omitempty"`
}

// Validate validates this ems event response records items0 message
func (m *EmsEventResponseRecordsItems0Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0Message) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var emsEventResponseRecordsItems0MessageTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emergency","alert","error","notice","informational","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsEventResponseRecordsItems0MessageTypeSeverityPropEnum = append(emsEventResponseRecordsItems0MessageTypeSeverityPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0Message
	// EmsEventResponseRecordsItems0Message
	// severity
	// Severity
	// emergency
	// END RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0MessageSeverityEmergency captures enum value "emergency"
	EmsEventResponseRecordsItems0MessageSeverityEmergency string = "emergency"

	// BEGIN RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0Message
	// EmsEventResponseRecordsItems0Message
	// severity
	// Severity
	// alert
	// END RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0MessageSeverityAlert captures enum value "alert"
	EmsEventResponseRecordsItems0MessageSeverityAlert string = "alert"

	// BEGIN RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0Message
	// EmsEventResponseRecordsItems0Message
	// severity
	// Severity
	// error
	// END RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0MessageSeverityError captures enum value "error"
	EmsEventResponseRecordsItems0MessageSeverityError string = "error"

	// BEGIN RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0Message
	// EmsEventResponseRecordsItems0Message
	// severity
	// Severity
	// notice
	// END RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0MessageSeverityNotice captures enum value "notice"
	EmsEventResponseRecordsItems0MessageSeverityNotice string = "notice"

	// BEGIN RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0Message
	// EmsEventResponseRecordsItems0Message
	// severity
	// Severity
	// informational
	// END RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0MessageSeverityInformational captures enum value "informational"
	EmsEventResponseRecordsItems0MessageSeverityInformational string = "informational"

	// BEGIN RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0Message
	// EmsEventResponseRecordsItems0Message
	// severity
	// Severity
	// debug
	// END RIPPY DEBUGGING
	// EmsEventResponseRecordsItems0MessageSeverityDebug captures enum value "debug"
	EmsEventResponseRecordsItems0MessageSeverityDebug string = "debug"
)

// prop value enum
func (m *EmsEventResponseRecordsItems0Message) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsEventResponseRecordsItems0MessageTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0Message) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("message"+"."+"severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems event response records items0 message based on the context it is used
func (m *EmsEventResponseRecordsItems0Message) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0Message) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0Message) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message"+"."+"name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0Message) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message"+"."+"severity", "body", string(m.Severity)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0Message) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0MessageLinks ems event response records items0 message links
//
// swagger:model EmsEventResponseRecordsItems0MessageLinks
type EmsEventResponseRecordsItems0MessageLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems event response records items0 message links
func (m *EmsEventResponseRecordsItems0MessageLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0MessageLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems event response records items0 message links based on the context it is used
func (m *EmsEventResponseRecordsItems0MessageLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0MessageLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0MessageLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0MessageLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0MessageLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0Node ems event response records items0 node
//
// swagger:model EmsEventResponseRecordsItems0Node
type EmsEventResponseRecordsItems0Node struct {

	// links
	Links *EmsEventResponseRecordsItems0NodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this ems event response records items0 node
func (m *EmsEventResponseRecordsItems0Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0Node) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems event response records items0 node based on the context it is used
func (m *EmsEventResponseRecordsItems0Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0Node) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0Node) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0NodeLinks ems event response records items0 node links
//
// swagger:model EmsEventResponseRecordsItems0NodeLinks
type EmsEventResponseRecordsItems0NodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems event response records items0 node links
func (m *EmsEventResponseRecordsItems0NodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0NodeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems event response records items0 node links based on the context it is used
func (m *EmsEventResponseRecordsItems0NodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0NodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0NodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0NodeLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0NodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0ParametersItems0 ems event response records items0 parameters items0
//
// swagger:model EmsEventResponseRecordsItems0ParametersItems0
type EmsEventResponseRecordsItems0ParametersItems0 struct {

	// Name of parameter
	// Example: numOps
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Value of parameter
	// Example: 123
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this ems event response records items0 parameters items0
func (m *EmsEventResponseRecordsItems0ParametersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems event response records items0 parameters items0 based on the context it is used
func (m *EmsEventResponseRecordsItems0ParametersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0ParametersItems0) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0ParametersItems0) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", string(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0ParametersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0ParametersItems0) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0ParametersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY