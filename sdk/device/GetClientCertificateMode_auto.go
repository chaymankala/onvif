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

// Call_GetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetClientCertificateModeResponse.
func Call_GetClientCertificateMode(ctx context.Context, dev *onvif.Device, request device.GetClientCertificateMode) (device.GetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetClientCertificateModeResponse device.GetClientCertificateModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetClientCertificateModeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetClientCertificateMode")
		return reply.Body.GetClientCertificateModeResponse, errors.Annotate(err, "reply")
	}
}