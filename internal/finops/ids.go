package finops

import (
	"os"
	"strings"
)

const (
	SeedAccountEnv = "CIFRA_SEED_AWS_ACCOUNT_ID"

	PrimaryTenantSlug   = "demo"
	PrimaryTenantName   = "Demo"
	PrimaryAccountAlias = "principal"
	DefaultRegion       = "us-east-1"

	AuthModeDefaultChain = "default_chain"
	AuthModeAccessKeys   = "access_keys"

	RoleOwner  = "owner"
	RoleMember = "member"

	SourceEstimate = "estimate"
	SourceCE       = "ce"

	SyncOK     = "ok"
	SyncFailed = "failed"

	FindingCEDenied      = "ce_denied"
	FindingUnattachedIP  = "unattached_ip"
	FindingUnknownS3Size = "unknown_s3_size"
	FindingStoppedBill   = "stopped_instance_billed"
)

func SeedAWSAccountID() string {
	return strings.TrimSpace(os.Getenv(SeedAccountEnv))
}

func ValidAWSAccountID(id string) bool {
	if len(id) != 12 {
		return false
	}
	for _, c := range id {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
