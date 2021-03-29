// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NvmeInterfaceGetReader is a Reader for the NvmeInterfaceGet structure.
type NvmeInterfaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeInterfaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeInterfaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeInterfaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeInterfaceGetOK creates a NvmeInterfaceGetOK with default headers values
func NewNvmeInterfaceGetOK() *NvmeInterfaceGetOK {
	return &NvmeInterfaceGetOK{}
}

/* NvmeInterfaceGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeInterfaceGetOK struct {
	Payload *models.NvmeInterface
}

func (o *NvmeInterfaceGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nvme/interfaces/{uuid}][%d] nvmeInterfaceGetOK  %+v", 200, o.Payload)
}
func (o *NvmeInterfaceGetOK) GetPayload() *models.NvmeInterface {
	return o.Payload
}

func (o *NvmeInterfaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeInterfaceGetDefault creates a NvmeInterfaceGetDefault with default headers values
func NewNvmeInterfaceGetDefault(code int) *NvmeInterfaceGetDefault {
	return &NvmeInterfaceGetDefault{
		_statusCode: code,
	}
}

/* NvmeInterfaceGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 2621462 | The supplied SVM does not exist. |

*/
type NvmeInterfaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme interface get default response
func (o *NvmeInterfaceGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeInterfaceGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nvme/interfaces/{uuid}][%d] nvme_interface_get default  %+v", o._statusCode, o.Payload)
}
func (o *NvmeInterfaceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeInterfaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
