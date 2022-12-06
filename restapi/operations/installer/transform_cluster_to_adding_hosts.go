// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TransformClusterToAddingHostsHandlerFunc turns a function with the right signature into a transform cluster to adding hosts handler
type TransformClusterToAddingHostsHandlerFunc func(TransformClusterToAddingHostsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn TransformClusterToAddingHostsHandlerFunc) Handle(params TransformClusterToAddingHostsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// TransformClusterToAddingHostsHandler interface for that can handle valid transform cluster to adding hosts params
type TransformClusterToAddingHostsHandler interface {
	Handle(TransformClusterToAddingHostsParams, interface{}) middleware.Responder
}

// NewTransformClusterToAddingHosts creates a new http.Handler for the transform cluster to adding hosts operation
func NewTransformClusterToAddingHosts(ctx *middleware.Context, handler TransformClusterToAddingHostsHandler) *TransformClusterToAddingHosts {
	return &TransformClusterToAddingHosts{Context: ctx, Handler: handler}
}

/* TransformClusterToAddingHosts swagger:route POST /v2/clusters/{cluster_id}/actions/allow-add-hosts installer transformClusterToAddingHosts

Transforms installed cluster to a state which allows adding hosts.

*/
type TransformClusterToAddingHosts struct {
	Context *middleware.Context
	Handler TransformClusterToAddingHostsHandler
}

func (o *TransformClusterToAddingHosts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTransformClusterToAddingHostsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
