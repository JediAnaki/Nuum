# Видеоплатформа

Современная видеоплатформа для русскоязычной аудитории с поддержкой загрузки, обработки и стриминга видео.

## Стек технологий

### Backend
- **Go** - основной язык для API и worker
- **Fiber** - web framework
- **PostgreSQL** - база данных
- **Redis** - кэш и очереди задач
- **MinIO** - S3-совместимое хранилище (или Backblaze B2)
- **FFmpeg** - обработка видео

### Frontend
- **Next.js 14** - React framework с App Router
- **TypeScript** - типизация
- **TailwindCSS** - стилизация
- **TanStack Query** - управление состоянием сервера
- **Axios** - HTTP клиент

### Инфраструктура
- **Docker & Docker Compose** - контейнеризация
- **Nginx** - reverse proxy (в продакшене)

## Архитектура

```
┌─────────────┐     ┌─────────────┐     ┌──────────────┐
│   Frontend  │────▶│   Backend   │────▶│  PostgreSQL  │
│  (Next.js)  │     │    (Go)     │     │              │
└─────────────┘     └─────────────┘     └──────────────┘
                           │
                           ├────▶┌──────────────┐
                           │     │    Redis     │
                           │     └──────────────┘
                           │
                           ├────▶┌──────────────┐
                           │     │    MinIO     │
                           │     └──────────────┘
                           │
                           ▼
                    ┌─────────────┐
                    │   Worker    │
                    │ (Video Proc)│
                    └─────────────┘
```

## Требования

### Для Docker Compose
- **Docker** >= 20.10
- **Docker Compose** >= 2.0

### Для Kubernetes (рекомендуется)
- **minikube** >= 1.30
- **kubectl** >= 1.28
- **Docker** >= 20.10

### Для локальной разработки
- **Go** >= 1.22
- **Node.js** >= 20

## Варианты запуска

Платформу можно запустить двумя способами:

### 🐳 Docker Compose (простой способ)
Подходит для быстрого тестирования и разработки.

### ☸️ Kubernetes (рекомендуется)
Подходит для изолированной разработки и близко к production окружению.
- Полная изоляция от системы
- Легкая очистка (`make k8s-delete` удалит всё)
- Автоматический restart при сбоях
- Близко к production окружению

Подробная инструкция: **[K8S_GUIDE.md](K8S_GUIDE.md)**

---

## Быстрый старт (Kubernetes)

### 1. Инициализируйте minikube

```bash
make k8s-init
```

### 2. Соберите Docker образы

```bash
make k8s-build
```

### 3. Разверните приложение

```bash
make k8s-deploy
```

### 4. Откройте frontend

```bash
make k8s-frontend
```

### Полезные команды

```bash
make k8s-status      # Статус всех сервисов
make k8s-logs        # Просмотр логов
make k8s-dashboard   # Kubernetes dashboard
make k8s-clean       # Удалить все ресурсы
make k8s-delete      # Полностью удалить кластер
```

📖 **Подробная документация**: [K8S_GUIDE.md](K8S_GUIDE.md)

---

## Быстрый старт (Docker Compose)

### 1. Настройте переменные окружения

```bash
# Backend
cp backend/.env.example backend/.env

# Frontend
cp frontend/.env.example frontend/.env

# Worker
cp worker/.env.example worker/.env
```

### 2. Запустите через Docker Compose

```bash
make start
# или
docker-compose up -d
```

Сервисы будут доступны по адресам:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **MinIO Console**: http://localhost:9001 (admin/minioadmin)
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

### 3. Проверьте статус

```bash
make logs
# или
docker-compose ps
```

## Локальная разработка

### Backend

```bash
cd backend

# Установите зависимости
go mod download

# Запустите сервер
go run cmd/api/main.go
```

### Worker

```bash
cd worker

# Установите зависимости
go mod download

# Запустите worker
go run cmd/worker/main.go
```

### Frontend

