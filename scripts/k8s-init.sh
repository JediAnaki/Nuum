#!/bin/bash

set -e

echo "=== Инициализация minikube кластера ==="

# Проверка установки minikube
if ! command -v minikube &> /dev/null; then
    echo "❌ minikube не установлен"
    echo "Установите minikube: https://minikube.sigs.k8s.io/docs/start/"
    exit 1
fi

# Проверка установки kubectl
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl не установлен"
    echo "Установите kubectl: https://kubernetes.io/docs/tasks/tools/"
    exit 1
fi

# Проверка запущен ли minikube
if minikube status &> /dev/null; then
    echo "✅ minikube уже запущен"
else
    echo "🚀 Запуск minikube..."
    minikube start \
        --cpus=4 \
        --memory=8192 \
        --disk-size=40g \
        --driver=docker
fi

# Включение addons
echo "🔧 Включение необходимых addons..."
minikube addons enable ingress
minikube addons enable metrics-server
minikube addons enable dashboard

# Настройка Docker environment
echo "🐳 Настройка Docker environment для minikube..."
eval $(minikube docker-env)

echo ""
echo "✅ minikube успешно инициализирован!"
echo ""
echo "📝 Полезные команды:"
echo "  • minikube status          - статус кластера"
echo "  • minikube dashboard       - открыть dashboard"
echo "  • minikube stop            - остановить кластер"
echo "  • minikube delete          - удалить кластер"
echo ""
