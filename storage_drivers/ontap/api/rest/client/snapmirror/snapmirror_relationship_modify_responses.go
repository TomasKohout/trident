// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnapmirrorRelationshipModifyReader is a Reader for the SnapmirrorRelationshipModify structure.
type SnapmirrorRelationshipModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorRelationshipModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSnapmirrorRelationshipModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorRelationshipModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorRelationshipModifyAccepted creates a SnapmirrorRelationshipModifyAccepted with default headers values
func NewSnapmirrorRelationshipModifyAccepted() *SnapmirrorRelationshipModifyAccepted {
	return &SnapmirrorRelationshipModifyAccepted{}
}

/* SnapmirrorRelationshipModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorRelationshipModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SnapmirrorRelationshipModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipModifyAccepted  %+v", 202, o.Payload)
}
func (o *SnapmirrorRelationshipModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipModifyDefault creates a SnapmirrorRelationshipModifyDefault with default headers values
func NewSnapmirrorRelationshipModifyDefault(code int) *SnapmirrorRelationshipModifyDefault {
	return &SnapmirrorRelationshipModifyDefault{
		_statusCode: code,
	}
}

/* SnapmirrorRelationshipModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 13303825    | Could not retrieve information for the SnapMirror policy type |
| 13303817    | Unknown value for the SnapMirror state |
| 13303829    | Invalid state |
| 13303830    | Transient state |
| 13303831    | Invalid state for async SnapMirror relationship |
| 13303834    | Given input valid only for FlexGroup SnapMirror relationship |
| 13303835    | Given flag is valid only when PATCH state is broken_off |
| 13303836    | Given flag is valid only when PATCH state is snapmirrored or in_sync |
| 13303818    | Invalid state transition requested |
| 13303828    | Given state change is not possible for SVM SnapMirror relationship |
| 13303833    | Requested state change is not possible |
| 13303832    | SnapMirror relationship is already initialized |
| 13303824    | Quiescing the SnapMirror relationship has failed |
| 13303826    | Required environment variables are not set |
| 13303827    | Internal Error |
| 13303823    | Quiesce operation timed out |
| 13303821    | Invalid SnapMirror policy name/UUID |
| 13303819    | Could not retrieve SnapMirror policy information |
| 13303851    | Cannot modify attributes of SnapMirror restore relationship |
| 13303816    | Could not retrieve state or status values |
| 13303837    | Given flags are valid only if SnapMirror state change is requested |
| 6619546     | Destination must be a dp volume |
| 13303808    | Transition to broken_off state failed |
| 13303809    | Transition to paused state failed |
| 13303810    | Transition to snapmirrored state failed |
| 13303811    | Transition from paused state failed |
| 13303820    | SnapMirror policy was successfully updated, state transition failed |
| 13303856    | SVM is not configured with any data protocol |
| 13303857    | SVM is not configured with any network interface |
| 13303858    | Internal error. Failed to check LIF and protocols details for SVM |
| 13303859    | Internal error. SVM Failover operation failed. SVM operational state is unavailable. |
| 13303865    | Modifying the specified SnapMirror policy is not supported. |
| 13303866    | Cannot use the specified policy to modify the policy of the relationship. |
| 13303867    | Modifying the policy of an async-mirror or a vault relationship is not supported. |
| 13303884    | LIF and protocols details are configured incorrectly for SVM. |

*/
type SnapmirrorRelationshipModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapmirror relationship modify default response
func (o *SnapmirrorRelationshipModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorRelationshipModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirror_relationship_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SnapmirrorRelationshipModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
