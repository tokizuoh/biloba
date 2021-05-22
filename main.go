package main

import (
	"awsce"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func generateTextBlockObjects(output awsce.CostOutput) []*slack.TextBlockObject {
	var objects []*slack.TextBlockObject
	for _, cost := range output.Costs {
		costStr := fmt.Sprintf("%v (%v)", cost.Amount, cost.Unit)
		object := slack.TextBlockObject{
			Text: "*" + cost.Service + "*\n" + costStr,
			Type: "mrkdwn",
		}
		objects = append(objects, &object)
	}
	return objects
}

func generateDateStr(start, end string) string {
	return fmt.Sprintf("*PERIOD*\nEnd: %v \nStart: %v", end, start)
}

func main() {
	costOutput, err := awsce.FetchTotalCost()
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	tkn := os.Getenv("SLACK_BOT_TOKEN")
	c := slack.New(tkn)

	respChannel, respTimestamp, err := c.PostMessage(os.Getenv("SLACK_CHANNEL_ID"), slack.MsgOptionBlocks(
		&slack.SectionBlock{
			Type: slack.MBTSection,
			Text: &slack.TextBlockObject{
				Type: "mrkdwn",
				Text: "Notify AWS costs by AWS Cost Explorer.",
			},
		},
		slack.NewDividerBlock(),
		slack.NewSectionBlock(
			&slack.TextBlockObject{
				Type: "mrkdwn",
				Text: generateDateStr(costOutput.TimePeriod.Start, costOutput.TimePeriod.End),
			},
			generateTextBlockObjects(costOutput),
			slack.NewAccessory(
				slack.NewImageBlockElement(os.Getenv("AWS_COST_EXPLORER_IMG_PATH"), "aws_ce"),
			),
		),
	))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Post Slack is completed!")
	log.Printf("[CHANNEL] %v", respChannel)
	log.Printf("[TIME_STAMP] %v", respTimestamp)
}
