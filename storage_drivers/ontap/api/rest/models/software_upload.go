// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SoftwareUpload software upload
//
// swagger:model software_upload
type SoftwareUpload struct {

	// Package file on a local file system
	// Example: base64 encoded package file content
	File string `json:"file,omitempty"`
}

// Validate validates this software upload
func (m *SoftwareUpload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this software upload based on context it is used
func (m *SoftwareUpload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareUpload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareUpload) UnmarshalBinary(b []byte) error {
	var res SoftwareUpload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
