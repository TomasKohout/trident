// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// EmsFilterRuleModifyReader is a Reader for the EmsFilterRuleModify structure.
type EmsFilterRuleModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsFilterRuleModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsFilterRuleModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsFilterRuleModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsFilterRuleModifyOK creates a EmsFilterRuleModifyOK with default headers values
func NewEmsFilterRuleModifyOK() *EmsFilterRuleModifyOK {
	return &EmsFilterRuleModifyOK{}
}

/* EmsFilterRuleModifyOK describes a response with status code 200, with default header values.

OK
*/
type EmsFilterRuleModifyOK struct {
}

func (o *EmsFilterRuleModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /support/ems/filters/{name}/rules/{index}][%d] emsFilterRuleModifyOK ", 200)
}

func (o *EmsFilterRuleModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmsFilterRuleModifyDefault creates a EmsFilterRuleModifyDefault with default headers values
func NewEmsFilterRuleModifyDefault(code int) *EmsFilterRuleModifyDefault {
	return &EmsFilterRuleModifyDefault{
		_statusCode: code,
	}
}

/* EmsFilterRuleModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 983092     | The index of the rule provided is outside the allowed range for the filter provided |
| 983095     | The rule index provided is invalid for the filter provided |
| 983113     | Default filters cannot be modified or removed |
| 983126     | A rule requires at least one name_pattern, severities, or snmp_trap_types to be defined |
| 983127     | A property cannot contain a combination of the wildcard characters and other values. |
| 983128     | An invalid value is provided for the property 'snmp_trap_types' |
| 983146     | An invalid value is provided for the property 'severities' |
| 983147     | The severity levels provided are not supported |
| 983155     | The provided severities property does not match that of the name_pattern |
| 983156     | The provided snmp_trap_types property does not match that of the name_pattern |
| 983157     | The provided severities and snmp_trap_types do not match those of the name_pattern |
| 983158     | The name_pattern provided does not exist |

*/
type EmsFilterRuleModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ems filter rule modify default response
func (o *EmsFilterRuleModifyDefault) Code() int {
	return o._statusCode
}

func (o *EmsFilterRuleModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /support/ems/filters/{name}/rules/{index}][%d] ems_filter_rule_modify default  %+v", o._statusCode, o.Payload)
}
func (o *EmsFilterRuleModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsFilterRuleModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
