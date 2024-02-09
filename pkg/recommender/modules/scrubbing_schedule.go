package recommendermodules

import (
	"context"
	"slices"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var scrubbingScheduleMinimumConfigOptions = []string{
	"osd_scrub_begin_hour",
	"osd_scrub_end_hour",
	"osd_scrub_begin_week_day",
	"osd_scrub_end_week_day",
}

func init() {
	moduleFactories["scrubbing_schedule"] = NewScrubbingSchedule
}

func NewScrubbingSchedule() (Module, error) {
	return &ScrubbingSchedule{}, nil
}

type ScrubbingSchedule struct {
	Module
}

func (m *ScrubbingSchedule) GetName() string {
	return "scrubbing_schedule"
}

func (m *ScrubbingSchedule) Run(ctx context.Context, p *Params) ([]*cephv1.ClusterRecommendation, error) {
	recs := []*cephv1.ClusterRecommendation{}

	clusters, err := p.K8S.GetRookClient().CephV1().CephClusters(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range clusters.Items {
		if cfg, ok := item.Spec.CephConfig["global"]; ok {
			// Collect the set global config options so we can check if a scrubbing schedule is set
			configOpts := []string{}
			for opt := range cfg {
				configOpts = append(configOpts, opt)
			}

			found := true
			for _, opt := range scrubbingScheduleMinimumConfigOptions {
				if !slices.Contains(configOpts, opt) {
					found = false
					break
				}
			}

			if found {
				break
			}
		}

		recs = append(recs, &cephv1.ClusterRecommendation{
			Title:       "No Custom Scrubbing Schedule set",
			Description: "It is recommended to set up a custom OSD scrubbing schedule to ensure the cluster will scrub during off hours.",
			Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_INFORMAL,
			Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_OSD,
		})
	}

	return recs, nil
}
