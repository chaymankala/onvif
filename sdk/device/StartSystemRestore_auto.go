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

// Call_StartSystemRestore forwards the call to dev.CallMethod() then parses the payload of the reply as a StartSystemRestoreResponse.
func Call_StartSystemRestore(ctx context.Context, dev *onvif.Device, request device.StartSystemRestore) (device.StartSystemRestoreResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartSystemRestoreResponse device.StartSystemRestoreResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "StartSystemRestore")
		return reply.Body.StartSystemRestoreResponse, errors.Annotate(err, "reply")
	}
}
