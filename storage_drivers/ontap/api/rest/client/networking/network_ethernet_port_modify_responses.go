// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkEthernetPortModifyReader is a Reader for the NetworkEthernetPortModify structure.
type NetworkEthernetPortModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetPortModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetPortModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetPortModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetPortModifyOK creates a NetworkEthernetPortModifyOK with default headers values
func NewNetworkEthernetPortModifyOK() *NetworkEthernetPortModifyOK {
	return &NetworkEthernetPortModifyOK{}
}

/* NetworkEthernetPortModifyOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetPortModifyOK struct {
}

func (o *NetworkEthernetPortModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ethernet/ports/{uuid}][%d] networkEthernetPortModifyOK ", 200)
}

func (o *NetworkEthernetPortModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetPortModifyDefault creates a NetworkEthernetPortModifyDefault with default headers values
func NewNetworkEthernetPortModifyDefault(code int) *NetworkEthernetPortModifyDefault {
	return &NetworkEthernetPortModifyDefault{
		_statusCode: code,
	}
}

/* NetworkEthernetPortModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1376361 | Port is already a member of a LAG. |
| 1377562 | Port cannot be used because it is currently the home port or current port of an interface. |
| 1377563 | Port is already a member of a LAG. |
| 1967087 | The specified broadcast domain UUID is not valid. |
| 1967088 | The specified broadcast domain name does not exist in the specified IPspace. |
| 1967089 | The specified broadcast domain UUID, name and IPspace name do not match. |
| 1967094 | The specified LAG member port UUID is not valid. |
| 1967095 | The specified LAG member port name and node name combination is not valid. |
| 1967096 | The specified node does not match the specified LAG member port node. |
| 1967097 | The specified LAG member ports UUID, name, and node name do not match. |
| 1967148 | Failure to remove port from broadcast domain. |
| 1967149 | Failure to add port to broadcast domain. |

*/
type NetworkEthernetPortModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ethernet port modify default response
func (o *NetworkEthernetPortModifyDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetPortModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /network/ethernet/ports/{uuid}][%d] network_ethernet_port_modify default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkEthernetPortModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetPortModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
