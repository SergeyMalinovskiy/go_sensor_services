
### Введение

Тестовый проект

---
### Установка пакетов

- Go Version **1.22**

##### 1. Загрузка зависимостей
```sh
go mod download
```

##### 2. Настройка конфига
```shell
cp .env.example .env
```
Заполните все необходимые поля в `.env`

Отредактируйте необходимые параметры в `config/your-config.yaml`

Добавьте параметр запуска `CONFIG_PATH=config/your-config.yaml` свою IDE или экспортируйте его в shell

```shell
export CONFIG_PATH=config/your-config.yaml
```

##### 3. Запуск приложения

```shell
go run cmd/sensor-services/main.go
```


---
### Разработка

---