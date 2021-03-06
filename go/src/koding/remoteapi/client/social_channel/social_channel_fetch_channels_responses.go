package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SocialChannelFetchChannelsReader is a Reader for the SocialChannelFetchChannels structure.
type SocialChannelFetchChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialChannelFetchChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialChannelFetchChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialChannelFetchChannelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialChannelFetchChannelsOK creates a SocialChannelFetchChannelsOK with default headers values
func NewSocialChannelFetchChannelsOK() *SocialChannelFetchChannelsOK {
	return &SocialChannelFetchChannelsOK{}
}

/*SocialChannelFetchChannelsOK handles this case with default header values.

Request processed successfully
*/
type SocialChannelFetchChannelsOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialChannelFetchChannelsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.fetchChannels][%d] socialChannelFetchChannelsOK  %+v", 200, o.Payload)
}

func (o *SocialChannelFetchChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialChannelFetchChannelsUnauthorized creates a SocialChannelFetchChannelsUnauthorized with default headers values
func NewSocialChannelFetchChannelsUnauthorized() *SocialChannelFetchChannelsUnauthorized {
	return &SocialChannelFetchChannelsUnauthorized{}
}

/*SocialChannelFetchChannelsUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialChannelFetchChannelsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialChannelFetchChannelsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.fetchChannels][%d] socialChannelFetchChannelsUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialChannelFetchChannelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
