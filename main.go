package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func main() {
	start := "2020-06-01"
	end := "2021-05-01"
	granularity := "MONTHLY"
	metrics := []string{
		"BlendedCost",
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	svc := costexplorer.New(
		sess,
	)

	result, err := svc.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Granularity: aws.String(granularity),
		GroupBy: []*costexplorer.GroupDefinition{
			&costexplorer.GroupDefinition{
				Key:  aws.String("SERVICE"),
				Type: aws.String("DIMENSION"),
			},
		},
		Metrics: aws.StringSlice(metrics),
	})
	if err != nil {
		log.Fatalf("Unable to generate report, %v", err)
	}

	log.Println("Cost Report:", result.ResultsByTime)
}