```bash
cd frontend

# Установите зависимости
npm install

# Запустите dev сервер
npm run dev
```

## API Endpoints

### Аутентификация

- `POST /api/v1/auth/register` - регистрация пользователя
- `POST /api/v1/auth/login` - авторизация
- `GET /api/v1/auth/me` - получить текущего пользователя (требует auth)

### Видео

- `GET /api/v1/videos` - список видео (с пагинацией)
- `GET /api/v1/videos/:id` - получить видео по ID
- `POST /api/v1/videos` - загрузить видео (требует auth)
- `DELETE /api/v1/videos/:id` - удалить видео (требует auth)

### Пример запроса

```bash
# Регистрация
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "full_name": "Test User"
  }'

# Загрузка видео
curl -X POST http://localhost:8080/api/v1/videos \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -F "video=@/path/to/video.mp4" \
  -F "title=My Video" \
  -F "description=Video description"
```

## Обработка видео

Worker автоматически обрабатывает загруженные видео:

1. Получает задачу из Redis Stream
2. Транскодирует видео в разные качества:
   - 360p (640x360, 800kbps)
   - 480p (854x480, 1400kbps)
   - 720p (1280x720, 2800kbps)
   - 1080p (1920x1080, 5000kbps)
3. Генерирует превью (thumbnail)
4. Сохраняет результаты в БД
5. Обновляет статус видео на "ready"

## Структура проекта

```
.
├── backend/                 # Go API
│   ├── cmd/
│   │   └── api/            # Точка входа API
│   ├── internal/
│   │   ├── config/         # Конфигурация
│   │   ├── database/       # Подключения к БД
│   │   ├── handlers/       # HTTP handlers
│   │   ├── middleware/     # Middleware
│   │   ├── models/         # Модели данных
│   │   └── services/       # Бизнес-логика
│   └── pkg/
│       └── logger/         # Логирование
│
├── worker/                 # Video processor
│   ├── cmd/
│   │   └── worker/        # Точка входа worker
│   └── internal/
│       ├── processor/     # Обработка видео
│       └── queue/         # Redis queue
│
├── frontend/              # Next.js frontend
│   ├── app/              # App Router pages
│   ├── components/       # React компоненты
│   ├── lib/             # Утилиты
│   └── types/           # TypeScript типы
│
└── docker-compose.yml   # Оркестрация сервисов
```

## Развертывание

### На Hetzner VPS (~€20/мес)

1. Создайте VPS (4 vCPU, 16GB RAM, 160GB SSD)

2. Установите Docker и Docker Compose

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
```

3. Склонируйте проект и настройте .env

4. Запустите через Docker Compose

```bash
docker-compose -f docker-compose.prod.yml up -d
```

5. Настройте Nginx как reverse proxy с SSL

### Оптимизация затрат

- **Хранилище**: Используйте Backblaze B2 ($5/TB вместо $23/TB на AWS S3)
- **CDN**: Bunny CDN ($1/TB трафика) или CloudFlare
- **VPS**: Hetzner дешевле AWS/GCP в 3-4 раза

## Мониторинг

В продакшене рекомендуется добавить:
- **Prometheus + Grafana** - метрики
- **Loki** - логи
- **Sentry** - отслеживание ошибок

## Следующие шаги

### Базовые фичи
- [ ] Комментарии к видео
- [ ] Система лайков/дизлайков
- [ ] Подписки на каналы
- [ ] Уведомления
- [ ] Поиск по видео

### Монетизация
- [ ] Интеграция рекламы
- [ ] Система подписок
- [ ] Донаты для авторов

### Улучшения
- [ ] Адаптивный стриминг (HLS/DASH)
- [ ] Live стриминг
- [ ] Мобильные приложения (React Native)
- [ ] Рекомендательная система на ML
- [ ] Автоматическая модерация контента

## Лицензия

MIT

## Поддержка

Для вопросов и предложений создавайте issue в репозитории.
