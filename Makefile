.PHONY: help start stop logs build clean dev-backend dev-frontend dev-worker \
	k8s-init k8s-build k8s-deploy k8s-clean k8s-delete k8s-status k8s-logs \
	k8s-dashboard k8s-tunnel k8s-restart k8s-frontend k8s-backend

help: ## Показать помощь
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

start: ## Запустить все сервисы
	docker-compose up -d

stop: ## Остановить все сервисы
	docker-compose down

logs: ## Показать логи всех сервисов
	docker-compose logs -f

build: ## Пересобрать все Docker образы
	docker-compose build

clean: ## Удалить все контейнеры и volumes
	docker-compose down -v
	rm -rf uploads/ processed_videos/ thumbnails/

dev-backend: ## Запустить backend локально
	cd backend && go run cmd/api/main.go

dev-worker: ## Запустить worker локально
	cd worker && go run cmd/worker/main.go

dev-frontend: ## Запустить frontend локально
	cd frontend && npm run dev

install-backend: ## Установить зависимости backend
	cd backend && go mod download

install-worker: ## Установить зависимости worker
	cd worker && go mod download

install-frontend: ## Установить зависимости frontend
	cd frontend && npm install

test-backend: ## Запустить тесты backend
	cd backend && go test ./...

test-worker: ## Запустить тесты worker
	cd worker && go test ./...

# ===== Kubernetes команды =====

k8s-init: ## Инициализировать minikube кластер
	@./scripts/k8s-init.sh

k8s-build: ## Собрать Docker образы и загрузить в minikube
	@./scripts/k8s-build-images.sh

k8s-deploy: ## Развернуть все сервисы в Kubernetes
	@./scripts/k8s-deploy.sh

k8s-clean: ## Удалить все ресурсы из кластера
	kubectl delete namespace video-platform || true

k8s-delete: ## Полностью удалить minikube кластер
	minikube delete

k8s-status: ## Показать статус pods и services
	@echo "=== Namespace ==="
	@kubectl get namespace video-platform 2>/dev/null || echo "Namespace не создан"
	@echo "\n=== Pods ==="
	@kubectl get pods -n video-platform 2>/dev/null || echo "Pods не найдены"
	@echo "\n=== Services ==="
	@kubectl get services -n video-platform 2>/dev/null || echo "Services не найдены"
	@echo "\n=== PVCs ==="
	@kubectl get pvc -n video-platform 2>/dev/null || echo "PVCs не найдены"

k8s-logs: ## Показать логи всех pods
	@echo "Выберите pod для просмотра логов:"
	@kubectl get pods -n video-platform -o name 2>/dev/null || echo "Pods не найдены"

k8s-dashboard: ## Открыть Kubernetes dashboard
	minikube dashboard

k8s-tunnel: ## Создать туннель для доступа к сервисам (запустить в отдельном терминале)
	minikube tunnel

k8s-restart: ## Перезапустить все deployments
	kubectl rollout restart deployment -n video-platform

k8s-frontend: ## Открыть frontend в браузере
	@minikube service frontend-service -n video-platform

k8s-backend: ## Получить URL backend API
	@minikube service backend-service -n video-platform --url

k8s-minio: ## Открыть MinIO console
	@minikube service minio-service -n video-platform
