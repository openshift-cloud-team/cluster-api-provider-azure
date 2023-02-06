/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	infrav1exp "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this AzureManagedControlPlane to the Hub version (v1beta1).
func (src *AzureManagedControlPlane) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1exp.AzureManagedControlPlane)
	if err := Convert_v1alpha3_AzureManagedControlPlane_To_v1beta1_AzureManagedControlPlane(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1exp.AzureManagedControlPlane{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	dst.Spec.IdentityRef = restored.Spec.IdentityRef
	dst.Spec.SKU = restored.Spec.SKU
	dst.Spec.LoadBalancerProfile = restored.Spec.LoadBalancerProfile
	dst.Spec.APIServerAccessProfile = restored.Spec.APIServerAccessProfile
	dst.Spec.AddonProfiles = restored.Spec.AddonProfiles
	dst.Spec.VirtualNetwork.ResourceGroup = restored.Spec.VirtualNetwork.ResourceGroup
	dst.Spec.VirtualNetwork.Subnet.ServiceEndpoints = restored.Spec.VirtualNetwork.Subnet.ServiceEndpoints
	dst.Spec.AutoScalerProfile = restored.Spec.AutoScalerProfile

	dst.Status.LongRunningOperationStates = restored.Status.LongRunningOperationStates
	dst.Status.Conditions = restored.Status.Conditions

	return nil
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *AzureManagedControlPlane) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1exp.AzureManagedControlPlane)

	if err := Convert_v1beta1_AzureManagedControlPlane_To_v1alpha3_AzureManagedControlPlane(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	return utilconversion.MarshalData(src, dst)
}

// Convert_v1beta1_AzureManagedControlPlaneSpec_To_v1alpha3_AzureManagedControlPlaneSpec is an autogenerated conversion function.
func Convert_v1beta1_AzureManagedControlPlaneSpec_To_v1alpha3_AzureManagedControlPlaneSpec(in *infrav1exp.AzureManagedControlPlaneSpec, out *AzureManagedControlPlaneSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_AzureManagedControlPlaneSpec_To_v1alpha3_AzureManagedControlPlaneSpec(in, out, s)
}

// Convert_v1beta1_AzureManagedControlPlaneStatus_To_v1alpha3_AzureManagedControlPlaneStatus is an autogenerated conversion function.
func Convert_v1beta1_AzureManagedControlPlaneStatus_To_v1alpha3_AzureManagedControlPlaneStatus(in *infrav1exp.AzureManagedControlPlaneStatus, out *AzureManagedControlPlaneStatus, s apiconversion.Scope) error {
	return autoConvert_v1beta1_AzureManagedControlPlaneStatus_To_v1alpha3_AzureManagedControlPlaneStatus(in, out, s)
}

// ConvertTo converts this AzureManagedControlPlane to the Hub version (v1beta1).
func (src *AzureManagedControlPlaneList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1exp.AzureManagedControlPlaneList)
	return Convert_v1alpha3_AzureManagedControlPlaneList_To_v1beta1_AzureManagedControlPlaneList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *AzureManagedControlPlaneList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1exp.AzureManagedControlPlaneList)
	return Convert_v1beta1_AzureManagedControlPlaneList_To_v1alpha3_AzureManagedControlPlaneList(src, dst, nil)
}

// Convert_v1beta1_ManagedControlPlaneVirtualNetwork_To_v1alpha3_ManagedControlPlaneVirtualNetwork converts v1beta1 ManagedControlPlaneVirtualNetwork to v1alpha3 ManagedControlPlaneVirtualNetwork.
func Convert_v1beta1_ManagedControlPlaneVirtualNetwork_To_v1alpha3_ManagedControlPlaneVirtualNetwork(in *infrav1exp.ManagedControlPlaneVirtualNetwork, out *ManagedControlPlaneVirtualNetwork, s apiconversion.Scope) error {
	out.Name = in.Name
	out.Subnet.Name = in.Subnet.Name
	out.Subnet.CIDRBlock = in.Subnet.CIDRBlock
	out.CIDRBlock = in.CIDRBlock

	return nil
}
