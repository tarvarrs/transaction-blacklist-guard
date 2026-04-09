# transaction-blacklist-guard

Бэкенд, который принимает решение о пропуске транзакций на основании справочника недобросовестных ID.

## Запуск проекта

1. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/tarvarrs/transaction-blacklist-guard.git
    cd transaction-blacklist-guard
    ```

2. Загрузите зависимости:

    ```bash
    go mod download
    ```

3. Создайте файл с переменными окружения .env на основе примера .env.example

4. Запустите БД:

    ```bash
    make up
    ```

5. Запустите приложение:

    ```bash
    make run
    ```

6. API будет доступен по адресу `http://localhost:8080`

    Пример request:
    ```bash
    curl --location 'localhost:8080/wallet-operation' \
    --header 'Content-Type: application/json' \
    --data '{
    "from_ID": "1",
    "to_ID": "5",
    "amount":10
    }'
    ```

    Пример response:
    ```json
    {
        "status": "Cancel"
    }
    ```

---

Примечание: в справочнике недобросовестных ID находятся значения с 1 по 5
