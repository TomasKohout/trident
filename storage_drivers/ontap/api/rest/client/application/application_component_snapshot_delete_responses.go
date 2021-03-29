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

// ApplicationComponentSnapshotDeleteReader is a Reader for the ApplicationComponentSnapshotDelete structure.
type ApplicationComponentSnapshotDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationComponentSnapshotDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewApplicationComponentSnapshotDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationComponentSnapshotDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationComponentSnapshotDeleteAccepted creates a ApplicationComponentSnapshotDeleteAccepted with default headers values
func NewApplicationComponentSnapshotDeleteAccepted() *ApplicationComponentSnapshotDeleteAccepted {
	return &ApplicationComponentSnapshotDeleteAccepted{}
}

/* ApplicationComponentSnapshotDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationComponentSnapshotDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *ApplicationComponentSnapshotDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] applicationComponentSnapshotDeleteAccepted  %+v", 202, o.Payload)
}
func (o *ApplicationComponentSnapshotDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationComponentSnapshotDeleteDefault creates a ApplicationComponentSnapshotDeleteDefault with default headers values
func NewApplicationComponentSnapshotDeleteDefault(code int) *ApplicationComponentSnapshotDeleteDefault {
	return &ApplicationComponentSnapshotDeleteDefault{
		_statusCode: code,
	}
}

/* ApplicationComponentSnapshotDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationComponentSnapshotDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application component snapshot delete default response
func (o *ApplicationComponentSnapshotDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationComponentSnapshotDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] application_component_snapshot_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationComponentSnapshotDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
