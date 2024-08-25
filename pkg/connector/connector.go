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
)

type GVConnector struct {
	Bridge *bridgev2.Bridge
	Config Config
}

var _ bridgev2.NetworkConnector = (*GVConnector)(nil)

func (gv *GVConnector) Init(bridge *bridgev2.Bridge) {
	gv.Bridge = bridge
}

func (gv *GVConnector) Start(ctx context.Context) error {
	return nil
}

func (gv *GVConnector) GetName() bridgev2.BridgeName {
	return bridgev2.BridgeName{
		DisplayName:          "Google Voice",
		NetworkURL:           "https://voice.google.com",
		NetworkIcon:          "mxc://maunium.net/VOPtYGBzHLRfPTEzGgNMpeKo",
		NetworkID:            "gvoice",
		BeeperBridgeType:     "gvoice",
		DefaultPort:          29338,
		DefaultCommandPrefix: "!gv",
	}
}
