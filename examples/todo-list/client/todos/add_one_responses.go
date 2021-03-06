package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/examples/todo-list/models"
)

type AddOneReader struct {
	formats strfmt.Registry
}

func (o *AddOneReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddOneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddOneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewAddOneCreated creates a AddOneCreated with default headers values
func NewAddOneCreated() *AddOneCreated {
	return &AddOneCreated{}
}

/*AddOneCreated

Created
*/
type AddOneCreated struct {
	Payload *models.Item
}

func (o *AddOneCreated) Error() string {
	return fmt.Sprintf("[POST /][%d] addOneCreated  %+v", 201, o.Payload)
}

func (o *AddOneCreated) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOneDefault creates a AddOneDefault with default headers values
func NewAddOneDefault(code int) *AddOneDefault {
	return &AddOneDefault{
		_statusCode: code,
	}
}

/*AddOneDefault

error
*/
type AddOneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add one default response
func (o *AddOneDefault) Code() int {
	return o._statusCode
}

func (o *AddOneDefault) Error() string {
	return fmt.Sprintf("[POST /][%d] addOne default  %+v", o._statusCode, o.Payload)
}

func (o *AddOneDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
