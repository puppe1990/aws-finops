# Cifra — FinOps AWS

Ferramenta de FinOps multi-tenant para AWS, gerada com [Cais](https://github.com/puppe1990/cais) (Go + Inertia + Svelte + SQLite).

## O que ela faz

- Isola clientes em workspaces (tenants): contas, inventário, orçamentos e achados não vazam
- Sincroniza Cost Explorer quando a IAM permite `ce:GetCostAndUsage`
- Sem CE, estima o burn pelo inventário: Lightsail (instâncias + IPs) e S3
- Mostra a policy IAM mínima para destravar billing real
- Permite vincular outras contas AWS por workspace (cadeia padrão ou access keys cifradas)

Login de desenvolvimento: `demo@example.com` / `password`.

## Stack

- Go 1.26 + Cais + gonertia
- Svelte 5 + Tailwind
- SQLite
- AWS SDK v2 (Cost Explorer, Lightsail, S3, STS)

## Quick start

```bash
export PATH="$HOME/go/bin:$PATH"
export LOCALE=pt
cp .env.example .env   # preencha CIFRA_SEED_AWS_ACCOUNT_ID se quiser seedar uma conta
cais install
cais dev               # http://localhost:8080
```

A primeira visita ao dashboard sincroniza as contas do workspace ativo (cadeia `~/.aws` ou access keys do tenant).

## Seed da sua conta (opcional, local)

Não commite account IDs. No `.env` (gitignored):

```
CIFRA_SEED_AWS_ACCOUNT_ID=xxxxxxxxxxxx
```

Sem essa variável o workspace Demo nasce vazio e você vincula a conta pela UI.

## IAM recomendada

Cole a policy em `/settings` numa role de leitura de FinOps. Sem `ce:GetCostAndUsage` o Cifra estima pelo inventário.

## Testes

```bash
cais test
go test ./... -count=1
npm run test:fe
```
