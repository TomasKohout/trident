// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityKeyManagerModifyReader is a Reader for the SecurityKeyManagerModify structure.
type SecurityKeyManagerModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerModifyOK creates a SecurityKeyManagerModifyOK with default headers values
func NewSecurityKeyManagerModifyOK() *SecurityKeyManagerModifyOK {
	return &SecurityKeyManagerModifyOK{}
}

/* SecurityKeyManagerModifyOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerModifyOK struct {
}

func (o *SecurityKeyManagerModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}][%d] securityKeyManagerModifyOK ", 200)
}

func (o *SecurityKeyManagerModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeyManagerModifyDefault creates a SecurityKeyManagerModifyDefault with default headers values
func NewSecurityKeyManagerModifyDefault(code int) *SecurityKeyManagerModifyDefault {
	return &SecurityKeyManagerModifyDefault{
		_statusCode: code,
	}
}

/* SecurityKeyManagerModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 65536139 | The existing passphrase value provided does not match the configured passphrase. |
| 65536150 | The new passphrase is same as old passphrase. |
| 65536404 | The passphrase does not match the accepted length. |
| 65536406 | The change of passphrase failed. |
| 65536407 | The passphrase update failed on some nodes. |
| 65536408 | The passphrase update failed on some nodes. |
| 65536802 | The passphrase does not match the accepted length in common criteria mode. |
| 65536821 | The certificate is not installed. |
| 65536822 | Multitenant key management is not supported in the current cluster version. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536850 | New client certificate public or private keys are different from the existing client certificate. |
| 65536852 | Failed to query supported KMIP protocol versions. |
| 65536917 | Updating an onboard passhrase requires both new and existing cluster passphrase. |
| 65537242 | The Onboard Key Manager existing_passphrase must be provided when performing a PATCH/synchronize operation. |
| 65537243 | The Onboard Key Manager passphrase must not be provided when performing a PATCH/synchronize operation. |
| 66060338 | Failed to establish secure connection for a key management server due to incorrect server_ca certificates. |
| 66060339 | Failed to establish secure connection for a key management server due to incorrect client certificates. |
| 66060340 | Failed to establish secure connection for a key management server due to Cryptsoft error. |
| 66060341 | Failed to establish secure connection for a key management server due to network configuration issues. |

*/
type SecurityKeyManagerModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security key manager modify default response
func (o *SecurityKeyManagerModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}][%d] security_key_manager_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SecurityKeyManagerModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
