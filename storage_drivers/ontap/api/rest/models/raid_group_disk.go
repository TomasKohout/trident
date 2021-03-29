// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RaidGroupDisk raid group disk
//
// swagger:model raid_group_disk
type RaidGroupDisk struct {

	// disk
	Disk *RaidGroupDiskDisk `json:"disk,omitempty"`

	// The position of the disk within the RAID group.
	// Read Only: true
	// Enum: [data parity dparity tparity copy]
	Position string `json:"position,omitempty"`

	// The state of the disk within the RAID group.
	// Read Only: true
	// Enum: [normal failed zeroing copy replacing evacuating prefail offline reconstructing]
	State string `json:"state,omitempty"`

	// Disk interface type
	// Example: ssd
	// Read Only: true
	// Enum: [ata bsas fcal fsas lun sas msata ssd vmdisk unknown ssd_nvm]
	Type string `json:"type,omitempty"`

	// Size in bytes that is usable by the aggregate.
	// Example: 947912704
	// Read Only: true
	UsableSize int64 `json:"usable_size,omitempty"`
}

// Validate validates this raid group disk
func (m *RaidGroupDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *RaidGroupDisk) validateDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {
		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

var raidGroupDiskTypePositionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["data","parity","dparity","tparity","copy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		raidGroupDiskTypePositionPropEnum = append(raidGroupDiskTypePositionPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// position
	// Position
	// data
	// END RIPPY DEBUGGING
	// RaidGroupDiskPositionData captures enum value "data"
	RaidGroupDiskPositionData string = "data"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// position
	// Position
	// parity
	// END RIPPY DEBUGGING
	// RaidGroupDiskPositionParity captures enum value "parity"
	RaidGroupDiskPositionParity string = "parity"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// position
	// Position
	// dparity
	// END RIPPY DEBUGGING
	// RaidGroupDiskPositionDparity captures enum value "dparity"
	RaidGroupDiskPositionDparity string = "dparity"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// position
	// Position
	// tparity
	// END RIPPY DEBUGGING
	// RaidGroupDiskPositionTparity captures enum value "tparity"
	RaidGroupDiskPositionTparity string = "tparity"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// position
	// Position
	// copy
	// END RIPPY DEBUGGING
	// RaidGroupDiskPositionCopy captures enum value "copy"
	RaidGroupDiskPositionCopy string = "copy"
)

// prop value enum
func (m *RaidGroupDisk) validatePositionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, raidGroupDiskTypePositionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RaidGroupDisk) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionEnum("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

var raidGroupDiskTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","failed","zeroing","copy","replacing","evacuating","prefail","offline","reconstructing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		raidGroupDiskTypeStatePropEnum = append(raidGroupDiskTypeStatePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// normal
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateNormal captures enum value "normal"
	RaidGroupDiskStateNormal string = "normal"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// failed
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateFailed captures enum value "failed"
	RaidGroupDiskStateFailed string = "failed"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// zeroing
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateZeroing captures enum value "zeroing"
	RaidGroupDiskStateZeroing string = "zeroing"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// copy
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateCopy captures enum value "copy"
	RaidGroupDiskStateCopy string = "copy"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// replacing
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateReplacing captures enum value "replacing"
	RaidGroupDiskStateReplacing string = "replacing"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// evacuating
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateEvacuating captures enum value "evacuating"
	RaidGroupDiskStateEvacuating string = "evacuating"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// prefail
	// END RIPPY DEBUGGING
	// RaidGroupDiskStatePrefail captures enum value "prefail"
	RaidGroupDiskStatePrefail string = "prefail"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// offline
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateOffline captures enum value "offline"
	RaidGroupDiskStateOffline string = "offline"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// state
	// State
	// reconstructing
	// END RIPPY DEBUGGING
	// RaidGroupDiskStateReconstructing captures enum value "reconstructing"
	RaidGroupDiskStateReconstructing string = "reconstructing"
)

// prop value enum
func (m *RaidGroupDisk) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, raidGroupDiskTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RaidGroupDisk) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var raidGroupDiskTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ata","bsas","fcal","fsas","lun","sas","msata","ssd","vmdisk","unknown","ssd_nvm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		raidGroupDiskTypeTypePropEnum = append(raidGroupDiskTypeTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// ata
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeAta captures enum value "ata"
	RaidGroupDiskTypeAta string = "ata"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// bsas
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeBsas captures enum value "bsas"
	RaidGroupDiskTypeBsas string = "bsas"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// fcal
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeFcal captures enum value "fcal"
	RaidGroupDiskTypeFcal string = "fcal"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// fsas
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeFsas captures enum value "fsas"
	RaidGroupDiskTypeFsas string = "fsas"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// lun
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeLun captures enum value "lun"
	RaidGroupDiskTypeLun string = "lun"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// sas
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeSas captures enum value "sas"
	RaidGroupDiskTypeSas string = "sas"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// msata
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeMsata captures enum value "msata"
	RaidGroupDiskTypeMsata string = "msata"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// ssd
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeSsd captures enum value "ssd"
	RaidGroupDiskTypeSsd string = "ssd"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// vmdisk
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeVmdisk captures enum value "vmdisk"
	RaidGroupDiskTypeVmdisk string = "vmdisk"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// unknown
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeUnknown captures enum value "unknown"
	RaidGroupDiskTypeUnknown string = "unknown"

	// BEGIN RIPPY DEBUGGING
	// raid_group_disk
	// RaidGroupDisk
	// type
	// Type
	// ssd_nvm
	// END RIPPY DEBUGGING
	// RaidGroupDiskTypeSsdNvm captures enum value "ssd_nvm"
	RaidGroupDiskTypeSsdNvm string = "ssd_nvm"
)

// prop value enum
func (m *RaidGroupDisk) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, raidGroupDiskTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RaidGroupDisk) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this raid group disk based on the context it is used
func (m *RaidGroupDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsableSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RaidGroupDisk) contextValidateDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Disk != nil {
		if err := m.Disk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *RaidGroupDisk) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "position", "body", string(m.Position)); err != nil {
		return err
	}

	return nil
}

func (m *RaidGroupDisk) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *RaidGroupDisk) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *RaidGroupDisk) contextValidateUsableSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "usable_size", "body", int64(m.UsableSize)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RaidGroupDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RaidGroupDisk) UnmarshalBinary(b []byte) error {
	var res RaidGroupDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RaidGroupDiskDisk Disk
//
// swagger:model RaidGroupDiskDisk
type RaidGroupDiskDisk struct {

	// links
	Links *RaidGroupDiskDiskLinks `json:"_links,omitempty"`

	// name
	// Example: 1.0.1
	Name string `json:"name,omitempty"`
}

// Validate validates this raid group disk disk
func (m *RaidGroupDiskDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RaidGroupDiskDisk) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this raid group disk disk based on the context it is used
func (m *RaidGroupDiskDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RaidGroupDiskDisk) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RaidGroupDiskDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RaidGroupDiskDisk) UnmarshalBinary(b []byte) error {
	var res RaidGroupDiskDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RaidGroupDiskDiskLinks raid group disk disk links
//
// swagger:model RaidGroupDiskDiskLinks
type RaidGroupDiskDiskLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this raid group disk disk links
func (m *RaidGroupDiskDiskLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RaidGroupDiskDiskLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this raid group disk disk links based on the context it is used
func (m *RaidGroupDiskDiskLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RaidGroupDiskDiskLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RaidGroupDiskDiskLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RaidGroupDiskDiskLinks) UnmarshalBinary(b []byte) error {
	var res RaidGroupDiskDiskLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
