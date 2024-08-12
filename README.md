# Тестовое задание для стажёра Backend
## Сервис баннеров

### Запуск приложения:
```bash
make app_start
```

### Запуск интеграционных тестов:
```bash
make integration_tests_run
```

### Запуск программы, имитирующей 1k PPS
```bash
make client_start
```

## Примеры запросов curl


```bash
curl -X GET "http://127.0.0.1:8080/user_banner?tag_id=1&feature_id=8&use_last_revision=true" -H "Token: user_token"
```

```bash
curl -X GET "http://127.0.0.1:8080/banner?tag_id=123&limit=1&offset=3" -H "Token: admin_token"
```

```bash
curl -X POST "http://127.0.0.1:8080/banner" -H "Content-Type: application/json" -H "Token: admin_token" -d '{
  "tag_ids": [123, 12, 1],
  "feature_id": 8,
  "content": "{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}",
  "is_active": false
}'
```

```bash
curl -X PATCH "http://127.0.0.1:8080/banner/10" -H "Content-Type: application/json" -H "Token: admin_token" -d '{
  "tag_ids": [31, 22],
  "feature_id": 9,
  "content": "{\"title6\": \"new_title7\", \"text3\": \"new_text3\", \"url\": \"new_url2\"}"
}'
```

```bash
curl -X DELETE "http://127.0.0.1:8080/banner/6" -H "Token: admin_token"
```

```bash
curl -X GET "http://127.0.0.1:8080/banner/9/versions" -H "Token: admin_token"
```


```bash
curl -X POST "http://127.0.0.1:8080/banner/versions/5" -H "Token: admin_token"
```

```bash
curl -X DELETE "http://127.0.0.1:8080/banners?feature_id=9" -H "Token: admin_token"
```