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

package ring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteRingAcceptedCode is the HTTP code returned for type DeleteRingAccepted
const DeleteRingAcceptedCode int = 202

/*DeleteRingAccepted Configuration change accepted and reload requested

swagger:response deleteRingAccepted
*/
type DeleteRingAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteRingAccepted creates DeleteRingAccepted with default headers values
func NewDeleteRingAccepted() *DeleteRingAccepted {

	return &DeleteRingAccepted{}
}

// WithReloadID adds the reloadId to the delete ring accepted response
func (o *DeleteRingAccepted) WithReloadID(reloadID string) *DeleteRingAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete ring accepted response
func (o *DeleteRingAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteRingAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteRingNoContentCode is the HTTP code returned for type DeleteRingNoContent
const DeleteRingNoContentCode int = 204

/*DeleteRingNoContent Ring deleted

swagger:response deleteRingNoContent
*/
type DeleteRingNoContent struct {
}

// NewDeleteRingNoContent creates DeleteRingNoContent with default headers values
func NewDeleteRingNoContent() *DeleteRingNoContent {

	return &DeleteRingNoContent{}
}

// WriteResponse to the client
func (o *DeleteRingNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteRingNotFoundCode is the HTTP code returned for type DeleteRingNotFound
const DeleteRingNotFoundCode int = 404

/*DeleteRingNotFound The specified resource was not found

swagger:response deleteRingNotFound
*/
type DeleteRingNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteRingNotFound creates DeleteRingNotFound with default headers values
func NewDeleteRingNotFound() *DeleteRingNotFound {

	return &DeleteRingNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete ring not found response
func (o *DeleteRingNotFound) WithConfigurationVersion(configurationVersion string) *DeleteRingNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete ring not found response
func (o *DeleteRingNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete ring not found response
func (o *DeleteRingNotFound) WithPayload(payload *models.Error) *DeleteRingNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ring not found response
func (o *DeleteRingNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteRingDefault General Error

swagger:response deleteRingDefault
*/
type DeleteRingDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteRingDefault creates DeleteRingDefault with default headers values
func NewDeleteRingDefault(code int) *DeleteRingDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteRingDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete ring default response
func (o *DeleteRingDefault) WithStatusCode(code int) *DeleteRingDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete ring default response
func (o *DeleteRingDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete ring default response
func (o *DeleteRingDefault) WithConfigurationVersion(configurationVersion string) *DeleteRingDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete ring default response
func (o *DeleteRingDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete ring default response
func (o *DeleteRingDefault) WithPayload(payload *models.Error) *DeleteRingDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ring default response
func (o *DeleteRingDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRingDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
