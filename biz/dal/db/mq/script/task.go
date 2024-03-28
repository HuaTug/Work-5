package script //nolint:gofmt

import (
	"context" //nolint:gofmt
	"log"

	"Hertz_refactored/biz/dal/db/mq/task" //nolint:goimports
)

func TaskCreateSync(ctx context.Context) {
	Sync := new(task.SyncTask)
	err := Sync.RunTaskCreate(ctx)
	if err != nil {
		log.Printf("RunTaskCreate: %s", err)

	}
}
func LoadingScript() {
	ctx := context.Background()
	go TaskCreateSync(ctx)
}
