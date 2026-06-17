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

const infinibandConnectedPlugAppArmor = `
# Description: Allow access to RDMA/InfiniBand devices and their associated
# sysfs attributes. This gives access to high-speed network fabric devices.

# InfiniBand userspace verbs (/dev/infiniband/uverbs*) and MAD devices
# (/dev/infiniband/umad*), and the RDMA CM device (/dev/infiniband/rdma_cm).
/dev/infiniband/* rw,

# Some userspace libraries access the same device via /dev/char/<major>:<minor>.
# InfiniBand uverbs devices use major 231.
/run/udev/data/c231:* r,

# Discovery and enumeration of IB HCAs via sysfs
/sys/class/infiniband/ r,
/sys/class/infiniband/** r,

# Userspace verbs and MAD device sysfs nodes
/sys/class/infiniband_verbs/ r,
/sys/class/infiniband_verbs/** r,
/sys/class/infiniband_mad/ r,
/sys/class/infiniband_mad/** r,

# PCI device sysfs – needed to resolve HCA port to PCI function
/sys/bus/pci/devices/**/infiniband/ r,
/sys/bus/pci/devices/**/infiniband/** r,

# Network interface attributes used to correlate IB ports with net devices
# (e.g. to map ens16f0np0 → mlx5_0 port 1 for RDMA CM path selection)
/sys/class/net/*/phys_port_name r,
/sys/class/net/*/phys_switch_id r,

# CPU topology is read by RDMA libraries to set affinity
/sys/devices/system/cpu/possible r,
/sys/devices/system/cpu/online r,

# IOMMU groups are inspected by RDMA libraries for memory registration
/sys/kernel/iommu_groups/{,**} r,

# Temporary storage used by RDMA applications (e.g. for memory-mapped buffers)
/var/tmp/** rw,
`

var infinibandConnectedPlugUDev = []string{
	// uverbs devices: /dev/infiniband/uverbs*
	`SUBSYSTEM=="infiniband_verbs"`,
	// umad devices: /dev/infiniband/umad*
	`SUBSYSTEM=="infiniband_mad"`,
	// RDMA CM device: /dev/infiniband/rdma_cm
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
