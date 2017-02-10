package dhcp_reservations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rackn/rocket-skates/models"
)

// POSTDhcpReservationHandlerFunc turns a function with the right signature into a p o s t dhcp reservation handler
type POSTDhcpReservationHandlerFunc func(POSTDhcpReservationParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn POSTDhcpReservationHandlerFunc) Handle(params POSTDhcpReservationParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// POSTDhcpReservationHandler interface for that can handle valid p o s t dhcp reservation params
type POSTDhcpReservationHandler interface {
	Handle(POSTDhcpReservationParams, *models.Principal) middleware.Responder
}

// NewPOSTDhcpReservation creates a new http.Handler for the p o s t dhcp reservation operation
func NewPOSTDhcpReservation(ctx *middleware.Context, handler POSTDhcpReservationHandler) *POSTDhcpReservation {
	return &POSTDhcpReservation{Context: ctx, Handler: handler}
}

/*POSTDhcpReservation swagger:route POST /reservations Dhcp reservations pOSTDhcpReservation

Create DHCP Reservation

*/
type POSTDhcpReservation struct {
	Context *middleware.Context
	Handler POSTDhcpReservationHandler
}

func (o *POSTDhcpReservation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPOSTDhcpReservationParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
