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

// ExportRule export rule
//
// swagger:model export_rule
type ExportRule struct {

	// links
	Links *ExportRuleLinks `json:"_links,omitempty"`

	// User ID To Which Anonymous Users Are Mapped.
	AnonymousUser *string `json:"anonymous_user,omitempty"`

	// Array of client matches
	Clients []*ExportClient `json:"clients,omitempty"`

	// Index of the rule within the export policy.
	//
	// Read Only: true
	Index int64 `json:"index,omitempty"`

	// protocols
	Protocols []*string `json:"protocols,omitempty"`

	// Authentication flavors that the read-only access rule governs
	//
	RoRule []ExportAuthenticationFlavor `json:"ro_rule,omitempty"`

	// Authentication flavors that the read/write access rule governs
	//
	RwRule []ExportAuthenticationFlavor `json:"rw_rule,omitempty"`

	// Authentication flavors that the superuser security type governs
	//
	Superuser []ExportAuthenticationFlavor `json:"superuser,omitempty"`
}

// Validate validates this export rule
func (m *ExportRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocols(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRwRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuperuser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRule) validateLinks(formats strfmt.Registry) error {
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

func (m *ExportRule) validateClients(formats strfmt.Registry) error {
	if swag.IsZero(m.Clients) { // not required
		return nil
	}

	for i := 0; i < len(m.Clients); i++ {
		if swag.IsZero(m.Clients[i]) { // not required
			continue
		}

		if m.Clients[i] != nil {
			if err := m.Clients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var exportRuleProtocolsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","nfs","nfs3","nfs4","cifs","flexcache"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportRuleProtocolsItemsEnum = append(exportRuleProtocolsItemsEnum, v)
	}
}

func (m *ExportRule) validateProtocolsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportRuleProtocolsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportRule) validateProtocols(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocols) { // not required
		return nil
	}

	for i := 0; i < len(m.Protocols); i++ {
		if swag.IsZero(m.Protocols[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateProtocolsItemsEnum("protocols"+"."+strconv.Itoa(i), "body", *m.Protocols[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExportRule) validateRoRule(formats strfmt.Registry) error {
	if swag.IsZero(m.RoRule) { // not required
		return nil
	}

	for i := 0; i < len(m.RoRule); i++ {

		if err := m.RoRule[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ro_rule" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ExportRule) validateRwRule(formats strfmt.Registry) error {
	if swag.IsZero(m.RwRule) { // not required
		return nil
	}

	for i := 0; i < len(m.RwRule); i++ {

		if err := m.RwRule[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rw_rule" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ExportRule) validateSuperuser(formats strfmt.Registry) error {
	if swag.IsZero(m.Superuser) { // not required
		return nil
	}

	for i := 0; i < len(m.Superuser); i++ {

		if err := m.Superuser[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("superuser" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this export rule based on the context it is used
func (m *ExportRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRwRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuperuser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ExportRule) contextValidateClients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clients); i++ {

		if m.Clients[i] != nil {
			if err := m.Clients[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", int64(m.Index)); err != nil {
		return err
	}

	return nil
}

func (m *ExportRule) contextValidateRoRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoRule); i++ {

		if err := m.RoRule[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ro_rule" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ExportRule) contextValidateRwRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RwRule); i++ {

		if err := m.RwRule[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rw_rule" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ExportRule) contextValidateSuperuser(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Superuser); i++ {

		if err := m.Superuser[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("superuser" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExportRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRule) UnmarshalBinary(b []byte) error {
	var res ExportRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExportRuleLinks export rule links
//
// swagger:model ExportRuleLinks
type ExportRuleLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this export rule links
func (m *ExportRuleLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this export rule links based on the context it is used
func (m *ExportRuleLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ExportRuleLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRuleLinks) UnmarshalBinary(b []byte) error {
	var res ExportRuleLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
