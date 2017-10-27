# Release script
# Install github-release => go get github.com/aktau/github-release

RELEASE=0.1.0
USERNAME=owainlewis
REPO=relay

if git rev-parse "$RELEASE" >/dev/null 2>&1; then
    echo "Tag $RELEASE already exists. Doing nothing";
else
    echo "Creating new release $RELEASE"
    git tag -a "$RELEASE" -m "Relase version: $RELEASE"
    git push --tags

    github-release release \
    --user $USERNAME \
    --repo $REPO \
    --tag $RELEASE \
    --name "Relay $RELEASE" \
    --description "Release version $RELEASE"

    github-release upload \
    --user $USERNAME \
    --repo $REPO \
    --tag $RELEASE \
    --name "relay" \
    --file dist/relay
fi
