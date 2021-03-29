// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NodeGetReader is a Reader for the NodeGet structure.
type NodeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeGetOK creates a NodeGetOK with default headers values
func NewNodeGetOK() *NodeGetOK {
	return &NodeGetOK{}
}

/* NodeGetOK describes a response with status code 200, with default header values.

OK
*/
type NodeGetOK struct {
	Payload *models.Node
}

func (o *NodeGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/nodes/{uuid}][%d] nodeGetOK  %+v", 200, o.Payload)
}
func (o *NodeGetOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *NodeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeGetDefault creates a NodeGetDefault with default headers values
func NewNodeGetDefault(code int) *NodeGetDefault {
	return &NodeGetDefault{
		_statusCode: code,
	}
}

/* NodeGetDefault describes a response with status code -1, with default header values.

Error
*/
type NodeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the node get default response
func (o *NodeGetDefault) Code() int {
	return o._statusCode
}

func (o *NodeGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/nodes/{uuid}][%d] node_get default  %+v", o._statusCode, o.Payload)
}
func (o *NodeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
