-- up
CREATE TABLE IF NOT EXISTS tenants (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD COLUMN active_tenant_id INTEGER REFERENCES tenants(id);

CREATE TABLE IF NOT EXISTS tenant_members (
    tenant_id INTEGER NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (tenant_id, user_id)
);

CREATE TABLE IF NOT EXISTS cloud_accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    aws_account_id TEXT NOT NULL,
    alias TEXT NOT NULL,
    region TEXT NOT NULL,
    auth_mode TEXT NOT NULL,
    access_key_id TEXT NOT NULL DEFAULT '',
    secret_cipher TEXT NOT NULL DEFAULT '',
    is_primary INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (tenant_id, aws_account_id)
);

CREATE TABLE IF NOT EXISTS cloud_resources (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cloud_account_id INTEGER NOT NULL REFERENCES cloud_accounts(id) ON DELETE CASCADE,
    kind TEXT NOT NULL,
    name TEXT NOT NULL,
    region TEXT NOT NULL,
    state TEXT NOT NULL DEFAULT '',
    monthly_cents INTEGER NOT NULL DEFAULT 0,
    source TEXT NOT NULL,
    external_id TEXT NOT NULL DEFAULT '',
    meta_json TEXT NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS cost_lines (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cloud_account_id INTEGER NOT NULL REFERENCES cloud_accounts(id) ON DELETE CASCADE,
    service TEXT NOT NULL,
    monthly_cents INTEGER NOT NULL DEFAULT 0,
    source TEXT NOT NULL,
    period_start TEXT NOT NULL DEFAULT '',
    period_end TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS findings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cloud_account_id INTEGER NOT NULL REFERENCES cloud_accounts(id) ON DELETE CASCADE,
    kind TEXT NOT NULL,
    severity TEXT NOT NULL,
    title TEXT NOT NULL,
    detail TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS budgets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    cloud_account_id INTEGER REFERENCES cloud_accounts(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    amount_cents INTEGER NOT NULL,
    period TEXT NOT NULL DEFAULT 'monthly'
);

CREATE TABLE IF NOT EXISTS sync_runs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cloud_account_id INTEGER NOT NULL REFERENCES cloud_accounts(id) ON DELETE CASCADE,
    status TEXT NOT NULL,
    source TEXT NOT NULL DEFAULT '',
    error TEXT NOT NULL DEFAULT '',
    warning TEXT NOT NULL DEFAULT '',
    started_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    finished_at DATETIME
);

-- down
DROP TABLE IF EXISTS sync_runs;
DROP TABLE IF EXISTS budgets;
DROP TABLE IF EXISTS findings;
DROP TABLE IF EXISTS cost_lines;
DROP TABLE IF EXISTS cloud_resources;
DROP TABLE IF EXISTS cloud_accounts;
DROP TABLE IF EXISTS tenant_members;
DROP TABLE IF EXISTS tenants;
