// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// HTTPProxyGetReader is a Reader for the HTTPProxyGet structure.
type HTTPProxyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPProxyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHTTPProxyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHTTPProxyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHTTPProxyGetOK creates a HTTPProxyGetOK with default headers values
func NewHTTPProxyGetOK() *HTTPProxyGetOK {
	return &HTTPProxyGetOK{}
}

/* HTTPProxyGetOK describes a response with status code 200, with default header values.

OK
*/
type HTTPProxyGetOK struct {
	Payload *models.NetworkHTTPProxy
}

func (o *HTTPProxyGetOK) Error() string {
	return fmt.Sprintf("[GET /network/http-proxy/{uuid}][%d] httpProxyGetOK  %+v", 200, o.Payload)
}
func (o *HTTPProxyGetOK) GetPayload() *models.NetworkHTTPProxy {
	return o.Payload
}

func (o *HTTPProxyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkHTTPProxy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHTTPProxyGetDefault creates a HTTPProxyGetDefault with default headers values
func NewHTTPProxyGetDefault(code int) *HTTPProxyGetDefault {
	return &HTTPProxyGetDefault{
		_statusCode: code,
	}
}

/* HTTPProxyGetDefault describes a response with status code -1, with default header values.

Error
*/
type HTTPProxyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the http proxy get default response
func (o *HTTPProxyGetDefault) Code() int {
	return o._statusCode
}

func (o *HTTPProxyGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/http-proxy/{uuid}][%d] http_proxy_get default  %+v", o._statusCode, o.Payload)
}
func (o *HTTPProxyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HTTPProxyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
