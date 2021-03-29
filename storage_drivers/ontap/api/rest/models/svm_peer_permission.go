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

// SvmPeerPermission Manage SVM peer permissions.
//
// swagger:model svm_peer_permission
type SvmPeerPermission struct {

	// links
	Links *SvmPeerPermissionLinks `json:"_links,omitempty"`

	// A list of applications for an SVM peer relation.
	// Example: ["snapmirror","flexcache"]
	Applications []SvmPeerPermissionApplications `json:"applications,omitempty"`

	// cluster peer
	ClusterPeer *SvmPeerPermissionClusterPeer `json:"cluster_peer,omitempty"`

	// svm
	Svm *SvmPeerPermissionSvm `json:"svm,omitempty"`
}

// Validate validates this svm peer permission
func (m *SvmPeerPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterPeer(formats); err != nil {
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

func (m *SvmPeerPermission) validateLinks(formats strfmt.Registry) error {
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

func (m *SvmPeerPermission) validateApplications(formats strfmt.Registry) error {
	if swag.IsZero(m.Applications) { // not required
		return nil
	}

	for i := 0; i < len(m.Applications); i++ {

		if err := m.Applications[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applications" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SvmPeerPermission) validateClusterPeer(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterPeer) { // not required
		return nil
	}

	if m.ClusterPeer != nil {
		if err := m.ClusterPeer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerPermission) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer permission based on the context it is used
func (m *SvmPeerPermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterPeer(ctx, formats); err != nil {
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

func (m *SvmPeerPermission) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmPeerPermission) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Applications); i++ {

		if err := m.Applications[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applications" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SvmPeerPermission) contextValidateClusterPeer(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterPeer != nil {
		if err := m.ClusterPeer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerPermission) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermission) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionClusterPeer Peer cluster details
//
// swagger:model SvmPeerPermissionClusterPeer
type SvmPeerPermissionClusterPeer struct {

	// links
	Links *SvmPeerPermissionClusterPeerLinks `json:"_links,omitempty"`

	// name
	// Example: cluster2
	Name string `json:"name,omitempty"`

	// uuid
	// Example: ebe27c49-1adf-4496-8335-ab862aebebf2
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this svm peer permission cluster peer
func (m *SvmPeerPermissionClusterPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionClusterPeer) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer permission cluster peer based on the context it is used
func (m *SvmPeerPermissionClusterPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionClusterPeer) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPermissionClusterPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionClusterPeer) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionClusterPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionClusterPeerLinks svm peer permission cluster peer links
//
// swagger:model SvmPeerPermissionClusterPeerLinks
type SvmPeerPermissionClusterPeerLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer permission cluster peer links
func (m *SvmPeerPermissionClusterPeerLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionClusterPeerLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer permission cluster peer links based on the context it is used
func (m *SvmPeerPermissionClusterPeerLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionClusterPeerLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPermissionClusterPeerLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionClusterPeerLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionClusterPeerLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionLinks svm peer permission links
//
// swagger:model SvmPeerPermissionLinks
type SvmPeerPermissionLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer permission links
func (m *SvmPeerPermissionLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer permission links based on the context it is used
func (m *SvmPeerPermissionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerPermissionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionSvm Local SVM permitted for peer relation. To create peer permissions for all SVMs, specify the SVM name as "*".
//
// swagger:model SvmPeerPermissionSvm
type SvmPeerPermissionSvm struct {

	// links
	Links *SvmPeerPermissionSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this svm peer permission svm
func (m *SvmPeerPermissionSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer permission svm based on the context it is used
func (m *SvmPeerPermissionSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerPermissionSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionSvm) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionSvmLinks svm peer permission svm links
//
// swagger:model SvmPeerPermissionSvmLinks
type SvmPeerPermissionSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer permission svm links
func (m *SvmPeerPermissionSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionSvmLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer permission svm links based on the context it is used
func (m *SvmPeerPermissionSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerPermissionSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionSvmLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
