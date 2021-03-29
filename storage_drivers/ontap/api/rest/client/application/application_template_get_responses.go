// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ApplicationTemplateGetReader is a Reader for the ApplicationTemplateGet structure.
type ApplicationTemplateGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationTemplateGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationTemplateGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationTemplateGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationTemplateGetOK creates a ApplicationTemplateGetOK with default headers values
func NewApplicationTemplateGetOK() *ApplicationTemplateGetOK {
	return &ApplicationTemplateGetOK{}
}

/* ApplicationTemplateGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationTemplateGetOK struct {
	Payload *models.ApplicationTemplate
}

func (o *ApplicationTemplateGetOK) Error() string {
	return fmt.Sprintf("[GET /application/templates/{name}][%d] applicationTemplateGetOK  %+v", 200, o.Payload)
}
func (o *ApplicationTemplateGetOK) GetPayload() *models.ApplicationTemplate {
	return o.Payload
}

func (o *ApplicationTemplateGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationTemplateGetDefault creates a ApplicationTemplateGetDefault with default headers values
func NewApplicationTemplateGetDefault(code int) *ApplicationTemplateGetDefault {
	return &ApplicationTemplateGetDefault{
		_statusCode: code,
	}
}

/* ApplicationTemplateGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationTemplateGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application template get default response
func (o *ApplicationTemplateGetDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationTemplateGetDefault) Error() string {
	return fmt.Sprintf("[GET /application/templates/{name}][%d] application_template_get default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationTemplateGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationTemplateGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
