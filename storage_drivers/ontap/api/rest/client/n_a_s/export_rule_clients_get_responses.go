// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ExportRuleClientsGetReader is a Reader for the ExportRuleClientsGet structure.
type ExportRuleClientsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleClientsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRuleClientsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleClientsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleClientsGetOK creates a ExportRuleClientsGetOK with default headers values
func NewExportRuleClientsGetOK() *ExportRuleClientsGetOK {
	return &ExportRuleClientsGetOK{}
}

/* ExportRuleClientsGetOK describes a response with status code 200, with default header values.

OK
*/
type ExportRuleClientsGetOK struct {
	Payload *models.ExportClientResponse
}

func (o *ExportRuleClientsGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] exportRuleClientsGetOK  %+v", 200, o.Payload)
}
func (o *ExportRuleClientsGetOK) GetPayload() *models.ExportClientResponse {
	return o.Payload
}

func (o *ExportRuleClientsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportClientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportRuleClientsGetDefault creates a ExportRuleClientsGetDefault with default headers values
func NewExportRuleClientsGetDefault(code int) *ExportRuleClientsGetDefault {
	return &ExportRuleClientsGetDefault{
		_statusCode: code,
	}
}

/* ExportRuleClientsGetDefault describes a response with status code -1, with default header values.

Error
*/
type ExportRuleClientsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export rule clients get default response
func (o *ExportRuleClientsGetDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleClientsGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] export_rule_clients_get default  %+v", o._statusCode, o.Payload)
}
func (o *ExportRuleClientsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleClientsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
