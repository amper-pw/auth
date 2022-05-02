Создание базы данных. Логин ``postgres`` пароль ``qwerty``
``docker run --name todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres``
Запуск миграций
``go run cmd/migrate.go``
Генерация ключей
`openssl ecparam -name prime256v1 -genkey -noout -out configs/jwt/private.pem`
`openssl ec -in configs/jwt/private.pem -pubout -out configs/jwt/public.pem`
Пример использвоания 
`go/pkg/mod/github.com/dgrijalva/jwt-go@v3.2.0+incompatible/ecdsa_test.go`
