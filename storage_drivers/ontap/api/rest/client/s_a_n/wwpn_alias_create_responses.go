// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// WwpnAliasCreateReader is a Reader for the WwpnAliasCreate structure.
type WwpnAliasCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WwpnAliasCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWwpnAliasCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWwpnAliasCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWwpnAliasCreateCreated creates a WwpnAliasCreateCreated with default headers values
func NewWwpnAliasCreateCreated() *WwpnAliasCreateCreated {
	return &WwpnAliasCreateCreated{}
}

/* WwpnAliasCreateCreated describes a response with status code 201, with default header values.

Created
*/
type WwpnAliasCreateCreated struct {
	Payload *models.WwpnAliasResponse
}

func (o *WwpnAliasCreateCreated) Error() string {
	return fmt.Sprintf("[POST /network/fc/wwpn-aliases][%d] wwpnAliasCreateCreated  %+v", 201, o.Payload)
}
func (o *WwpnAliasCreateCreated) GetPayload() *models.WwpnAliasResponse {
	return o.Payload
}

func (o *WwpnAliasCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WwpnAliasResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWwpnAliasCreateDefault creates a WwpnAliasCreateDefault with default headers values
func NewWwpnAliasCreateDefault(code int) *WwpnAliasCreateDefault {
	return &WwpnAliasCreateDefault{
		_statusCode: code,
	}
}

/* WwpnAliasCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1254317    | The alias already exists. |
| 1260882    | The supplied SVM does not exist. |
| 2621462    | The supplied SVM does not exist. |
| 2621706    | Both the SVM UUID and SVM name were supplied, but do not refer to the same SVM. |
| 2621707    | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5373982    | An invalid WWPN was supplied. The valid WWN format is XX:XX:XX:XX:XX:XX:XX:XX, where X is a hexadecimal digit. Example: "01:02:03:04:0a:0b:0c:0d". |

*/
type WwpnAliasCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the wwpn alias create default response
func (o *WwpnAliasCreateDefault) Code() int {
	return o._statusCode
}

func (o *WwpnAliasCreateDefault) Error() string {
	return fmt.Sprintf("[POST /network/fc/wwpn-aliases][%d] wwpn_alias_create default  %+v", o._statusCode, o.Payload)
}
func (o *WwpnAliasCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WwpnAliasCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
