# Development Workflow

## Git Branching Strategy

Мы используем **Git Flow** с некоторыми модификациями:

```
main (production)
  │
  └── develop (integration branch)
       ├── feature/001-auth-system
       ├── feature/002-video-player
       ├── feature/your-feature
       ├── bugfix/fix-upload-timeout
       └── hotfix/critical-security-fix
```

### Типы веток

#### main
- **Production-ready** код
- Только stable releases
- Protected branch (требует PR + approvals)
- Автоматический deploy на production

#### develop
- **Integration** ветка для разработки
- Все фичи мержатся сюда
- Protected branch
- Автоматический deploy на staging

#### feature/*
- Новая функциональность
- Создается от: `develop`
- Мержится в: `develop`
- Naming: `feature/123-short-description` или `feature/auth-system`

#### bugfix/*
- Исправление багов в develop
- Создается от: `develop`
- Мержится в: `develop`
- Naming: `bugfix/123-fix-description`

#### hotfix/*
- Критические баги в production
- Создается от: `main`
- Мержится в: `main` И `develop`
- Naming: `hotfix/critical-issue-description`

#### release/*
- Подготовка к релизу
- Создается от: `develop`
- Мержится в: `main` и `develop`
- Naming: `release/v1.2.0`

## Workflow для разных сценариев

### Сценарий 1: Новая фича

```bash
# 1. Обновите develop
git checkout develop
git pull origin develop

# 2. Создайте ветку фичи
git checkout -b feature/video-comments

# 3. Разрабатывайте
# ... make changes ...
git add .
git commit -m "feat(comments): add comments API"

# 4. Регулярно синхронизируйте с develop
git fetch origin develop
git rebase origin/develop

# 5. Push ветки
git push origin feature/video-comments

# 6. Создайте Pull Request в develop
# GitHub UI: feature/video-comments → develop

# 7. После review и approve - squash merge
# Ветка автоматически удалится
```

### Сценарий 2: Работа над одной из 36 фич

У нас уже созданы ветки для всех фич из design-features:

```bash
# 1. Посмотрите список веток фич
git branch -r | grep feature/

# 2. Checkout нужной фичи
git checkout feature/001-auth-system

# 3. Обновите от develop
git pull origin develop

# 4. Разрабатывайте
# ... make changes ...

# 5. Commit и push
git add .
git commit -m "feat(auth): implement JWT tokens"
git push origin feature/001-auth-system

# 6. Когда фича готова - создайте PR в develop
```

### Сценарий 3: Исправление бага

```bash
# 1. Создайте bugfix ветку
git checkout develop
git pull origin develop
git checkout -b bugfix/upload-timeout

# 2. Исправьте баг
# ... fix bug ...
git add .
git commit -m "fix(upload): increase timeout to 5 minutes

Fixes #234"

# 3. Push и создайте PR
git push origin bugfix/upload-timeout
# PR → develop
```

### Сценарий 4: Hotfix для production

```bash
# 1. Создайте hotfix от main
git checkout main
git pull origin main
git checkout -b hotfix/security-patch

# 2. Исправьте критический баг
# ... emergency fix ...
git add .
git commit -m "fix(security): patch XSS vulnerability

SECURITY: Fixes critical XSS in video titles"

# 3. Push
git push origin hotfix/security-patch

# 4. Создайте 2 PR:
# - hotfix/security-patch → main (priority!)
# - hotfix/security-patch → develop

# 5. После merge в main - создайте tag
git checkout main
git pull origin main
git tag -a v1.2.1 -m "Security hotfix v1.2.1"
git push origin v1.2.1
```

### Сценарий 5: Release

```bash
# 1. Создайте release ветку от develop
git checkout develop
git pull origin develop
git checkout -b release/v1.3.0

# 2. Подготовьте релиз
# - Обновите версию в package.json, go.mod
# - Обновите CHANGELOG.md
# - Последние bug fixes
# - Документацию

git add .
git commit -m "chore(release): prepare v1.3.0"

# 3. Push
git push origin release/v1.3.0

# 4. Тестирование на staging

# 5. Если все ОК - создайте 2 PR:
# - release/v1.3.0 → main
# - release/v1.3.0 → develop

# 6. После merge в main - создайте tag
git checkout main
git pull origin main
git tag -a v1.3.0 -m "Release v1.3.0"
git push origin v1.3.0

# 7. GitHub Release
# Создайте release в GitHub UI с notes
```

## Commit Message Convention

