// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/chaymankala/onvif"
	"github.com/chaymankala/onvif/sdk"
	"github.com/chaymankala/onvif/device"
)

// Call_GetIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a GetIPAddressFilterResponse.
func Call_GetIPAddressFilter(ctx context.Context, dev *onvif.Device, request device.GetIPAddressFilter) (device.GetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetIPAddressFilterResponse device.GetIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetIPAddressFilterResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetIPAddressFilter")
		return reply.Body.GetIPAddressFilterResponse, errors.Annotate(err, "reply")
	}
}
