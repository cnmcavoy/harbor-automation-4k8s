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

	"github.com/szlabs/harbor-automation-4k8s/pkg/sdk/harbor/models"
)

// NewPostReplicationPoliciesParams creates a new PostReplicationPoliciesParams object
// with the default values initialized.
func NewPostReplicationPoliciesParams() *PostReplicationPoliciesParams {
	var ()
	return &PostReplicationPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostReplicationPoliciesParamsWithTimeout creates a new PostReplicationPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostReplicationPoliciesParamsWithTimeout(timeout time.Duration) *PostReplicationPoliciesParams {
	var ()
	return &PostReplicationPoliciesParams{

		timeout: timeout,
	}
}

// NewPostReplicationPoliciesParamsWithContext creates a new PostReplicationPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostReplicationPoliciesParamsWithContext(ctx context.Context) *PostReplicationPoliciesParams {
	var ()
	return &PostReplicationPoliciesParams{

		Context: ctx,
	}
}

// NewPostReplicationPoliciesParamsWithHTTPClient creates a new PostReplicationPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostReplicationPoliciesParamsWithHTTPClient(client *http.Client) *PostReplicationPoliciesParams {
	var ()
	return &PostReplicationPoliciesParams{
		HTTPClient: client,
	}
}

/*PostReplicationPoliciesParams contains all the parameters to send to the API endpoint
for the post replication policies operation typically these are written to a http.Request
*/
type PostReplicationPoliciesParams struct {

	/*Policy
	  The policy model.

	*/
	Policy *models.ReplicationPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post replication policies params
func (o *PostReplicationPoliciesParams) WithTimeout(timeout time.Duration) *PostReplicationPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post replication policies params
func (o *PostReplicationPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post replication policies params
func (o *PostReplicationPoliciesParams) WithContext(ctx context.Context) *PostReplicationPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post replication policies params
func (o *PostReplicationPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post replication policies params
func (o *PostReplicationPoliciesParams) WithHTTPClient(client *http.Client) *PostReplicationPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post replication policies params
func (o *PostReplicationPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the post replication policies params
func (o *PostReplicationPoliciesParams) WithPolicy(policy *models.ReplicationPolicy) *PostReplicationPoliciesParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the post replication policies params
func (o *PostReplicationPoliciesParams) SetPolicy(policy *models.ReplicationPolicy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *PostReplicationPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