Используем **Conventional Commits**:

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: новая функциональность
- `fix`: исправление бага
- `docs`: изменения в документации
- `style`: форматирование, пропущенные точки с запятой
- `refactor`: рефакторинг кода
- `perf`: улучшение производительности
- `test`: добавление тестов
- `chore`: обновление зависимостей, конфигов
- `ci`: изменения в CI/CD
- `build`: изменения в build системе

### Scopes (примеры)

- `auth`: аутентификация
- `video`: работа с видео
- `upload`: загрузка файлов
- `player`: видео плеер
- `comments`: комментарии
- `api`: API changes
- `db`: database changes
- `ui`: UI components

### Примеры

```bash
# Новая фича
git commit -m "feat(auth): add Google OAuth login

- Added Google OAuth provider
- Created callback endpoint
- Updated user model with google_id

Closes #123"

# Bug fix
git commit -m "fix(upload): resolve timeout on large files

- Increased upload timeout to 5 minutes
- Added progress indicator
- Improved error messages

Fixes #234"

# Breaking change
git commit -m "feat(api): redesign video upload API

BREAKING CHANGE: Upload endpoint now requires multipart form data
instead of JSON. Update clients accordingly.

Closes #345"

# Multiple changes
git commit -m "refactor(player): improve video player performance

- Lazy load player component
- Optimize buffer management
- Reduce memory usage by 30%
- Add performance metrics

Related to #456"
```

## Pull Request Process

### 1. Перед созданием PR

Checklist:
```bash
# Убедитесь что код обновлен
git fetch origin develop
git rebase origin/develop

# Запустите тесты
make test

# Запустите линтеры
make lint

# Проверьте что все работает
make start
# Ручное тестирование

# Self-review изменений
git diff develop...HEAD
```

### 2. Создание PR

```bash
# Push ветки
git push origin feature/your-feature

# В GitHub UI:
# 1. Нажмите "New Pull Request"
# 2. Base: develop ← Compare: feature/your-feature
# 3. Заполните PR template
# 4. Назначьте reviewers
# 5. Добавьте labels
# 6. Link issues (Closes #123)
# 7. Создайте PR
```

### 3. PR Template

PR должен содержать:

```markdown
## Description
Краткое описание изменений

## Related Issues
Closes #123
Related to #456

## Changes
- Added feature X
- Fixed bug Y
- Improved Z

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Tested manually
- [ ] All tests pass

## Screenshots (if UI changes)
[Скриншоты до/после]

## Deployment Notes
- Database migrations: yes/no
- Config changes: yes/no
- Breaking changes: yes/no
```

### 4. Code Review

**Для автора:**
- Отвечайте на комментарии в течение 24 часов
- Вносите правки promptly
- Mark conversations as resolved
- Re-request review после правок

**Для ревьюера:**
- Review в течение 24 часов
- Используйте constructive feedback
- Проверьте:
  - Логика корректна
  - Тесты покрывают изменения
  - Код читаем и maintainable
  - Нет security issues
  - Performance OK
- Approve или Request Changes

### 5. Merge

После получения approvals:

```bash
# В GitHub UI нажмите "Squash and merge"
# Это создаст один clean commit в develop

# Commit message будет:
feat(scope): PR title (#123)

* commit 1
* commit 2
* commit 3

# Ветка автоматически удалится
```

## Branch Protection Rules

### main branch

```yaml
Required:
  - At least 2 approvals
  - All CI checks must pass
  - No direct pushes
  - Require linear history
  - Include administrators

Protections:
  - Require status checks:
    - tests
    - lint
    - build
  - Require branches to be up to date
  - Require conversation resolution
```

### develop branch

```yaml
Required:
  - At least 1 approval
  - All CI checks must pass
  - No direct pushes

Protections:
  - Require status checks:
    - tests
    - lint
```

## CI/CD Pipeline

### On Pull Request

```yaml
name: PR Checks

on:
  pull_request:
    branches: [develop, main]

jobs:
  lint:
    - Backend: golangci-lint
    - Frontend: eslint, prettier

  test:
    - Backend: go test with coverage
    - Frontend: jest, playwright
    - Integration tests

  build:
    - Backend: go build
    - Frontend: npm run build
    - Docker images

  security:
    - Dependency check
    - SAST scanning
    - Secrets scanning
```

### On Merge to develop

```yaml
name: Deploy Staging

on:
  push:
    branches: [develop]

jobs:
  deploy:
    - Build Docker images
    - Push to registry
    - Deploy to staging environment
    - Run smoke tests
    - Notify team
```

