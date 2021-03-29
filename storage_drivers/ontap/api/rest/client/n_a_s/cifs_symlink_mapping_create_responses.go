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

// CifsSymlinkMappingCreateReader is a Reader for the CifsSymlinkMappingCreate structure.
type CifsSymlinkMappingCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSymlinkMappingCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCifsSymlinkMappingCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSymlinkMappingCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSymlinkMappingCreateCreated creates a CifsSymlinkMappingCreateCreated with default headers values
func NewCifsSymlinkMappingCreateCreated() *CifsSymlinkMappingCreateCreated {
	return &CifsSymlinkMappingCreateCreated{}
}

/* CifsSymlinkMappingCreateCreated describes a response with status code 201, with default header values.

Created
*/
type CifsSymlinkMappingCreateCreated struct {
	Payload *models.CifsSymlinkMappingResponse
}

func (o *CifsSymlinkMappingCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/unix-symlink-mapping][%d] cifsSymlinkMappingCreateCreated  %+v", 201, o.Payload)
}
func (o *CifsSymlinkMappingCreateCreated) GetPayload() *models.CifsSymlinkMappingResponse {
	return o.Payload
}

func (o *CifsSymlinkMappingCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSymlinkMappingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsSymlinkMappingCreateDefault creates a CifsSymlinkMappingCreateDefault with default headers values
func NewCifsSymlinkMappingCreateDefault(code int) *CifsSymlinkMappingCreateDefault {
	return &CifsSymlinkMappingCreateDefault{
		_statusCode: code,
	}
}

/* CifsSymlinkMappingCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 655654     | Must specify the target CIFS share while creating path mapping entries with localities "local" or "widelink" |
| 655572     | The target path contains illegal characters or is too long |
| 655574     | The target server contains illegal characters or is too long |
| 655436     | If the locality is "local", the target server must be blank or must match the CIFS NetBIOS name for given SVM |
| 655439     | The Specified target server is local CIFS server for given SVM but the locality is specified as "widelink" |
| 655546     | Failed to create symlink mapping becasue administrative share cannot be used as target share |
| 655437     | Failed to create the symlink mapping with locality "local" because the target share does not exist for specified SVM |
| 655429     | UNIX path must begin and end with a "/" |
| 655430     | Target path must begin and end with a "/" |
| 655399     | Failed to get the CIFS server for specified SVM |

*/
type CifsSymlinkMappingCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs symlink mapping create default response
func (o *CifsSymlinkMappingCreateDefault) Code() int {
	return o._statusCode
}

func (o *CifsSymlinkMappingCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/unix-symlink-mapping][%d] cifs_symlink_mapping_create default  %+v", o._statusCode, o.Payload)
}
func (o *CifsSymlinkMappingCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSymlinkMappingCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
