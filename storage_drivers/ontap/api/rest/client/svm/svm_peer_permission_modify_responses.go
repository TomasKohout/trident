// Code generated by go-swagger; DO NOT EDIT.

package svm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SvmPeerPermissionModifyReader is a Reader for the SvmPeerPermissionModify structure.
type SvmPeerPermissionModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerPermissionModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmPeerPermissionModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerPermissionModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerPermissionModifyOK creates a SvmPeerPermissionModifyOK with default headers values
func NewSvmPeerPermissionModifyOK() *SvmPeerPermissionModifyOK {
	return &SvmPeerPermissionModifyOK{}
}

/* SvmPeerPermissionModifyOK describes a response with status code 200, with default header values.

OK
*/
type SvmPeerPermissionModifyOK struct {
	Payload *models.SvmPeerPermission
}

func (o *SvmPeerPermissionModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /svm/peer-permissions/{cluster_peer.uuid}/{svm.uuid}][%d] svmPeerPermissionModifyOK  %+v", 200, o.Payload)
}
func (o *SvmPeerPermissionModifyOK) GetPayload() *models.SvmPeerPermission {
	return o.Payload
}

func (o *SvmPeerPermissionModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmPeerPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerPermissionModifyDefault creates a SvmPeerPermissionModifyDefault with default headers values
func NewSvmPeerPermissionModifyDefault(code int) *SvmPeerPermissionModifyDefault {
	return &SvmPeerPermissionModifyDefault{
		_statusCode: code,
	}
}

/* SvmPeerPermissionModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 26345572    | {field} is a required field. |
| 26345574    | Failed to find the SVM or volume name with UUID. |
```
<br/>

*/
type SvmPeerPermissionModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the svm peer permission modify default response
func (o *SvmPeerPermissionModifyDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerPermissionModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /svm/peer-permissions/{cluster_peer.uuid}/{svm.uuid}][%d] svm_peer_permission_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SvmPeerPermissionModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerPermissionModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
