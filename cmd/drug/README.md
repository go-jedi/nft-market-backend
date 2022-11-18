# http API

```sh
# dev
http://localhost:8080/
```

### Регистрация пользователя

````sh
Url: http://localhost:8080/api-v1/user/registration
```js
{
  "body": {
    "tele_id": int, <- Telegram id пользователя из таблицы users поле tele_id
    "tele_name": int <- Telegram name мероприятия из таблицы ad поле tele_name
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешная регистрация пользователя" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка получения из базы данных, {err}";
  -message: "ошибка выполнения функции uid из базы данных, {err}";
  -message: "ошибка выполнения функции user_create из базы данных, {err}";
````
