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

// CifsShare CIFS share is a named access point in a volume. Before users and applications can access data on the CIFS server over SMB,
// a CIFS share must be created with sufficient share permission. CIFS shares are tied to the CIFS server on the SVM.
// When a CIFS share is created, ONTAP creates a default ACL for the share with Full Control permissions for Everyone.
//
//
// swagger:model cifs_share
type CifsShare struct {

	// links
	Links *CifsShareLinks `json:"_links,omitempty"`

	// If enabled, all folders inside this share are visible to a user based on that individual user access right; prevents
	// the display of folders or other shared resources that the user does not have access to.
	//
	AccessBasedEnumeration *bool `json:"access_based_enumeration,omitempty"`

	// acls
	Acls []*CifsShareACL `json:"acls,omitempty"`

	// Specifies whether CIFS clients can request for change notifications for directories on this share.
	ChangeNotify *bool `json:"change_notify,omitempty"`

	// Specify the CIFS share descriptions.
	// Example: HR Department Share
	// Max Length: 256
	// Min Length: 1
	Comment string `json:"comment,omitempty"`

	// Specifies that SMB encryption must be used when accessing this share. Clients that do not support encryption are not
	// able to access this share.
	//
	Encryption *bool `json:"encryption,omitempty"`

	// Specifies whether or not the share is a home directory share, where the share and path names are dynamic.
	// ONTAP home directory functionality automatically offer each user a dynamic share to their home directory without creating an
	// individual SMB share for each user.
	// The ONTAP CIFS home directory feature enable us to configure a share that maps to
	// different directories based on the user that connects to it. Instead of creating a separate shares for each user,
	// a single share with a home directory parameters can be created.
	// In a home directory share, ONTAP dynamically generates the share-name and share-path by substituting
	// %w, %u, and %d variables with the corresponding Windows user name, UNIX user name, and domain name, respectively.
	//
	HomeDirectory *bool `json:"home_directory,omitempty"`

	// Specifies the name of the CIFS share that you want to create. If this
	// is a home directory share then the share name includes the pattern as
	// %w (Windows user name), %u (UNIX user name) and %d (Windows domain name)
	// variables in any combination with this parameter to generate shares dynamically.
	//
	// Example: HR_SHARE
	// Max Length: 80
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// Specify whether opportunistic locks are enabled on this share. "Oplocks" allow clients to lock files and cache content locally,
	// which can increase performance for file operations.
	//
	Oplocks *bool `json:"oplocks,omitempty"`

	// The fully-qualified pathname in the owning SVM namespace that is shared through this share.
	// If this is a home directory share then the path should be dynamic by specifying the pattern
	// %w (Windows user name), %u (UNIX user name), or %d (domain name) variables in any combination.
	// ONTAP generates the path dynamically for the connected user and this path is appended to each
	// search path to find the full Home Directory path.
	//
	// Example: /volume_1/eng_vol/
	// Max Length: 256
	// Min Length: 1
	Path string `json:"path,omitempty"`

	// svm
	Svm *CifsShareSvm `json:"svm,omitempty"`

	// Controls the access of UNIX symbolic links to CIFS clients.
	// The supported values are:
	//     * local - Enables only local symbolic links which is within the same CIFS share.
	//     * widelink - Enables both local symlinks and widelinks.
	//     * disable - Disables local symlinks and widelinks.
	//
	// Enum: [local widelink disable]
	UnixSymlink *string `json:"unix_symlink,omitempty"`

	// volume
	Volume *CifsShareVolume `json:"volume,omitempty"`
}

// Validate validates this cifs share
func (m *CifsShare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnixSymlink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShare) validateLinks(formats strfmt.Registry) error {
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

func (m *CifsShare) validateAcls(formats strfmt.Registry) error {
	if swag.IsZero(m.Acls) { // not required
		return nil
	}

	for i := 0; i < len(m.Acls); i++ {
		if swag.IsZero(m.Acls[i]) { // not required
			continue
		}

		if m.Acls[i] != nil {
			if err := m.Acls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CifsShare) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", m.Comment, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", m.Comment, 256); err != nil {
		return err
	}

	return nil
}

func (m *CifsShare) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 80); err != nil {
		return err
	}

	return nil
}

