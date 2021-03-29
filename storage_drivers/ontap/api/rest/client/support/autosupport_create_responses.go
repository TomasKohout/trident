// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AutosupportCreateReader is a Reader for the AutosupportCreate structure.
type AutosupportCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutosupportCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAutosupportCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutosupportCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutosupportCreateCreated creates a AutosupportCreateCreated with default headers values
func NewAutosupportCreateCreated() *AutosupportCreateCreated {
	return &AutosupportCreateCreated{}
}

/* AutosupportCreateCreated describes a response with status code 201, with default header values.

Created
*/
type AutosupportCreateCreated struct {
	Payload *AutosupportCreateCreatedBody
}

func (o *AutosupportCreateCreated) Error() string {
	return fmt.Sprintf("[POST /support/autosupport/messages][%d] autosupportCreateCreated  %+v", 201, o.Payload)
}
func (o *AutosupportCreateCreated) GetPayload() *AutosupportCreateCreatedBody {
	return o.Payload
}

func (o *AutosupportCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AutosupportCreateCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutosupportCreateDefault creates a AutosupportCreateDefault with default headers values
func NewAutosupportCreateDefault(code int) *AutosupportCreateDefault {
	return &AutosupportCreateDefault{
		_statusCode: code,
	}
}

/* AutosupportCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 53149748   | The destination URI provided for the invoked AutoSupport is invalid |
| 655294464  | The message parameter is not supported with performance AutoSupports |

*/
type AutosupportCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the autosupport create default response
func (o *AutosupportCreateDefault) Code() int {
	return o._statusCode
}

func (o *AutosupportCreateDefault) Error() string {
	return fmt.Sprintf("[POST /support/autosupport/messages][%d] autosupport_create default  %+v", o._statusCode, o.Payload)
}
func (o *AutosupportCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutosupportCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AutosupportCreateCreatedBody List of messages invoked on the node
swagger:model AutosupportCreateCreatedBody
*/
type AutosupportCreateCreatedBody struct {

	// links
	Links *AutosupportCreateCreatedBodyLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 3
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*AutosupportCreateCreatedBodyRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this autosupport create created body
func (o *AutosupportCreateCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autosupportCreateCreated" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *AutosupportCreateCreatedBody) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.Records) { // not required
		return nil
	}

	for i := 0; i < len(o.Records); i++ {
		if swag.IsZero(o.Records[i]) { // not required
			continue
		}

		if o.Records[i] != nil {
			if err := o.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("autosupportCreateCreated" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this autosupport create created body based on the context it is used
func (o *AutosupportCreateCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autosupportCreateCreated" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *AutosupportCreateCreatedBody) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Records); i++ {

		if o.Records[i] != nil {
			if err := o.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("autosupportCreateCreated" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutosupportCreateCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutosupportCreateCreatedBody) UnmarshalBinary(b []byte) error {
	var res AutosupportCreateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AutosupportCreateCreatedBodyLinks autosupport create created body links
swagger:model AutosupportCreateCreatedBodyLinks
*/
type AutosupportCreateCreatedBodyLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this autosupport create created body links
func (o *AutosupportCreateCreatedBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autosupportCreateCreated" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this autosupport create created body links based on the context it is used
func (o *AutosupportCreateCreatedBodyLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autosupportCreateCreated" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyLinks) UnmarshalBinary(b []byte) error {
	var res AutosupportCreateCreatedBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AutosupportCreateCreatedBodyRecordsItems0 autosupport create created body records items0
swagger:model AutosupportCreateCreatedBodyRecordsItems0
*/
type AutosupportCreateCreatedBodyRecordsItems0 struct {

	// links
	Links *AutosupportCreateCreatedBodyRecordsItems0Links `json:"_links,omitempty"`

	// Sequence number of the generated AutoSupport
	// Example: 9
	Index int64 `json:"index,omitempty"`

	// node
	Node *AutosupportCreateCreatedBodyRecordsItems0Node `json:"node,omitempty"`
}

// Validate validates this autosupport create created body records items0
func (o *AutosupportCreateCreatedBodyRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(o.Node) { // not required
		return nil
	}

	if o.Node != nil {
		if err := o.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this autosupport create created body records items0 based on the context it is used
func (o *AutosupportCreateCreatedBodyRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if o.Node != nil {
		if err := o.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0) UnmarshalBinary(b []byte) error {
	var res AutosupportCreateCreatedBodyRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AutosupportCreateCreatedBodyRecordsItems0Links autosupport create created body records items0 links
swagger:model AutosupportCreateCreatedBodyRecordsItems0Links
*/
type AutosupportCreateCreatedBodyRecordsItems0Links struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this autosupport create created body records items0 links
func (o *AutosupportCreateCreatedBodyRecordsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this autosupport create created body records items0 links based on the context it is used
func (o *AutosupportCreateCreatedBodyRecordsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0Links) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0Links) UnmarshalBinary(b []byte) error {
	var res AutosupportCreateCreatedBodyRecordsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AutosupportCreateCreatedBodyRecordsItems0Node autosupport create created body records items0 node
swagger:model AutosupportCreateCreatedBodyRecordsItems0Node
*/
type AutosupportCreateCreatedBodyRecordsItems0Node struct {

	// links
	Links *AutosupportCreateCreatedBodyRecordsItems0NodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this autosupport create created body records items0 node
func (o *AutosupportCreateCreatedBodyRecordsItems0Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0Node) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this autosupport create created body records items0 node based on the context it is used
func (o *AutosupportCreateCreatedBodyRecordsItems0Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0Node) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0Node) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0Node) UnmarshalBinary(b []byte) error {
	var res AutosupportCreateCreatedBodyRecordsItems0Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AutosupportCreateCreatedBodyRecordsItems0NodeLinks autosupport create created body records items0 node links
swagger:model AutosupportCreateCreatedBodyRecordsItems0NodeLinks
*/
type AutosupportCreateCreatedBodyRecordsItems0NodeLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this autosupport create created body records items0 node links
func (o *AutosupportCreateCreatedBodyRecordsItems0NodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0NodeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this autosupport create created body records items0 node links based on the context it is used
func (o *AutosupportCreateCreatedBodyRecordsItems0NodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutosupportCreateCreatedBodyRecordsItems0NodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0NodeLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutosupportCreateCreatedBodyRecordsItems0NodeLinks) UnmarshalBinary(b []byte) error {
	var res AutosupportCreateCreatedBodyRecordsItems0NodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
