![Go Report](https://goreportcard.com/badge/github.com/denis-shcherbinin/spbpu-software-design-project)
![Repository Top Language](https://img.shields.io/github/languages/top/denis-shcherbinin/spbpu-software-design-project)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/denis-shcherbinin/spbpu-software-design-project)
![Lines of code](https://img.shields.io/tokei/lines/github/denis-shcherbinin/spbpu-software-design-project)
![Github Repository Size](https://img.shields.io/github/repo-size/evt/rest-api-example)

# Проект "todo-app" в рамках курса "Конструирование программного обеспечения" в СПбПУ.  

**Денис Щербинин, гр. 3530202/90201.**  

## ToDo App REST API  

<img align="right" width="32%" src="./images/gopher-big-slice.png" alt="">  

Приложение REST API для управления повседневными задачами.  
Возможности API:
1. Профиль пользователя:
    - Регистрация
    - Вход в профиль
2. Списки задач ***(todo-list)***
3. Задачи к определённому списку ***(todo-item)*** 

## Архитектура


## Тестирование:
**Структура и логика приложения очень простая. Поэтому были написаны unit-тесты на основной функционал.**  
**Также было проведено ручное тестирование через Postman**
1. Чтобы запустить тесты:
    ```
    make gotest
    ```

## Запуск приложения 

### Требования:
* Docker, docker-compose
* Поддержка Makefile

### Старт приложения: 
1. Создать `.env` файл. Для тестирования можно взять данные из `.env-example`
2. Из корня проекта запустить: 
    ```
    make run
    ```
При первом запуске, некоторые контейнеры могут перезапускаться, это нормально.
