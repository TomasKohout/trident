// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FlexcacheOriginGetReader is a Reader for the FlexcacheOriginGet structure.
type FlexcacheOriginGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlexcacheOriginGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlexcacheOriginGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFlexcacheOriginGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFlexcacheOriginGetOK creates a FlexcacheOriginGetOK with default headers values
func NewFlexcacheOriginGetOK() *FlexcacheOriginGetOK {
	return &FlexcacheOriginGetOK{}
}

/* FlexcacheOriginGetOK describes a response with status code 200, with default header values.

OK
*/
type FlexcacheOriginGetOK struct {
	Payload *models.FlexcacheOrigin
}

func (o *FlexcacheOriginGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/flexcache/origins/{uuid}][%d] flexcacheOriginGetOK  %+v", 200, o.Payload)
}
func (o *FlexcacheOriginGetOK) GetPayload() *models.FlexcacheOrigin {
	return o.Payload
}

func (o *FlexcacheOriginGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlexcacheOrigin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlexcacheOriginGetDefault creates a FlexcacheOriginGetDefault with default headers values
func NewFlexcacheOriginGetDefault(code int) *FlexcacheOriginGetDefault {
	return &FlexcacheOriginGetDefault{
		_statusCode: code,
	}
}

/* FlexcacheOriginGetDefault describes a response with status code -1, with default header values.

Error
*/
type FlexcacheOriginGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the flexcache origin get default response
func (o *FlexcacheOriginGetDefault) Code() int {
	return o._statusCode
}

func (o *FlexcacheOriginGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/flexcache/origins/{uuid}][%d] flexcache_origin_get default  %+v", o._statusCode, o.Payload)
}
func (o *FlexcacheOriginGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FlexcacheOriginGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
