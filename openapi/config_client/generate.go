package generate

//go:generate sh -c "mkdir -p ../../gens/voipbin_client && go run -mod=mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config config.generate.yaml -o ../../gens/voipbin_client/gen.go ../../../monorepo/bin-api-manager/openapi/openapi.yaml"
