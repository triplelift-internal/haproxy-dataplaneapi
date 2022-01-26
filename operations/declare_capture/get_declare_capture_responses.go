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

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// GetDeclareCaptureOKCode is the HTTP code returned for type GetDeclareCaptureOK
const GetDeclareCaptureOKCode int = 200

/*GetDeclareCaptureOK Successful operation

swagger:response getDeclareCaptureOK
*/
type GetDeclareCaptureOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetDeclareCaptureOKBody `json:"body,omitempty"`
}

// NewGetDeclareCaptureOK creates GetDeclareCaptureOK with default headers values
func NewGetDeclareCaptureOK() *GetDeclareCaptureOK {

	return &GetDeclareCaptureOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get declare capture o k response
func (o *GetDeclareCaptureOK) WithConfigurationVersion(configurationVersion string) *GetDeclareCaptureOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get declare capture o k response
func (o *GetDeclareCaptureOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get declare capture o k response
func (o *GetDeclareCaptureOK) WithPayload(payload *GetDeclareCaptureOKBody) *GetDeclareCaptureOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get declare capture o k response
func (o *GetDeclareCaptureOK) SetPayload(payload *GetDeclareCaptureOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeclareCaptureOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDeclareCaptureNotFoundCode is the HTTP code returned for type GetDeclareCaptureNotFound
const GetDeclareCaptureNotFoundCode int = 404

/*GetDeclareCaptureNotFound The specified resource already exists

swagger:response getDeclareCaptureNotFound
*/
type GetDeclareCaptureNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDeclareCaptureNotFound creates GetDeclareCaptureNotFound with default headers values
func NewGetDeclareCaptureNotFound() *GetDeclareCaptureNotFound {

	return &GetDeclareCaptureNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get declare capture not found response
func (o *GetDeclareCaptureNotFound) WithConfigurationVersion(configurationVersion string) *GetDeclareCaptureNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get declare capture not found response
func (o *GetDeclareCaptureNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get declare capture not found response
func (o *GetDeclareCaptureNotFound) WithPayload(payload *models.Error) *GetDeclareCaptureNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get declare capture not found response
func (o *GetDeclareCaptureNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeclareCaptureNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetDeclareCaptureDefault General Error

swagger:response getDeclareCaptureDefault
*/
type GetDeclareCaptureDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDeclareCaptureDefault creates GetDeclareCaptureDefault with default headers values
func NewGetDeclareCaptureDefault(code int) *GetDeclareCaptureDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDeclareCaptureDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get declare capture default response
func (o *GetDeclareCaptureDefault) WithStatusCode(code int) *GetDeclareCaptureDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get declare capture default response
func (o *GetDeclareCaptureDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get declare capture default response
func (o *GetDeclareCaptureDefault) WithConfigurationVersion(configurationVersion string) *GetDeclareCaptureDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get declare capture default response
func (o *GetDeclareCaptureDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get declare capture default response
func (o *GetDeclareCaptureDefault) WithPayload(payload *models.Error) *GetDeclareCaptureDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get declare capture default response
func (o *GetDeclareCaptureDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeclareCaptureDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
