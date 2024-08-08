# 使い方

### ネットワークを作成

```
docker network create hentai_salon
```

### 起動

```
docker compose up -d
```

### 終了

```
dokcer compose stop

```

### generate

```
docker compose exec server-app go generate ./infrastructure/ent
```

### マイグレーションフィアル生成

```
docker compose exec server-app go run -mod=mod ./infrastructure/ent/migrate/main.go {{ マイグレーション名 }}
```

### マイグレーション

```
docker compose exec server-app atlas migrate apply \
  --dir "file://./infrastructure/ent/migrate/migrations" \
  --url "mysql://hentai_salon:H3nt@1_Sa!0n_2024@mysql:3306/hentai_salon"
```

### シード

```
docker compose exec server-app go run ./infrastructure/ent/seed/main.go
```
