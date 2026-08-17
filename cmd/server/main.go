package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/puppe1990/aws-finops/internal/app"
	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/crypto"
	"github.com/puppe1990/aws-finops/internal/db"
	appi18n "github.com/puppe1990/aws-finops/internal/i18n"
	"github.com/puppe1990/aws-finops/internal/store"
	"github.com/puppe1990/aws-finops/internal/syncer"
	"github.com/puppe1990/aws-finops/web"
	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/boot"
	"github.com/puppe1990/cais/pkg/cais/meta"
	inertia "github.com/romsar/gonertia/v3"
)

func main() {
	loadDotEnv(".env")
	cfg := cais.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}
	preferredPort := cfg.Port
	port, shifted, err := cais.ResolvePort(cfg.Port, cfg.Env)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Port = port

	a, err := bootstrapWithConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	shiftedFrom := ""
	if shifted {
		shiftedFrom = preferredPort
	}
	boot.Print(os.Stdout, boot.Options{
		AppName:         "Cifra",
		Config:          cfg,
		Version:         boot.CaisVersion(),
		PortShiftedFrom: shiftedFrom,
	})
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

func bootstrap() (*app.App, error) {
	return bootstrapWithConfig(cais.Load())
}

func bootstrapWithConfig(cfg cais.Config) (*app.App, error) {
	tmplFS, err := fs.Sub(web.Templates, "templates")
	if err != nil {
		return nil, fmt.Errorf("templates: %w", err)
	}

	catalog := appi18n.NewCatalog(cfg.Locale)
	templatesDir, err := cais.ResolveWebDir("templates", cfg.TemplatesDir)
	if err != nil {
		templatesDir = ""
	}
	renderer, err := cais.NewRendererForEnv(cfg, tmplFS, templatesDir, catalog)
	if err != nil {
		return nil, fmt.Errorf("renderer: %w", err)
	}

	s, err := store.NewSQLiteStore(cfg.DBPath, cfg.Env)
	if err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}
	if err := db.RunSeeds(s); err != nil {
		_ = s.Close()
		return nil, fmt.Errorf("seed: %w", err)
	}

	appSecret, err := app.ResolveAppSecret(cfg.Env, os.Getenv("APP_SECRET"))
	if err != nil {
		_ = s.Close()
		return nil, err
	}
	sync := syncer.New(s, awsinv.NewLive()).WithDecrypt(func(cipher string) (string, error) {
		return crypto.Decrypt(appSecret, cipher)
	})

	staticDir, err := cais.ResolveWebDir("static", cfg.StaticDir)
	if err != nil {
		_ = s.Close()
		return nil, err
	}

	inertiaI, err := inertia.NewFromFileFS(tmplFS, "app.html")
	if err != nil {
		_ = s.Close()
		return nil, fmt.Errorf("inertia root: %w", err)
	}

	return app.New(cfg, app.Deps{
		Renderer:  renderer,
		Store:     s,
		StaticDir: staticDir,
		Site:      meta.SiteFrom("Cifra", cfg.AppURL),
		Catalog:   catalog,
		Inertia:   inertiaI,
		AppSecret: appSecret,
		Syncer:    sync,
	})
}
