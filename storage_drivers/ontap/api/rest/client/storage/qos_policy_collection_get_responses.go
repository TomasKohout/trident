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

// QosPolicyCollectionGetReader is a Reader for the QosPolicyCollectionGet structure.
type QosPolicyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQosPolicyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyCollectionGetOK creates a QosPolicyCollectionGetOK with default headers values
func NewQosPolicyCollectionGetOK() *QosPolicyCollectionGetOK {
	return &QosPolicyCollectionGetOK{}
}

/* QosPolicyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type QosPolicyCollectionGetOK struct {
	Payload *models.QosPolicyResponse
}

func (o *QosPolicyCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/qos/policies][%d] qosPolicyCollectionGetOK  %+v", 200, o.Payload)
}
func (o *QosPolicyCollectionGetOK) GetPayload() *models.QosPolicyResponse {
	return o.Payload
}

func (o *QosPolicyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QosPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyCollectionGetDefault creates a QosPolicyCollectionGetDefault with default headers values
func NewQosPolicyCollectionGetDefault(code int) *QosPolicyCollectionGetDefault {
	return &QosPolicyCollectionGetDefault{
		_statusCode: code,
	}
}

/* QosPolicyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type QosPolicyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the qos policy collection get default response
func (o *QosPolicyCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *QosPolicyCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/qos/policies][%d] qos_policy_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *QosPolicyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
