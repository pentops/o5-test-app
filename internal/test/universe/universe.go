package universe

import (
	"context"

	"github.com/pentops/flowtest"
	"github.com/pentops/flowtest/runner/testclient"
)

type Universe struct {
	Client         *testclient.API
	MetaClient     *testclient.API
	ReplayQueueURL string
}

type APIConfig struct {
	APIRoot        string
	MetaRoot       string
	BearerToken    string
	ReplayQueueURL string
}

func UniverseWrapper(cfg *APIConfig, callback func(flow flowtest.StepSetter, uu *Universe)) func(flow flowtest.StepSetter) {
	return func(flow flowtest.StepSetter) {

		universe := &Universe{}

		flow.Setup(func(ctx context.Context, t flowtest.Asserter) error {
			t.Log("SETUP")
			client, err := testclient.NewAPI(cfg.APIRoot)
			if err != nil {
				return err
			}

			client.Auth = testclient.BearerToken(cfg.BearerToken)

			client.Logger = LogAPIRequest(flow)
			universe.Client = client

			metaClient, err := testclient.NewAPI(cfg.MetaRoot)
			if err != nil {
				return err
			}

			metaClient.Auth = testclient.BearerToken(cfg.BearerToken)

			metaClient.Logger = LogAPIRequest(flow)
			universe.MetaClient = metaClient
			universe.ReplayQueueURL = cfg.ReplayQueueURL

			return nil
		})

		callback(flow, universe)
	}
}

func LogAPIRequest(flow flowtest.StepSetter) func(req *testclient.RequestLog) {
	return func(req *testclient.RequestLog) {
		flow.Log(req)
	}
}
