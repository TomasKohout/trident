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

// ZappNvmeComponentsSubsystem components.subsystem
//
// swagger:model zapp_nvme_components_subsystem
type ZappNvmeComponentsSubsystem struct {

	// hosts
	Hosts []*ZappNvmeComponentsSubsystemHostsItems0 `json:"hosts,omitempty"`

	// The name of the subsystem accessing the component. If neither the name nor the UUID is provided, the name defaults to &lt;application-name&gt;_&lt;component-name&gt;, whether that subsystem already exists or not.
	Name string `json:"name,omitempty"`

	// The name of the host OS accessing the component. The default value is the host OS that is running the application.
	// Enum: [linux vmware windows]
	OsType string `json:"os_type,omitempty"`

	// The UUID of an existing subsystem to be granted access to the component. Usage: &lt;UUID&gt;
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this zapp nvme components subsystem
func (m *ZappNvmeComponentsSubsystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappNvmeComponentsSubsystem) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	for i := 0; i < len(m.Hosts); i++ {
		if swag.IsZero(m.Hosts[i]) { // not required
			continue
		}

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var zappNvmeComponentsSubsystemTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["linux","vmware","windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zappNvmeComponentsSubsystemTypeOsTypePropEnum = append(zappNvmeComponentsSubsystemTypeOsTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// zapp_nvme_components_subsystem
	// ZappNvmeComponentsSubsystem
	// os_type
	// OsType
	// linux
	// END RIPPY DEBUGGING
	// ZappNvmeComponentsSubsystemOsTypeLinux captures enum value "linux"
	ZappNvmeComponentsSubsystemOsTypeLinux string = "linux"

	// BEGIN RIPPY DEBUGGING
	// zapp_nvme_components_subsystem
	// ZappNvmeComponentsSubsystem
	// os_type
	// OsType
	// vmware
	// END RIPPY DEBUGGING
	// ZappNvmeComponentsSubsystemOsTypeVmware captures enum value "vmware"
	ZappNvmeComponentsSubsystemOsTypeVmware string = "vmware"

	// BEGIN RIPPY DEBUGGING
	// zapp_nvme_components_subsystem
	// ZappNvmeComponentsSubsystem
	// os_type
	// OsType
	// windows
	// END RIPPY DEBUGGING
	// ZappNvmeComponentsSubsystemOsTypeWindows captures enum value "windows"
	ZappNvmeComponentsSubsystemOsTypeWindows string = "windows"
)

// prop value enum
func (m *ZappNvmeComponentsSubsystem) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, zappNvmeComponentsSubsystemTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ZappNvmeComponentsSubsystem) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this zapp nvme components subsystem based on the context it is used
func (m *ZappNvmeComponentsSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappNvmeComponentsSubsystem) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hosts); i++ {

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZappNvmeComponentsSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappNvmeComponentsSubsystem) UnmarshalBinary(b []byte) error {
	var res ZappNvmeComponentsSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ZappNvmeComponentsSubsystemHostsItems0 zapp nvme components subsystem hosts items0
//
// swagger:model ZappNvmeComponentsSubsystemHostsItems0
type ZappNvmeComponentsSubsystemHostsItems0 struct {

	// The host NQN.
	Nqn string `json:"nqn,omitempty"`
}

// Validate validates this zapp nvme components subsystem hosts items0
func (m *ZappNvmeComponentsSubsystemHostsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this zapp nvme components subsystem hosts items0 based on context it is used
func (m *ZappNvmeComponentsSubsystemHostsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZappNvmeComponentsSubsystemHostsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappNvmeComponentsSubsystemHostsItems0) UnmarshalBinary(b []byte) error {
	var res ZappNvmeComponentsSubsystemHostsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
