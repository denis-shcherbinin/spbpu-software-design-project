![Go Report](https://goreportcard.com/badge/github.com/denis-shcherbinin/spbpu-software-design-project)
![Repository Top Language](https://img.shields.io/github/languages/top/denis-shcherbinin/spbpu-software-design-project)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/denis-shcherbinin/spbpu-software-design-project)
![Lines of code](https://img.shields.io/tokei/lines/github/denis-shcherbinin/spbpu-software-design-project)
![Github Repository Size](https://img.shields.io/github/repo-size/evt/rest-api-example)

# Проект "todo-app" в рамках курса "Конструирование программного обеспечения" в СПбПУ.  

**Денис Щербинин, гр. 3530202/90201.**  

<img align="right" width="35%" src="./images/gopher-big-slice.png" alt="">  

## ToDo App REST API  

Приложение REST API для управления повседневными задачами.  
Возможности API:
1. Профиль пользователя:
    - Регистрация
    - Вход в профиль
2. Списки задач ***(todo-list)***
3. Задачи к определённому списку ***(todo-item)*** 

## Запуск приложения 

Требования:
* Docker, docker-compose
* Поддержка Makefile

Шаги: 
1. Создать `.env` файл. Для тестирования можно взять данные из `.env-example`
2. Из корня проекта запустить: 
    ```
    make run
    ```
При первом запуске, некоторые контейнеры могут перезапускаться, это нормально.
