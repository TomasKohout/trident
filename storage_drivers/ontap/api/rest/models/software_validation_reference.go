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

// SoftwareValidationReference software validation reference
//
// swagger:model software_validation_reference
type SoftwareValidationReference struct {

	// action
	Action *SoftwareValidationReferenceAction `json:"action,omitempty"`

	// issue
	Issue *SoftwareValidationReferenceIssue `json:"issue,omitempty"`

	// Status of this update check.
	// Example: warning
	// Read Only: true
	// Enum: [warning error]
	Status string `json:"status,omitempty"`

	// Name of the update check to be validated.
	// Example: nfs_mounts
	// Read Only: true
	UpdateCheck string `json:"update_check,omitempty"`
}

// Validate validates this software validation reference
func (m *SoftwareValidationReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareValidationReference) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareValidationReference) validateIssue(formats strfmt.Registry) error {
	if swag.IsZero(m.Issue) { // not required
		return nil
	}

	if m.Issue != nil {
		if err := m.Issue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

var softwareValidationReferenceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["warning","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareValidationReferenceTypeStatusPropEnum = append(softwareValidationReferenceTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// software_validation_reference
	// SoftwareValidationReference
	// status
	// Status
	// warning
	// END RIPPY DEBUGGING
	// SoftwareValidationReferenceStatusWarning captures enum value "warning"
	SoftwareValidationReferenceStatusWarning string = "warning"

	// BEGIN RIPPY DEBUGGING
	// software_validation_reference
	// SoftwareValidationReference
	// status
	// Status
	// error
	// END RIPPY DEBUGGING
	// SoftwareValidationReferenceStatusError captures enum value "error"
	SoftwareValidationReferenceStatusError string = "error"
)

// prop value enum
func (m *SoftwareValidationReference) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwareValidationReferenceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwareValidationReference) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this software validation reference based on the context it is used
func (m *SoftwareValidationReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateCheck(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareValidationReference) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {
		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareValidationReference) contextValidateIssue(ctx context.Context, formats strfmt.Registry) error {

	if m.Issue != nil {
		if err := m.Issue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareValidationReference) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareValidationReference) contextValidateUpdateCheck(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "update_check", "body", string(m.UpdateCheck)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareValidationReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareValidationReference) UnmarshalBinary(b []byte) error {
	var res SoftwareValidationReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareValidationReferenceAction software validation reference action
//
// swagger:model SoftwareValidationReferenceAction
type SoftwareValidationReferenceAction struct {

	// Specifies the corrective action to be taken to resolve a validation error
	Message string `json:"message,omitempty"`
}

// Validate validates this software validation reference action
func (m *SoftwareValidationReferenceAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software validation reference action based on the context it is used
func (m *SoftwareValidationReferenceAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareValidationReferenceAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareValidationReferenceAction) UnmarshalBinary(b []byte) error {
	var res SoftwareValidationReferenceAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareValidationReferenceIssue software validation reference issue
//
// swagger:model SoftwareValidationReferenceIssue
type SoftwareValidationReferenceIssue struct {

	// Details of the error or warning encountered by the update checks
	// Example: Validation error: Cluster HA is not configured in the cluster
	Message string `json:"message,omitempty"`
}

// Validate validates this software validation reference issue
func (m *SoftwareValidationReferenceIssue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software validation reference issue based on the context it is used
func (m *SoftwareValidationReferenceIssue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareValidationReferenceIssue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareValidationReferenceIssue) UnmarshalBinary(b []byte) error {
	var res SoftwareValidationReferenceIssue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
