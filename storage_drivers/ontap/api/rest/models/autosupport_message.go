// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutosupportMessage autosupport message
//
// swagger:model autosupport_message
type AutosupportMessage struct {

	// Destination for the AutoSupport
	// Example: http
	// Read Only: true
	// Enum: [smtp http noteto retransmit]
	Destination string `json:"destination,omitempty"`

	// error
	Error *AutosupportMessageError `json:"error,omitempty"`

	// Date and Time of AutoSupport generation in ISO-8601 format
	// Example: 2019-03-25 21:30:04
	// Read Only: true
	// Format: date-time
	GeneratedOn *strfmt.DateTime `json:"generated_on,omitempty"`

	// Sequence number of the AutoSupport
	// Example: 9
	// Read Only: true
	Index int64 `json:"index,omitempty"`

	// Message included in the AutoSupport subject
	// Example: invoked_test_autosupport_rest
	Message string `json:"message,omitempty"`

	// node
	Node *AutosupportMessageNode `json:"node,omitempty"`

	// State of AutoSupport delivery
	// Example: sent_successful
	// Read Only: true
	// Enum: [initializing collection_failed collection_in_progress queued transmitting sent_successful ignore re_queued transmission_failed ondemand_ignore cancelled]
	State string `json:"state,omitempty"`

	// Subject line for the AutoSupport
	// Example: WEEKLY_LOG
	// Read Only: true
	Subject string `json:"subject,omitempty"`

	// Type of AutoSupport collection to issue
	// Example: test
	// Enum: [test performance all]
	Type *string `json:"type,omitempty"`

	// Alternate destination for the AutoSupport
	// Example: http://1.2.3.4/delivery_uri
	// Format: uri
	URI strfmt.URI `json:"uri,omitempty"`
}

// Validate validates this autosupport message
func (m *AutosupportMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneratedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var autosupportMessageTypeDestinationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["smtp","http","noteto","retransmit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autosupportMessageTypeDestinationPropEnum = append(autosupportMessageTypeDestinationPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// destination
	// Destination
	// smtp
	// END RIPPY DEBUGGING
	// AutosupportMessageDestinationSMTP captures enum value "smtp"
	AutosupportMessageDestinationSMTP string = "smtp"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// destination
	// Destination
	// http
	// END RIPPY DEBUGGING
	// AutosupportMessageDestinationHTTP captures enum value "http"
	AutosupportMessageDestinationHTTP string = "http"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// destination
	// Destination
	// noteto
	// END RIPPY DEBUGGING
	// AutosupportMessageDestinationNoteto captures enum value "noteto"
	AutosupportMessageDestinationNoteto string = "noteto"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// destination
	// Destination
	// retransmit
	// END RIPPY DEBUGGING
	// AutosupportMessageDestinationRetransmit captures enum value "retransmit"
	AutosupportMessageDestinationRetransmit string = "retransmit"
)

// prop value enum
func (m *AutosupportMessage) validateDestinationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autosupportMessageTypeDestinationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutosupportMessage) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	// value enum
	if err := m.validateDestinationEnum("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessage) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *AutosupportMessage) validateGeneratedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.GeneratedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("generated_on", "body", "date-time", m.GeneratedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessage) validateNode(formats strfmt.Registry) error {
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

var autosupportMessageTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["initializing","collection_failed","collection_in_progress","queued","transmitting","sent_successful","ignore","re_queued","transmission_failed","ondemand_ignore","cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autosupportMessageTypeStatePropEnum = append(autosupportMessageTypeStatePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// initializing
	// END RIPPY DEBUGGING
	// AutosupportMessageStateInitializing captures enum value "initializing"
	AutosupportMessageStateInitializing string = "initializing"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// collection_failed
	// END RIPPY DEBUGGING
	// AutosupportMessageStateCollectionFailed captures enum value "collection_failed"
	AutosupportMessageStateCollectionFailed string = "collection_failed"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// collection_in_progress
	// END RIPPY DEBUGGING
	// AutosupportMessageStateCollectionInProgress captures enum value "collection_in_progress"
	AutosupportMessageStateCollectionInProgress string = "collection_in_progress"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// queued
	// END RIPPY DEBUGGING
	// AutosupportMessageStateQueued captures enum value "queued"
	AutosupportMessageStateQueued string = "queued"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// transmitting
	// END RIPPY DEBUGGING
	// AutosupportMessageStateTransmitting captures enum value "transmitting"
	AutosupportMessageStateTransmitting string = "transmitting"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// sent_successful
	// END RIPPY DEBUGGING
	// AutosupportMessageStateSentSuccessful captures enum value "sent_successful"
	AutosupportMessageStateSentSuccessful string = "sent_successful"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// ignore
	// END RIPPY DEBUGGING
	// AutosupportMessageStateIgnore captures enum value "ignore"
	AutosupportMessageStateIgnore string = "ignore"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// re_queued
	// END RIPPY DEBUGGING
	// AutosupportMessageStateReQueued captures enum value "re_queued"
	AutosupportMessageStateReQueued string = "re_queued"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// transmission_failed
	// END RIPPY DEBUGGING
	// AutosupportMessageStateTransmissionFailed captures enum value "transmission_failed"
	AutosupportMessageStateTransmissionFailed string = "transmission_failed"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// ondemand_ignore
	// END RIPPY DEBUGGING
	// AutosupportMessageStateOndemandIgnore captures enum value "ondemand_ignore"
	AutosupportMessageStateOndemandIgnore string = "ondemand_ignore"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// state
	// State
	// cancelled
	// END RIPPY DEBUGGING
	// AutosupportMessageStateCancelled captures enum value "cancelled"
	AutosupportMessageStateCancelled string = "cancelled"
)

// prop value enum
func (m *AutosupportMessage) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autosupportMessageTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutosupportMessage) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var autosupportMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["test","performance","all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autosupportMessageTypeTypePropEnum = append(autosupportMessageTypeTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// type
	// Type
	// test
	// END RIPPY DEBUGGING
	// AutosupportMessageTypeTest captures enum value "test"
	AutosupportMessageTypeTest string = "test"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// type
	// Type
	// performance
	// END RIPPY DEBUGGING
	// AutosupportMessageTypePerformance captures enum value "performance"
	AutosupportMessageTypePerformance string = "performance"

	// BEGIN RIPPY DEBUGGING
	// autosupport_message
	// AutosupportMessage
	// type
	// Type
	// all
	// END RIPPY DEBUGGING
	// AutosupportMessageTypeAll captures enum value "all"
	AutosupportMessageTypeAll string = "all"
)

