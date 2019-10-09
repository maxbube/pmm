// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddRemoteNodeReader is a Reader for the AddRemoteNode structure.
type AddRemoteNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRemoteNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRemoteNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddRemoteNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRemoteNodeOK creates a AddRemoteNodeOK with default headers values
func NewAddRemoteNodeOK() *AddRemoteNodeOK {
	return &AddRemoteNodeOK{}
}

/*AddRemoteNodeOK handles this case with default header values.

A successful response.
*/
type AddRemoteNodeOK struct {
	Payload *AddRemoteNodeOKBody
}

func (o *AddRemoteNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddRemote][%d] addRemoteNodeOk  %+v", 200, o.Payload)
}

func (o *AddRemoteNodeOK) GetPayload() *AddRemoteNodeOKBody {
	return o.Payload
}

func (o *AddRemoteNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRemoteNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRemoteNodeDefault creates a AddRemoteNodeDefault with default headers values
func NewAddRemoteNodeDefault(code int) *AddRemoteNodeDefault {
	return &AddRemoteNodeDefault{
		_statusCode: code,
	}
}

/*AddRemoteNodeDefault handles this case with default header values.

An error response.
*/
type AddRemoteNodeDefault struct {
	_statusCode int

	Payload *AddRemoteNodeDefaultBody
}

// Code gets the status code for the add remote node default response
func (o *AddRemoteNodeDefault) Code() int {
	return o._statusCode
}

func (o *AddRemoteNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddRemote][%d] AddRemoteNode default  %+v", o._statusCode, o.Payload)
}

func (o *AddRemoteNodeDefault) GetPayload() *AddRemoteNodeDefaultBody {
	return o.Payload
}

func (o *AddRemoteNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRemoteNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddRemoteNodeBody add remote node body
swagger:model AddRemoteNodeBody
*/
type AddRemoteNodeBody struct {

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add remote node body
func (o *AddRemoteNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteNodeBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteNodeDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddRemoteNodeDefaultBody
*/
type AddRemoteNodeDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add remote node default body
func (o *AddRemoteNodeDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteNodeOKBody add remote node OK body
swagger:model AddRemoteNodeOKBody
*/
type AddRemoteNodeOKBody struct {

	// remote
	Remote *AddRemoteNodeOKBodyRemote `json:"remote,omitempty"`
}

// Validate validates this add remote node OK body
func (o *AddRemoteNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRemoteNodeOKBody) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	if o.Remote != nil {
		if err := o.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRemoteNodeOk" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteNodeOKBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteNodeOKBodyRemote RemoteNode represents generic remote Node. Agents can't run on Remote Nodes.
swagger:model AddRemoteNodeOKBodyRemote
*/
type AddRemoteNodeOKBodyRemote struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add remote node OK body remote
func (o *AddRemoteNodeOKBodyRemote) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteNodeOKBodyRemote) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteNodeOKBodyRemote) UnmarshalBinary(b []byte) error {
	var res AddRemoteNodeOKBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
