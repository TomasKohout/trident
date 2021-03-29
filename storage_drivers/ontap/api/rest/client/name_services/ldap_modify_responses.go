// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LdapModifyReader is a Reader for the LdapModify structure.
type LdapModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LdapModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLdapModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLdapModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLdapModifyOK creates a LdapModifyOK with default headers values
func NewLdapModifyOK() *LdapModifyOK {
	return &LdapModifyOK{}
}

/* LdapModifyOK describes a response with status code 200, with default header values.

OK
*/
type LdapModifyOK struct {
}

func (o *LdapModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/ldap/{svm.uuid}][%d] ldapModifyOK ", 200)
}

func (o *LdapModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLdapModifyDefault creates a LdapModifyDefault with default headers values
func NewLdapModifyDefault(code int) *LdapModifyDefault {
	return &LdapModifyDefault{
		_statusCode: code,
	}
}

/* LdapModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 262186     | LDAP Servers cannot be used with Active Directory domain and/or preferred Active Directory servers |
| 2621488    | Invalid SVM context |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
| 4915203    | The specified LDAP schema does not exist |
| 4915208    | The specified LDAP servers or preferred Active Directory servers contain duplicate server entries |
| 4915229    | DNS resolution failed due to an internal error. Contact technical support if this issue persists |
| 4915231    | DNS resolution failed for one or more of the specified LDAP servers. Verify that a valid DNS server is configured |
| 23724132   | DNS resolution failed for all the specified LDAP servers. Verify that a valid DNS server is configured |
| 4915234    | The specified LDAP server or preferred Active Directory server is not supported because it is one of the following: multicast, loopback, 0.0.0.0, or broadcast |
| 4915248    | LDAP servers cannot be empty or "-". Specified Active Directory domain is invalid because it is empty or "-" or it contains either the special characters or "-" at the start or end of the domain. |
| 4915251    | STARTTLS and LDAPS cannot be used together |
| 4915257    | The LDAP configuration is invalid. Verify that the distinguished names and bind password are correct |
| 4915258    | The LDAP configuration is invalid. Verify that the Active Directory domain or servers are reachable and that the network configuration is correct |
| 4915259    | LDAP configurations with Active Directory domains are not supported on admin SVM. |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 LIFs |

*/
type LdapModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ldap modify default response
func (o *LdapModifyDefault) Code() int {
	return o._statusCode
}

func (o *LdapModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /name-services/ldap/{svm.uuid}][%d] ldap_modify default  %+v", o._statusCode, o.Payload)
}
func (o *LdapModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LdapModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
