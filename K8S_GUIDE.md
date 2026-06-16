# Руководство по развертыванию в Kubernetes

Подробная инструкция по запуску видеоплатформы в Kubernetes (minikube).

## Предварительные требования

### 1. Установка необходимых инструментов

**minikube**
```bash
# macOS
brew install minikube

# Linux
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

**kubectl**
```bash
# macOS
brew install kubectl

# Linux
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install kubectl /usr/local/bin/kubectl
```

**Docker**
- Убедитесь, что Docker установлен и запущен
- minikube будет использовать Docker driver по умолчанию

## Быстрый старт

### 1. Инициализация кластера

```bash
make k8s-init
```

Эта команда:
- Проверит наличие minikube и kubectl
- Запустит minikube с 4 CPU, 8GB RAM, 40GB disk
- Включит необходимые addons (ingress, metrics-server, dashboard)

### 2. Сборка Docker образов

```bash
make k8s-build
```

Эта команда:
- Настроит Docker environment для minikube
- Соберет образы backend, worker и frontend
- Образы будут доступны внутри minikube

### 3. Развертывание приложения

```bash
make k8s-deploy
```

Эта команда:
- Создаст namespace `video-platform`
- Развернет все сервисы по порядку:
  1. Secrets и ConfigMaps
  2. PersistentVolumeClaims
  3. PostgreSQL и Redis
  4. MinIO
  5. Backend API
  6. Worker
  7. Frontend
- Дождется готовности каждого сервиса

## Доступ к сервисам

### Frontend

```bash
make k8s-frontend
# или
minikube service frontend-service -n video-platform
```

Откроет frontend в браузере.

### Backend API

```bash
make k8s-backend
# или
minikube service backend-service -n video-platform --url
```

Получить URL backend API для тестирования.

### MinIO Console

```bash
make k8s-minio
# или
minikube service minio-service -n video-platform
```

Логин: `minioadmin` / `minioadmin`

## Мониторинг и отладка

### Статус всех сервисов

```bash
make k8s-status
```

Покажет состояние:
- Namespace
- Pods
- Services
- PersistentVolumeClaims

### Просмотр логов

**Все pods:**
```bash
kubectl get pods -n video-platform
```

**Логи конкретного pod:**
```bash
kubectl logs -f <pod-name> -n video-platform

# Примеры:
kubectl logs -f backend-xxxxx -n video-platform
kubectl logs -f worker-xxxxx -n video-platform
kubectl logs -f frontend-xxxxx -n video-platform
```

**Логи предыдущего запуска (если pod перезапустился):**
```bash
kubectl logs <pod-name> -n video-platform --previous
```

### Dashboard

```bash
make k8s-dashboard
```

Откроет Kubernetes Dashboard с визуальным интерфейсом.

### Подключение к pod

```bash
kubectl exec -it <pod-name> -n video-platform -- /bin/sh
```

## Управление

### Перезапуск сервисов

**Все deployments:**
```bash
make k8s-restart
```

**Конкретный deployment:**
```bash
kubectl rollout restart deployment/<name> -n video-platform

# Примеры:
kubectl rollout restart deployment/backend -n video-platform
kubectl rollout restart deployment/frontend -n video-platform
```

### Масштабирование

```bash
# Увеличить количество реплик backend
kubectl scale deployment/backend --replicas=3 -n video-platform

# Уменьшить до 1
kubectl scale deployment/backend --replicas=1 -n video-platform
```

### Обновление образов

После изменения кода:

```bash
# 1. Пересобрать образы
make k8s-build

