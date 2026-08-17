package awsinv

const FinOpsIAMPolicy = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "CostExplorerAndBudgets",
      "Effect": "Allow",
      "Action": [
        "ce:GetCostAndUsage",
        "ce:GetCostForecast",
        "ce:GetRightsizingRecommendation",
        "ce:GetAnomalies",
        "budgets:ViewBudget"
      ],
      "Resource": "*"
    },
    {
      "Sid": "InventoryRead",
      "Effect": "Allow",
      "Action": [
        "sts:GetCallerIdentity",
        "s3:ListAllMyBuckets",
        "s3:GetBucketLocation",
        "s3:ListBucket",
        "lightsail:GetInstances",
        "lightsail:GetBundles",
        "lightsail:GetStaticIps",
        "lightsail:GetDisks",
        "lightsail:GetInstanceSnapshots",
        "lightsail:GetRelationalDatabases",
        "cloudwatch:GetMetricStatistics"
      ],
      "Resource": "*"
    }
  ]
}`
