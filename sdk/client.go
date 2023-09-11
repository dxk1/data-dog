/**
 * @Package sdk
 * @Author: tbb
 * @Date: 2023/9/9 15:07
 */
package sdk

import (
	"context"
	"data-dog/log"
	"data-dog/pb"
	"data-dog/sdk/api"
	"fmt"
	"git.code.oa.com/tencentcloud-serverless/scf_common/polarissdk"
	"time"
)

type ReportClient struct {
	ServiceConfig *ServiceConfig
	Stop          chan struct{}

	LocalIp string
}

func Init(config *ServiceConfig, localIp string) *ReportClient {
	return &ReportClient{
		ServiceConfig: config,
		LocalIp:       localIp,
	}
}

func (r *ReportClient) ReportInvokeMetric(ctx context.Context, funId, ty string, files map[string]string) error {
	if api.CountingPass(fmt.Sprintf("%s:%s", funId, ty), r.ServiceConfig.SwitchConfig) {
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
	}
	return nil
}

func (r *ReportClient) ReportLocalMetric(ctx context.Context, metric string) error {
	initReq := &polarissdk.GetOneInstanceRequest{}
	initReq.Namespace = r.ServiceConfig.Namespace
	initReq.Service = r.ServiceConfig.Service
	conn, err := polarissdk.GetGrpcConn(initReq)
	if err != nil {
		return err
	}

	c := pb.NewFuncMetricReportClient(conn.Conn.Value())
	_, err = c.ReportLocalMetric(ctx, &pb.LocalMetricRequest{
		LocalIp:          r.LocalIp,
		SecondOverMetric: metric,
	})
	return err
}

func (r *ReportClient) ReportLocalMetricTick() {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ticker.C:
			ctx, _ := context.WithTimeout(context.Background(), time.Duration(200*time.Millisecond))
			err := r.ReportLocalMetric(ctx, api.ReportOver())
			if err != nil {
				log.Log.Error().Msgf("ReportLocalMetricTick ReportLocalMetric Err:%s", err)
			}
		case <-r.Stop:
			ticker.Stop()
			return
		}
	}
}

// ServiceConfig ...
type ServiceConfig struct {
	Namespace    string            `yaml:"namespace" json:"namespace"`
	Service      string            `yaml:"service" json:"service"`
	SwitchConfig *api.SwitchConfig `yaml:"switch_config" json:"switch_config"`
}
