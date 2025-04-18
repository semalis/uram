#!/usr/bin/env bash

set -eo pipefail

if ! [ -x "$(command -v swagger-combine)" ]; then
	echo 'Error: swagger-combine is not installed. Install with $~ npm i -g swagger-combine' >&2
	exit 1
fi

mkdir -p ./tmp-swagger-gen
cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
	# generate swagger files (filter query files)
	query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
	if [[ -n $query_file ]]; then
		buf generate --template buf.gen.swagger.yaml "$query_file"
	fi
done

cd ..
# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging

arkeoTemplate="{uram-version}"
arkeoVersion=$(echo $(git describe --tags) | sed 's/^v//')
sed "s/$arkeoTemplate/$arkeoVersion/g" ./docs/proto/config.json > ./tmp-swagger-gen/config.json

# Fetching the cosmos-sdk version to use the appropriate swagger file
sdkTemplate="{sdk-version}"
sdkVersion=$(go list -m -f '{{ .Version }}' github.com/cosmos/cosmos-sdk)
sed -i "s/$sdkTemplate/$sdkVersion/g" ./tmp-swagger-gen/config.json

swagger-combine ./tmp-swagger-gen/config.json -o ./docs/static/swagger.yaml -f yaml --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen
