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

// Call_UpgradeSystemFirmware forwards the call to dev.CallMethod() then parses the payload of the reply as a UpgradeSystemFirmwareResponse.
func Call_UpgradeSystemFirmware(ctx context.Context, dev *onvif.Device, request device.UpgradeSystemFirmware) (device.UpgradeSystemFirmwareResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UpgradeSystemFirmwareResponse device.UpgradeSystemFirmwareResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "UpgradeSystemFirmware")
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Annotate(err, "reply")
	}
}