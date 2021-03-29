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

// FlexcacheDeleteReader is a Reader for the FlexcacheDelete structure.
type FlexcacheDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlexcacheDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewFlexcacheDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFlexcacheDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFlexcacheDeleteAccepted creates a FlexcacheDeleteAccepted with default headers values
func NewFlexcacheDeleteAccepted() *FlexcacheDeleteAccepted {
	return &FlexcacheDeleteAccepted{}
}

/* FlexcacheDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FlexcacheDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *FlexcacheDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches/{uuid}][%d] flexcacheDeleteAccepted  %+v", 202, o.Payload)
}
func (o *FlexcacheDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FlexcacheDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlexcacheDeleteDefault creates a FlexcacheDeleteDefault with default headers values
func NewFlexcacheDeleteDefault(code int) *FlexcacheDeleteDefault {
	return &FlexcacheDeleteDefault{
		_statusCode: code,
	}
}

/* FlexcacheDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 66846879   | The specified volume UUID is not a FlexCache volume |
| 66846731   | Failed to delete the FlexCache volume |
| 524546     | Failed to delete the FlexCache volume because the FlexCache volume is not unmounted |

*/
type FlexcacheDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the flexcache delete default response
func (o *FlexcacheDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FlexcacheDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches/{uuid}][%d] flexcache_delete default  %+v", o._statusCode, o.Payload)
}
func (o *FlexcacheDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FlexcacheDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
