// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ScheduleDeleteReader is a Reader for the ScheduleDelete structure.
type ScheduleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleDeleteOK creates a ScheduleDeleteOK with default headers values
func NewScheduleDeleteOK() *ScheduleDeleteOK {
	return &ScheduleDeleteOK{}
}

/* ScheduleDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ScheduleDeleteOK struct {
}

func (o *ScheduleDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/schedules/{uuid}][%d] scheduleDeleteOK ", 200)
}

func (o *ScheduleDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleDeleteDefault creates a ScheduleDeleteDefault with default headers values
func NewScheduleDeleteDefault(code int) *ScheduleDeleteDefault {
	return &ScheduleDeleteDefault{
		_statusCode: code,
	}
}

/* ScheduleDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 459758 | Cannot delete a job schedule that is in use. Remove all references to the schedule, and then try to delete again. |
| 459761 | Schedule cannot be deleted on this cluster because it is replicated from the remote cluster. |
| 459762 | The schedule cannot be deleted because it is a system-level schedule. |

*/
type ScheduleDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the schedule delete default response
func (o *ScheduleDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ScheduleDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /cluster/schedules/{uuid}][%d] schedule_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ScheduleDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ScheduleDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
