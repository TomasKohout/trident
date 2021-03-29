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

// SnapmirrorRelationship SnapMirror relationship information
//
// swagger:model snapmirror_relationship
type SnapmirrorRelationship struct {

	// links
	Links *SnapmirrorRelationshipLinks `json:"_links,omitempty"`

	// create destination
	CreateDestination *SnapmirrorDestinationCreation `json:"create_destination,omitempty"`

	// This property is the destination endpoint of the relationship. The destination endpoint can be a FlexVol volume, FlexGroup volume, or SVM. For the POST request, the destination endpoint must be of type "DP" when the endpoint is a FlexVol volume or a FlexGroup volume. The POST request for SVM must have a destination endpoint of type "dp-destination". The destination endpoint path name must be specified in the "destination.path" property. The destination endpoint for FlexVol volume or FlexGroup volume will change to type "RW" when the relationship status is "broken_off" and will revert to type "DP" when the relationship status is "snapmirrored" using the PATCH request. The destination endpoint for SVM will change from "dp-destination" to type "default" when the relationship status is "broken_off" and will revert to type "dp-destination" when the relationship status is "snapmirrored" using the PATCH request.
	Destination *SnapmirrorEndpoint `json:"destination,omitempty"`

	// Snapshot copy exported to clients on destination.
	// Read Only: true
	ExportedSnapshot string `json:"exported_snapshot,omitempty"`

	// Is the relationship healthy?
	// Read Only: true
	Healthy *bool `json:"healthy,omitempty"`

	// Time since the exported Snapshot copy was created.
	// Example: PT8H35M42S
	// Read Only: true
	LagTime string `json:"lag_time,omitempty"`

	// policy
	Policy *SnapmirrorRelationshipPolicy `json:"policy,omitempty"`

	// Set to true on resync to preserve Snapshot copies on the destination that are newer than the latest common Snapshot copy. This property is applicable only for relationships with FlexGroup or FlexVol endpoints and when the PATCH state is being changed to "snapmirrored".
	Preserve *bool `json:"preserve,omitempty"`

	// Set to true to reduce resync time by not preserving storage efficiency. This property is applicable only for relationships with FlexVol endpoints and when the PATCH state is being changed to "snapmirrored".
	QuickResync *bool `json:"quick_resync,omitempty"`

	// Set to true to recover from a failed SnapMirror break operation on a FlexGroup relationship. This restores all destination FlexGroup constituents to the latest Snapshot copy, and any writes to the read-write constituents are lost. This property is applicable only for SnapMirror relationships with FlexGroup endpoints and when the PATCH state is being changed to "broken_off".
	RecoverAfterBreak *bool `json:"recover_after_break,omitempty"`

	// Set to true to create a relationship for restore. To trigger restore-transfer, use transfers POST on the restore relationship.
	Restore *bool `json:"restore,omitempty"`

	// Specifies the Snapshot copy to restore to on the destination during the break operation. This property is applicable only for SnapMirror relationships with FlexVol endpoints and when the PATCH state is being changed to "broken_off".
	RestoreToSnapshot string `json:"restore_to_snapshot,omitempty"`

	// This property is the source endpoint of the relationship. The source endpoint can be a FlexVol volume, FlexGroup volume, or SVM.
	Source *SnapmirrorEndpoint `json:"source,omitempty"`

	// State of the relationship. To initialize the relationship, PATCH the state to "snapmirrored" for relationships with a policy of type "async" or "in-sync" for relationships with a policy of type "sync". To break the relationship, PATCH the state to "broken_off". To resync the broken relationship, PATCH the state to "snapmirrored" for relationships with a policy of type "async" or "in_sync" for relationships with a policy of type "sync". To pause the relationship, suspending further transfers, PATCH the state to "paused". To resume transfers for a paused relationship, PATCH the state to "snapmirrored" or "in_sync". The entries "in_sync", "out_of_sync", and "synchronizing" are only applicable to relationships with a policy of type "sync". A PATCH call on the state change only triggers the transition to the specified state. You must poll on the "state", "healthy" and "unhealthy_reason" properties using a GET request to determine if the transition is successful. To automatically initialize the relationship when specifying “create_destination”, set the state to “snapmirrored” for relationships with a policy of type "async" or "in_sync" for relationships with a policy of type "sync".
	// Example: snapmirrored
	// Enum: [broken_off paused snapmirrored uninitialized in_sync out_of_sync synchronizing]
	State string `json:"state,omitempty"`

	// transfer
	Transfer *SnapmirrorRelationshipTransfer `json:"transfer,omitempty"`

	// Reason the relationship is not healthy. It is a concatenation of up to four levels of error messages.
	// Example: [{"code":"6621444","message":"Failed to complete update operation on one or more item relationships.","parameters":[]},{"code":"6621445","message":"Group Update failed","parameters":[]}]
	// Read Only: true
	UnhealthyReason []*SnapmirrorError `json:"unhealthy_reason,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this snapmirror relationship
func (m *SnapmirrorRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransfer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnhealthyReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationship) validateLinks(formats strfmt.Registry) error {
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

func (m *SnapmirrorRelationship) validateCreateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateDestination) { // not required
		return nil
	}

	if m.CreateDestination != nil {
		if err := m.CreateDestination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("create_destination")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

var snapmirrorRelationshipTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["broken_off","paused","snapmirrored","uninitialized","in_sync","out_of_sync","synchronizing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapmirrorRelationshipTypeStatePropEnum = append(snapmirrorRelationshipTypeStatePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// snapmirror_relationship
	// SnapmirrorRelationship
	// state
	// State
	// broken_off
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipStateBrokenOff captures enum value "broken_off"
	SnapmirrorRelationshipStateBrokenOff string = "broken_off"

	// BEGIN RIPPY DEBUGGING
	// snapmirror_relationship
	// SnapmirrorRelationship
	// state
	// State
	// paused
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipStatePaused captures enum value "paused"
	SnapmirrorRelationshipStatePaused string = "paused"

	// BEGIN RIPPY DEBUGGING
	// snapmirror_relationship
	// SnapmirrorRelationship
	// state
	// State
	// snapmirrored
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipStateSnapmirrored captures enum value "snapmirrored"
	SnapmirrorRelationshipStateSnapmirrored string = "snapmirrored"

	// BEGIN RIPPY DEBUGGING
	// snapmirror_relationship
	// SnapmirrorRelationship
	// state
	// State
	// uninitialized
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipStateUninitialized captures enum value "uninitialized"
	SnapmirrorRelationshipStateUninitialized string = "uninitialized"

	// BEGIN RIPPY DEBUGGING
	// snapmirror_relationship
	// SnapmirrorRelationship
	// state
	// State
	// in_sync
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipStateInSync captures enum value "in_sync"
	SnapmirrorRelationshipStateInSync string = "in_sync"

	// BEGIN RIPPY DEBUGGING
	// snapmirror_relationship
	// SnapmirrorRelationship
	// state
	// State
	// out_of_sync
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipStateOutOfSync captures enum value "out_of_sync"
	SnapmirrorRelationshipStateOutOfSync string = "out_of_sync"

	// BEGIN RIPPY DEBUGGING
	// snapmirror_relationship
	// SnapmirrorRelationship
	// state
	// State
	// synchronizing
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipStateSynchronizing captures enum value "synchronizing"
	SnapmirrorRelationshipStateSynchronizing string = "synchronizing"
)

// prop value enum
func (m *SnapmirrorRelationship) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapmirrorRelationshipTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapmirrorRelationship) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *SnapmirrorRelationship) validateTransfer(formats strfmt.Registry) error {
	if swag.IsZero(m.Transfer) { // not required
		return nil
	}

	if m.Transfer != nil {
		if err := m.Transfer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) validateUnhealthyReason(formats strfmt.Registry) error {
	if swag.IsZero(m.UnhealthyReason) { // not required
		return nil
	}

	for i := 0; i < len(m.UnhealthyReason); i++ {
		if swag.IsZero(m.UnhealthyReason[i]) { // not required
			continue
		}

		if m.UnhealthyReason[i] != nil {
			if err := m.UnhealthyReason[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unhealthy_reason" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapmirrorRelationship) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapmirror relationship based on the context it is used
func (m *SnapmirrorRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportedSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLagTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransfer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnhealthyReason(ctx, formats); err != nil {
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

func (m *SnapmirrorRelationship) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapmirrorRelationship) contextValidateCreateDestination(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateDestination != nil {
		if err := m.CreateDestination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("create_destination")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if m.Destination != nil {
		if err := m.Destination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateExportedSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exported_snapshot", "body", string(m.ExportedSnapshot)); err != nil {
		return err
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateHealthy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateLagTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lag_time", "body", string(m.LagTime)); err != nil {
		return err
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {
		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateTransfer(ctx context.Context, formats strfmt.Registry) error {

	if m.Transfer != nil {
		if err := m.Transfer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateUnhealthyReason(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "unhealthy_reason", "body", []*SnapmirrorError(m.UnhealthyReason)); err != nil {
		return err
	}

	for i := 0; i < len(m.UnhealthyReason); i++ {

		if m.UnhealthyReason[i] != nil {
			if err := m.UnhealthyReason[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unhealthy_reason" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapmirrorRelationship) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", strfmt.UUID(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorRelationship) UnmarshalBinary(b []byte) error {
	var res SnapmirrorRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorRelationshipLinks snapmirror relationship links
//
// swagger:model SnapmirrorRelationshipLinks
type SnapmirrorRelationshipLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror relationship links
func (m *SnapmirrorRelationshipLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snapmirror relationship links based on the context it is used
func (m *SnapmirrorRelationshipLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapmirrorRelationshipLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorRelationshipLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorRelationshipLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorRelationshipPolicy Basic policy information of the relationship.
//
// swagger:model SnapmirrorRelationshipPolicy
type SnapmirrorRelationshipPolicy struct {

	// links
	Links *SnapmirrorRelationshipPolicyLinks `json:"_links,omitempty"`

	// name
	// Example: Asynchronous
	Name string `json:"name,omitempty"`

	// type
	// Read Only: true
	// Enum: [async sync]
	Type string `json:"type,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this snapmirror relationship policy
func (m *SnapmirrorRelationshipPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipPolicy) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var snapmirrorRelationshipPolicyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["async","sync"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapmirrorRelationshipPolicyTypeTypePropEnum = append(snapmirrorRelationshipPolicyTypeTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipPolicy
	// SnapmirrorRelationshipPolicy
	// type
	// Type
	// async
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipPolicyTypeAsync captures enum value "async"
	SnapmirrorRelationshipPolicyTypeAsync string = "async"

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipPolicy
	// SnapmirrorRelationshipPolicy
	// type
	// Type
	// sync
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipPolicyTypeSync captures enum value "sync"
	SnapmirrorRelationshipPolicyTypeSync string = "sync"
)

// prop value enum
func (m *SnapmirrorRelationshipPolicy) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapmirrorRelationshipPolicyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapmirrorRelationshipPolicy) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("policy"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SnapmirrorRelationshipPolicy) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("policy"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapmirror relationship policy based on the context it is used
func (m *SnapmirrorRelationshipPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipPolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorRelationshipPolicy) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy"+"."+"type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorRelationshipPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorRelationshipPolicy) UnmarshalBinary(b []byte) error {
	var res SnapmirrorRelationshipPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorRelationshipPolicyLinks snapmirror relationship policy links
//
// swagger:model SnapmirrorRelationshipPolicyLinks
type SnapmirrorRelationshipPolicyLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror relationship policy links
func (m *SnapmirrorRelationshipPolicyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipPolicyLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapmirror relationship policy links based on the context it is used
func (m *SnapmirrorRelationshipPolicyLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipPolicyLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorRelationshipPolicyLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorRelationshipPolicyLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorRelationshipPolicyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorRelationshipTransfer Basic information on the current transfer.
//
// swagger:model SnapmirrorRelationshipTransfer
type SnapmirrorRelationshipTransfer struct {

	// links
	Links *SnapmirrorRelationshipTransferLinks `json:"_links,omitempty"`

	// Bytes transferred.
	BytesTransferred int64 `json:"bytes_transferred,omitempty"`

	// state
	// Enum: [aborted failed hard_aborted queued success transferring]
	State string `json:"state,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this snapmirror relationship transfer
func (m *SnapmirrorRelationshipTransfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipTransfer) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var snapmirrorRelationshipTransferTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aborted","failed","hard_aborted","queued","success","transferring"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapmirrorRelationshipTransferTypeStatePropEnum = append(snapmirrorRelationshipTransferTypeStatePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipTransfer
	// SnapmirrorRelationshipTransfer
	// state
	// State
	// aborted
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipTransferStateAborted captures enum value "aborted"
	SnapmirrorRelationshipTransferStateAborted string = "aborted"

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipTransfer
	// SnapmirrorRelationshipTransfer
	// state
	// State
	// failed
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipTransferStateFailed captures enum value "failed"
	SnapmirrorRelationshipTransferStateFailed string = "failed"

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipTransfer
	// SnapmirrorRelationshipTransfer
	// state
	// State
	// hard_aborted
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipTransferStateHardAborted captures enum value "hard_aborted"
	SnapmirrorRelationshipTransferStateHardAborted string = "hard_aborted"

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipTransfer
	// SnapmirrorRelationshipTransfer
	// state
	// State
	// queued
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipTransferStateQueued captures enum value "queued"
	SnapmirrorRelationshipTransferStateQueued string = "queued"

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipTransfer
	// SnapmirrorRelationshipTransfer
	// state
	// State
	// success
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipTransferStateSuccess captures enum value "success"
	SnapmirrorRelationshipTransferStateSuccess string = "success"

	// BEGIN RIPPY DEBUGGING
	// SnapmirrorRelationshipTransfer
	// SnapmirrorRelationshipTransfer
	// state
	// State
	// transferring
	// END RIPPY DEBUGGING
	// SnapmirrorRelationshipTransferStateTransferring captures enum value "transferring"
	SnapmirrorRelationshipTransferStateTransferring string = "transferring"
)

// prop value enum
func (m *SnapmirrorRelationshipTransfer) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapmirrorRelationshipTransferTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapmirrorRelationshipTransfer) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("transfer"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *SnapmirrorRelationshipTransfer) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("transfer"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapmirror relationship transfer based on the context it is used
func (m *SnapmirrorRelationshipTransfer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipTransfer) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorRelationshipTransfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorRelationshipTransfer) UnmarshalBinary(b []byte) error {
	var res SnapmirrorRelationshipTransfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorRelationshipTransferLinks snapmirror relationship transfer links
//
// swagger:model SnapmirrorRelationshipTransferLinks
type SnapmirrorRelationshipTransferLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror relationship transfer links
func (m *SnapmirrorRelationshipTransferLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipTransferLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapmirror relationship transfer links based on the context it is used
func (m *SnapmirrorRelationshipTransferLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorRelationshipTransferLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transfer" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorRelationshipTransferLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorRelationshipTransferLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorRelationshipTransferLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