// prop value enum
func (m *AutosupportMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autosupportMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutosupportMessage) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessage) validateURI(formats strfmt.Registry) error {
	if swag.IsZero(m.URI) { // not required
		return nil
	}

	if err := validate.FormatOf("uri", "body", "uri", m.URI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this autosupport message based on the context it is used
func (m *AutosupportMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeneratedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessage) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "destination", "body", string(m.Destination)); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessage) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *AutosupportMessage) contextValidateGeneratedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "generated_on", "body", m.GeneratedOn); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessage) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", int64(m.Index)); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessage) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AutosupportMessage) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessage) contextValidateSubject(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subject", "body", string(m.Subject)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutosupportMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportMessage) UnmarshalBinary(b []byte) error {
	var res AutosupportMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutosupportMessageError Last error during delivery attempt. Empty if "status=sent-successful".
//
// swagger:model AutosupportMessageError
type AutosupportMessageError struct {

	// Error code
	// Example: 53149746
	// Read Only: true
	Code int64 `json:"code,omitempty"`

	// Error message
	// Example: Could not resolve host: test.com
	// Read Only: true
	Message string `json:"message,omitempty"`
}

// Validate validates this autosupport message error
func (m *AutosupportMessageError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this autosupport message error based on the context it is used
func (m *AutosupportMessageError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessageError) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "error"+"."+"code", "body", int64(m.Code)); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportMessageError) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "error"+"."+"message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutosupportMessageError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportMessageError) UnmarshalBinary(b []byte) error {
	var res AutosupportMessageError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutosupportMessageNode autosupport message node
//
// swagger:model AutosupportMessageNode
type AutosupportMessageNode struct {

	// links
	Links *AutosupportMessageNodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this autosupport message node
func (m *AutosupportMessageNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessageNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this autosupport message node based on the context it is used
func (m *AutosupportMessageNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessageNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AutosupportMessageNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportMessageNode) UnmarshalBinary(b []byte) error {
	var res AutosupportMessageNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutosupportMessageNodeLinks autosupport message node links
//
// swagger:model AutosupportMessageNodeLinks
type AutosupportMessageNodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this autosupport message node links
func (m *AutosupportMessageNodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessageNodeLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this autosupport message node links based on the context it is used
func (m *AutosupportMessageNodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessageNodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AutosupportMessageNodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportMessageNodeLinks) UnmarshalBinary(b []byte) error {
	var res AutosupportMessageNodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
