pushd .. > /dev/null
go run cmd/download.go -y=2022 -l=rs -d="$1"
# shellcheck disable=SC2164
popd > /dev/null
