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

// SnapmirrorRelationshipCreateReader is a Reader for the SnapmirrorRelationshipCreate structure.
type SnapmirrorRelationshipCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorRelationshipCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSnapmirrorRelationshipCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorRelationshipCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorRelationshipCreateAccepted creates a SnapmirrorRelationshipCreateAccepted with default headers values
func NewSnapmirrorRelationshipCreateAccepted() *SnapmirrorRelationshipCreateAccepted {
	return &SnapmirrorRelationshipCreateAccepted{}
}

/* SnapmirrorRelationshipCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorRelationshipCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SnapmirrorRelationshipCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirrorRelationshipCreateAccepted  %+v", 202, o.Payload)
}
func (o *SnapmirrorRelationshipCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipCreateDefault creates a SnapmirrorRelationshipCreateDefault with default headers values
func NewSnapmirrorRelationshipCreateDefault(code int) *SnapmirrorRelationshipCreateDefault {
	return &SnapmirrorRelationshipCreateDefault{
		_statusCode: code,
	}
}

/* SnapmirrorRelationshipCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 6620374 | Internal error. Failed to get SVM information. |
| 6620478 | Internal error. Failed to check SnapMirror capability. |
| 13303819 | Could not retrieve SnapMirror policy information. |
| 13303821 | Invalid SnapMirror policy UUID. |
| 13303841 | This operation is not supported for SnapMirror relationships between these endpoints. |
| 13303852 | destination.path provided does not contain \\\":\\\". |
| 13303853 | Restore relationships are not supported for SVM-DR endpoints. |
| 13303868 | Create of destination endpoint and SnapMirror relationship failed. |
| 13303869 | Creating a destination endpoint is not supported for restore relationships. |
| 13303870 | A tiering policy cannot be specified if tiering is not being set to supported. |
| 13303871 | Storage service properties cannot be specified if the storage service is not being enabled. |
| 13303872 | Specified property requires a later effective cluster version. |
| 13303873 | Specifying a state when creating a relationship is only supported when creating a destination endpoint. |
| 13303874 | Specified state is not supported when creating this relationship. |
| 13303875 | Destination aggregates do not have sufficient space for hosting copies of source volumes. |
| 13303876 | Destination cluster does not have composite aggregates. |
| 13303877 | Source or destination cluster must be specified. |
| 13303878 | The specified fields do not match. |
| 13303879 | Source cluster name or UUID is needed to provision a destination SVM on the local cluster. |
| 13303880 | Source cluster must be remote for provisioning a destination SVM on the local cluster. |
| 13303881 | Network validation failed. |
| 13303882 | SVM validation failed. |
| 13303883 | Encryption is not enabled on the destination cluster. |
| 13303887 | Synchronous SnapMirror relationships between FlexGroup volumes are not supported. |
| 13303888 | Synchronous SnapMirror relationships require an effective cluster version of 9.5 or later on both the source and destination clusters. |
| 13303889 | Asynchronous SnapMirror relationships between FlexGroup volumes require an effective cluster version of 9.5 or later on both the source and destination clusters. |
| 13303890 | Asynchronous SnapMirror relationships between FlexVol volumes require an effective cluster version of 9.3, 9.5, or later on both the source and destination clusters. |
| 13303891 | Creating a destination endpoint with storage service requires an effective cluster version of 9.7 or later. |
| 13303892 | Fetching remote information from the destination cluster failed. |
| 13303893 | Updating job description failed. |
| 13303894 | Destination volume name is invalid. It must contain the source volume name and have a suffix when creating a destination endpoint on a cluster with an effective cluster version of 9.6 or earlier. |
| 13303895 | Operation on the remote destination cluster is not supported. |
| 13303916 | FlexGroup volumes are not supported on SnapLock aggregates. |
| 13303918 | No suitable destination aggregate type is available. |
| 13303919 | Only FabricPool enabled aggregates are available on the destination. |
| 13303920 | Only SnapLock aggregates are available on the destination. FlexGroup volumes are not supported on SnapLock aggregates. |
| 13303921 | Unable to retrieve the SnapMirror capabilities of the destination cluster. |
| 13303922 | Specified source SVM is not a data SVM. |
| 13303923 | Specified destination SVM is not a data SVM. |
| 13303924 | Source SVM has an invalid Snapshot copy policy. |
| 13303925 | SnapMirror validation has failed. |
| 13303930 | The specified tiering policy is not supported for destination volumes of Synchronous relationships. |
| 13303938 | Fetching information from the local cluster failed. |
| 13303939 | Could not create an SVM peer relationship. |
| 13303944 | An SVM-DR relationship is not supported because the source SVM has CIFS configured and the associated SnapMirror policy has either the identity_preservation property not set or set to exclude_network_and_protocol_config. |

*/
type SnapmirrorRelationshipCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapmirror relationship create default response
func (o *SnapmirrorRelationshipCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorRelationshipCreateDefault) Error() string {
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirror_relationship_create default  %+v", o._statusCode, o.Payload)
}
func (o *SnapmirrorRelationshipCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
