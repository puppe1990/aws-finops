package awsinv

import (
	"context"

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
