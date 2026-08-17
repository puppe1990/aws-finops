package syncer

import (
	"context"
	"fmt"
	"strings"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/store"
)

type Syncer struct {
	store     store.Store
	collector awsinv.Collector
	decrypt   func(cipher string) (string, error)
}

func New(s store.Store, collector awsinv.Collector) *Syncer {
	return &Syncer{store: s, collector: collector}
}

func (s *Syncer) WithDecrypt(fn func(string) (string, error)) *Syncer {
	s.decrypt = fn
	return s
}

func (s *Syncer) SyncAccount(ctx context.Context, accountID int64) (run storeRun, err error) {
	acc, err := s.store.FindCloudAccount(accountID)
	if err != nil {
		return storeRun{}, err
	}
	runID, err := s.store.StartSyncRun(accountID)
	if err != nil {
		return storeRun{}, err
	}

	creds := awsinv.Credentials{
		Mode:        acc.AuthMode,
		AccessKeyID: acc.AccessKeyID,
		Region:      acc.Region,
		AccountID:   acc.AWSAccountID,
	}
	if acc.AuthMode == finops.AuthModeAccessKeys && acc.SecretCipher != "" && s.decrypt != nil {
		secret, decErr := s.decrypt(acc.SecretCipher)
		if decErr != nil {
			_ = s.store.FinishSyncRun(runID, finops.SyncFailed, "", "", decErr.Error())
			return storeRun{Status: finops.SyncFailed, Error: decErr.Error()}, decErr
		}
		creds.SecretAccessKey = secret
	}

	inv, err := s.collector.Collect(ctx, creds)
	if err != nil {
		_ = s.store.FinishSyncRun(runID, finops.SyncFailed, "", "", err.Error())
		return storeRun{Status: finops.SyncFailed, Error: err.Error()}, err
	}
	if err := s.store.ReplaceResources(accountID, inv.Resources); err != nil {
		return storeRun{}, err
	}
	if err := s.store.ReplaceCostLines(accountID, inv.Lines); err != nil {
		return storeRun{}, err
	}
	if err := s.store.ReplaceFindings(accountID, inv.Findings); err != nil {
		return storeRun{}, err
	}
	warning := strings.Join(inv.Warnings, "; ")
	if err := s.store.FinishSyncRun(runID, finops.SyncOK, inv.Source, warning, ""); err != nil {
		return storeRun{}, err
	}
	return storeRun{ID: runID, Status: finops.SyncOK, Source: inv.Source, Warning: warning}, nil
}

func (s *Syncer) SyncTenant(ctx context.Context, tenantID int64) error {
	accounts, err := s.store.ListCloudAccounts(tenantID)
	if err != nil {
		return err
	}
	var first error
	for _, acc := range accounts {
		if _, err := s.SyncAccount(ctx, acc.ID); err != nil && first == nil {
			first = fmt.Errorf("sync %s: %w", acc.AWSAccountID, err)
		}
	}
	return first
}

type storeRun struct {
	ID      int64
	Status  string
	Source  string
	Warning string
	Error   string
}
