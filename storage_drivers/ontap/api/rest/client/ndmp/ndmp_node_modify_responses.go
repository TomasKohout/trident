// Code generated by go-swagger; DO NOT EDIT.

package ndmp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NdmpNodeModifyReader is a Reader for the NdmpNodeModify structure.
type NdmpNodeModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpNodeModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpNodeModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpNodeModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpNodeModifyOK creates a NdmpNodeModifyOK with default headers values
func NewNdmpNodeModifyOK() *NdmpNodeModifyOK {
	return &NdmpNodeModifyOK{}
}

/* NdmpNodeModifyOK describes a response with status code 200, with default header values.

OK
*/
type NdmpNodeModifyOK struct {
	Payload *models.NdmpNode
}

func (o *NdmpNodeModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/ndmp/nodes/{node.uuid}][%d] ndmpNodeModifyOK  %+v", 200, o.Payload)
}
func (o *NdmpNodeModifyOK) GetPayload() *models.NdmpNode {
	return o.Payload
}

func (o *NdmpNodeModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpNodeModifyDefault creates a NdmpNodeModifyDefault with default headers values
func NewNdmpNodeModifyDefault(code int) *NdmpNodeModifyDefault {
	return &NdmpNodeModifyDefault{
		_statusCode: code,
	}
}

/* NdmpNodeModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 68812800    | The user is required to enable NDMP on a node.|
| 68812801    | Node-scoped operations are not allowed in an SVM-scope.|
| 68812802    | The UUID is not valid.|

*/
type NdmpNodeModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ndmp node modify default response
func (o *NdmpNodeModifyDefault) Code() int {
	return o._statusCode
}

func (o *NdmpNodeModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/ndmp/nodes/{node.uuid}][%d] ndmp_node_modify default  %+v", o._statusCode, o.Payload)
}
func (o *NdmpNodeModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpNodeModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
