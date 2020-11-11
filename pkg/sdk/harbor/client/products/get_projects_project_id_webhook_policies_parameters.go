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
)

// NewGetProjectsProjectIDWebhookPoliciesParams creates a new GetProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized.
func NewGetProjectsProjectIDWebhookPoliciesParams() *GetProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &GetProjectsProjectIDWebhookPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsProjectIDWebhookPoliciesParamsWithTimeout creates a new GetProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectsProjectIDWebhookPoliciesParamsWithTimeout(timeout time.Duration) *GetProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &GetProjectsProjectIDWebhookPoliciesParams{

		timeout: timeout,
	}
}

// NewGetProjectsProjectIDWebhookPoliciesParamsWithContext creates a new GetProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectsProjectIDWebhookPoliciesParamsWithContext(ctx context.Context) *GetProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &GetProjectsProjectIDWebhookPoliciesParams{

		Context: ctx,
	}
}

// NewGetProjectsProjectIDWebhookPoliciesParamsWithHTTPClient creates a new GetProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectsProjectIDWebhookPoliciesParamsWithHTTPClient(client *http.Client) *GetProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &GetProjectsProjectIDWebhookPoliciesParams{
		HTTPClient: client,
	}
}

/*GetProjectsProjectIDWebhookPoliciesParams contains all the parameters to send to the API endpoint
for the get projects project ID webhook policies operation typically these are written to a http.Request
*/
type GetProjectsProjectIDWebhookPoliciesParams struct {

	/*ProjectID
	  Relevant project ID.

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) WithTimeout(timeout time.Duration) *GetProjectsProjectIDWebhookPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) WithContext(ctx context.Context) *GetProjectsProjectIDWebhookPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) WithHTTPClient(client *http.Client) *GetProjectsProjectIDWebhookPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) WithProjectID(projectID int64) *GetProjectsProjectIDWebhookPoliciesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get projects project ID webhook policies params
func (o *GetProjectsProjectIDWebhookPoliciesParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsProjectIDWebhookPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
