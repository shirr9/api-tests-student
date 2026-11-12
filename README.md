# API-тесты для проекта https://testboard.avito.com/

## Полезные ссылки

1. [Swagger-спецификация](https://testboard.avito.com/swagger/index.html)
2. [Техническое задание](https://docs.google.com/document/d/1IZ9SxEcePoQT6lyEmukycp1qNok4HZtok73UQnPwKSA/edit?usp=sharing)
3. [Библиотека для написания API-тестов](https://github.com/steinfletcher/apitest)
4. [Мощная библиотека анализа JSON для Go](https://github.com/tidwall/gjson)

## Запуск тестов

1. Установите любое IDE для разработки: [Jetbrains GoLand](https://www.jetbrains.com/ru-ru/go/) (Платное или по студенческой лицензии), [VS Code](https://code.visualstudio.com/) (бесплатное).
2. Установите [Git](https://git-scm.com/install/) на компьютер
3. [Установите](https://go.dev/doc/install) Go версии 1.25.1
4. [Настройте](https://docs.github.com/ru/authentication/connecting-to-github-with-ssh) доступ к своему Github-аккаунту
   по SSH-ключу.
5. Склонируйте себе репозиторий: `git clone <путь к вашему репозиторию>`
6. Подтяните необходимые зависимости через команду в корне проекта:

```bash
   go mod tidy
```

7. Создайте в корне `.env.override` с данными двух разных аккаунтов TestBoard. Укажите фамилию владельца и используйте покупателя с пустым избранным.
   Задайте `TEST_AD_TITLE` длиной до 39 символов и укажите абсолютный путь к JPEG или PNG. Второе фото необязательно.

```dotenv
TEST_OWNER_LOGIN=owner@example.com
TEST_OWNER_PASSWORD=your_owner_password
TEST_BUYER_LOGIN=buyer@example.com
TEST_BUYER_PASSWORD=your_buyer_password
TEST_AD_TITLE="Котик"
TEST_AD_PHOTO="/absolute/path/photo.jpeg"
# TEST_AD_SECOND_PHOTO="/absolute/path/second photo.jpeg"
```

Для теста лимита в каталоге нужны как минимум 50 чужих объявлений от продавцов с заполненной фамилией.
Для пакета `advertisementSearch` на стенде должно быть объявление «QA Practice — тестовое объявление 1».

8. Для запуска всех тестов достаточно выполнить команду в корне проекта:

```bash
  go test -p 1 -count=1 -v ./...
```

9. Отдельный пакет с тестами можно запустить командой:

```bash
  go test -v ./tests/scenarios/myAdvertisement/
```

## Логирование

Для включения логирования request/response в консоль: выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/)
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```