// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IpspaceGetReader is a Reader for the IpspaceGet structure.
type IpspaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpspaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpspaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpspaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpspaceGetOK creates a IpspaceGetOK with default headers values
func NewIpspaceGetOK() *IpspaceGetOK {
	return &IpspaceGetOK{}
}

/* IpspaceGetOK describes a response with status code 200, with default header values.

OK
*/
type IpspaceGetOK struct {
	Payload *models.Ipspace
}

func (o *IpspaceGetOK) Error() string {
	return fmt.Sprintf("[GET /network/ipspaces/{uuid}][%d] ipspaceGetOK  %+v", 200, o.Payload)
}
func (o *IpspaceGetOK) GetPayload() *models.Ipspace {
	return o.Payload
}

func (o *IpspaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ipspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpspaceGetDefault creates a IpspaceGetDefault with default headers values
func NewIpspaceGetDefault(code int) *IpspaceGetDefault {
	return &IpspaceGetDefault{
		_statusCode: code,
	}
}

/* IpspaceGetDefault describes a response with status code -1, with default header values.

Error
*/
type IpspaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ipspace get default response
func (o *IpspaceGetDefault) Code() int {
	return o._statusCode
}

func (o *IpspaceGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/ipspaces/{uuid}][%d] ipspace_get default  %+v", o._statusCode, o.Payload)
}
func (o *IpspaceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IpspaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
