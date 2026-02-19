# コマンド一覧を表示
default:
    @just --list

# バイナリをビルド
build:
    go build -o dbpilot ./cmd/mcp-db

# バイナリをビルドして /usr/local/bin にインストール
install: build
    mv dbpilot /usr/local/bin/dbpilot
    @echo "Installed dbpilot to /usr/local/bin/dbpilot"

# 設定ファイルのサンプルをコピー
setup:
    cp -n config.example.yaml config.yaml || true
    @echo "Created config.yaml. Please configure your database connection settings."
