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

// EmsFilterResponse ems filter response
//
// swagger:model ems_filter_response
type EmsFilterResponse struct {

	// links
	Links *EmsFilterResponseLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 3
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*EmsFilterResponseRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this ems filter response
func (m *EmsFilterResponse) Validate(formats strfmt.Registry) error {
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

func (m *EmsFilterResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *EmsFilterResponse) validateRecords(formats strfmt.Registry) error {
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

// ContextValidate validate this ems filter response based on the context it is used
func (m *EmsFilterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *EmsFilterResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsFilterResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsFilterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponse) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsFilterResponseLinks ems filter response links
//
// swagger:model EmsFilterResponseLinks
type EmsFilterResponseLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems filter response links
func (m *EmsFilterResponseLinks) Validate(formats strfmt.Registry) error {
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

func (m *EmsFilterResponseLinks) validateNext(formats strfmt.Registry) error {
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

func (m *EmsFilterResponseLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems filter response links based on the context it is used
func (m *EmsFilterResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *EmsFilterResponseLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsFilterResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsFilterResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponseLinks) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsFilterResponseRecordsItems0 ems filter response records items0
//
// swagger:model EmsFilterResponseRecordsItems0
type EmsFilterResponseRecordsItems0 struct {

	// links
	Links *EmsFilterResponseRecordsItems0Links `json:"_links,omitempty"`

	// Filter name
	// Example: snmp-traphost
	Name string `json:"name,omitempty"`

	// Array of event filter rules on which to match.
	Rules []*EmsFilterResponseRecordsItems0RulesItems0 `json:"rules,omitempty"`
}

// Validate validates this ems filter response records items0
func (m *EmsFilterResponseRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0) validateLinks(formats strfmt.Registry) error {
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

func (m *EmsFilterResponseRecordsItems0) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems filter response records items0 based on the context it is used
func (m *EmsFilterResponseRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsFilterResponseRecordsItems0) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {
			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponseRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsFilterResponseRecordsItems0Links ems filter response records items0 links
//
// swagger:model EmsFilterResponseRecordsItems0Links
type EmsFilterResponseRecordsItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems filter response records items0 links
func (m *EmsFilterResponseRecordsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0Links) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems filter response records items0 links based on the context it is used
func (m *EmsFilterResponseRecordsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsFilterResponseRecordsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0Links) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponseRecordsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsFilterResponseRecordsItems0RulesItems0 Rule for an event filter
//
// swagger:model EmsFilterResponseRecordsItems0RulesItems0
type EmsFilterResponseRecordsItems0RulesItems0 struct {

	// links
	Links *EmsFilterResponseRecordsItems0RulesItems0Links `json:"_links,omitempty"`

	// Rule index. Rules are evaluated in ascending order. If a rule's index order is not specified during creation, the rule is appended to the end of the list.
	// Example: 1
	Index int64 `json:"index,omitempty"`

	// message criteria
	MessageCriteria *EmsFilterResponseRecordsItems0RulesItems0MessageCriteria `json:"message_criteria,omitempty"`

	// Rule type
	// Example: include
	// Enum: [include exclude]
	Type string `json:"type,omitempty"`
}

// Validate validates this ems filter response records items0 rules items0
func (m *EmsFilterResponseRecordsItems0RulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0) validateLinks(formats strfmt.Registry) error {
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

func (m *EmsFilterResponseRecordsItems0RulesItems0) validateMessageCriteria(formats strfmt.Registry) error {
	if swag.IsZero(m.MessageCriteria) { // not required
		return nil
	}

	if m.MessageCriteria != nil {
		if err := m.MessageCriteria.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message_criteria")
			}
			return err
		}
	}

	return nil
}

var emsFilterResponseRecordsItems0RulesItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["include","exclude"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsFilterResponseRecordsItems0RulesItems0TypeTypePropEnum = append(emsFilterResponseRecordsItems0RulesItems0TypeTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// EmsFilterResponseRecordsItems0RulesItems0
	// EmsFilterResponseRecordsItems0RulesItems0
	// type
	// Type
	// include
	// END RIPPY DEBUGGING
	// EmsFilterResponseRecordsItems0RulesItems0TypeInclude captures enum value "include"
	EmsFilterResponseRecordsItems0RulesItems0TypeInclude string = "include"

	// BEGIN RIPPY DEBUGGING
	// EmsFilterResponseRecordsItems0RulesItems0
	// EmsFilterResponseRecordsItems0RulesItems0
	// type
	// Type
	// exclude
	// END RIPPY DEBUGGING
	// EmsFilterResponseRecordsItems0RulesItems0TypeExclude captures enum value "exclude"
	EmsFilterResponseRecordsItems0RulesItems0TypeExclude string = "exclude"
)

// prop value enum
func (m *EmsFilterResponseRecordsItems0RulesItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsFilterResponseRecordsItems0RulesItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems filter response records items0 rules items0 based on the context it is used
func (m *EmsFilterResponseRecordsItems0RulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessageCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsFilterResponseRecordsItems0RulesItems0) contextValidateMessageCriteria(ctx context.Context, formats strfmt.Registry) error {

	if m.MessageCriteria != nil {
		if err := m.MessageCriteria.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message_criteria")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0RulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0RulesItems0) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponseRecordsItems0RulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsFilterResponseRecordsItems0RulesItems0Links ems filter response records items0 rules items0 links
//
// swagger:model EmsFilterResponseRecordsItems0RulesItems0Links
type EmsFilterResponseRecordsItems0RulesItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems filter response records items0 rules items0 links
func (m *EmsFilterResponseRecordsItems0RulesItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0Links) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems filter response records items0 rules items0 links based on the context it is used
func (m *EmsFilterResponseRecordsItems0RulesItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsFilterResponseRecordsItems0RulesItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0RulesItems0Links) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponseRecordsItems0RulesItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsFilterResponseRecordsItems0RulesItems0MessageCriteria Matching message definitions for the filter. A property must be specified.
//
// swagger:model EmsFilterResponseRecordsItems0RulesItems0MessageCriteria
type EmsFilterResponseRecordsItems0RulesItems0MessageCriteria struct {

	// links
	Links *EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks `json:"_links,omitempty"`

	// Message name filter on which to match. Supports wildcards. Defaults to * if not specified.
	// Example: callhome.*
	NamePattern string `json:"name_pattern,omitempty"`

	// A comma-separated list of severities or a wildcard.
	// Example: error,informational
	Severities *string `json:"severities,omitempty"`

	// A comma separated list of snmp_trap_types or a wildcard.
	// Example: standard|built_in
	SnmpTrapTypes *string `json:"snmp_trap_types,omitempty"`
}

// Validate validates this ems filter response records items0 rules items0 message criteria
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteria) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteria) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message_criteria" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems filter response records items0 rules items0 message criteria based on the context it is used
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteria) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteria) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message_criteria" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteria) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponseRecordsItems0RulesItems0MessageCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks ems filter response records items0 rules items0 message criteria links
//
// swagger:model EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks
type EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks struct {

	// related
	Related *Href `json:"related,omitempty"`
}

// Validate validates this ems filter response records items0 rules items0 message criteria links
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks) validateRelated(formats strfmt.Registry) error {
	if swag.IsZero(m.Related) { // not required
		return nil
	}

	if m.Related != nil {
		if err := m.Related.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message_criteria" + "." + "_links" + "." + "related")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems filter response records items0 rules items0 message criteria links based on the context it is used
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRelated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks) contextValidateRelated(ctx context.Context, formats strfmt.Registry) error {

	if m.Related != nil {
		if err := m.Related.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message_criteria" + "." + "_links" + "." + "related")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks) UnmarshalBinary(b []byte) error {
	var res EmsFilterResponseRecordsItems0RulesItems0MessageCriteriaLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
