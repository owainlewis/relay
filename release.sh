echo "Creating release $1"

gox -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"

ghr -u owainlewis $1 build
