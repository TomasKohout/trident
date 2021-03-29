// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LicenseKeys License keys or NLF contents.
//
// swagger:model license_keys
type LicenseKeys struct {

	// keys
	Keys []string `json:"keys,omitempty"`
}

// Validate validates this license keys
func (m *LicenseKeys) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this license keys based on context it is used
func (m *LicenseKeys) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicenseKeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseKeys) UnmarshalBinary(b []byte) error {
	var res LicenseKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
