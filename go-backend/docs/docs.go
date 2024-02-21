// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/check_donation": {
            "post": {
                "description": "Загрузка справки пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Загрузить справку конкретного юзера",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Адрес электронной почты пользователя",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Тип донации пользователя",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Фотография пользователя",
                        "name": "photo",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Фотография успешно загружена",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось загрузить фотографию",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/clinic_create": {
            "post": {
                "description": "Создает новый аккаунт клиники с предоставленным email, паролем, именем и адрессом.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Создать новоый аккаунт клиники",
                "parameters": [
                    {
                        "description": "Запрос на создание пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ClinicCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Аккаунт клиники успешно создан",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось создать аккаунт клиники",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/create_other_info": {
            "post": {
                "description": "Создает анкету юзера для данных в профиле",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Создать анкету юзера",
                "parameters": [
                    {
                        "description": "Запрос на создание анкеты пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateUserOtherInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное создание анкеты",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось создать анкету",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/create_questionnaire": {
            "post": {
                "description": "Создает анкету пользователя или клиники",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Создать анкету",
                "parameters": [
                    {
                        "description": "Запрос на создание анкеты пользователя или клиники",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.QuestionnaireRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное создание анкеты",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось создать анкету",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/delete_application": {
            "post": {
                "description": "Удаляет запись по указанному айди.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Удалить выполненную заявку из БД",
                "parameters": [
                    {
                        "description": "Запрос на удаление заявки",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Заявка успешно удалена",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось удалить заявку",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/forgot": {
            "post": {
                "description": "Инициирует восстановление пароля по email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Восстановление",
                "parameters": [
                    {
                        "description": "Запрос на инициацию восстановления пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ForgotRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Процесс восстановления начат",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось начать процесс восстановления",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_all_clinic_cards": {
            "get": {
                "description": "Возвращает все карточеки клиник.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Все карточки клиник",
                "responses": {
                    "200": {
                        "description": "Карточки получены",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить карточки",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_all_donations": {
            "get": {
                "description": "Получает все справки о новых донациях юзеров.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Получить все справки о донациях юзеров",
                "responses": {
                    "200": {
                        "description": "Заявки успешно получены",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить заявки",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_all_user_cards": {
            "get": {
                "description": "Возвращает все карточки нуждающихся.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Все карточки юзеров",
                "responses": {
                    "200": {
                        "description": "Карточки получены",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить карточки",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_clinic_cards": {
            "get": {
                "description": "Возвращает 10 карточек клиник.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Первые 10 карточек клиник",
                "responses": {
                    "200": {
                        "description": "Карточки получены",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить карточки",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_user_cards": {
            "get": {
                "description": "Возвращает 10 карточек нуждающихся.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Первые 10 карточек юзеров",
                "responses": {
                    "200": {
                        "description": "Карточки получены",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить карточки",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/get_user_stats/{key}": {
            "get": {
                "description": "Статистика донаций пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Статистика конкретного юзера",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user's email key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Статистика получена",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить статистику",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/image/{key}": {
            "get": {
                "description": "Возвращает фото справки пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Фото справки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image id",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Фото получено",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить фото",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/login": {
            "post": {
                "description": "Авторизует пользователя с предоставленным email и паролем.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Логин",
                "parameters": [
                    {
                        "description": "Запрос на авторизацию пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пользователь авторизован",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось авторизовать пользователя",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/logout": {
            "get": {
                "description": "Выполняет выход из аккаунта с сбросом cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Выход из акканута",
                "responses": {
                    "200": {
                        "description": "Выход выполнен",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось выполнить выход из аккаунта пользователя",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/mark": {
            "post": {
                "description": "Помещает анкету в отдельную базу нуждающихся",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Пометить анкету как нуждающегося",
                "parameters": [
                    {
                        "description": "Запрос на перемещение анкеты пользователя в раздел нуждающихся",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное перемещение анкеты",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось переместить анкету",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/message": {
            "post": {
                "description": "Сообщение ассистенту.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Assistant"
                ],
                "summary": "Сообщение ассистенту",
                "parameters": [
                    {
                        "description": "Сообщение ассистенту",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MessageBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Сообщене получено",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить сообщение",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/newpass": {
            "get": {
                "description": "Новый пароль для восстановления пароля по email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Новый пароль",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Токен восстановления пароля",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Процесс восстановления завершен",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось завершить процесс восстановления",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Сохранение нового пароля .",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Добавление нового пароля в БД",
                "parameters": [
                    {
                        "description": "Новый пароль пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.NewPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пароль сохранен",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось сохранить пароль",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/other_info": {
            "get": {
                "description": "Получает анкету пользователя для профиля",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Получить анкету юзера",
                "parameters": [
                    {
                        "description": "Запрос на получение анкеты пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestUserOtherInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешно получена анкета",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить анкету",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/profile": {
            "get": {
                "description": "Возвращает профиль пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Профиль",
                "responses": {
                    "200": {
                        "description": "Профиль получен",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить профиль",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/profile/{key}": {
            "get": {
                "description": "Возвращает профиль другого пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Чужой Профиль",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user's email key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Профиль получен",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить профиль",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/questionnaire": {
            "get": {
                "description": "Возвращает анкету пользователя для старнички \"profile\".",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Анкета",
                "responses": {
                    "200": {
                        "description": "Анкета получена",
                        "schema": {
                            "$ref": "#/definitions/model.OneOfUserClinicQuestionnaireResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить анкету",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/update_user_stats": {
            "put": {
                "description": "Статистика донаций пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Обновить статистику конкретного юзера",
                "parameters": [
                    {
                        "description": "Запрос на получение статистики юзера",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestUpdateStat"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Статистика получена",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось получить статистику",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/user_create": {
            "post": {
                "description": "Создает нового пользователя с предоставленным email, паролем и именем.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Создать нового пользователя",
                "parameters": [
                    {
                        "description": "Запрос на создание пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пользователь успешно создан",
                        "schema": {
                            "$ref": "#/definitions/model.CodeResponse"
                        }
                    },
                    "400": {
                        "description": "Не удалось создать пользователя",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ClinicCreateRequest": {
            "type": "object",
            "required": [
                "Name",
                "address",
                "email",
                "password"
            ],
            "properties": {
                "Name": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.CodeResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.CreateUserOtherInfo": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "age": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.ForgotRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.MessageBody": {
            "type": "object",
            "properties": {
                "message_text": {
                    "type": "string"
                }
            }
        },
        "model.NewPassword": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.OneOfUserClinicQuestionnaireResponse": {
            "type": "object",
            "required": [
                "age",
                "bloodType",
                "bloodTypes",
                "breed",
                "contacts",
                "petName",
                "workHours"
            ],
            "properties": {
                "age": {
                    "type": "string"
                },
                "bloodType": {
                    "type": "string"
                },
                "bloodTypes": {
                    "type": "object",
                    "additionalProperties": true
                },
                "breed": {
                    "type": "string"
                },
                "contacts": {
                    "type": "string"
                },
                "petName": {
                    "type": "string"
                },
                "workHours": {
                    "type": "string"
                }
            }
        },
        "model.QuestionnaireRequest": {
            "type": "object"
        },
        "model.RequestBody": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.RequestUpdateStat": {
            "type": "object",
            "properties": {
                "donationType": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "model.RequestUserOtherInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "model.UserCreateRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
