#!/bin/bash

set -e

echo "=== Развертывание видеоплатформы в Kubernetes ==="

# Создание namespace
echo "📦 Создание namespace..."
kubectl apply -f k8s/namespace.yaml

# Ожидание создания namespace
sleep 2

# Создание secrets и configmaps
echo "🔐 Создание secrets и configmaps..."
kubectl apply -f k8s/secrets/
kubectl apply -f k8s/configmaps/

# Создание storage
echo "💾 Создание persistent volume claims..."
kubectl apply -f k8s/storage/postgres-pvc.yaml
kubectl apply -f k8s/storage/minio-pvc.yaml

# Развертывание баз данных
echo "🗄️  Развертывание PostgreSQL..."
kubectl apply -f k8s/databases/postgres-deployment.yaml
kubectl apply -f k8s/databases/postgres-service.yaml

echo "🗄️  Развертывание Redis..."
kubectl apply -f k8s/databases/redis-deployment.yaml
kubectl apply -f k8s/databases/redis-service.yaml

# Ожидание готовности БД
echo "⏳ Ожидание готовности баз данных..."
kubectl wait --for=condition=ready pod -l app=postgres -n video-platform --timeout=120s
kubectl wait --for=condition=ready pod -l app=redis -n video-platform --timeout=120s

# Развертывание MinIO
echo "📁 Развертывание MinIO..."
kubectl apply -f k8s/storage/minio-deployment.yaml
kubectl apply -f k8s/storage/minio-service.yaml

echo "⏳ Ожидание готовности MinIO..."
kubectl wait --for=condition=ready pod -l app=minio -n video-platform --timeout=120s

# Развертывание backend
echo "🚀 Развертывание Backend API..."
kubectl apply -f k8s/backend/

echo "⏳ Ожидание готовности Backend..."
kubectl wait --for=condition=ready pod -l app=backend -n video-platform --timeout=180s

# Развертывание worker
echo "⚙️  Развертывание Worker..."
kubectl apply -f k8s/worker/

# Развертывание frontend
echo "🎨 Развертывание Frontend..."
kubectl apply -f k8s/frontend/

echo "⏳ Ожидание готовности Frontend..."
kubectl wait --for=condition=ready pod -l app=frontend -n video-platform --timeout=180s

# Развертывание ingress (опционально)
echo "🌐 Развертывание Ingress..."
kubectl apply -f k8s/ingress.yaml || echo "⚠️  Ingress не развернут (это нормально если addon отключен)"

echo ""
echo "✅ Развертывание завершено!"
echo ""
echo "📊 Статус сервисов:"
kubectl get pods -n video-platform
echo ""
kubectl get services -n video-platform
echo ""
echo "🌐 Доступ к сервисам:"
echo "  • Frontend:     minikube service frontend-service -n video-platform"
echo "  • Backend API:  minikube service backend-service -n video-platform"
echo "  • MinIO:        minikube service minio-service -n video-platform"
echo ""
echo "📝 Полезные команды:"
echo "  • make k8s-status     - статус всех сервисов"
echo "  • make k8s-frontend   - открыть frontend"
echo "  • make k8s-dashboard  - открыть Kubernetes dashboard"
echo ""
