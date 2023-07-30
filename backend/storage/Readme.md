# Generate client

docker run --rm   -v ${PWD}:/local openapitools/openapi-generator-cli generate   -i /local/vault-spec.yaml   -g go   -o /local/client  -c /local/config.yaml 