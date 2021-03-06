package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new todos API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*AddOne add one API
 */
func (a *Client) AddOne(params *AddOneParams, authInfo client.AuthInfoWriter) (*AddOneCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOneParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "addOne",
		Method:      "POST",
		PathPattern: "/",
		Schemes:     []string{"http", "https"},
		Params:      params,
		Reader:      &AddOneReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOneCreated), nil
}

/*DestroyOne destroy one API
 */
func (a *Client) DestroyOne(params *DestroyOneParams, authInfo client.AuthInfoWriter) (*DestroyOneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDestroyOneParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "destroyOne",
		Method:      "DELETE",
		PathPattern: "/{id}",
		Schemes:     []string{"http", "https"},
		Params:      params,
		Reader:      &DestroyOneReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DestroyOneNoContent), nil
}

/*Find find API
 */
func (a *Client) Find(params *FindParams, authInfo client.AuthInfoWriter) (*FindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "find",
		Method:      "GET",
		PathPattern: "/",
		Schemes:     []string{"http", "https"},
		Params:      params,
		Reader:      &FindReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindOK), nil
}

/*UpdateOne update one API
 */
func (a *Client) UpdateOne(params *UpdateOneParams, authInfo client.AuthInfoWriter) (*UpdateOneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOneParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "updateOne",
		Method:      "PUT",
		PathPattern: "/{id}",
		Schemes:     []string{"http", "https"},
		Params:      params,
		Reader:      &UpdateOneReader{formats: a.formats},
		AuthInfo:    authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOneOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
