// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockComplianceClockGetReader is a Reader for the SnaplockComplianceClockGet structure.
type SnaplockComplianceClockGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockComplianceClockGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockComplianceClockGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockComplianceClockGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockComplianceClockGetOK creates a SnaplockComplianceClockGetOK with default headers values
func NewSnaplockComplianceClockGetOK() *SnaplockComplianceClockGetOK {
	return &SnaplockComplianceClockGetOK{}
}

/* SnaplockComplianceClockGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockComplianceClockGetOK struct {
	Payload *models.SnaplockComplianceClock
}

func (o *SnaplockComplianceClockGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks/{node.uuid}][%d] snaplockComplianceClockGetOK  %+v", 200, o.Payload)
}
func (o *SnaplockComplianceClockGetOK) GetPayload() *models.SnaplockComplianceClock {
	return o.Payload
}

func (o *SnaplockComplianceClockGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockComplianceClock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockComplianceClockGetDefault creates a SnaplockComplianceClockGetDefault with default headers values
func NewSnaplockComplianceClockGetDefault(code int) *SnaplockComplianceClockGetDefault {
	return &SnaplockComplianceClockGetDefault{
		_statusCode: code,
	}
}

/* SnaplockComplianceClockGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockComplianceClockGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock compliance clock get default response
func (o *SnaplockComplianceClockGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockComplianceClockGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks/{node.uuid}][%d] snaplock_compliance_clock_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockComplianceClockGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockComplianceClockGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
