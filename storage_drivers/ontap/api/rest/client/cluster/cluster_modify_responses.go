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

// ClusterModifyReader is a Reader for the ClusterModify structure.
type ClusterModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewClusterModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterModifyAccepted creates a ClusterModifyAccepted with default headers values
func NewClusterModifyAccepted() *ClusterModifyAccepted {
	return &ClusterModifyAccepted{}
}

/* ClusterModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *ClusterModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /cluster][%d] clusterModifyAccepted  %+v", 202, o.Payload)
}
func (o *ClusterModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterModifyDefault creates a ClusterModifyDefault with default headers values
func NewClusterModifyDefault(code int) *ClusterModifyDefault {
	return &ClusterModifyDefault{
		_statusCode: code,
	}
}

/* ClusterModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 3604491 | Updating timezone failed. |
| 3604520 | Internal error. System state is not correct to read or change timezone. |
| 8847361 | Too many DNS domains provided. |
| 8847362 | Too many name servers provided. |
| 9240587 | A name must be provided. |
| 12451843 | Certificate does not exist. |

*/
type ClusterModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster modify default response
func (o *ClusterModifyDefault) Code() int {
	return o._statusCode
}

func (o *ClusterModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /cluster][%d] cluster_modify default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
