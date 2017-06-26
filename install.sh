RELEASE=0.1

wget https://github.com/owainlewis/relay/releases/download/$RELEASE/relay

mv relay /usr/local/bin/ && \
chmod +x /usr/local/bin/relay
