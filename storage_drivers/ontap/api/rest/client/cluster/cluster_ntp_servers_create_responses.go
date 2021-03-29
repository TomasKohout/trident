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

// ClusterNtpServersCreateReader is a Reader for the ClusterNtpServersCreate structure.
type ClusterNtpServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewClusterNtpServersCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpServersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpServersCreateAccepted creates a ClusterNtpServersCreateAccepted with default headers values
func NewClusterNtpServersCreateAccepted() *ClusterNtpServersCreateAccepted {
	return &ClusterNtpServersCreateAccepted{}
}

/* ClusterNtpServersCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterNtpServersCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *ClusterNtpServersCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] clusterNtpServersCreateAccepted  %+v", 202, o.Payload)
}
func (o *ClusterNtpServersCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersCreateDefault creates a ClusterNtpServersCreateDefault with default headers values
func NewClusterNtpServersCreateDefault(code int) *ClusterNtpServersCreateDefault {
	return &ClusterNtpServersCreateDefault{
		_statusCode: code,
	}
}

/* ClusterNtpServersCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 2097163 | NTP server IPv4 address was invalid. |
| 2097164 | NTP server IPv6 address was invalid. |
| 2097165 | Cannot resolve NTP server name. |
| 2097166 | NTP server address query returned no valid IP addresses. |
| 2097167 | Failed to connect to NTP server. |
| 2097169 | NTP server provided was not synchronized with a clock or another NTP server. |
| 2097174 | NTP server provided had too high of root distance. |
| 2097177 | NTP server provided an invalid stratum. |
| 2097179 | Too many NTP servers have been configured. |
| 2097181 | NTP server address was invalid. It is a special purpose address such as loopback, multicast, or broadcast address. |
| 2097182 | NTP server address was invalid. The address is neither an IPv4 or IPv6. |
| 2097183 | NTP symmetric key authentication cannot be used for a node not in a cluster. |
| 2097185 | NTP key authentication failed for the provided key. |
| 2097193 | An unknown NTP key was provided. |

*/
type ClusterNtpServersCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster ntp servers create default response
func (o *ClusterNtpServersCreateDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpServersCreateDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] cluster_ntp_servers_create default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterNtpServersCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpServersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
