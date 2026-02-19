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

# 依存関係を整理
tidy:
    go mod tidy
