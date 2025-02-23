package main

//go:generate godotenv -f ./.env tern migrate --migrations ./internal/pgstore/migrations/ --config ./internal/pgstore/migrations/tern.conf
//go:generate sqlc generate -f .\internal\pgstore\sqlc.yaml
//go:generate goapi-gen --package=spec --out ./internal/api/spec/planner.gen.spec.go ./internal/api/spec/planner.spec.json