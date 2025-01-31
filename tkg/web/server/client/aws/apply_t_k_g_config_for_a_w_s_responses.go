// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// ApplyTKGConfigForAWSReader is a Reader for the ApplyTKGConfigForAWS structure.
type ApplyTKGConfigForAWSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplyTKGConfigForAWSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplyTKGConfigForAWSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplyTKGConfigForAWSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewApplyTKGConfigForAWSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplyTKGConfigForAWSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApplyTKGConfigForAWSOK creates a ApplyTKGConfigForAWSOK with default headers values
func NewApplyTKGConfigForAWSOK() *ApplyTKGConfigForAWSOK {
	return &ApplyTKGConfigForAWSOK{}
}

/*
ApplyTKGConfigForAWSOK handles this case with default header values.

Apply change to TKG configuration successfully
*/
type ApplyTKGConfigForAWSOK struct {
	Payload *models.ConfigFileInfo
}

func (o *ApplyTKGConfigForAWSOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/tkgconfig][%d] applyTKGConfigForAWSOK  %+v", 200, o.Payload)
}

func (o *ApplyTKGConfigForAWSOK) GetPayload() *models.ConfigFileInfo {
	return o.Payload
}

func (o *ApplyTKGConfigForAWSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigFileInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForAWSBadRequest creates a ApplyTKGConfigForAWSBadRequest with default headers values
func NewApplyTKGConfigForAWSBadRequest() *ApplyTKGConfigForAWSBadRequest {
	return &ApplyTKGConfigForAWSBadRequest{}
}

/*
ApplyTKGConfigForAWSBadRequest handles this case with default header values.

Bad request
*/
type ApplyTKGConfigForAWSBadRequest struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForAWSBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/tkgconfig][%d] applyTKGConfigForAWSBadRequest  %+v", 400, o.Payload)
}

func (o *ApplyTKGConfigForAWSBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForAWSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForAWSUnauthorized creates a ApplyTKGConfigForAWSUnauthorized with default headers values
func NewApplyTKGConfigForAWSUnauthorized() *ApplyTKGConfigForAWSUnauthorized {
	return &ApplyTKGConfigForAWSUnauthorized{}
}

/*
ApplyTKGConfigForAWSUnauthorized handles this case with default header values.

Incorrect credentials
*/
type ApplyTKGConfigForAWSUnauthorized struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForAWSUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/tkgconfig][%d] applyTKGConfigForAWSUnauthorized  %+v", 401, o.Payload)
}

func (o *ApplyTKGConfigForAWSUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForAWSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForAWSInternalServerError creates a ApplyTKGConfigForAWSInternalServerError with default headers values
func NewApplyTKGConfigForAWSInternalServerError() *ApplyTKGConfigForAWSInternalServerError {
	return &ApplyTKGConfigForAWSInternalServerError{}
}

/*
ApplyTKGConfigForAWSInternalServerError handles this case with default header values.

Internal server error
*/
type ApplyTKGConfigForAWSInternalServerError struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForAWSInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/tkgconfig][%d] applyTKGConfigForAWSInternalServerError  %+v", 500, o.Payload)
}

func (o *ApplyTKGConfigForAWSInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForAWSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
