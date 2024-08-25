// mautrix-gvoice - A Matrix-Google Voice puppeting bridge.
// Copyright (C) 2024 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package connector

import (
	"context"

	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/event"
)

var generalCaps = &bridgev2.NetworkGeneralCapabilities{
	DisappearingMessages: false,
	AggressiveUpdateInfo: false,
}

func (gv *GVConnector) GetCapabilities() *bridgev2.NetworkGeneralCapabilities {
	return generalCaps
}

var roomCaps = &bridgev2.NetworkRoomCapabilities{
	Captions:         true,
	MaxTextLength:    160,
	MaxCaptionLength: 0,
	DefaultFileRestriction: &bridgev2.FileRestriction{
		MaxSize:   0,
		MimeTypes: []string{},
	},
	Files: map[event.MessageType]bridgev2.FileRestriction{
		event.MsgImage: {
			MaxSize:   2 * 1024 * 1024,
			MimeTypes: []string{"image/png", "image/jpeg", "image/bmp", "image/tiff"},
		},
	},
}

func (gc *GVClient) GetCapabilities(ctx context.Context, portal *bridgev2.Portal) *bridgev2.NetworkRoomCapabilities {
	return roomCaps
}
