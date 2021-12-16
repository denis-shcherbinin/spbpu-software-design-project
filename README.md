![Go Report](https://goreportcard.com/badge/github.com/denis-shcherbinin/spbpu-software-design-project)
![Repository Top Language](https://img.shields.io/github/languages/top/denis-shcherbinin/spbpu-software-design-project)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/denis-shcherbinin/spbpu-software-design-project)
![Lines of code](https://img.shields.io/tokei/lines/github/denis-shcherbinin/spbpu-software-design-project)
![Github Repository Size](https://img.shields.io/github/repo-size/evt/rest-api-example)

# Проект "TODO App API" в рамках курса "Конструирование программного обеспечения" в СПбПУ.  

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
**Схема модели C4 для визуализации компонентов проекта:** 

![image](https://user-images.githubusercontent.com/61324182/146454196-ad7fce43-c99b-4a9a-b968-28c43b489229.png) 

**Архитектура приложения:** 

![image](https://user-images.githubusercontent.com/61324182/146454380-22761684-0bdf-4ab7-b5fd-b340b393034b.png)
* Слой `Handler` отвечает за логику запрос/ответ
* Слой `Service` отвечает за бизнес-логику приложения
* Слой `Repository` отвечает за логику работы с базой данных

При разработке приложения были соблюдены подход *Clean Architecture* и принципы *Solid*

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

### Запуск приложения: 
1. Создать `.env` файл. Для тестирования можно взять данные из `.env-example`
2. Команда запуска:: 
    ```
    make run
    ```
При первом запуске, некоторые контейнеры могут перезапускаться, это нормально.
