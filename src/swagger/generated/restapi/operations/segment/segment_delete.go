// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SegmentDeleteHandlerFunc turns a function with the right signature into a segment delete handler
type SegmentDeleteHandlerFunc func(SegmentDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SegmentDeleteHandlerFunc) Handle(params SegmentDeleteParams) middleware.Responder {
	return fn(params)
}

// SegmentDeleteHandler interface for that can handle valid segment delete params
type SegmentDeleteHandler interface {
	Handle(SegmentDeleteParams) middleware.Responder
}

// NewSegmentDelete creates a new http.Handler for the segment delete operation
func NewSegmentDelete(ctx *middleware.Context, handler SegmentDeleteHandler) *SegmentDelete {
	return &SegmentDelete{Context: ctx, Handler: handler}
}

/*
	SegmentDelete swagger:route DELETE /segment Segment segmentDelete

# Removes existing segment

Accepts segment slug and removes it
*/
type SegmentDelete struct {
	Context *middleware.Context
	Handler SegmentDeleteHandler
}

func (o *SegmentDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSegmentDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}