package jobs

import (
	"context"
	"fmt"
	"log"
	"strings"
	"task-scheduler-api/db"

	"github.com/robfig/cron/v3"
)

func loadEnabledJobsFromDB(repo *db.Queries, ctx context.Context) ([]db.Task, error) {
	tasks, err := repo.GetEnabledTask(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func convertToCronExpr(freqDate, freqTime string) string {
	parts := strings.Split(freqTime, ":")
	hour := parts[0]
	minute := parts[1]

	switch freqDate {
	case "daily":
		return fmt.Sprintf("%s %s * * *", minute, hour)
	case "weekly":
		return fmt.Sprintf("%s %s * * 0", minute, hour) // Sunday
	case "monthly":
		return fmt.Sprintf("%s %s 1 * *", minute, hour) // First day of month
	default:
		return ""
	}
}

func ScheduleAllJobs(c *cron.Cron, repo *db.Queries, ctx context.Context) {
	jobs, err := loadEnabledJobsFromDB(repo, ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, job := range jobs {
		cronExpr := convertToCronExpr(job.FrequencyDate.String, job.FrequencyTime.String)
		if cronExpr == "" {
			log.Printf("Invalid schedule for job %s", job.Code.String)
			continue
		}

		j := job

		id, err := c.AddFunc(cronExpr, func() {
			fmt.Printf("Running job: %s\n", j.Code.String)
			RunTask(j.Code.String)
		})
		if err != nil {
			log.Printf("Failed to schedule job %s: %v", j.Code.String, err)
			continue
		}
		log.Printf("Scheduled job %s with entry ID %d", j.Code.String, id)
	}
}

func RunTask(code string) {
	fmt.Printf(">> Running scheduled task: %s\n", code)
}
