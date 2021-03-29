// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3ServiceCollectionGetReader is a Reader for the S3ServiceCollectionGet structure.
type S3ServiceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3ServiceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3ServiceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3ServiceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3ServiceCollectionGetOK creates a S3ServiceCollectionGetOK with default headers values
func NewS3ServiceCollectionGetOK() *S3ServiceCollectionGetOK {
	return &S3ServiceCollectionGetOK{}
}

/* S3ServiceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type S3ServiceCollectionGetOK struct {
	Payload *models.S3ServiceResponse
}

func (o *S3ServiceCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services][%d] s3ServiceCollectionGetOK  %+v", 200, o.Payload)
}
func (o *S3ServiceCollectionGetOK) GetPayload() *models.S3ServiceResponse {
	return o.Payload
}

func (o *S3ServiceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3ServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3ServiceCollectionGetDefault creates a S3ServiceCollectionGetDefault with default headers values
func NewS3ServiceCollectionGetDefault(code int) *S3ServiceCollectionGetDefault {
	return &S3ServiceCollectionGetDefault{
		_statusCode: code,
	}
}

/* S3ServiceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3ServiceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 service collection get default response
func (o *S3ServiceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *S3ServiceCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services][%d] s3_service_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *S3ServiceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3ServiceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
