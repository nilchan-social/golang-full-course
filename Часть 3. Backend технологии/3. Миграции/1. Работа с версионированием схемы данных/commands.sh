# Создать файлы миграций для новой версии схемы данных
migrate create -ext sql -dir migrations -seq unique_phone

# Узнать текущую версию схемы данных (на месте {connection_string} строка подключения)
migrate -database {connection_string} "{connection_string}" version

# Накатить самую свежую версию схемы данных
migrate -database {connection_string} "{connection_string}" up

# Откатить схему данных в нулевую версию
migrate -database {connection_string} "{connection_string}" down

# Прыгнуть на схему данных, версия которой выше текущей на {n}
migrate -database {connection_string} "{connection_string}" up {n}

# Откатиться на схему данных, версия которой ниже текущей на {n}
migrate -database {connection_string} "{connection_string}" down {n}
