// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Audit Auditing for NAS events is a security measure that enables you to track and log certain CIFS and NFS events on SVMs.
//
// swagger:model audit
type Audit struct {

	// Specifies whether or not auditing is enabled on the SVM.
	Enabled *bool `json:"enabled,omitempty"`

	// events
	Events *AuditEvents `json:"events,omitempty"`

	// log
	Log *Log `json:"log,omitempty"`

	// The audit log destination path where consolidated audit logs are stored.
	LogPath string `json:"log_path,omitempty"`

	// svm
	Svm *AuditSvm `json:"svm,omitempty"`
}

// Validate validates this audit
func (m *Audit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
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

func (m *Audit) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	if m.Events != nil {
		if err := m.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("events")
			}
			return err
		}
	}

	return nil
}

func (m *Audit) validateLog(formats strfmt.Registry) error {
	if swag.IsZero(m.Log) { // not required
		return nil
	}

	if m.Log != nil {
		if err := m.Log.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

func (m *Audit) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this audit based on the context it is used
func (m *Audit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Audit) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	if m.Events != nil {
		if err := m.Events.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("events")
			}
			return err
		}
	}

	return nil
}

func (m *Audit) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

	if m.Log != nil {
		if err := m.Log.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

func (m *Audit) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Audit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Audit) UnmarshalBinary(b []byte) error {
	var res Audit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuditEvents audit events
//
// swagger:model AuditEvents
type AuditEvents struct {

	// Authorization policy change events
	AuthorizationPolicy *bool `json:"authorization_policy,omitempty"`

	// Central access policy staging events
	CapStaging *bool `json:"cap_staging,omitempty"`

	// CIFS logon and logoff events
	CifsLogonLogoff *bool `json:"cifs_logon_logoff,omitempty"`

	// File operation events
	FileOperations *bool `json:"file_operations,omitempty"`

	// File share category events
	FileShare *bool `json:"file_share,omitempty"`

	// Local security group management events
	SecurityGroup *bool `json:"security_group,omitempty"`

	// Local user account management events
	UserAccount *bool `json:"user_account,omitempty"`
}

// Validate validates this audit events
func (m *AuditEvents) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this audit events based on context it is used
func (m *AuditEvents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditEvents) UnmarshalBinary(b []byte) error {
	var res AuditEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuditSvm audit svm
//
// swagger:model AuditSvm
type AuditSvm struct {

	// links
	Links *AuditSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this audit svm
func (m *AuditSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this audit svm based on the context it is used
func (m *AuditSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditSvm) UnmarshalBinary(b []byte) error {
	var res AuditSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuditSvmLinks audit svm links
//
// swagger:model AuditSvmLinks
type AuditSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this audit svm links
func (m *AuditSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this audit svm links based on the context it is used
func (m *AuditSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditSvmLinks) UnmarshalBinary(b []byte) error {
	var res AuditSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