### On Merge to main

```yaml
name: Deploy Production

on:
  push:
    branches: [main]

jobs:
  deploy:
    - Build Docker images
    - Push to registry
    - Deploy to production
    - Run health checks
    - Create GitHub Release
    - Notify team
```

## Code Review Guidelines

### Что проверять

**Functionality**
- [ ] Код делает то, что заявлено
- [ ] Edge cases обработаны
- [ ] Error handling корректен

**Code Quality**
- [ ] Код читаем и понятен
- [ ] Нет дублирования
- [ ] Функции небольшие и focused
- [ ] Переменные названы понятно

**Testing**
- [ ] Тесты покрывают новый код
- [ ] Тесты проверяют edge cases
- [ ] Тесты читаемы
- [ ] Coverage >= 70%

**Performance**
- [ ] Нет N+1 запросов
- [ ] Индексы в БД на месте
- [ ] Кеширование используется
- [ ] Нет memory leaks

**Security**
- [ ] Input validation
- [ ] SQL injection protection
- [ ] XSS protection
- [ ] Authentication/Authorization
- [ ] Secrets не в коде

**Documentation**
- [ ] Комментарии для сложной логики
- [ ] API документация обновлена
- [ ] README обновлен если нужно

### Типы комментариев

Используйте префиксы:

```
[BLOCKING] - must be fixed before merge
[SUGGESTION] - nice to have, optional
[QUESTION] - need clarification
[NITPICK] - style preference, optional
[PRAISE] - хорошая работа!
```

Примеры:

```
[BLOCKING] This will cause a memory leak. Please add defer close()

[SUGGESTION] Consider extracting this to a separate function for reusability

[QUESTION] Why did you choose this approach instead of X?

[NITPICK] Can we rename this variable to be more descriptive?

[PRAISE] Great use of generics here!
```

## Troubleshooting

### Merge conflicts

```bash
# Обновите develop
git checkout develop
git pull origin develop

# Вернитесь в вашу ветку
git checkout feature/your-feature

# Rebase на develop
git rebase develop

# Если конфликты:
# 1. Откройте файлы с конфликтами
# 2. Разрешите конфликты
# 3. git add <resolved-files>
# 4. git rebase --continue

# Или отмените rebase
git rebase --abort

# Force push после rebase
git push origin feature/your-feature --force-with-lease
```

### Нужно изменить последний commit

```bash
# Внесите изменения
# ... edit files ...

# Amend commit
git add .
git commit --amend --no-edit

# Push
git push origin feature/your-feature --force-with-lease
```

### Нужно изменить commit message

```bash
# Последний commit
git commit --amend

# Предыдущие commits
git rebase -i HEAD~3  # последние 3 коммита
# Измените 'pick' на 'reword' для нужных commits

# Push
git push origin feature/your-feature --force-with-lease
```

### Случайно закоммитили в develop

```bash
# НЕ ДЕЛАЙТЕ ТАК! Но если случилось:

# Создайте ветку с текущими изменениями
git branch feature/my-changes

# Вернитесь к предыдущему состоянию develop
git reset --hard origin/develop

# Переключитесь на новую ветку
git checkout feature/my-changes

# Push новой ветки
git push origin feature/my-changes
```

## Best Practices

1. **Делайте маленькие PR**
   - Легче review
   - Быстрее merge
   - Меньше конфликтов

2. **Commit often**
   - Маленькие логичные commits
   - Легче откатить изменения
   - Лучше history

3. **Rebase, не merge**
   - Чистая linear history
   - Легче читать git log
   - Меньше merge commits

4. **Test before push**
   - Экономит время CI
   - Меньше сломанных builds
   - Лучше качество кода

5. **Review своего кода**
   - Смотрите diff перед PR
   - Находите собственные ошибки
   - Лучшее quality

6. **Держите ветки updated**
   - Регулярно rebase на develop
   - Меньше конфликтов
   - Легче merge

7. **Удаляйте merged ветки**
   - Чистый репозиторий
   - Легче навигация
   - Меньше путаницы

---

## Quick Reference

```bash
# Начать новую фичу
git checkout develop && git pull && git checkout -b feature/name

# Синхронизировать с develop
git fetch origin develop && git rebase origin/develop

# Создать PR
git push origin feature/name
# Затем в GitHub UI

# После merge - cleanup
git checkout develop && git pull && git branch -d feature/name
```
