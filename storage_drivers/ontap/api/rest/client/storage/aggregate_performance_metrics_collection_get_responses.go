// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AggregatePerformanceMetricsCollectionGetReader is a Reader for the AggregatePerformanceMetricsCollectionGet structure.
type AggregatePerformanceMetricsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatePerformanceMetricsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatePerformanceMetricsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregatePerformanceMetricsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregatePerformanceMetricsCollectionGetOK creates a AggregatePerformanceMetricsCollectionGetOK with default headers values
func NewAggregatePerformanceMetricsCollectionGetOK() *AggregatePerformanceMetricsCollectionGetOK {
	return &AggregatePerformanceMetricsCollectionGetOK{}
}

/* AggregatePerformanceMetricsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AggregatePerformanceMetricsCollectionGetOK struct {
	Payload *models.PerformanceMetricResponse
}

func (o *AggregatePerformanceMetricsCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}/metrics][%d] aggregatePerformanceMetricsCollectionGetOK  %+v", 200, o.Payload)
}
func (o *AggregatePerformanceMetricsCollectionGetOK) GetPayload() *models.PerformanceMetricResponse {
	return o.Payload
}

func (o *AggregatePerformanceMetricsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatePerformanceMetricsCollectionGetDefault creates a AggregatePerformanceMetricsCollectionGetDefault with default headers values
func NewAggregatePerformanceMetricsCollectionGetDefault(code int) *AggregatePerformanceMetricsCollectionGetDefault {
	return &AggregatePerformanceMetricsCollectionGetDefault{
		_statusCode: code,
	}
}

/* AggregatePerformanceMetricsCollectionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 8586225 | Encountered unexpected error in retrieving metrics for the requested aggregate. |

*/
type AggregatePerformanceMetricsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the aggregate performance metrics collection get default response
func (o *AggregatePerformanceMetricsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AggregatePerformanceMetricsCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}/metrics][%d] aggregate_performance_metrics_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *AggregatePerformanceMetricsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AggregatePerformanceMetricsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
