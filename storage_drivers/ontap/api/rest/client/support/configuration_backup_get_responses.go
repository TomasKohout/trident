// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ConfigurationBackupGetReader is a Reader for the ConfigurationBackupGet structure.
type ConfigurationBackupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigurationBackupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigurationBackupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigurationBackupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigurationBackupGetOK creates a ConfigurationBackupGetOK with default headers values
func NewConfigurationBackupGetOK() *ConfigurationBackupGetOK {
	return &ConfigurationBackupGetOK{}
}

/* ConfigurationBackupGetOK describes a response with status code 200, with default header values.

OK
*/
type ConfigurationBackupGetOK struct {
	Payload *models.ConfigurationBackup
}

func (o *ConfigurationBackupGetOK) Error() string {
	return fmt.Sprintf("[GET /support/configuration-backup][%d] configurationBackupGetOK  %+v", 200, o.Payload)
}
func (o *ConfigurationBackupGetOK) GetPayload() *models.ConfigurationBackup {
	return o.Payload
}

func (o *ConfigurationBackupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurationBackupGetDefault creates a ConfigurationBackupGetDefault with default headers values
func NewConfigurationBackupGetDefault(code int) *ConfigurationBackupGetDefault {
	return &ConfigurationBackupGetDefault{
		_statusCode: code,
	}
}

/* ConfigurationBackupGetDefault describes a response with status code -1, with default header values.

Error
*/
type ConfigurationBackupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the configuration backup get default response
func (o *ConfigurationBackupGetDefault) Code() int {
	return o._statusCode
}

func (o *ConfigurationBackupGetDefault) Error() string {
	return fmt.Sprintf("[GET /support/configuration-backup][%d] configuration_backup_get default  %+v", o._statusCode, o.Payload)
}
func (o *ConfigurationBackupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurationBackupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
