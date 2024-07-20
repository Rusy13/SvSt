
# Примеры запросов <a name="examples"></a>


* [Создание мока: POST http://localhost:8000/mocks с телом]
```
{
  "method": "GET",
  "url": "/api/resource",
  "request_body": "",
  "status_code": 200,
  "headers": {
    "Content-Type": "application/json"
  },
  "body": {
    "message": "Mocked response"    
  }
}
```


* [Список моков: GET http://localhost:8000/mocks]


* [Определенный мок: GET http://localhost:8000/mocks/{method}/{url:.*}]
```
Реальный запрос: http://localhost:8000/mocks/GET/api/resource
```




```
Процесс запуска: 

Для запуска БД: docker-compose up --build
Прогон миграций делается из: make migration-up (Нужно перед этим установить goose)
Для запуска приложения нужно перейти в директорию cmd/main и запустить команду go run main.go

```
