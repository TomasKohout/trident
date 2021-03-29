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

// ApplicationSnapshot application snapshot
//
// swagger:model application_snapshot
type ApplicationSnapshot struct {

	// links
	Links *ApplicationSnapshotLinks `json:"_links,omitempty"`

	// application
	Application *ApplicationSnapshotApplication `json:"application,omitempty"`

	// Comment. Valid in POST.
	// Max Length: 255
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// components
	// Read Only: true
	Components []*ApplicationSnapshotComponentsItems0 `json:"components,omitempty"`

	// Consistency type. This is for categorization purposes only. A Snapshot copy should not be set to 'application consistent' unless the host application is quiesced for the Snapshot copy. Valid in POST.
	// Enum: [crash application]
	ConsistencyType string `json:"consistency_type,omitempty"`

	// Creation time
	// Read Only: true
	CreateTime string `json:"create_time,omitempty"`

	// A partial Snapshot copy means that not all volumes in an application component were included in the Snapshot copy.
	// Read Only: true
	IsPartial *bool `json:"is_partial,omitempty"`

	// The Snapshot copy name. Valid in POST.
	Name string `json:"name,omitempty"`

	// svm
	Svm *ApplicationSnapshotSvm `json:"svm,omitempty"`

	// The Snapshot copy UUID. Valid in URL.
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this application snapshot
func (m *ApplicationSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshot) validateLinks(formats strfmt.Registry) error {
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

func (m *ApplicationSnapshot) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSnapshot) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 255); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSnapshot) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var applicationSnapshotTypeConsistencyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["crash","application"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationSnapshotTypeConsistencyTypePropEnum = append(applicationSnapshotTypeConsistencyTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// application_snapshot
	// ApplicationSnapshot
	// consistency_type
	// ConsistencyType
	// crash
	// END RIPPY DEBUGGING
	// ApplicationSnapshotConsistencyTypeCrash captures enum value "crash"
	ApplicationSnapshotConsistencyTypeCrash string = "crash"

	// BEGIN RIPPY DEBUGGING
	// application_snapshot
	// ApplicationSnapshot
	// consistency_type
	// ConsistencyType
	// application
	// END RIPPY DEBUGGING
	// ApplicationSnapshotConsistencyTypeApplication captures enum value "application"
	ApplicationSnapshotConsistencyTypeApplication string = "application"
)

// prop value enum
func (m *ApplicationSnapshot) validateConsistencyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationSnapshotTypeConsistencyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationSnapshot) validateConsistencyType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConsistencyTypeEnum("consistency_type", "body", m.ConsistencyType); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSnapshot) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application snapshot based on the context it is used
func (m *ApplicationSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsPartial(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshot) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationSnapshot) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {
		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSnapshot) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "components", "body", []*ApplicationSnapshotComponentsItems0(m.Components)); err != nil {
		return err
	}

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {
			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationSnapshot) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "create_time", "body", string(m.CreateTime)); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSnapshot) contextValidateIsPartial(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_partial", "body", m.IsPartial); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSnapshot) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSnapshot) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSnapshot) UnmarshalBinary(b []byte) error {
	var res ApplicationSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSnapshotApplication application snapshot application
//
// swagger:model ApplicationSnapshotApplication
type ApplicationSnapshotApplication struct {

	// links
	Links *ApplicationSnapshotApplicationLinks `json:"_links,omitempty"`

	// Application name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The application UUID. Valid in URL.
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this application snapshot application
func (m *ApplicationSnapshotApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotApplication) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application snapshot application based on the context it is used
func (m *ApplicationSnapshotApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotApplication) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSnapshotApplication) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "application"+"."+"name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSnapshotApplication) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "application"+"."+"uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSnapshotApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSnapshotApplication) UnmarshalBinary(b []byte) error {
	var res ApplicationSnapshotApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSnapshotApplicationLinks application snapshot application links
//
// swagger:model ApplicationSnapshotApplicationLinks
type ApplicationSnapshotApplicationLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this application snapshot application links
func (m *ApplicationSnapshotApplicationLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotApplicationLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application snapshot application links based on the context it is used
func (m *ApplicationSnapshotApplicationLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotApplicationLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSnapshotApplicationLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSnapshotApplicationLinks) UnmarshalBinary(b []byte) error {
	var res ApplicationSnapshotApplicationLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSnapshotComponentsItems0 application snapshot components items0
//
// swagger:model ApplicationSnapshotComponentsItems0
type ApplicationSnapshotComponentsItems0 struct {

	// links
	Links *ApplicationSnapshotComponentsItems0Links `json:"_links,omitempty"`

	// Component name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Component UUID
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this application snapshot components items0
func (m *ApplicationSnapshotComponentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotComponentsItems0) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this application snapshot components items0 based on the context it is used
func (m *ApplicationSnapshotComponentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotComponentsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationSnapshotComponentsItems0) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSnapshotComponentsItems0) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSnapshotComponentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSnapshotComponentsItems0) UnmarshalBinary(b []byte) error {
	var res ApplicationSnapshotComponentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSnapshotComponentsItems0Links application snapshot components items0 links
//
// swagger:model ApplicationSnapshotComponentsItems0Links
type ApplicationSnapshotComponentsItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this application snapshot components items0 links
func (m *ApplicationSnapshotComponentsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotComponentsItems0Links) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this application snapshot components items0 links based on the context it is used
func (m *ApplicationSnapshotComponentsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotComponentsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ApplicationSnapshotComponentsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSnapshotComponentsItems0Links) UnmarshalBinary(b []byte) error {
	var res ApplicationSnapshotComponentsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSnapshotLinks application snapshot links
//
// swagger:model ApplicationSnapshotLinks
type ApplicationSnapshotLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this application snapshot links
func (m *ApplicationSnapshotLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this application snapshot links based on the context it is used
func (m *ApplicationSnapshotLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ApplicationSnapshotLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSnapshotLinks) UnmarshalBinary(b []byte) error {
	var res ApplicationSnapshotLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSnapshotSvm application snapshot svm
//
// swagger:model ApplicationSnapshotSvm
type ApplicationSnapshotSvm struct {

	// SVM name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// SVM UUID
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this application snapshot svm
func (m *ApplicationSnapshotSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this application snapshot svm based on the context it is used
func (m *ApplicationSnapshotSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSnapshotSvm) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "svm"+"."+"name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSnapshotSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "svm"+"."+"uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSnapshotSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSnapshotSvm) UnmarshalBinary(b []byte) error {
	var res ApplicationSnapshotSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
