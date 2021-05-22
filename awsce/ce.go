package awsce

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

type CostOutput struct {
	Service    string
	Amount     string
	Unit       string
	TimePeriod TimePeriod
}

type TimePeriod struct {
	end   string
	start string
}

func FetchTotalCost() ([]CostOutput, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().UTC().In(jst)
	dayBefore := now.AddDate(0, 0, -1)
	dayBeforeYesterday := now.AddDate(0, 0, -2)

	dayBeforeStr := dayBefore.Format("2006-01-02")
	dayBeforeYesterdayStr := dayBeforeYesterday.Format("2006-01-02")

	start := dayBeforeYesterdayStr
	end := dayBeforeStr
	granularity := "DAILY"
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
		return nil, err
	}

	var costOutputs = []CostOutput{}
	for _, rbt := range result.ResultsByTime {
		tp := TimePeriod{
			end:   *rbt.TimePeriod.End,
			start: *rbt.TimePeriod.Start,
		}
		for _, g := range rbt.Groups {
			co := CostOutput{
				Service:    *g.Keys[0],
				Amount:     *g.Metrics["BlendedCost"].Amount,
				Unit:       *g.Metrics["BlendedCost"].Unit,
				TimePeriod: tp,
			}
			costOutputs = append(costOutputs, co)
		}
	}

	return costOutputs, nil
}
