# Onboarding Guide для новых разработчиков

Добро пожаловать в команду Nuum! Этот гайд поможет вам быстро влиться в проект.

## День 1: Настройка окружения

### Шаг 1: Получение доступов

- [ ] Доступ к GitHub репозиторию
- [ ] Добавлен в команду на GitHub
- [ ] Добавлен в GitHub Project
- [ ] Доступ к дев стенду (если есть)
- [ ] Доступ к документации
- [ ] Доступ к коммуникационным каналам

### Шаг 2: Установка необходимых инструментов

#### Обязательные инструменты

```bash
# Git
git --version  # должно быть >= 2.30

# Docker & Docker Compose
docker --version  # >= 20.10
docker-compose --version  # >= 2.0

# Kubernetes tools (для K8s окружения)
minikube version  # >= 1.30
kubectl version   # >= 1.28
```

#### Инструменты для Backend (Go)

```bash
# Go
go version  # >= 1.22

# Go tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/swaggo/swag/cmd/swag@latest

# FFmpeg (для работы с видео)
ffmpeg -version  # >= 4.4
```

#### Инструменты для Frontend (Node.js)

```bash
# Node.js и npm
node --version  # >= 20
npm --version   # >= 10

# или pnpm (рекомендуется)
npm install -g pnpm
pnpm --version
```

### Шаг 3: Клонирование и настройка проекта

```bash
# 1. Клонируйте репозиторий
git clone git@github.com:JediAnaki/Nuum.git
cd Nuum

# 2. Настройте Git
git config user.name "Your Name"
git config user.email "your.email@example.com"

# 3. Создайте develop ветку
git checkout develop
git pull origin develop

# 4. Скопируйте .env файлы
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
cp worker/.env.example worker/.env

# 5. Запустите проект (выберите вариант)

# Вариант A: Kubernetes (рекомендуется)
make k8s-init
make k8s-build
make k8s-deploy
make k8s-frontend  # Откроет frontend в браузере

# Вариант B: Docker Compose
make start
```

### Шаг 4: Проверка работоспособности

1. Откройте frontend: http://localhost:3000
2. Проверьте API: http://localhost:8080/health
3. Проверьте что можете:
   - Зарегистрироваться
   - Войти в систему
   - Загрузить тестовое видео
   - Просмотреть видео

## День 2: Изучение проекта

### Структура проекта

```
Nuum/
├── backend/                    # Go API сервер
│   ├── cmd/api/               # Точка входа
│   ├── internal/
│   │   ├── config/           # Конфигурация
│   │   ├── database/         # DB подключения
│   │   ├── handlers/         # HTTP handlers
│   │   ├── middleware/       # Middleware (auth, cors, etc.)
│   │   ├── models/           # Data models
│   │   └── services/         # Бизнес-логика
│   └── design-features/      # Документация фич
│       ├── features/         # 36 дизайн документов фич
│       └── implementation/   # Гайды по имплементации
│
├── worker/                    # Video processing worker
│   ├── cmd/worker/           # Точка входа
│   └── internal/
│       ├── processor/        # Обработка видео (FFmpeg)
│       └── queue/            # Redis queue
│
├── frontend/                  # Next.js приложение
│   ├── app/                  # App Router pages
│   ├── components/           # React компоненты
│   ├── lib/                  # Утилиты (API client, etc.)
│   └── types/                # TypeScript types
│
├── k8s/                      # Kubernetes манифесты
├── docs/                     # Документация
└── scripts/                  # Utility scripts
```

### Ключевые концепции

#### Архитектура

1. **Backend (Go + Fiber)**
   - RESTful API
   - JWT аутентификация
   - PostgreSQL для данных
   - Redis для кеша и очередей
   - MinIO/S3 для хранения видео

2. **Worker (Go + FFmpeg)**
   - Обрабатывает видео асинхронно
   - Создает разные качества (360p, 480p, 720p, 1080p)
   - Генерирует thumbnails
   - Обновляет статус в БД

3. **Frontend (Next.js 14)**
   - Server-side rendering
   - TanStack Query для state management
   - TailwindCSS для стилей
   - TypeScript для type safety

#### Поток данных

```
User → Frontend → Backend API → PostgreSQL
                    ↓
                 Redis Queue
                    ↓
                  Worker → FFmpeg → MinIO/S3
                    ↓
              Update DB status
```

### Изучите design docs

Все фичи задокументированы в `backend/design-features/features/`:

**Базовые фичи (начните с них):**
- `001-auth-system.md` - Аутентификация и авторизация
- `002-video-player.md` - Видео плеер
- `003-home-page-feed.md` - Главная страница
- `005-video-upload.md` - Загрузка видео

**Читайте по мере необходимости:**
- Остальные 32+ фичи в той же папке

### API endpoints

Основные endpoints (см. полную документацию в Swagger):

**Auth:**
```
POST /api/v1/auth/register  - Регистрация
POST /api/v1/auth/login     - Вход
GET  /api/v1/auth/me        - Текущий пользователь
```

**Videos:**
```
GET    /api/v1/videos       - Список видео
GET    /api/v1/videos/:id   - Одно видео
POST   /api/v1/videos       - Загрузка видео
DELETE /api/v1/videos/:id   - Удаление видео
PATCH  /api/v1/videos/:id   - Обновление метаданных
```

