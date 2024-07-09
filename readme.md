# 使い方

### 起動

```
docker compose up -d
```

### 終了

```
dokcer compose down

```

### マイグレーション

```
docker compose exec server atlas migrate apply \
  --dir "file://./infrastructure/ent/migrate/migrations" \
  --url "mysql://hentai_salon:password@192.168.10.20:3306/hentai_salon"
```

### シード

```
docker compose exec server go run ./infrastructure/ent/seed/main.go
```
