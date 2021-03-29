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

// SvmPeer An SVM peer relation object.
//
// swagger:model svm_peer
type SvmPeer struct {

	// links
	Links *SvmPeerLinks `json:"_links,omitempty"`

	// A list of applications for an SVM peer relation.
	// Example: ["snapmirror","lun_copy"]
	Applications []SvmPeerApplications `json:"applications,omitempty"`

	// A peer SVM alias name to avoid a name conflict on the local cluster.
	Name string `json:"name,omitempty"`

	// peer
	Peer *SvmPeerPeer `json:"peer,omitempty"`

	// SVM peering state. To accept a pending SVM peer request, PATCH the state to "peered". To reject a pending SVM peer request, PATCH the state to "rejected". To suspend a peered SVM peer relation, PATCH the state to "suspended". To resume a suspended SVM peer relation, PATCH the state to "peered". The states "initiated", "pending", and "initializing" are system-generated and cannot be used for PATCH.
	// Example: peered
	// Enum: [peered rejected suspended initiated pending initializing]
	State string `json:"state,omitempty"`

	// svm
	Svm *SvmPeerSvm `json:"svm,omitempty"`

	// SVM peer relation UUID
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this svm peer
func (m *SvmPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *SvmPeer) validateLinks(formats strfmt.Registry) error {
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

func (m *SvmPeer) validateApplications(formats strfmt.Registry) error {
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

func (m *SvmPeer) validatePeer(formats strfmt.Registry) error {
	if swag.IsZero(m.Peer) { // not required
		return nil
	}

	if m.Peer != nil {
		if err := m.Peer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer")
			}
			return err
		}
	}

	return nil
}

var svmPeerTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["peered","rejected","suspended","initiated","pending","initializing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		svmPeerTypeStatePropEnum = append(svmPeerTypeStatePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// svm_peer
	// SvmPeer
	// state
	// State
	// peered
	// END RIPPY DEBUGGING
	// SvmPeerStatePeered captures enum value "peered"
	SvmPeerStatePeered string = "peered"

	// BEGIN RIPPY DEBUGGING
	// svm_peer
	// SvmPeer
	// state
	// State
	// rejected
	// END RIPPY DEBUGGING
	// SvmPeerStateRejected captures enum value "rejected"
	SvmPeerStateRejected string = "rejected"

	// BEGIN RIPPY DEBUGGING
	// svm_peer
	// SvmPeer
	// state
	// State
	// suspended
	// END RIPPY DEBUGGING
	// SvmPeerStateSuspended captures enum value "suspended"
	SvmPeerStateSuspended string = "suspended"

	// BEGIN RIPPY DEBUGGING
	// svm_peer
	// SvmPeer
	// state
	// State
	// initiated
	// END RIPPY DEBUGGING
	// SvmPeerStateInitiated captures enum value "initiated"
	SvmPeerStateInitiated string = "initiated"

	// BEGIN RIPPY DEBUGGING
	// svm_peer
	// SvmPeer
	// state
	// State
	// pending
	// END RIPPY DEBUGGING
	// SvmPeerStatePending captures enum value "pending"
	SvmPeerStatePending string = "pending"

	// BEGIN RIPPY DEBUGGING
	// svm_peer
	// SvmPeer
	// state
	// State
	// initializing
	// END RIPPY DEBUGGING
	// SvmPeerStateInitializing captures enum value "initializing"
	SvmPeerStateInitializing string = "initializing"
)

// prop value enum
func (m *SvmPeer) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, svmPeerTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SvmPeer) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *SvmPeer) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer based on the context it is used
func (m *SvmPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeer(ctx, formats); err != nil {
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

func (m *SvmPeer) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmPeer) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmPeer) contextValidatePeer(ctx context.Context, formats strfmt.Registry) error {

	if m.Peer != nil {
		if err := m.Peer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeer) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmPeer) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeer) UnmarshalBinary(b []byte) error {
	var res SvmPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerLinks svm peer links
//
// swagger:model SvmPeerLinks
type SvmPeerLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer links
func (m *SvmPeerLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer links based on the context it is used
func (m *SvmPeerLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPeer Details for a peer SVM object.
//
// swagger:model SvmPeerPeer
type SvmPeerPeer struct {

	// cluster
	Cluster *SvmPeerPeerCluster `json:"cluster,omitempty"`

	// svm
	Svm *SvmPeerPeerSvm `json:"svm,omitempty"`
}

// Validate validates this svm peer peer
func (m *SvmPeerPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
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

func (m *SvmPeerPeer) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerPeer) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer peer based on the context it is used
func (m *SvmPeerPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
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

func (m *SvmPeerPeer) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerPeer) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPeer) UnmarshalBinary(b []byte) error {
	var res SvmPeerPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPeerCluster svm peer peer cluster
//
// swagger:model SvmPeerPeerCluster
type SvmPeerPeerCluster struct {

	// links
	Links *SvmPeerPeerClusterLinks `json:"_links,omitempty"`

	// name
	// Example: cluster2
	Name string `json:"name,omitempty"`

	// uuid
	// Example: ebe27c49-1adf-4496-8335-ab862aebebf2
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this svm peer peer cluster
func (m *SvmPeerPeerCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer peer cluster based on the context it is used
func (m *SvmPeerPeerCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPeerCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPeerCluster) UnmarshalBinary(b []byte) error {
	var res SvmPeerPeerCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPeerClusterLinks svm peer peer cluster links
//
// swagger:model SvmPeerPeerClusterLinks
type SvmPeerPeerClusterLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer peer cluster links
func (m *SvmPeerPeerClusterLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerClusterLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer peer cluster links based on the context it is used
func (m *SvmPeerPeerClusterLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerClusterLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPeerClusterLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPeerClusterLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPeerClusterLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPeerSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model SvmPeerPeerSvm
type SvmPeerPeerSvm struct {

	// links
	Links *SvmPeerPeerSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this svm peer peer svm
func (m *SvmPeerPeerSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer peer svm based on the context it is used
func (m *SvmPeerPeerSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPeerSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPeerSvm) UnmarshalBinary(b []byte) error {
	var res SvmPeerPeerSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPeerSvmLinks svm peer peer svm links
//
// swagger:model SvmPeerPeerSvmLinks
type SvmPeerPeerSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer peer svm links
func (m *SvmPeerPeerSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer peer svm links based on the context it is used
func (m *SvmPeerPeerSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPeerSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPeerSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPeerSvmLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPeerSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerSvm Local SVM details
//
// swagger:model SvmPeerSvm
type SvmPeerSvm struct {

	// links
	Links *SvmPeerSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this svm peer svm
func (m *SvmPeerSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer svm based on the context it is used
func (m *SvmPeerSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerSvm) UnmarshalBinary(b []byte) error {
	var res SvmPeerSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerSvmLinks svm peer svm links
//
// swagger:model SvmPeerSvmLinks
type SvmPeerSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer svm links
func (m *SvmPeerSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerSvmLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer svm links based on the context it is used
func (m *SvmPeerSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerSvmLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
