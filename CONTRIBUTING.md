# Contributing to Nuum

Спасибо за интерес к проекту! Эта инструкция поможет вам быстро начать работу.

## Быстрый старт для новых разработчиков

### 1. Первоначальная настройка

```bash
# Клонируйте репозиторий
git clone git@github.com:JediAnaki/Nuum.git
cd Nuum

# Создайте свою feature ветку от develop
git checkout develop
git pull origin develop
git checkout -b feature/your-feature-name
```

### 2. Запуск окружения

Выберите один из вариантов:

#### Вариант A: Kubernetes (рекомендуется)
```bash
make k8s-init    # Инициализация minikube
make k8s-build   # Сборка образов
make k8s-deploy  # Деплой
```

#### Вариант B: Docker Compose
```bash
# Скопируйте .env файлы
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
cp worker/.env.example worker/.env

# Запустите
make start
```

### 3. Проверка работоспособности

- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- API docs: http://localhost:8080/swagger (если настроен)

## Git Workflow

Мы используем **Git Flow** модель:

```
main       - production-ready код
  └── develop - основная ветка разработки
      ├── feature/001-auth-system
      ├── feature/002-video-player
      └── feature/your-feature-name
```

### Работа над фичей

1. **Создайте ветку от develop**
```bash
git checkout develop
git pull origin develop
git checkout -b feature/short-description
```

2. **Разработка**
   - Делайте небольшие, логичные коммиты
   - Пишите понятные commit messages на английском
   - Регулярно синхронизируйте с develop

3. **Перед созданием PR**
```bash
# Обновите develop
git checkout develop
git pull origin develop

# Вернитесь в свою ветку и сделайте rebase
git checkout feature/your-feature-name
git rebase develop

# Убедитесь что тесты проходят
make test

# Проверьте линтеры
make lint
```

4. **Создайте Pull Request**
   - В GitHub создайте PR из вашей ветки в `develop`
   - Заполните template PR (описание, тесты, скриншоты)
   - Назначьте ревьюеров
   - Свяжите с issue и GitHub Project

### Commit Messages

Используйте формат:

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: новая функциональность
- `fix`: исправление бага
- `docs`: изменения в документации
- `style`: форматирование, пропущенные точки с запятой и т.д.
- `refactor`: рефакторинг кода
- `test`: добавление тестов
- `chore`: обновление конфигурации, зависимостей

**Примеры:**
```
feat(auth): add Google OAuth integration

- Added Google OAuth provider
- Created callback handler
- Updated user model with google_id

Closes #123
```

```
fix(video): resolve upload timeout issue

- Increased upload timeout to 5 minutes
- Added progress indicator
- Improved error handling

Fixes #456
```

## Стандарты кода

### Go (Backend & Worker)

```bash
# Форматирование
gofmt -w .

# Линтинг
golangci-lint run

# Тесты
go test ./... -v -cover
```

**Правила:**
- Используйте `gofmt` для форматирования
- Следуйте [Effective Go](https://golang.org/doc/effective_go.html)
- Покрытие тестами минимум 70%
- Все экспортируемые функции должны иметь комментарии

### TypeScript/React (Frontend)

```bash
# Линтинг
npm run lint

# Форматирование
npm run format

# Тесты
npm run test
```

**Правила:**
- Используйте TypeScript для всех новых файлов
- Компоненты должны быть функциональными
- Используйте React hooks
- Следуйте [Airbnb React Style Guide](https://github.com/airbnb/javascript/tree/master/react)

## Code Review Process

### Для автора PR

1. **Self-review** - проверьте свой код перед созданием PR
2. Убедитесь что:
   - Все тесты проходят
   - Код отформатирован
   - Нет console.log / debug statements
   - Документация обновлена
3. Назначьте минимум 2 ревьюеров
4. Отвечайте на комментарии в течение 24 часов
5. После approve - сделайте squash merge в develop

### Для ревьюера

1. Проверьте PR в течение 24 часов
2. Используйте constructive feedback
3. Проверьте:
   - Логика корректна
   - Тесты покрывают новую функциональность
   - Нет дублирования кода
   - Производительность не пострадала
   - Безопасность (SQL injection, XSS, etc.)
4. Approve только если уверены в коде

## Тестирование

### Backend
```bash
cd backend
go test ./... -v -cover
```

### Frontend
```bash
cd frontend
npm run test
npm run test:e2e  # E2E тесты (Playwright)
```

### Integration Tests
```bash
make test-integration
```

## Документация

При добавлении новой фичи обновите:

1. **API документацию** - `/docs/api/`
2. **README.md** - если изменился setup
3. **Swagger/OpenAPI** - для новых endpoints
4. **Комментарии в коде** - для сложной логики

## База знаний

### Полезные ссылки

- [Структура проекта](docs/PROJECT_STRUCTURE.md)
- [Архитектура](docs/ARCHITECTURE.md)
- [API документация](docs/api/README.md)
- [Deployment](docs/DEPLOYMENT.md)
- [Troubleshooting](docs/TROUBLESHOOTING.md)

### Где найти информацию

- **Design docs** - `backend/design-features/features/`
- **Implementation guides** - `backend/design-features/implementation/`
- **GitHub Projects** - планирование и tracking
- **Issues** - баги и feature requests
- **Discussions** - вопросы и идеи

## Коммуникация

- **GitHub Issues** - для багов и feature requests
- **GitHub Discussions** - для вопросов и обсуждений
- **Pull Requests** - для code review
- **GitHub Projects** - для tracking прогресса

## Вопросы?

Если что-то непонятно:

1. Проверьте [FAQ](docs/FAQ.md)
2. Поищите в GitHub Discussions
3. Создайте новый Discussion
4. Спросите в PR комментариях

Спасибо за ваш вклад! ❤️
