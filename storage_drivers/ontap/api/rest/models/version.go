// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Version This returns the cluster version information.  When the cluster has more than one node, the cluster version is equivalent to the lowest of generation, major, and minor versions on all nodes.
//
// swagger:model version
type Version struct {

	// The full cluster version string.
	// Example: NetApp Release 9.4.0: Sun Nov 05 18:20:57 UTC 2017
	// Read Only: true
	Full string `json:"full,omitempty"`

	// The generation portion of the version.
	// Example: 9
	// Read Only: true
	Generation int64 `json:"generation,omitempty"`

	// The major portion of the version.
	// Example: 4
	// Read Only: true
	Major int64 `json:"major,omitempty"`

	// The minor portion of the version.
	// Example: 0
	// Read Only: true
	Minor int64 `json:"minor,omitempty"`
}

// Validate validates this version
func (m *Version) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this version based on the context it is used
func (m *Version) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFull(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeneration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMajor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Version) contextValidateFull(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "full", "body", string(m.Full)); err != nil {
		return err
	}

	return nil
}

func (m *Version) contextValidateGeneration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "generation", "body", int64(m.Generation)); err != nil {
		return err
	}

	return nil
}

func (m *Version) contextValidateMajor(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "major", "body", int64(m.Major)); err != nil {
		return err
	}

	return nil
}

func (m *Version) contextValidateMinor(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "minor", "body", int64(m.Minor)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Version) UnmarshalBinary(b []byte) error {
	var res Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