func (m *CifsShare) validatePath(formats strfmt.Registry) error {
	if swag.IsZero(m.Path) { // not required
		return nil
	}

	if err := validate.MinLength("path", "body", m.Path, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("path", "body", m.Path, 256); err != nil {
		return err
	}

	return nil
}

func (m *CifsShare) validateSvm(formats strfmt.Registry) error {
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

var cifsShareTypeUnixSymlinkPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["local","widelink","disable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsShareTypeUnixSymlinkPropEnum = append(cifsShareTypeUnixSymlinkPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// cifs_share
	// CifsShare
	// unix_symlink
	// UnixSymlink
	// local
	// END RIPPY DEBUGGING
	// CifsShareUnixSymlinkLocal captures enum value "local"
	CifsShareUnixSymlinkLocal string = "local"

	// BEGIN RIPPY DEBUGGING
	// cifs_share
	// CifsShare
	// unix_symlink
	// UnixSymlink
	// widelink
	// END RIPPY DEBUGGING
	// CifsShareUnixSymlinkWidelink captures enum value "widelink"
	CifsShareUnixSymlinkWidelink string = "widelink"

	// BEGIN RIPPY DEBUGGING
	// cifs_share
	// CifsShare
	// unix_symlink
	// UnixSymlink
	// disable
	// END RIPPY DEBUGGING
	// CifsShareUnixSymlinkDisable captures enum value "disable"
	CifsShareUnixSymlinkDisable string = "disable"
)

// prop value enum
func (m *CifsShare) validateUnixSymlinkEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsShareTypeUnixSymlinkPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsShare) validateUnixSymlink(formats strfmt.Registry) error {
	if swag.IsZero(m.UnixSymlink) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnixSymlinkEnum("unix_symlink", "body", *m.UnixSymlink); err != nil {
		return err
	}

	return nil
}

func (m *CifsShare) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs share based on the context it is used
func (m *CifsShare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAcls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShare) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CifsShare) contextValidateAcls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Acls); i++ {

		if m.Acls[i] != nil {
			if err := m.Acls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CifsShare) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CifsShare) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsShare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShare) UnmarshalBinary(b []byte) error {
	var res CifsShare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareLinks cifs share links
//
// swagger:model CifsShareLinks
type CifsShareLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs share links
func (m *CifsShareLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs share links based on the context it is used
func (m *CifsShareLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsShareLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareLinks) UnmarshalBinary(b []byte) error {
	var res CifsShareLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareSvm cifs share svm
//
// swagger:model CifsShareSvm
type CifsShareSvm struct {

	// links
	Links *CifsShareSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this cifs share svm
func (m *CifsShareSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs share svm based on the context it is used
func (m *CifsShareSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsShareSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareSvm) UnmarshalBinary(b []byte) error {
	var res CifsShareSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareSvmLinks cifs share svm links
//
// swagger:model CifsShareSvmLinks
type CifsShareSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs share svm links
func (m *CifsShareSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareSvmLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs share svm links based on the context it is used
func (m *CifsShareSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsShareSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareSvmLinks) UnmarshalBinary(b []byte) error {
	var res CifsShareSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareVolume cifs share volume
//
// swagger:model CifsShareVolume
type CifsShareVolume struct {

	// links
	Links *CifsShareVolumeLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this cifs share volume
func (m *CifsShareVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs share volume based on the context it is used
func (m *CifsShareVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsShareVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareVolume) UnmarshalBinary(b []byte) error {
	var res CifsShareVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareVolumeLinks cifs share volume links
//
// swagger:model CifsShareVolumeLinks
type CifsShareVolumeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs share volume links
func (m *CifsShareVolumeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareVolumeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs share volume links based on the context it is used
func (m *CifsShareVolumeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareVolumeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsShareVolumeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareVolumeLinks) UnmarshalBinary(b []byte) error {
	var res CifsShareVolumeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
