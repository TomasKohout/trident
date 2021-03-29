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

// SnapshotCreateReader is a Reader for the SnapshotCreate structure.
type SnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSnapshotCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotCreateAccepted creates a SnapshotCreateAccepted with default headers values
func NewSnapshotCreateAccepted() *SnapshotCreateAccepted {
	return &SnapshotCreateAccepted{}
}

/* SnapshotCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapshotCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SnapshotCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/snapshots][%d] snapshotCreateAccepted  %+v", 202, o.Payload)
}
func (o *SnapshotCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapshotCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotCreateDefault creates a SnapshotCreateDefault with default headers values
func NewSnapshotCreateDefault(code int) *SnapshotCreateDefault {
	return &SnapshotCreateDefault{
		_statusCode: code,
	}
}

/* SnapshotCreateDefault describes a response with status code -1, with default header values.

Error
*/
type SnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot create default response
func (o *SnapshotCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/snapshots][%d] snapshot_create default  %+v", o._statusCode, o.Payload)
}
func (o *SnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
