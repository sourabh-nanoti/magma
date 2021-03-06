// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

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

// NewGetCwfNetworkIDGatewaysParams creates a new GetCwfNetworkIDGatewaysParams object
// with the default values initialized.
func NewGetCwfNetworkIDGatewaysParams() *GetCwfNetworkIDGatewaysParams {
	var ()
	return &GetCwfNetworkIDGatewaysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDGatewaysParamsWithTimeout creates a new GetCwfNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCwfNetworkIDGatewaysParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysParams {
	var ()
	return &GetCwfNetworkIDGatewaysParams{

		timeout: timeout,
	}
}

// NewGetCwfNetworkIDGatewaysParamsWithContext creates a new GetCwfNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCwfNetworkIDGatewaysParamsWithContext(ctx context.Context) *GetCwfNetworkIDGatewaysParams {
	var ()
	return &GetCwfNetworkIDGatewaysParams{

		Context: ctx,
	}
}

// NewGetCwfNetworkIDGatewaysParamsWithHTTPClient creates a new GetCwfNetworkIDGatewaysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCwfNetworkIDGatewaysParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysParams {
	var ()
	return &GetCwfNetworkIDGatewaysParams{
		HTTPClient: client,
	}
}

/*GetCwfNetworkIDGatewaysParams contains all the parameters to send to the API endpoint
for the get cwf network ID gateways operation typically these are written to a http.Request
*/
type GetCwfNetworkIDGatewaysParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) WithContext(ctx context.Context) *GetCwfNetworkIDGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) WithNetworkID(networkID string) *GetCwfNetworkIDGatewaysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID gateways params
func (o *GetCwfNetworkIDGatewaysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
