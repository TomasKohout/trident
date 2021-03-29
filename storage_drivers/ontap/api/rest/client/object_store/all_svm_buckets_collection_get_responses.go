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

// AllSvmBucketsCollectionGetReader is a Reader for the AllSvmBucketsCollectionGet structure.
type AllSvmBucketsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllSvmBucketsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllSvmBucketsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAllSvmBucketsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllSvmBucketsCollectionGetOK creates a AllSvmBucketsCollectionGetOK with default headers values
func NewAllSvmBucketsCollectionGetOK() *AllSvmBucketsCollectionGetOK {
	return &AllSvmBucketsCollectionGetOK{}
}

/* AllSvmBucketsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AllSvmBucketsCollectionGetOK struct {
	Payload *models.S3BucketResponse
}

func (o *AllSvmBucketsCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/buckets][%d] allSvmBucketsCollectionGetOK  %+v", 200, o.Payload)
}
func (o *AllSvmBucketsCollectionGetOK) GetPayload() *models.S3BucketResponse {
	return o.Payload
}

func (o *AllSvmBucketsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllSvmBucketsCollectionGetDefault creates a AllSvmBucketsCollectionGetDefault with default headers values
func NewAllSvmBucketsCollectionGetDefault(code int) *AllSvmBucketsCollectionGetDefault {
	return &AllSvmBucketsCollectionGetDefault{
		_statusCode: code,
	}
}

/* AllSvmBucketsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type AllSvmBucketsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the all svm buckets collection get default response
func (o *AllSvmBucketsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AllSvmBucketsCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/buckets][%d] all_svm_buckets_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *AllSvmBucketsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AllSvmBucketsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