**Comments:**
```
GET  /api/v1/videos/:id/comments     - Комментарии
POST /api/v1/videos/:id/comments     - Добавить комментарий
```

## День 3: Первая задача

### Выберите простую задачу

1. Зайдите в GitHub Projects
2. Найдите issues с метками:
   - `good first issue`
   - `help wanted`
   - `difficulty: easy`

3. Типичные первые задачи:
   - Исправить UI баг
   - Добавить валидацию
   - Написать unit test
   - Улучшить error message
   - Обновить документацию

### Workflow первой задачи

```bash
# 1. Создайте ветку от develop
git checkout develop
git pull origin develop
git checkout -b feature/your-task-name

# 2. Внесите изменения

# 3. Запустите тесты
make test

# 4. Закоммитьте
git add .
git commit -m "feat(scope): add feature X

- Details
- More details

Closes #123"

# 5. Запушьте
git push origin feature/your-task-name

# 6. Создайте PR
# - Заполните PR template
# - Назначьте ревьюеров
# - Свяжите с issue
# - Добавьте скриншоты (если UI)

# 7. Отреагируйте на ревью
# - Внесите правки
# - Ответьте на комментарии
# - Push изменения

# 8. После approve - merge
```

## Неделя 1: Полезные практики

### Daily routine

1. **Начало дня:**
   ```bash
   git checkout develop
   git pull origin develop
   # Убедитесь что окружение работает
   make test
   ```

2. **Во время работы:**
   - Делайте небольшие коммиты часто
   - Пишите тесты для нового кода
   - Проверяйте линтеры перед push
   - Синхронизируйте с develop регулярно

3. **Перед созданием PR:**
   - Самостоятельно проверьте diff
   - Запустите все тесты
   - Проверьте что не забыли удалить debug код
   - Убедитесь что документация обновлена

### Как искать информацию

1. **В коде:**
   ```bash
   # Найти использование функции
   grep -r "functionName" .

   # Найти все TODO
   grep -r "TODO" backend/

   # Посмотреть историю файла
   git log -p backend/internal/handlers/auth_handler.go
   ```

2. **В документации:**
   - `backend/design-features/` - дизайн фич
   - `docs/` - общая документация
   - `README.md` - setup инструкции
   - GitHub Wiki - дополнительные гайды

3. **У команды:**
   - Создайте Discussion в GitHub
   - Оставьте комментарий в PR
   - Напишите в issue

### Типичные проблемы и решения

#### Проблема: Docker контейнеры не запускаются

```bash
# Проверьте логи
docker-compose logs

# Пересоздайте контейнеры
docker-compose down -v
docker-compose up -d --build

# Очистите все
docker system prune -a
```

#### Проблема: Тесты падают локально

```bash
# Убедитесь что зависимости актуальны
cd backend && go mod download
cd frontend && npm install

# Очистите кеш
go clean -testcache
```

#### Проблема: Git конфликты

```bash
# Обновите develop
git checkout develop
git pull origin develop

# Вернитесь в ветку и сделайте rebase
git checkout feature/your-branch
git rebase develop

# Разрешите конфликты
# ... edit files ...
git add .
git rebase --continue
```

## Полезные команды

### Makefile commands

```bash
make help           # Показать все доступные команды
make start          # Запустить Docker Compose
make stop           # Остановить Docker Compose
make logs           # Показать логи
make test           # Запустить все тесты
make lint           # Запустить линтеры

# Kubernetes
make k8s-init       # Инициализировать minikube
make k8s-build      # Собрать образы
make k8s-deploy     # Задеплоить
make k8s-status     # Статус сервисов
make k8s-logs       # Логи
make k8s-clean      # Очистить ресурсы
```

### Git commands

```bash
# Посмотреть все ветки
git branch -a

# Посмотреть статус
git status

# Посмотреть diff
git diff

# Stash изменения
git stash
git stash pop

# Изменить последний коммит
git commit --amend

# Откатить изменения
git reset --hard HEAD
```

## Чеклист первой недели

- [ ] Настроено окружение
- [ ] Проект запускается локально
- [ ] Изучена архитектура
- [ ] Прочитаны базовые design docs
- [ ] Создан первый PR
- [ ] Пройден code review
- [ ] Слит первый PR
- [ ] Понятен workflow команды
- [ ] Знаете где искать информацию
- [ ] Знаете к кому обратиться за помощью

## Следующие шаги

После первой недели:

1. **Углубитесь в свою зону ответственности:**
   - Frontend разработчики: изучите Next.js patterns, component library
   - Backend разработчики: изучите Go best practices, database design
   - DevOps: изучите K8s setup, CI/CD pipelines

2. **Изучите advanced темы:**
   - Видео processing (FFmpeg)
   - Scalability patterns
   - Security best practices
   - Performance optimization

3. **Начните работать над большими фичами:**
   - Выберите фичу из GitHub Project
   - Изучите design doc
   - Составьте plan
   - Разбейте на подзадачи
   - Реализуйте пошагово

Удачи! Если есть вопросы - не стесняйтесь спрашивать! 🚀
