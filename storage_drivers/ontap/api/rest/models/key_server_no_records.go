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
	"github.com/go-openapi/validate"
)

// KeyServerNoRecords key server no records
//
// swagger:model key_server_no_records
type KeyServerNoRecords struct {

	// links
	Links *KeyServerNoRecordsLinks `json:"_links,omitempty"`

	// connectivity
	Connectivity *KeyServerNoRecordsConnectivity `json:"connectivity,omitempty"`

	// Password credentials for connecting with the key server. This is not audited.
	// Example: password
	// Format: password
	Password strfmt.Password `json:"password,omitempty"`

	// External key server for key management. If no port is provided, a default port of 5696 is used. Not valid in POST if `records` is provided.
	// Example: keyserver1.com:5698
	Server string `json:"server,omitempty"`

	// I/O timeout in seconds for communicating with the key server.
	// Example: 60
	// Maximum: 60
	// Minimum: 1
	Timeout int64 `json:"timeout,omitempty"`

	// KMIP username credentials for connecting with the key server.
	// Example: username
	Username string `json:"username,omitempty"`
}

// Validate validates this key server no records
func (m *KeyServerNoRecords) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyServerNoRecords) validateLinks(formats strfmt.Registry) error {
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

func (m *KeyServerNoRecords) validateConnectivity(formats strfmt.Registry) error {
	if swag.IsZero(m.Connectivity) { // not required
		return nil
	}

	if m.Connectivity != nil {
		if err := m.Connectivity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity")
			}
			return err
		}
	}

	return nil
}

func (m *KeyServerNoRecords) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KeyServerNoRecords) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout", "body", m.Timeout, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timeout", "body", m.Timeout, 60, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this key server no records based on the context it is used
func (m *KeyServerNoRecords) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyServerNoRecords) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *KeyServerNoRecords) contextValidateConnectivity(ctx context.Context, formats strfmt.Registry) error {

	if m.Connectivity != nil {
		if err := m.Connectivity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyServerNoRecords) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyServerNoRecords) UnmarshalBinary(b []byte) error {
	var res KeyServerNoRecords
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KeyServerNoRecordsConnectivity This property returns the key server connectivity state on all nodes of the cluster. The state is returned for a node only if the connectivity is not in an available state on that node.
// This is an advanced property; there is an added cost to retrieving its value. The property is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter or GET for all advanced properties is enabled.
//
//
// swagger:model KeyServerNoRecordsConnectivity
type KeyServerNoRecordsConnectivity struct {

	// Set to true when key server connectivity state is available on all nodes of the cluster.
	// Read Only: true
	ClusterAvailability *bool `json:"cluster_availability,omitempty"`

	// An array of key server connectivity states for each node.
	//
	// Read Only: true
	Records []*KeyServerState `json:"records,omitempty"`
}

// Validate validates this key server no records connectivity
func (m *KeyServerNoRecordsConnectivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyServerNoRecordsConnectivity) validateRecords(formats strfmt.Registry) error {
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
					return ve.ValidateName("connectivity" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this key server no records connectivity based on the context it is used
func (m *KeyServerNoRecordsConnectivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterAvailability(ctx, formats); err != nil {
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

func (m *KeyServerNoRecordsConnectivity) contextValidateClusterAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connectivity"+"."+"cluster_availability", "body", m.ClusterAvailability); err != nil {
		return err
	}

	return nil
}

func (m *KeyServerNoRecordsConnectivity) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connectivity"+"."+"records", "body", []*KeyServerState(m.Records)); err != nil {
		return err
	}

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectivity" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyServerNoRecordsConnectivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyServerNoRecordsConnectivity) UnmarshalBinary(b []byte) error {
	var res KeyServerNoRecordsConnectivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KeyServerNoRecordsLinks key server no records links
//
// swagger:model KeyServerNoRecordsLinks
type KeyServerNoRecordsLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this key server no records links
func (m *KeyServerNoRecordsLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyServerNoRecordsLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this key server no records links based on the context it is used
func (m *KeyServerNoRecordsLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyServerNoRecordsLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *KeyServerNoRecordsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyServerNoRecordsLinks) UnmarshalBinary(b []byte) error {
	var res KeyServerNoRecordsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
