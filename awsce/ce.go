package awsce

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

type CostOutput struct {
	Service string
	Amount  string
	Unit    string
}

func fetchTotalCost() []CostOutput {
	start := "2021-05-20"
	end := "2021-05-21"
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

	var costOutputs = []CostOutput{}
	for _, rbt := range result.ResultsByTime {
		for _, g := range rbt.Groups {
			co := CostOutput{
				Service: *g.Keys[0],
				Amount:  *g.Metrics["BlendedCost"].Amount,
				Unit:    *g.Metrics["BlendedCost"].Unit,
			}
			costOutputs = append(costOutputs, co)
		}
	}

	return costOutputs
}
