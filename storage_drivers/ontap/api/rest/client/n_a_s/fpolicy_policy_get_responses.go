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

// FpolicyPolicyGetReader is a Reader for the FpolicyPolicyGet structure.
type FpolicyPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyPolicyGetOK creates a FpolicyPolicyGetOK with default headers values
func NewFpolicyPolicyGetOK() *FpolicyPolicyGetOK {
	return &FpolicyPolicyGetOK{}
}

/* FpolicyPolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyPolicyGetOK struct {
	Payload *models.FpolicyPolicy
}

func (o *FpolicyPolicyGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/policies/{name}][%d] fpolicyPolicyGetOK  %+v", 200, o.Payload)
}
func (o *FpolicyPolicyGetOK) GetPayload() *models.FpolicyPolicy {
	return o.Payload
}

func (o *FpolicyPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyPolicyGetDefault creates a FpolicyPolicyGetDefault with default headers values
func NewFpolicyPolicyGetDefault(code int) *FpolicyPolicyGetDefault {
	return &FpolicyPolicyGetDefault{
		_statusCode: code,
	}
}

/* FpolicyPolicyGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyPolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy policy get default response
func (o *FpolicyPolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyPolicyGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/policies/{name}][%d] fpolicy_policy_get default  %+v", o._statusCode, o.Payload)
}
func (o *FpolicyPolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
