#!/bin/bash

set -e

echo "=== Сборка Docker образов для minikube ==="

# Настройка Docker environment для minikube
echo "🐳 Настройка Docker environment..."
eval $(minikube docker-env)

# Сборка backend
echo ""
echo "📦 Сборка backend образа..."
docker build -t video-platform-backend:latest ./backend

# Сборка worker
echo ""
echo "📦 Сборка worker образа..."
docker build -t video-platform-worker:latest ./worker

# Сборка frontend
echo ""
echo "📦 Сборка frontend образа..."
docker build -t video-platform-frontend:latest ./frontend

echo ""
echo "✅ Все образы успешно собраны!"
echo ""
echo "📋 Список образов:"
docker images | grep video-platform

echo ""
echo "💡 Теперь можно развернуть приложение: make k8s-deploy"
echo ""
