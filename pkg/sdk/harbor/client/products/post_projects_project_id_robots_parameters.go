// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/szlabs/harbor-automation-4k8s/pkg/sdk/harbor/models"
)

// NewPostProjectsProjectIDRobotsParams creates a new PostProjectsProjectIDRobotsParams object
// with the default values initialized.
func NewPostProjectsProjectIDRobotsParams() *PostProjectsProjectIDRobotsParams {
	var ()
	return &PostProjectsProjectIDRobotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProjectsProjectIDRobotsParamsWithTimeout creates a new PostProjectsProjectIDRobotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProjectsProjectIDRobotsParamsWithTimeout(timeout time.Duration) *PostProjectsProjectIDRobotsParams {
	var ()
	return &PostProjectsProjectIDRobotsParams{

		timeout: timeout,
	}
}

// NewPostProjectsProjectIDRobotsParamsWithContext creates a new PostProjectsProjectIDRobotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProjectsProjectIDRobotsParamsWithContext(ctx context.Context) *PostProjectsProjectIDRobotsParams {
	var ()
	return &PostProjectsProjectIDRobotsParams{

		Context: ctx,
	}
}

// NewPostProjectsProjectIDRobotsParamsWithHTTPClient creates a new PostProjectsProjectIDRobotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProjectsProjectIDRobotsParamsWithHTTPClient(client *http.Client) *PostProjectsProjectIDRobotsParams {
	var ()
	return &PostProjectsProjectIDRobotsParams{
		HTTPClient: client,
	}
}

/*PostProjectsProjectIDRobotsParams contains all the parameters to send to the API endpoint
for the post projects project ID robots operation typically these are written to a http.Request
*/
type PostProjectsProjectIDRobotsParams struct {

	/*ProjectID
	  Relevant project ID.

	*/
	ProjectID int64
	/*Robot
	  Request body of creating a robot account.

	*/
	Robot *models.RobotAccountCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) WithTimeout(timeout time.Duration) *PostProjectsProjectIDRobotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) WithContext(ctx context.Context) *PostProjectsProjectIDRobotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) WithHTTPClient(client *http.Client) *PostProjectsProjectIDRobotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) WithProjectID(projectID int64) *PostProjectsProjectIDRobotsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithRobot adds the robot to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) WithRobot(robot *models.RobotAccountCreate) *PostProjectsProjectIDRobotsParams {
	o.SetRobot(robot)
	return o
}

// SetRobot adds the robot to the post projects project ID robots params
func (o *PostProjectsProjectIDRobotsParams) SetRobot(robot *models.RobotAccountCreate) {
	o.Robot = robot
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectsProjectIDRobotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if o.Robot != nil {
		if err := r.SetBodyParam(o.Robot); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
