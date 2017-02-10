package dhcp_subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// POSTDhcpSubnetReader is a Reader for the POSTDhcpSubnet structure.
type POSTDhcpSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTDhcpSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPOSTDhcpSubnetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTDhcpSubnetCreated creates a POSTDhcpSubnetCreated with default headers values
func NewPOSTDhcpSubnetCreated() *POSTDhcpSubnetCreated {
	return &POSTDhcpSubnetCreated{}
}

/*POSTDhcpSubnetCreated handles this case with default header values.

POSTDhcpSubnetCreated p o s t dhcp subnet created
*/
type POSTDhcpSubnetCreated struct {
	Payload *models.DhcpSubnetInput
}

func (o *POSTDhcpSubnetCreated) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] pOSTDhcpSubnetCreated  %+v", 201, o.Payload)
}

func (o *POSTDhcpSubnetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DhcpSubnetInput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
