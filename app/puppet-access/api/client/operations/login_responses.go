// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/puppetlabs/pe-sdk-go/app/puppet-access/api/models"
)

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/* LoginOK describes a response with status code 200, with default header values.

Login to generate a token in RBAC
*/
type LoginOK struct {
	Payload *LoginOKBody
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[POST /auth/token][%d] loginOK  %+v", 200, o.Payload)
}
func (o *LoginOK) GetPayload() *LoginOKBody {
	return o.Payload
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LoginOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginDefault creates a LoginDefault with default headers values
func NewLoginDefault(code int) *LoginDefault {
	return &LoginDefault{
		_statusCode: code,
	}
}

/* LoginDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type LoginDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the login default response
func (o *LoginDefault) Code() int {
	return o._statusCode
}

func (o *LoginDefault) Error() string {
	return fmt.Sprintf("[POST /auth/token][%d] login default  %+v", o._statusCode, o.Payload)
}
func (o *LoginDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *LoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LoginOKBody puppet access login
swagger:model LoginOKBody
*/
type LoginOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this login o k body
func (o *LoginOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this login o k body based on context it is used
func (o *LoginOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBody) UnmarshalBinary(b []byte) error {
	var res LoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
