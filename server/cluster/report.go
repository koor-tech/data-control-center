package cluster

import (
	"bytes"
	"context"
	"embed"

	"html/template"

	"connectrpc.com/connect"
	"github.com/Masterminds/sprig/v3"
	clusterpb "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
	"github.com/koor-tech/data-control-center/pkg/version"
)

var (
	//go:embed *.tpl
	reportTemplates embed.FS
)

type ReportContent struct {
	Name        string
	Description string
	Content     string
	Error       error
}

func (s *Server) GetTroubleshootReport(ctx context.Context, req *connect.Request[clusterpb.GetTroubleshootReportRequest]) (*connect.Response[clusterpb.GetTroubleshootReportResponse], error) {
	reportContent := []ReportContent{}

	kCli := s.k.GetClient()

	k8sVersion, err := kCli.ServerVersion()
	reportContent = append(reportContent, ReportContent{
		Name:    "Kubernetes Version",
		Content: k8sVersion.String(),
		Error:   err,
	})

	rookOperator, err := s.k.GetOperatorVersion(ctx, s.Namespace)
	reportContent = append(reportContent, ReportContent{
		Name:    "Operator Version",
		Content: rookOperator,
		Error:   err,
	})

	// TODO get rook ceph pods, `ceph versions` info and health status + message of custom resources

	data := struct {
		Version  string
		Contents []ReportContent
	}{
		Version:  version.Version,
		Contents: reportContent,
	}

	tpl := template.Must(
		template.New("report").Funcs(sprig.FuncMap()).ParseFS(reportTemplates, "*.tpl"),
	)

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		return nil, err
	}

	res := connect.NewResponse(&clusterpb.GetTroubleshootReportResponse{
		Report: buf.String(),
	})
	return res, nil
}
