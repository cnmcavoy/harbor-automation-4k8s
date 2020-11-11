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

// GetUsersCurrentPermissionsReader is a Reader for the GetUsersCurrentPermissions structure.
type GetUsersCurrentPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersCurrentPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersCurrentPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersCurrentPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersCurrentPermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersCurrentPermissionsOK creates a GetUsersCurrentPermissionsOK with default headers values
func NewGetUsersCurrentPermissionsOK() *GetUsersCurrentPermissionsOK {
	return &GetUsersCurrentPermissionsOK{}
}

/*GetUsersCurrentPermissionsOK handles this case with default header values.

Get current user permission successfully.
*/
type GetUsersCurrentPermissionsOK struct {
	Payload []*models.Permission
}

func (o *GetUsersCurrentPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /users/current/permissions][%d] getUsersCurrentPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetUsersCurrentPermissionsOK) GetPayload() []*models.Permission {
	return o.Payload
}

func (o *GetUsersCurrentPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersCurrentPermissionsUnauthorized creates a GetUsersCurrentPermissionsUnauthorized with default headers values
func NewGetUsersCurrentPermissionsUnauthorized() *GetUsersCurrentPermissionsUnauthorized {
	return &GetUsersCurrentPermissionsUnauthorized{}
}

/*GetUsersCurrentPermissionsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetUsersCurrentPermissionsUnauthorized struct {
}

func (o *GetUsersCurrentPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/permissions][%d] getUsersCurrentPermissionsUnauthorized ", 401)
}

func (o *GetUsersCurrentPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersCurrentPermissionsInternalServerError creates a GetUsersCurrentPermissionsInternalServerError with default headers values
func NewGetUsersCurrentPermissionsInternalServerError() *GetUsersCurrentPermissionsInternalServerError {
	return &GetUsersCurrentPermissionsInternalServerError{}
}

/*GetUsersCurrentPermissionsInternalServerError handles this case with default header values.

Internal errors.
*/
type GetUsersCurrentPermissionsInternalServerError struct {
}

func (o *GetUsersCurrentPermissionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/permissions][%d] getUsersCurrentPermissionsInternalServerError ", 500)
}

func (o *GetUsersCurrentPermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
