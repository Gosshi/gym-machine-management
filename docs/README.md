# Gym Machine Management 設計ドキュメント

このリポジトリでは、ジムのマシン管理サービスの設計に関するドキュメントを管理します。

## 📌 ドキュメント一覧
- [要件定義](requirements.md)
- [システムアーキテクチャ](system-architecture.md)
- [データベース設計](database-schema.md)
- [API 設計](api-design.md)
- [UI ワイヤーフレーム](ui-wireframe.md)
- [将来的な拡張計画](future-plans.md)


Go の import 文はプロジェクトの module path に基づいて解決されるため、
backend/ の相対パスではなく、Go モジュールのルート (github.com/gosshi/gym-machine-management) からのパスで指定する必要があります。