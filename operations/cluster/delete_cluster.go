// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteClusterHandlerFunc turns a function with the right signature into a delete cluster handler
type DeleteClusterHandlerFunc func(DeleteClusterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteClusterHandlerFunc) Handle(params DeleteClusterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteClusterHandler interface for that can handle valid delete cluster params
type DeleteClusterHandler interface {
	Handle(DeleteClusterParams, interface{}) middleware.Responder
}

// NewDeleteCluster creates a new http.Handler for the delete cluster operation
func NewDeleteCluster(ctx *middleware.Context, handler DeleteClusterHandler) *DeleteCluster {
	return &DeleteCluster{Context: ctx, Handler: handler}
}

/*DeleteCluster swagger:route DELETE /cluster Cluster deleteCluster

Delete cluster settings

Delete cluster settings and move the node back to single mode

*/
type DeleteCluster struct {
	Context *middleware.Context
	Handler DeleteClusterHandler
}

func (o *DeleteCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteClusterParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
