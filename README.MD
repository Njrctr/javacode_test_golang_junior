#                               **Walletter REST Api**

### Для запуска приложения:
Сначала скачайте проект к себе в окружение:
```
git clone http://github.com/Njrctr/javacode_test_golang_junior && cd javacode_test_golang_junior
```
Запуск:
```
make build && make run
```
Альтернативный запуск без утилиты make:
```
docker-compose build walletter-backend && docker-compose up walletter-backend --force-recreate
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate_up
```

### В данном проекте был реализован REST Api сервис для работы с кошельками:
* Весь функционал можно протестировать в SWAGGER документации: http://localhost:8080/swagger/
* Реализован функционал Аутентификации и Авторизации на основе JWT токена
* Разделены хендлеры Админа и обычного юзера (подогнал под требования ТЗ)
