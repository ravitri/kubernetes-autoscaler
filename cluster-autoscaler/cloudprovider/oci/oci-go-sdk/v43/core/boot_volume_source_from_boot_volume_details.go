// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"encoding/json"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/oci-go-sdk/v43/common"
)

// BootVolumeSourceFromBootVolumeDetails Specifies the source boot volume.
type BootVolumeSourceFromBootVolumeDetails struct {

	// The OCID of the boot volume.
	Id *string `mandatory:"true" json:"id"`
}

func (m BootVolumeSourceFromBootVolumeDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m BootVolumeSourceFromBootVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBootVolumeSourceFromBootVolumeDetails BootVolumeSourceFromBootVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeBootVolumeSourceFromBootVolumeDetails
	}{
		"bootVolume",
		(MarshalTypeBootVolumeSourceFromBootVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