# 2. Перезапустить deployments
make k8s-restart
```

## Очистка

### Удалить все ресурсы, сохранить кластер

```bash
make k8s-clean
```

Удалит namespace `video-platform` со всеми ресурсами.

### Полностью удалить minikube кластер

```bash
make k8s-delete
```

Удалит весь minikube кластер (освободит ~8GB RAM и ~40GB disk).

## Структура ресурсов

```
k8s/
├── namespace.yaml                    # Namespace для изоляции
├── configmaps/
│   ├── backend-config.yaml          # Переменные окружения backend
│   └── worker-config.yaml           # Переменные окружения worker
├── secrets/
│   └── app-secrets.yaml             # Чувствительные данные
├── storage/
│   ├── postgres-pvc.yaml            # 5Gi для PostgreSQL
│   ├── minio-pvc.yaml               # 20Gi для видео/файлов
│   ├── minio-deployment.yaml
│   └── minio-service.yaml
├── databases/
│   ├── postgres-deployment.yaml     # PostgreSQL 16
│   ├── postgres-service.yaml
│   ├── redis-deployment.yaml        # Redis 7
│   └── redis-service.yaml
├── backend/
│   ├── backend-deployment.yaml      # Go API
│   └── backend-service.yaml
├── worker/
│   └── worker-deployment.yaml       # Video processor
├── frontend/
│   ├── frontend-deployment.yaml     # Next.js
│   └── frontend-service.yaml
└── ingress.yaml                      # Ingress controller
```

## Resource Limits

Каждый сервис настроен с лимитами ресурсов:

| Сервис     | CPU Request | CPU Limit | Memory Request | Memory Limit |
|------------|-------------|-----------|----------------|--------------|
| PostgreSQL | 250m        | 500m      | 256Mi          | 512Mi        |
| Redis      | 100m        | 200m      | 128Mi          | 256Mi        |
| MinIO      | 200m        | 400m      | 256Mi          | 512Mi        |
| Backend    | 200m        | 500m      | 256Mi          | 512Mi        |
| Worker     | 500m        | 1000m     | 512Mi          | 1Gi          |
| Frontend   | 200m        | 400m      | 256Mi          | 512Mi        |

**Итого:** ~1.5 CPU, ~2GB RAM minimum

## Troubleshooting

### Pod не запускается

```bash
# Проверить статус
kubectl describe pod <pod-name> -n video-platform

# Посмотреть события
kubectl get events -n video-platform --sort-by='.lastTimestamp'
```

### ImagePullBackOff

Если образ не найден:
```bash
# Пересобрать образы в minikube environment
eval $(minikube docker-env)
make k8s-build
```

### Недостаточно ресурсов

```bash
# Увеличить ресурсы minikube
minikube stop
minikube delete
minikube start --cpus=6 --memory=12288
```

### База данных не готова

```bash
# Проверить логи PostgreSQL
kubectl logs -f postgres-xxxxx -n video-platform

# Проверить PVC
kubectl get pvc -n video-platform
```

### Backend не может подключиться к БД

Проверьте что initContainers завершились успешно:
```bash
kubectl describe pod backend-xxxxx -n video-platform
```

## Production Considerations

Для production развертывания рекомендуется:

1. **Использовать реальный Kubernetes кластер** (не minikube):
   - Google GKE
   - Amazon EKS
   - Azure AKS
   - DigitalOcean Kubernetes

2. **Настроить Helm charts** для упрощения управления

3. **Использовать внешние управляемые сервисы**:
   - Managed PostgreSQL (RDS, Cloud SQL)
   - Managed Redis (ElastiCache, MemoryStore)
   - S3-совместимое хранилище (AWS S3, Backblaze B2)

4. **Настроить мониторинг**:
   - Prometheus + Grafana
   - Loki для логов
   - Jaeger для tracing

5. **Настроить автомасштабирование**:
   - Horizontal Pod Autoscaler (HPA)
   - Vertical Pod Autoscaler (VPA)
   - Cluster Autoscaler

6. **Настроить безопасность**:
   - NetworkPolicies
   - Pod Security Policies
   - Secrets management (Sealed Secrets, External Secrets)
   - TLS/SSL сертификаты

7. **Настроить backup**:
   - Velero для backup кластера
   - Регулярные снимки БД
   - Репликация хранилища

## Полезные ссылки

- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [minikube Documentation](https://minikube.sigs.k8s.io/docs/)
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
