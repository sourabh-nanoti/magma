// Code generated by go-swagger; DO NOT EDIT.

package network_probes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParams creates a new DeleteLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized.
func NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParams() *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &DeleteLTENetworkIDNetworkProbeTasksTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParamsWithTimeout creates a new DeleteLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParamsWithTimeout(timeout time.Duration) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &DeleteLTENetworkIDNetworkProbeTasksTaskIDParams{

		timeout: timeout,
	}
}

// NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParamsWithContext creates a new DeleteLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParamsWithContext(ctx context.Context) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &DeleteLTENetworkIDNetworkProbeTasksTaskIDParams{

		Context: ctx,
	}
}

// NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParamsWithHTTPClient creates a new DeleteLTENetworkIDNetworkProbeTasksTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParamsWithHTTPClient(client *http.Client) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	var ()
	return &DeleteLTENetworkIDNetworkProbeTasksTaskIDParams{
		HTTPClient: client,
	}
}

/*DeleteLTENetworkIDNetworkProbeTasksTaskIDParams contains all the parameters to send to the API endpoint
for the delete LTE network ID network probe tasks task ID operation typically these are written to a http.Request
*/
type DeleteLTENetworkIDNetworkProbeTasksTaskIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*TaskID
	  Network Probe Task ID

	*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) WithTimeout(timeout time.Duration) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) WithContext(ctx context.Context) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) WithHTTPClient(client *http.Client) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) WithNetworkID(networkID string) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTaskID adds the taskID to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) WithTaskID(taskID string) *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the delete LTE network ID network probe tasks task ID params
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
