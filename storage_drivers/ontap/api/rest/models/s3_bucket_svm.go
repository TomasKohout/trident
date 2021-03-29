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

// S3BucketSvm A bucket is a container of objects. Each bucket defines an object namespace. S3 requests specify objects using a bucket-name and object-name pair. An object resides within a bucket.
//
// swagger:model s3_bucket_svm
type S3BucketSvm struct {

	// A list of aggregates for FlexGroup volume constituents where the bucket is hosted. If this option is not specified, the bucket is auto-provisioned as a FlexGroup volume.
	Aggregates []*S3BucketSvmAggregatesItems0 `json:"aggregates,omitempty"`

	// Can contain any additional information about the bucket being created or modified.
	// Example: S3 bucket.
	// Max Length: 256
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// Specifies the number of constituents or FlexVol volumes per aggregate. A FlexGroup volume consisting of all such constituents across all specified aggregates is created. This option is used along with the aggregates option and cannot be used independently.
	// Example: 4
	// Maximum: 1000
	// Minimum: 1
	ConstituentsPerAggregate int64 `json:"constituents_per_aggregate,omitempty"`

	// encryption
	Encryption *S3BucketSvmEncryption `json:"encryption,omitempty"`

	// Specifies the bucket logical used size up to this point.
	// Read Only: true
	LogicalUsedSize int64 `json:"logical_used_size,omitempty"`

	// Specifies the name of the bucket. Bucket name is a string that can only contain the following combination of ASCII-range alphanumeric characters 0-9, a-z, ".", and "-".
	// Example: bucket-1
	// Max Length: 63
	// Min Length: 3
	Name string `json:"name,omitempty"`

	// Specifies the bucket size in bytes; ranges from 80MB to 64TB.
	// Example: 1677721600
	// Maximum: 7.0368744177664e+13
	// Minimum: 8.388608e+07
	Size int64 `json:"size,omitempty"`

	// svm
	Svm *S3BucketSvmSvm `json:"svm,omitempty"`

	// Specifies the unique identifier of the bucket.
	// Example: 414b29a1-3b26-11e9-bd58-0050568ea055
	// Read Only: true
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`

	// volume
	Volume *S3BucketSvmVolume `json:"volume,omitempty"`
}

// Validate validates this s3 bucket svm
func (m *S3BucketSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituentsPerAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
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

func (m *S3BucketSvm) validateAggregates(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregates) { // not required
		return nil
	}

	for i := 0; i < len(m.Aggregates); i++ {
		if swag.IsZero(m.Aggregates[i]) { // not required
			continue
		}

		if m.Aggregates[i] != nil {
			if err := m.Aggregates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3BucketSvm) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 256); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketSvm) validateConstituentsPerAggregate(formats strfmt.Registry) error {
	if swag.IsZero(m.ConstituentsPerAggregate) { // not required
		return nil
	}

	if err := validate.MinimumInt("constituents_per_aggregate", "body", m.ConstituentsPerAggregate, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("constituents_per_aggregate", "body", m.ConstituentsPerAggregate, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketSvm) validateEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	if m.Encryption != nil {
		if err := m.Encryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketSvm) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 63); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketSvm) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := validate.MinimumInt("size", "body", m.Size, 8.388608e+07, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("size", "body", m.Size, 7.0368744177664e+13, false); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketSvm) validateSvm(formats strfmt.Registry) error {
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

func (m *S3BucketSvm) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketSvm) validateVolume(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket svm based on the context it is used
func (m *S3BucketSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogicalUsedSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
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

func (m *S3BucketSvm) contextValidateAggregates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Aggregates); i++ {

		if m.Aggregates[i] != nil {
			if err := m.Aggregates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3BucketSvm) contextValidateEncryption(ctx context.Context, formats strfmt.Registry) error {

	if m.Encryption != nil {
		if err := m.Encryption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketSvm) contextValidateLogicalUsedSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "logical_used_size", "body", int64(m.LogicalUsedSize)); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketSvm) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *S3BucketSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", strfmt.UUID(m.UUID)); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketSvm) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvm) UnmarshalBinary(b []byte) error {
	var res S3BucketSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketSvmAggregatesItems0 Aggregate
//
// swagger:model S3BucketSvmAggregatesItems0
type S3BucketSvmAggregatesItems0 struct {

	// links
	Links *S3BucketSvmAggregatesItems0Links `json:"_links,omitempty"`

	// name
	// Example: aggr1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this s3 bucket svm aggregates items0
func (m *S3BucketSvmAggregatesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmAggregatesItems0) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket svm aggregates items0 based on the context it is used
func (m *S3BucketSvmAggregatesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmAggregatesItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *S3BucketSvmAggregatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvmAggregatesItems0) UnmarshalBinary(b []byte) error {
	var res S3BucketSvmAggregatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketSvmAggregatesItems0Links s3 bucket svm aggregates items0 links
//
// swagger:model S3BucketSvmAggregatesItems0Links
type S3BucketSvmAggregatesItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket svm aggregates items0 links
func (m *S3BucketSvmAggregatesItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmAggregatesItems0Links) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket svm aggregates items0 links based on the context it is used
func (m *S3BucketSvmAggregatesItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmAggregatesItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketSvmAggregatesItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvmAggregatesItems0Links) UnmarshalBinary(b []byte) error {
	var res S3BucketSvmAggregatesItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketSvmEncryption s3 bucket svm encryption
//
// swagger:model S3BucketSvmEncryption
type S3BucketSvmEncryption struct {

	// Specifies whether encryption is enabled on the bucket. By default, encryption is disabled on a bucket.
	Enabled *bool `json:"enabled,omitempty"`
}

// Validate validates this s3 bucket svm encryption
func (m *S3BucketSvmEncryption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this s3 bucket svm encryption based on the context it is used
func (m *S3BucketSvmEncryption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketSvmEncryption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvmEncryption) UnmarshalBinary(b []byte) error {
	var res S3BucketSvmEncryption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketSvmSvm s3 bucket svm svm
//
// swagger:model S3BucketSvmSvm
type S3BucketSvmSvm struct {

	// links
	Links *S3BucketSvmSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this s3 bucket svm svm
func (m *S3BucketSvmSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket svm svm based on the context it is used
func (m *S3BucketSvmSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketSvmSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvmSvm) UnmarshalBinary(b []byte) error {
	var res S3BucketSvmSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketSvmSvmLinks s3 bucket svm svm links
//
// swagger:model S3BucketSvmSvmLinks
type S3BucketSvmSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket svm svm links
func (m *S3BucketSvmSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmSvmLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket svm svm links based on the context it is used
func (m *S3BucketSvmSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketSvmSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvmSvmLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketSvmSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketSvmVolume Specifies the FlexGroup volume name and UUID where the bucket is hosted.
//
// swagger:model S3BucketSvmVolume
type S3BucketSvmVolume struct {

	// links
	Links *S3BucketSvmVolumeLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this s3 bucket svm volume
func (m *S3BucketSvmVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmVolume) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket svm volume based on the context it is used
func (m *S3BucketSvmVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketSvmVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvmVolume) UnmarshalBinary(b []byte) error {
	var res S3BucketSvmVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketSvmVolumeLinks s3 bucket svm volume links
//
// swagger:model S3BucketSvmVolumeLinks
type S3BucketSvmVolumeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket svm volume links
func (m *S3BucketSvmVolumeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmVolumeLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket svm volume links based on the context it is used
func (m *S3BucketSvmVolumeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketSvmVolumeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketSvmVolumeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketSvmVolumeLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketSvmVolumeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY
