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

// SecurityKeyManagerKeyServersDeleteReader is a Reader for the SecurityKeyManagerKeyServersDelete structure.
type SecurityKeyManagerKeyServersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerKeyServersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerKeyServersDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerKeyServersDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerKeyServersDeleteOK creates a SecurityKeyManagerKeyServersDeleteOK with default headers values
func NewSecurityKeyManagerKeyServersDeleteOK() *SecurityKeyManagerKeyServersDeleteOK {
	return &SecurityKeyManagerKeyServersDeleteOK{}
}

/* SecurityKeyManagerKeyServersDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerKeyServersDeleteOK struct {
}

func (o *SecurityKeyManagerKeyServersDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}/key-servers/{server}][%d] securityKeyManagerKeyServersDeleteOK ", 200)
}

func (o *SecurityKeyManagerKeyServersDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeyManagerKeyServersDeleteDefault creates a SecurityKeyManagerKeyServersDeleteDefault with default headers values
func NewSecurityKeyManagerKeyServersDeleteDefault(code int) *SecurityKeyManagerKeyServersDeleteDefault {
	return &SecurityKeyManagerKeyServersDeleteDefault{
		_statusCode: code,
	}
}

/* SecurityKeyManagerKeyServersDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 65536700 | The key server contains keys that are currently in use and not available from any other configured key server in the SVM. |
| 65536822 | Multitenant key management is not supported in the current cluster version. |
| 65536824 | Multitenant key management is not supported in MetroCluster configurations. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536843 | The key management server is not configured for the SVM. |

*/
type SecurityKeyManagerKeyServersDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security key manager key servers delete default response
func (o *SecurityKeyManagerKeyServersDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerKeyServersDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}/key-servers/{server}][%d] security_key_manager_key_servers_delete default  %+v", o._statusCode, o.Payload)
}
func (o *SecurityKeyManagerKeyServersDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
