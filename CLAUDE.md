# CLAUDE.md

このファイルは Claude Code (claude.ai/code) がこのリポジトリで作業する際のガイドです。

## プロジェクト概要

**dbpilot** は Claude Code から MySQL/PostgreSQL データベースと連携するための MCP (Model Context Protocol) サーバーです。

## 実装計画

段階的な実装計画は `tmp/planning/` ディレクトリを参照してください。

- [実装計画概要](tmp/planning/00_overview.md)
- Phase 1 (MVP): ステップ 01-06
- Phase 2: セキュリティ強化
- Phase 3: 拡張機能 (分析・DDL)
- Phase 4: テスト・ドキュメント・リリース

## 参考資料

- 設計ドキュメント: `tmp/mcp-db-server-design.md`
