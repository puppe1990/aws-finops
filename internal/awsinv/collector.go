package awsinv

import (
	"context"
	"time"

	"github.com/puppe1990/aws-finops/internal/models"
)

type Credentials struct {
	Mode            string
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	AccountID       string
}

type Inventory struct {
	Source    string
	Resources []models.CloudResource
	Lines     []models.CostLine
	Findings  []models.Finding
	Warnings  []string
}

type Collector interface {
	Collect(ctx context.Context, creds Credentials) (Inventory, error)
}

type MonthCoster interface {
	CostForMonth(ctx context.Context, creds Credentials, period time.Time) ([]models.CostLine, error)
}

type CostForecaster interface {
	ForecastForMonth(ctx context.Context, creds Credentials, period time.Time) (int64, error)
}
