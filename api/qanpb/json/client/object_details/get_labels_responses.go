// Code generated by go-swagger; DO NOT EDIT.

package object_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLabelsReader is a Reader for the GetLabels structure.
type GetLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLabelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLabelsOK creates a GetLabelsOK with default headers values
func NewGetLabelsOK() *GetLabelsOK {
	return &GetLabelsOK{}
}

/*GetLabelsOK handles this case with default header values.

A successful response.
*/
type GetLabelsOK struct {
	Payload *GetLabelsOKBody
}

func (o *GetLabelsOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetLabels][%d] getLabelsOk  %+v", 200, o.Payload)
}

func (o *GetLabelsOK) GetPayload() *GetLabelsOKBody {
	return o.Payload
}

func (o *GetLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLabelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsDefault creates a GetLabelsDefault with default headers values
func NewGetLabelsDefault(code int) *GetLabelsDefault {
	return &GetLabelsDefault{
		_statusCode: code,
	}
}

/*GetLabelsDefault handles this case with default header values.

An error response.
*/
type GetLabelsDefault struct {
	_statusCode int

	Payload *GetLabelsDefaultBody
}

// Code gets the status code for the get labels default response
func (o *GetLabelsDefault) Code() int {
	return o._statusCode
}

func (o *GetLabelsDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetLabels][%d] GetLabels default  %+v", o._statusCode, o.Payload)
}

func (o *GetLabelsDefault) GetPayload() *GetLabelsDefaultBody {
	return o.Payload
}

func (o *GetLabelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLabelsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLabelsBody ObjectDetailsLabelsRequest defines filtering of object detail's labels for specific value of
// dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
swagger:model GetLabelsBody
*/
type GetLabelsBody struct {

	// period start from
	// Format: date-time
	PeriodStartFrom strfmt.DateTime `json:"period_start_from,omitempty"`

	// period start to
	// Format: date-time
	PeriodStartTo strfmt.DateTime `json:"period_start_to,omitempty"`

	// dimension value: ex: queryid - 1D410B4BE5060972.
	FilterBy string `json:"filter_by,omitempty"`

	// one of dimension: queryid | host ...
	GroupBy string `json:"group_by,omitempty"`
}

// Validate validates this get labels body
func (o *GetLabelsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePeriodStartFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeriodStartTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLabelsBody) validatePeriodStartFrom(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_from", "body", "date-time", o.PeriodStartFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetLabelsBody) validatePeriodStartTo(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartTo) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_to", "body", "date-time", o.PeriodStartTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsBody) UnmarshalBinary(b []byte) error {
	var res GetLabelsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetLabelsDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model GetLabelsDefaultBody
*/
type GetLabelsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get labels default body
func (o *GetLabelsDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetLabelsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetLabelsOKBody ObjectDetailsLabelsReply is a map of labels names as keys and labels values as a list.
swagger:model GetLabelsOKBody
*/
type GetLabelsOKBody struct {

	// labels
	Labels map[string]LabelsAnon `json:"labels,omitempty"`
}

// Validate validates this get labels OK body
func (o *GetLabelsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLabelsOKBody) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for k := range o.Labels {

		if swag.IsZero(o.Labels[k]) { // not required
			continue
		}
		if val, ok := o.Labels[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLabelsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLabelsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLabelsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LabelsAnon ListLabelValues is list of label's values.
swagger:model LabelsAnon
*/
type LabelsAnon struct {

	// values
	Values []string `json:"values"`
}

// Validate validates this labels anon
func (o *LabelsAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LabelsAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LabelsAnon) UnmarshalBinary(b []byte) error {
	var res LabelsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
