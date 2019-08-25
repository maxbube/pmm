// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// StartPostgreSQLShowCreateTableActionReader is a Reader for the StartPostgreSQLShowCreateTableAction structure.
type StartPostgreSQLShowCreateTableActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPostgreSQLShowCreateTableActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPostgreSQLShowCreateTableActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartPostgreSQLShowCreateTableActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPostgreSQLShowCreateTableActionOK creates a StartPostgreSQLShowCreateTableActionOK with default headers values
func NewStartPostgreSQLShowCreateTableActionOK() *StartPostgreSQLShowCreateTableActionOK {
	return &StartPostgreSQLShowCreateTableActionOK{}
}

/*StartPostgreSQLShowCreateTableActionOK handles this case with default header values.

A successful response.
*/
type StartPostgreSQLShowCreateTableActionOK struct {
	Payload *StartPostgreSQLShowCreateTableActionOKBody
}

func (o *StartPostgreSQLShowCreateTableActionOK) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartPostgreSQLShowCreateTable][%d] startPostgreSqlShowCreateTableActionOk  %+v", 200, o.Payload)
}

func (o *StartPostgreSQLShowCreateTableActionOK) GetPayload() *StartPostgreSQLShowCreateTableActionOKBody {
	return o.Payload
}

func (o *StartPostgreSQLShowCreateTableActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPostgreSQLShowCreateTableActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPostgreSQLShowCreateTableActionDefault creates a StartPostgreSQLShowCreateTableActionDefault with default headers values
func NewStartPostgreSQLShowCreateTableActionDefault(code int) *StartPostgreSQLShowCreateTableActionDefault {
	return &StartPostgreSQLShowCreateTableActionDefault{
		_statusCode: code,
	}
}

/*StartPostgreSQLShowCreateTableActionDefault handles this case with default header values.

An error response.
*/
type StartPostgreSQLShowCreateTableActionDefault struct {
	_statusCode int

	Payload *StartPostgreSQLShowCreateTableActionDefaultBody
}

// Code gets the status code for the start postgre SQL show create table action default response
func (o *StartPostgreSQLShowCreateTableActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPostgreSQLShowCreateTableActionDefault) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartPostgreSQLShowCreateTable][%d] StartPostgreSQLShowCreateTableAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartPostgreSQLShowCreateTableActionDefault) GetPayload() *StartPostgreSQLShowCreateTableActionDefaultBody {
	return o.Payload
}

func (o *StartPostgreSQLShowCreateTableActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPostgreSQLShowCreateTableActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartPostgreSQLShowCreateTableActionBody start postgre SQL show create table action body
swagger:model StartPostgreSQLShowCreateTableActionBody
*/
type StartPostgreSQLShowCreateTableActionBody struct {

	// Database name. Required if not given in the table_name field.
	Database string `json:"database,omitempty"`

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Table name. Required. May additionally contain a database name.
	TableName string `json:"table_name,omitempty"`
}

// Validate validates this start postgre SQL show create table action body
func (o *StartPostgreSQLShowCreateTableActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowCreateTableActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPostgreSQLShowCreateTableActionDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model StartPostgreSQLShowCreateTableActionDefaultBody
*/
type StartPostgreSQLShowCreateTableActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this start postgre SQL show create table action default body
func (o *StartPostgreSQLShowCreateTableActionDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowCreateTableActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPostgreSQLShowCreateTableActionOKBody start postgre SQL show create table action OK body
swagger:model StartPostgreSQLShowCreateTableActionOKBody
*/
type StartPostgreSQLShowCreateTableActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start postgre SQL show create table action OK body
func (o *StartPostgreSQLShowCreateTableActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPostgreSQLShowCreateTableActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPostgreSQLShowCreateTableActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
