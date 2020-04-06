/*
Package jobs is for all scheduled jobs (tasks) to run
*/
package jobs

import (
	"fmt"

	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-logger"
)

// exampleJob is an example job
func exampleJob() {

	logger.Data(2, logger.DEBUG, "starting job...")

	// Do something

	// Do something else

	logger.Data(2, logger.DEBUG, "job complete!")
}

// RunExampleJob will run the job every X minutes
func RunExampleJob(runNow bool, andEveryXMinutes int) {
	if runNow {
		exampleJob()
	}
	_, err := config.Values.Scheduler.AddJob("example-job", fmt.Sprintf("@every %dm", andEveryXMinutes), exampleJob)
	if err != nil {
		logger.Data(2, logger.ERROR, "error adding job: RunExampleJob: "+err.Error())
	}
}
