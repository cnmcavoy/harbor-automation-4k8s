// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectsProjectIDWebhookPoliciesPolicyIDReader is a Reader for the DeleteProjectsProjectIDWebhookPoliciesPolicyID structure.
type DeleteProjectsProjectIDWebhookPoliciesPolicyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDOK creates a DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK with default headers values
func NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDOK() *DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK {
	return &DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK{}
}

/*DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK handles this case with default header values.

Delete webhook policy successfully.
*/
type DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK struct {
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/webhook/policies/{policy_id}][%d] deleteProjectsProjectIdWebhookPoliciesPolicyIdOK ", 200)
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest creates a DeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest with default headers values
func NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest() *DeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest {
	return &DeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest{}
}

/*DeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type DeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest struct {
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/webhook/policies/{policy_id}][%d] deleteProjectsProjectIdWebhookPoliciesPolicyIdBadRequest ", 400)
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized creates a DeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized with default headers values
func NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized() *DeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized {
	return &DeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized{}
}

/*DeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type DeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized struct {
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/webhook/policies/{policy_id}][%d] deleteProjectsProjectIdWebhookPoliciesPolicyIdUnauthorized ", 401)
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden creates a DeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden with default headers values
func NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden() *DeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden {
	return &DeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden{}
}

/*DeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden handles this case with default header values.

User have no permission to delete webhook policy of the project.
*/
type DeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden struct {
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/webhook/policies/{policy_id}][%d] deleteProjectsProjectIdWebhookPoliciesPolicyIdForbidden ", 403)
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound creates a DeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound with default headers values
func NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound() *DeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound {
	return &DeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound{}
}

/*DeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound handles this case with default header values.

Webhook policy ID does not exist.
*/
type DeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound struct {
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/webhook/policies/{policy_id}][%d] deleteProjectsProjectIdWebhookPoliciesPolicyIdNotFound ", 404)
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError creates a DeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError with default headers values
func NewDeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError() *DeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError {
	return &DeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError{}
}

/*DeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError handles this case with default header values.

Internal server errors.
*/
type DeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError struct {
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}/webhook/policies/{policy_id}][%d] deleteProjectsProjectIdWebhookPoliciesPolicyIdInternalServerError ", 500)
}

func (o *DeleteProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
