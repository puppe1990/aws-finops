package awsinv

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	cetypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func collectCostAnomalies(ctx context.Context, cfg aws.Config, from, to time.Time) ([]CostAnomaly, error) {
	out, err := costexplorer.NewFromConfig(cfg).GetAnomalies(ctx, &costexplorer.GetAnomaliesInput{
		DateInterval: &cetypes.AnomalyDateInterval{
			StartDate: aws.String(from.UTC().Format("2006-01-02")),
			EndDate:   aws.String(to.UTC().Format("2006-01-02")),
		},
	})
	if err != nil {
		return nil, err
	}
	var items []CostAnomaly
	for _, a := range out.Anomalies {
		impact, score := 0.0, 0.0
		if a.Impact != nil {
			impact = a.Impact.TotalImpact
		}
		if a.AnomalyScore != nil {
			score = a.AnomalyScore.CurrentScore
		}
		items = append(items, MapCostAnomaly(
			aws.ToString(a.AnomalyId),
			aws.ToString(a.DimensionValue),
			aws.ToString(a.AnomalyStartDate),
			aws.ToString(a.AnomalyEndDate),
			impact,
			score,
		))
	}
	return items, nil
}
