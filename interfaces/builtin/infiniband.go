// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2026 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

const infinibandSummary = `allows access to RDMA/InfiniBand devices`

const infinibandBaseDeclarationSlots = `
  infiniband:
    allow-installation:
      slot-snap-type:
        - core
    deny-auto-connection: true
`

// Allow access to RDMA/InfiniBand devices
const infinibandConnectedPlugAppArmor = `
/dev/infiniband/* rw,
`

var infinibandConnectedPlugUDev = []string{
	`SUBSYSTEM=="infiniband_verbs"`,
	`SUBSYSTEM=="infiniband_mad"`,
	`SUBSYSTEM=="rdma_cm"`,
}

func init() {
	registerIface(&commonInterface{
		name:                  "infiniband",
		summary:               infinibandSummary,
		implicitOnCore:        true,
		implicitOnClassic:     true,
		baseDeclarationSlots:  infinibandBaseDeclarationSlots,
		connectedPlugAppArmor: infinibandConnectedPlugAppArmor,
		connectedPlugUDev:     infinibandConnectedPlugUDev,
	})
}
