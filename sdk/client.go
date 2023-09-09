/**
 * @Package sdk
 * @Author: tbb
 * @Date: 2023/9/9 15:07
 */
package sdk

import (
	"context"
	"data-dog/pb"
	"git.code.oa.com/tencentcloud-serverless/scf_common/polarissdk"
)

type ReportClient struct {
	ServiceConfig *ServiceConfig
}

func Init(config *ServiceConfig) *ReportClient {
	return &ReportClient{
		ServiceConfig: config,
	}
}

func (r *ReportClient) ReportInvokeMetric(ctx context.Context, files map[string]string) error {
	initReq := &polarissdk.GetOneInstanceRequest{}
	initReq.Namespace = r.ServiceConfig.Namespace
	initReq.Service = r.ServiceConfig.Service
	conn, err := polarissdk.GetGrpcConn(initReq)
	if err != nil {
		return err
	}

	c := pb.NewFuncMetricReportClient(conn.Conn.Value())
	_, err = c.InvokeMetricReport(ctx, &pb.InvokeMetricRequest{
		CustomFields: files,
	})
	if err != nil {
		return err
	}

	return nil
}

// ServiceConfig ...
type ServiceConfig struct {
	Namespace    string        `yaml:"namespace" json:"namespace"`
	Service      string        `yaml:"service" json:"service"`
	SwitchConfig *SwitchConfig `yaml:"switch_config" json:"switch_config"`
}

// SwitchConfig force pass limit switch config
type SwitchConfig struct {
	ForceIgnore         bool  `yaml:"force_ignore" json:"force_ignore"`
	OverSecondQpsIgnore int64 `yaml:"over_second_qps_ignore" json:"over_second_qps_ignore"`
	OverMinuteQpsIgnore int64 `yaml:"over_minute_qps_ignore" json:"over_minute_qps_ignore"`
}
