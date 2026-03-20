# stock-web

## アプリ概要（アイデア）
- 投資をアシストしてくれるアプリ
- 英語学習をアシストしてくれるアプリ

---

## 本リポジトリのひな型（Boilerplate）について
本リポジトリは、以下の構成を持つWebアプリケーションの基盤（ひな型）として完全にセットアップされています。

### 構成技術スタック
- **Backend**: Go, Gin, Ent, Golembic
- **Frontend**: TypeScript, React, Vite, react-admin
- **Infrastructure**: Traefik, PostgreSQL, Docker Compose
- **Observability**: Loki, Grafana
- **CI/CD & Prov**: GitHub Actions, Ansible, Makefile

### ひな型の詳細構成
- **backend/**
  - `Go` + `Gin` で構築されたWebAPIサーバーです。
  - `ent/schema` 配下で DB スキーマを管理・生成します。
  - Docker上にてビルドされ、`/api` エンドポイントとして Traefik 経由で公開されます。
- **frontend/**
  - `Vite` ベースの React + TypeScript アプリケーションです。
  - `react-admin` が導入されており、素早く管理画面・CRUD UI を構築できます。
  - Nginxコンテナで静的配信され、ルート (`/`) アクセスとして Traefik に紐付きます。
- **Docker Compose & Traefik** (`docker-compose.yml`)
  - **Traefik** がリバースプロキシとして働き、`http://localhost` へのアクセスをWebフロントエンドへ、`http://localhost/api` 以下への通信をバックエンドサーバーへ自動ルーティングします。
- **モニタリング環境** (`loki/`, `grafana/`)
  - アプリケーションやコンテナのログを Loki で収集し、`http://localhost:3000` (Grafana) で解析・閲覧できる基盤です。

---

## 利用方法（ローカル環境の起動）
1. Makefileを利用してDocker Composeで関連コンテナを一括起動します。
   ```bash
   make up
   ```
2. ブラウザから以下のURLにアクセスして動作を確認します。
   - **Frontend (Web UI)**: http://localhost
   - **Backend API (API応答確認)**: http://localhost/api/ping
   - **Grafana (ログ可視化)**: http://localhost:3000
3. 開発を終了してコンテナをすべて停止・削除する場合：
   ```bash
   make down
   ```

---

## 今後の拡張方法（開発の進め方）

### 1. データベースのテーブル（モデル）を追加する
バックエンドの `backend/` 階層にて、以下の `ent` コマンドを実行し新規テーブルのスキーマを作成します。
```bash
cd backend
go run -mod=mod entgo.io/ent/cmd/ent init User
```
生成された `ent/schema/user.go` に必要なカラム（フィールド）を定義した後、以下のコマンドでGoのエンティティコードを自動生成させます:
```bash
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
```

### 2. 新しいAPIエンドポイントを追加する
テーブルとデータの準備ができたら、`backend/main.go` （または別ファイルへの分割）に対して新しいルート（例: `r.GET("/api/users", ...)`）を登録し、DBからデータを取得してJSONで返す処理を実装します。Goの実装を変更した後はコンテナをリビルド（`make build` して `make up`）します。

### 3. フロントエンドに管理画面 (CRUD) を追加する
`frontend/src/App.tsx` の `<Admin>` コンポーネントの中に、連携したいエンドポイントのパスを持つ `<Resource>` を追加するだけで、即座に一覧や作成画面を組み込めます。
```tsx
import { Admin, Resource, ListGuesser } from "react-admin";
import simpleRestProvider from "ra-data-simple-rest";

const dataProvider = simpleRestProvider("http://localhost/api");

const App = () => (
  <Admin dataProvider={dataProvider}>
    {/* /api/users エンドポイントのデータ一覧画面を自動生成 */}
    <Resource name="users" list={ListGuesser} />
  </Admin>
);

export default App;
```
フロントエンド側のみのUI開発・調整を行う際は、コンテナではなくホスト側で `cd frontend && npm run dev` を実行すると、コードの変更がブラウザへ即座に反映される（ホットリロード）ため、効率的に開発を進めることができます。