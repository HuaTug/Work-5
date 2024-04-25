package sentinels

import (
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "Default",
			Threshold:              10,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       1000,
		},
	})
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}
}
