// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationBackingStorage application backing storage
//
// swagger:model application_backing_storage
type ApplicationBackingStorage struct {

	// luns
	Luns []*ApplicationLunObject `json:"luns,omitempty"`

	// namespaces
	Namespaces []*ApplicationNamespaceObject `json:"namespaces,omitempty"`

	// volumes
	Volumes []*ApplicationVolumeObject `json:"volumes,omitempty"`
}

// Validate validates this application backing storage
func (m *ApplicationBackingStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationBackingStorage) validateLuns(formats strfmt.Registry) error {
	if swag.IsZero(m.Luns) { // not required
		return nil
	}

	for i := 0; i < len(m.Luns); i++ {
		if swag.IsZero(m.Luns[i]) { // not required
			continue
		}

		if m.Luns[i] != nil {
			if err := m.Luns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("luns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationBackingStorage) validateNamespaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Namespaces); i++ {
		if swag.IsZero(m.Namespaces[i]) { // not required
			continue
		}

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationBackingStorage) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this application backing storage based on the context it is used
func (m *ApplicationBackingStorage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationBackingStorage) contextValidateLuns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Luns); i++ {

		if m.Luns[i] != nil {
			if err := m.Luns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("luns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationBackingStorage) contextValidateNamespaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Namespaces); i++ {

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationBackingStorage) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationBackingStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationBackingStorage) UnmarshalBinary(b []byte) error {
	var res ApplicationBackingStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
