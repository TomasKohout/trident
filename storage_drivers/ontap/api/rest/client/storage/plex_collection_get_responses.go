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

// PlexCollectionGetReader is a Reader for the PlexCollectionGet structure.
type PlexCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlexCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlexCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPlexCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPlexCollectionGetOK creates a PlexCollectionGetOK with default headers values
func NewPlexCollectionGetOK() *PlexCollectionGetOK {
	return &PlexCollectionGetOK{}
}

/* PlexCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PlexCollectionGetOK struct {
	Payload *models.PlexResponse
}

func (o *PlexCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes][%d] plexCollectionGetOK  %+v", 200, o.Payload)
}
func (o *PlexCollectionGetOK) GetPayload() *models.PlexResponse {
	return o.Payload
}

func (o *PlexCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlexResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlexCollectionGetDefault creates a PlexCollectionGetDefault with default headers values
func NewPlexCollectionGetDefault(code int) *PlexCollectionGetDefault {
	return &PlexCollectionGetDefault{
		_statusCode: code,
	}
}

/* PlexCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type PlexCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the plex collection get default response
func (o *PlexCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PlexCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes][%d] plex_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *PlexCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PlexCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
