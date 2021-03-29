// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// RoleGetReader is a Reader for the RoleGet structure.
type RoleGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoleGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRoleGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRoleGetOK creates a RoleGetOK with default headers values
func NewRoleGetOK() *RoleGetOK {
	return &RoleGetOK{}
}

/* RoleGetOK describes a response with status code 200, with default header values.

OK
*/
type RoleGetOK struct {
	Payload *models.RoleResponse
}

func (o *RoleGetOK) Error() string {
	return fmt.Sprintf("[GET /security/roles/{owner.uuid}/{name}][%d] roleGetOK  %+v", 200, o.Payload)
}
func (o *RoleGetOK) GetPayload() *models.RoleResponse {
	return o.Payload
}

func (o *RoleGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleGetDefault creates a RoleGetDefault with default headers values
func NewRoleGetDefault(code int) *RoleGetDefault {
	return &RoleGetDefault{
		_statusCode: code,
	}
}

/* RoleGetDefault describes a response with status code -1, with default header values.

Error
*/
type RoleGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the role get default response
func (o *RoleGetDefault) Code() int {
	return o._statusCode
}

func (o *RoleGetDefault) Error() string {
	return fmt.Sprintf("[GET /security/roles/{owner.uuid}/{name}][%d] role_get default  %+v", o._statusCode, o.Payload)
}
func (o *RoleGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RoleGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
