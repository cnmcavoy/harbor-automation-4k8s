// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/szlabs/harbor-automation-4k8s/pkg/sdk/harbor/models"
)

// GetRetentionsIDReader is a Reader for the GetRetentionsID structure.
type GetRetentionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRetentionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRetentionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRetentionsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRetentionsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRetentionsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRetentionsIDOK creates a GetRetentionsIDOK with default headers values
func NewGetRetentionsIDOK() *GetRetentionsIDOK {
	return &GetRetentionsIDOK{}
}

/*GetRetentionsIDOK handles this case with default header values.

Get Retention Policy successfully.
*/
type GetRetentionsIDOK struct {
	Payload *models.RetentionPolicy
}

func (o *GetRetentionsIDOK) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionsIdOK  %+v", 200, o.Payload)
}

func (o *GetRetentionsIDOK) GetPayload() *models.RetentionPolicy {
	return o.Payload
}

func (o *GetRetentionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetentionPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRetentionsIDUnauthorized creates a GetRetentionsIDUnauthorized with default headers values
func NewGetRetentionsIDUnauthorized() *GetRetentionsIDUnauthorized {
	return &GetRetentionsIDUnauthorized{}
}

/*GetRetentionsIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetRetentionsIDUnauthorized struct {
}

func (o *GetRetentionsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionsIdUnauthorized ", 401)
}

func (o *GetRetentionsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRetentionsIDForbidden creates a GetRetentionsIDForbidden with default headers values
func NewGetRetentionsIDForbidden() *GetRetentionsIDForbidden {
	return &GetRetentionsIDForbidden{}
}

/*GetRetentionsIDForbidden handles this case with default header values.

User have no permission.
*/
type GetRetentionsIDForbidden struct {
}

func (o *GetRetentionsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionsIdForbidden ", 403)
}

func (o *GetRetentionsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRetentionsIDInternalServerError creates a GetRetentionsIDInternalServerError with default headers values
func NewGetRetentionsIDInternalServerError() *GetRetentionsIDInternalServerError {
	return &GetRetentionsIDInternalServerError{}
}

/*GetRetentionsIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetRetentionsIDInternalServerError struct {
}

func (o *GetRetentionsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionsIdInternalServerError ", 500)
}

func (o *GetRetentionsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
