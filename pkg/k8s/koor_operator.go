package k8s

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	kopv1 "github.com/koor-tech/koor-operator/api/v1alpha1"
	"google.golang.org/protobuf/types/known/wrapperspb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	koorv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/koor/v1"
)

func (k *K8s) GetKoorCluster(ctx context.Context, namespace string) (*koorv1.KoorCluster, error) {
	koorClusters, err := ListCustomResource[kopv1.KoorCluster](ctx, k.client, kopv1.GroupVersion.WithResource("koorclusters"), namespace, metav1.ListOptions{})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if len(koorClusters) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("koor cluster Not found"))
	}

	return convertKoorCluster(koorClusters[0]), nil
}

func convertKoorCluster(koorCluster *kopv1.KoorCluster) *koorv1.KoorCluster {
	res := &koorv1.KoorCluster{
		Name:      koorCluster.Name,
		Namespace: koorCluster.Namespace,
		Spec: &koorv1.KoorClusterSpec{
			UseAllDevices:         ConvertBoolValue(koorCluster.Spec.UseAllDevices),
			MonitoringEnabled:     ConvertBoolValue(koorCluster.Spec.MonitoringEnabled),
			DashboardEnabled:      ConvertBoolValue(koorCluster.Spec.DashboardEnabled),
			ToolboxEnabled:        ConvertBoolValue(koorCluster.Spec.ToolboxEnabled),
			UpgradeOptions:        convertUpgradeOptions(&koorCluster.Spec.UpgradeOptions),
			KsdReleaseName:        koorCluster.Spec.KsdReleaseName,
			KsdClusterReleaseName: koorCluster.Spec.KsdClusterReleaseName,
		},
		Status: &koorv1.KoorClusterStatus{
			TotalResources: &koorv1.ClusterResources{
				Nodes:   koorCluster.Status.TotalResources.Nodes.String(),
				Storage: koorCluster.Status.TotalResources.Storage.String(),
				Cpu:     koorCluster.Status.TotalResources.Cpu.String(),
				Memory:  koorCluster.Status.TotalResources.Memory.String(),
			},
			MeetsMinimumResources: koorCluster.Status.MeetsMinimumResources,
			CurrentVersions: &koorv1.ProductVersions{
				Kube:         koorCluster.Status.CurrentVersions.Kube,
				KoorOperator: koorCluster.Status.CurrentVersions.KoorOperator,
				Ksd:          koorCluster.Status.CurrentVersions.Ksd,
				Ceph:         koorCluster.Status.CurrentVersions.Ceph,
			},
		},
	}

	if koorCluster.Status.LatestVersions != nil {
		res.Status.LatestVersions = &koorv1.DetailedProductVersions{
			KoorOperator: convertDetailedVersion(koorCluster.Status.LatestVersions.KoorOperator),
			Ksd:          convertDetailedVersion(koorCluster.Status.LatestVersions.Ksd),
			Ceph:         convertDetailedVersion(koorCluster.Status.LatestVersions.Ceph),
		}
	}
	return res
}

func ConvertBoolValue(val *bool) *wrapperspb.BoolValue {
	if val != nil {
		return &wrapperspb.BoolValue{
			Value: *val,
		}
	}
	return nil
}

func convertUpgradeOptions(upgradeOptions *kopv1.UpgradeOptions) *koorv1.UpgradeOptions {
	if upgradeOptions == nil {
		return nil
	}

	res := &koorv1.UpgradeOptions{
		Endpoint: upgradeOptions.Endpoint,
		Schedule: upgradeOptions.Schedule,
	}

	switch upgradeOptions.Mode {
	case kopv1.UpgradeModeDisabled:
		res.Mode = koorv1.UpgradeMode_UPGRADE_MODE_DISABLED
	case kopv1.UpgradeModeNotify:
		res.Mode = koorv1.UpgradeMode_UPGRADE_MODE_NOTIFY
	case kopv1.UpgradeModeUpgrade:
		res.Mode = koorv1.UpgradeMode_UPGRADE_MODE_UPGRADE
	default:
		res.Mode = koorv1.UpgradeMode_UPGRADE_MODE_UNSPECIFIED
	}

	return res
}

func convertDetailedVersion(v *kopv1.DetailedVersion) *koorv1.DetailedVersion {
	if v == nil {
		return nil
	}

	return &koorv1.DetailedVersion{
		Version:        v.Version,
		ImageUri:       v.ImageUri,
		ImageHash:      v.ImageHash,
		HelmRepository: v.HelmRepository,
		HelmChart:      v.HelmChart,
	}
}
