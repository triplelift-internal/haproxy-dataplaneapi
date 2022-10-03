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

package process_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateProgramCreatedCode is the HTTP code returned for type CreateProgramCreated
const CreateProgramCreatedCode int = 201

/*
CreateProgramCreated Program created

swagger:response createProgramCreated
*/
type CreateProgramCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Program `json:"body,omitempty"`
}

// NewCreateProgramCreated creates CreateProgramCreated with default headers values
func NewCreateProgramCreated() *CreateProgramCreated {

	return &CreateProgramCreated{}
}

// WithPayload adds the payload to the create program created response
func (o *CreateProgramCreated) WithPayload(payload *models.Program) *CreateProgramCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create program created response
func (o *CreateProgramCreated) SetPayload(payload *models.Program) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProgramCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProgramAcceptedCode is the HTTP code returned for type CreateProgramAccepted
const CreateProgramAcceptedCode int = 202

/*
CreateProgramAccepted Configuration change accepted and reload requested

swagger:response createProgramAccepted
*/
type CreateProgramAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Program `json:"body,omitempty"`
}

// NewCreateProgramAccepted creates CreateProgramAccepted with default headers values
func NewCreateProgramAccepted() *CreateProgramAccepted {

	return &CreateProgramAccepted{}
}

// WithReloadID adds the reloadId to the create program accepted response
func (o *CreateProgramAccepted) WithReloadID(reloadID string) *CreateProgramAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create program accepted response
func (o *CreateProgramAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create program accepted response
func (o *CreateProgramAccepted) WithPayload(payload *models.Program) *CreateProgramAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create program accepted response
func (o *CreateProgramAccepted) SetPayload(payload *models.Program) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProgramAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProgramBadRequestCode is the HTTP code returned for type CreateProgramBadRequest
const CreateProgramBadRequestCode int = 400

/*
CreateProgramBadRequest Bad request

swagger:response createProgramBadRequest
*/
type CreateProgramBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateProgramBadRequest creates CreateProgramBadRequest with default headers values
func NewCreateProgramBadRequest() *CreateProgramBadRequest {

	return &CreateProgramBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create program bad request response
func (o *CreateProgramBadRequest) WithConfigurationVersion(configurationVersion string) *CreateProgramBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create program bad request response
func (o *CreateProgramBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create program bad request response
func (o *CreateProgramBadRequest) WithPayload(payload *models.Error) *CreateProgramBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create program bad request response
func (o *CreateProgramBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProgramBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProgramConflictCode is the HTTP code returned for type CreateProgramConflict
const CreateProgramConflictCode int = 409

/*
CreateProgramConflict The specified resource already exists

swagger:response createProgramConflict
*/
type CreateProgramConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateProgramConflict creates CreateProgramConflict with default headers values
func NewCreateProgramConflict() *CreateProgramConflict {

	return &CreateProgramConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create program conflict response
func (o *CreateProgramConflict) WithConfigurationVersion(configurationVersion string) *CreateProgramConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create program conflict response
func (o *CreateProgramConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create program conflict response
func (o *CreateProgramConflict) WithPayload(payload *models.Error) *CreateProgramConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create program conflict response
func (o *CreateProgramConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProgramConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateProgramDefault General Error

swagger:response createProgramDefault
*/
type CreateProgramDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateProgramDefault creates CreateProgramDefault with default headers values
func NewCreateProgramDefault(code int) *CreateProgramDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateProgramDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create program default response
func (o *CreateProgramDefault) WithStatusCode(code int) *CreateProgramDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create program default response
func (o *CreateProgramDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create program default response
func (o *CreateProgramDefault) WithConfigurationVersion(configurationVersion string) *CreateProgramDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create program default response
func (o *CreateProgramDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create program default response
func (o *CreateProgramDefault) WithPayload(payload *models.Error) *CreateProgramDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create program default response
func (o *CreateProgramDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProgramDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
