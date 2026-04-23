#!/bin/bash
set -e

APP_NAME="weather-app"
VERSION="1.0.0"
BUILD_DIR="${APP_NAME}_${VERSION}-1"

echo "Сборка Weather App deb-пакета"

# Очистка
rm -rf ${BUILD_DIR}
rm -f ${BUILD_DIR}.deb

# Компиляция Go
echo "Компиляция..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o ${APP_NAME} \
    ./cmd/linux/cli \  

if [ ! -f ${APP_NAME} ]; then
    echo "Ошибка компиляции! Проверь путь к main.go"
    echo "Ожидался файл: ./cmd/${APP_NAME}/main.go"
    exit 1
fi

# Копирование структуры из шаблона debian/
echo "Копирование файлов..."
cp -r debian ${BUILD_DIR}

# Копирование скомпилированного бинарника
cp ${APP_NAME} ${BUILD_DIR}/usr/local/bin/
chmod 755 ${BUILD_DIR}/usr/local/bin/${APP_NAME}

# Сборка deb-пакета
echo "Сборка пакета..."
dpkg-deb --build ${BUILD_DIR}

if [ -f ${BUILD_DIR}.deb ]; then
    echo "Готово! Пакет создан: ${BUILD_DIR}.deb"
    ls -lh ${BUILD_DIR}.deb
else
    echo "Ошибка сборки"
    exit 1
fi
