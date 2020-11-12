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

// GetProjectsProjectIDWebhookEventsReader is a Reader for the GetProjectsProjectIDWebhookEvents structure.
type GetProjectsProjectIDWebhookEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDWebhookEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDWebhookEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetProjectsProjectIDWebhookEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDWebhookEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDWebhookEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsProjectIDWebhookEventsOK creates a GetProjectsProjectIDWebhookEventsOK with default headers values
func NewGetProjectsProjectIDWebhookEventsOK() *GetProjectsProjectIDWebhookEventsOK {
	return &GetProjectsProjectIDWebhookEventsOK{}
}

/*GetProjectsProjectIDWebhookEventsOK handles this case with default header values.

Success
*/
type GetProjectsProjectIDWebhookEventsOK struct {
	Payload *models.SupportedWebhookEventTypes
}

func (o *GetProjectsProjectIDWebhookEventsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/events][%d] getProjectsProjectIdWebhookEventsOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDWebhookEventsOK) GetPayload() *models.SupportedWebhookEventTypes {
	return o.Payload
}

func (o *GetProjectsProjectIDWebhookEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedWebhookEventTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDWebhookEventsUnauthorized creates a GetProjectsProjectIDWebhookEventsUnauthorized with default headers values
func NewGetProjectsProjectIDWebhookEventsUnauthorized() *GetProjectsProjectIDWebhookEventsUnauthorized {
	return &GetProjectsProjectIDWebhookEventsUnauthorized{}
}

/*GetProjectsProjectIDWebhookEventsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDWebhookEventsUnauthorized struct {
}

func (o *GetProjectsProjectIDWebhookEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/events][%d] getProjectsProjectIdWebhookEventsUnauthorized ", 401)
}

func (o *GetProjectsProjectIDWebhookEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookEventsForbidden creates a GetProjectsProjectIDWebhookEventsForbidden with default headers values
func NewGetProjectsProjectIDWebhookEventsForbidden() *GetProjectsProjectIDWebhookEventsForbidden {
	return &GetProjectsProjectIDWebhookEventsForbidden{}
}

/*GetProjectsProjectIDWebhookEventsForbidden handles this case with default header values.

User have no permission to list webhook jobs of the project.
*/
type GetProjectsProjectIDWebhookEventsForbidden struct {
}

func (o *GetProjectsProjectIDWebhookEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/events][%d] getProjectsProjectIdWebhookEventsForbidden ", 403)
}

func (o *GetProjectsProjectIDWebhookEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookEventsInternalServerError creates a GetProjectsProjectIDWebhookEventsInternalServerError with default headers values
func NewGetProjectsProjectIDWebhookEventsInternalServerError() *GetProjectsProjectIDWebhookEventsInternalServerError {
	return &GetProjectsProjectIDWebhookEventsInternalServerError{}
}

/*GetProjectsProjectIDWebhookEventsInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDWebhookEventsInternalServerError struct {
}

func (o *GetProjectsProjectIDWebhookEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/events][%d] getProjectsProjectIdWebhookEventsInternalServerError ", 500)
}

func (o *GetProjectsProjectIDWebhookEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}