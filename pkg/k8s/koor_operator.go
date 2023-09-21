package k8s

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	kopv1 "github.com/koor-tech/koor-operator/api/v1alpha1"
	"google.golang.org/protobuf/types/known/wrapperspb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/koor-tech/data-control-center/gen/go/api/resources/stats"
)

func (k *K8s) GetKoorCluster(ctx context.Context, namespace string) (*stats.KoorCluster, error) {
	koorClusters, err := ListCustomResource[kopv1.KoorCluster](ctx, k.client, kopv1.GroupVersion.WithResource("koorclusters"), namespace, metav1.ListOptions{})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if len(koorClusters) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("koor cluster Not found"))
	}

	return convertKoorCluster(koorClusters[0]), nil
}

func convertKoorCluster(koorCluster *kopv1.KoorCluster) *stats.KoorCluster {
	res := &stats.KoorCluster{
		Name:      koorCluster.Name,
		Namespace: koorCluster.Namespace,
		Spec: &stats.KoorClusterSpec{
			UseAllDevices:         ConvertBoolValue(koorCluster.Spec.UseAllDevices),
			MonitoringEnabled:     ConvertBoolValue(koorCluster.Spec.MonitoringEnabled),
			DashboardEnabled:      ConvertBoolValue(koorCluster.Spec.DashboardEnabled),
			ToolboxEnabled:        ConvertBoolValue(koorCluster.Spec.ToolboxEnabled),
			UpgradeOptions:        convertUpgradeOptions(&koorCluster.Spec.UpgradeOptions),
			KsdReleaseName:        koorCluster.Spec.KsdReleaseName,
			KsdClusterReleaseName: koorCluster.Spec.KsdClusterReleaseName,
		},
		Status: &stats.KoorClusterStatus{
			TotalResources: &stats.ClusterResources{
				Nodes:   koorCluster.Status.TotalResources.Nodes.String(),
				Storage: koorCluster.Status.TotalResources.Storage.String(),
				Cpu:     koorCluster.Status.TotalResources.Cpu.String(),
				Memory:  koorCluster.Status.TotalResources.Memory.String(),
			},
			MeetsMinimumResources: koorCluster.Status.MeetsMinimumResources,
			CurrentVersions: &stats.ProductVersions{
				Kube:         koorCluster.Status.CurrentVersions.Kube,
				KoorOperator: koorCluster.Status.CurrentVersions.KoorOperator,
				Ksd:          koorCluster.Status.CurrentVersions.Ksd,
				Ceph:         koorCluster.Status.CurrentVersions.Ceph,
			},
		},
	}

    if koorCluster.Status.LatestVersions != nil {
        res.Status.LatestVersions = &stats.DetailedProductVersions{
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

func convertUpgradeOptions(upgradeOptions *kopv1.UpgradeOptions) *stats.UpgradeOptions {
	if upgradeOptions == nil {
		return nil
	}

	res := &stats.UpgradeOptions{
		Endpoint: upgradeOptions.Endpoint,
		Schedule: upgradeOptions.Schedule,
	}

	switch upgradeOptions.Mode {
	case kopv1.UpgradeModeDisabled:
		res.Mode = stats.UpgradeMode_UPGRADE_MODE_DISABLED
	case kopv1.UpgradeModeNotify:
		res.Mode = stats.UpgradeMode_UPGRADE_MODE_NOTIFY
	case kopv1.UpgradeModeUpgrade:
		res.Mode = stats.UpgradeMode_UPGRADE_MODE_UPGRADE
	default:
		res.Mode = stats.UpgradeMode_UPGRADE_MODE_UNSPECIFIED
	}

	return res
}

func convertDetailedVersion(v *kopv1.DetailedVersion) *stats.DetailedVersion {
	if v == nil {
		return nil
	}

	return &stats.DetailedVersion{
		Version:        v.Version,
		ImageUri:       v.ImageUri,
		ImageHash:      v.ImageHash,
		HelmRepository: v.HelmRepository,
		HelmChart:      v.HelmChart,
	}
}
