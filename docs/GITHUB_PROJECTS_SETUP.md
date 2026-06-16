# GitHub Projects - Setup & Workflow

## Создание GitHub Project

### Шаг 1: Создайте новый Project

1. Перейдите в репозиторий: https://github.com/JediAnaki/Nuum
2. Нажмите на вкладку **Projects**
3. Нажмите **New project**
4. Выберите шаблон **Board** или **Table** (рекомендую Board)
5. Назовите проект: **Nuum Development**

### Шаг 2: Настройте колонки (Board view)

Создайте следующие колонки:

1. **📋 Backlog**
   - Все задачи, которые планируются в будущем
   - Feature requests
   - Идеи

2. **🎯 Ready**
   - Задачи готовые к работе
   - Fully defined
   - Имеют все необходимые детали

3. **🚧 In Progress**
   - Задачи в работе
   - Assigned to someone
   - Активная разработка

4. **👀 In Review**
   - PR создан
   - Ждет code review
   - Testing

5. **✅ Done**
   - Завершенные задачи
   - Merged to develop
   - Deployed

### Шаг 3: Настройте поля (Custom fields)

Добавьте custom fields для лучшего tracking:

**Priority** (Single select)
- 🔴 Critical
- 🟠 High
- 🟡 Medium
- 🟢 Low

**Size** (Single select)
- XS (1-2 hours)
- S (2-4 hours)
- M (1-2 days)
- L (3-5 days)
- XL (1-2 weeks)

**Type** (Single select)
- Feature
- Bug
- Enhancement
- Documentation
- Refactoring
- Testing

**Component** (Single select)
- Backend
- Frontend
- Worker
- DevOps
- Database
- Documentation

**Sprint** (Text)
- Sprint 1, Sprint 2, etc.

**Assignee** (People) - уже есть

**Status** (Single select) - уже есть

### Шаг 4: Создайте Views

#### View 1: Board (по умолчанию)
- Группировка по Status
- Сортировка по Priority

#### View 2: By Component
- Группировка по Component
- Фильтр: Status != Done

#### View 3: Current Sprint
- Фильтр: Sprint = "Current Sprint"
- Группировка по Assignee

#### View 4: Bugs
- Фильтр: Type = Bug
- Сортировка по Priority

#### View 5: Features Roadmap
- Группировка по Feature Category
- Показать только Type = Feature

## Структура Issues

### Issue Template

Создайте `.github/ISSUE_TEMPLATE/feature.md`:

```markdown
---
name: Feature Request
about: Propose a new feature
labels: feature
---

## Feature Description

Brief description of the feature

## User Story

As a [type of user]
I want [goal]
So that [benefit]

## Design Doc

Link to design doc: `backend/design-features/features/XXX-feature-name.md`

## Acceptance Criteria

- [ ] Criterion 1
- [ ] Criterion 2
- [ ] Criterion 3

## Technical Details

### Backend Changes
- API endpoints
- Database migrations
- Services/handlers

### Frontend Changes
- New components
- API integration
- UI/UX

### Testing
- [ ] Unit tests
- [ ] Integration tests
- [ ] E2E tests

## Dependencies

- Depends on #123
- Blocks #456

## Estimated Size

- [ ] XS (1-2h)
- [ ] S (2-4h)
- [ ] M (1-2 days)
- [ ] L (3-5 days)
- [ ] XL (1-2 weeks)
```

Создайте `.github/ISSUE_TEMPLATE/bug.md`:

```markdown
---
name: Bug Report
about: Report a bug
labels: bug
---

## Bug Description

Clear description of the bug

## Steps to Reproduce

1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

## Expected Behavior

What should happen

## Actual Behavior

What actually happens

## Screenshots

If applicable

## Environment

- OS: [e.g. macOS, Ubuntu]
- Browser: [e.g. Chrome, Safari]
- Version: [e.g. 22]

## Logs

```
Paste relevant logs here
```

## Possible Solution

If you have ideas

## Priority

- [ ] Critical (blocks development)
- [ ] High (important bug)
- [ ] Medium (annoying but workable)
- [ ] Low (minor issue)
```

### Pull Request Template

Создайте `.github/PULL_REQUEST_TEMPLATE.md`:

```markdown
## Description

Brief description of changes

## Related Issues

Closes #123
Related to #456

## Type of Change

- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update
- [ ] Refactoring
- [ ] Performance improvement

## Changes

### Backend
- Change 1
- Change 2

### Frontend
- Change 1
- Change 2

### Database
- [ ] Migrations included
- [ ] Seed data updated

## Testing

### Unit Tests
- [ ] All existing tests pass
- [ ] New tests added
- [ ] Coverage >= 70%

### Manual Testing
- [ ] Tested locally
- [ ] Tested in Docker
- [ ] Tested edge cases

### Checklist
- [ ] Code follows style guidelines
- [ ] Self-reviewed code
- [ ] Commented hard-to-understand code
- [ ] Updated documentation
- [ ] No new warnings
- [ ] Added tests
- [ ] All tests pass
- [ ] No breaking changes (or documented)

## Screenshots (if applicable)

Before:
[Screenshot]

After:
[Screenshot]

## Performance Impact

- [ ] No impact
- [ ] Improves performance
- [ ] Degrades performance (explain why acceptable)

## Deployment Notes

Any special deployment considerations:
- Environment variables
- Database migrations
- Cache clearing
- etc.

## Reviewer Notes

@reviewer1 - please check XYZ
@reviewer2 - FYI on ABC changes
```

## Workflow для команды

### Создание новой фичи

1. **Tech Lead создает Issue**
   ```markdown
   Title: [Feature] Add video comments functionality
   Labels: feature, component:backend, component:frontend
   Project: Nuum Development
   Milestone: Sprint 5
   ```

2. **Заполняет детали**
   - Link to design doc
   - Acceptance criteria
   - Size estimate
   - Priority

3. **Добавляет в Project**
   - Status: Backlog
   - Priority: Medium
   - Size: L
   - Component: Backend, Frontend

4. **На planning meeting**
   - Переносим в Ready
   - Назначаем developer
   - Уточняем детали

### Работа разработчика

1. **Берет задачу из Ready**
   - Проверяет что все понятно
   - Задает вопросы если нужно
   - Переводит в In Progress

2. **Создает ветку**
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b feature/001-auth-system
   ```

3. **Разрабатывает**
   - Делает коммиты
   - Обновляет Issue с прогрессом
   - Пишет тесты

4. **Создает PR**
   - Заполняет template
   - Links issue: "Closes #123"
   - Назначает reviewers
   - Переводит в In Review

5. **Code Review**
   - Отвечает на комментарии
   - Вносит правки
   - Re-request review

6. **Merge**
   - После approve
   - Squash and merge в develop
   - Issue автоматически закроется
   - Переместится в Done

## Автоматизация с GitHub Actions

### Auto-assign to Project

Создайте `.github/workflows/add-to-project.yml`:

```yaml
name: Add to Project

on:
  issues:
    types: [opened]
  pull_request:
    types: [opened]

jobs:
  add-to-project:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/add-to-project@v0.4.0
        with:
          project-url: https://github.com/orgs/JediAnaki/projects/1
          github-token: ${{ secrets.GITHUB_TOKEN }}
```

### Auto-move cards

```yaml
name: Move PR to Review

on:
  pull_request:
    types: [opened, ready_for_review]

jobs:
  move-to-review:
    runs-on: ubuntu-latest
    steps:
      - uses: alex-page/github-project-automation-plus@v0.8.1
        with:
          project: Nuum Development
          column: In Review
          repo-token: ${{ secrets.GITHUB_TOKEN }}
```

## Best Practices

### Для Issue

1. **Descriptive titles**
   - ✅ "Add pagination to video feed"
   - ❌ "Fix bug"

2. **Use labels**
   - Type: feature, bug, enhancement
   - Component: backend, frontend, worker
   - Priority: critical, high, medium, low
   - Difficulty: easy, medium, hard

3. **Link everything**
   - Link to design doc
   - Link to related issues
   - Link to PRs
   - Link to discussions

4. **Keep updated**
   - Add comments with progress
   - Update estimates if needed
   - Close when done

### Для Pull Requests

1. **Small PRs**
   - Легче review
   - Быстрее merge
   - Меньше конфликтов

2. **Descriptive titles**
   - ✅ "feat(auth): add Google OAuth integration"
   - ❌ "Update files"

3. **Good descriptions**
   - What changed
   - Why it changed
   - How to test
   - Screenshots (if UI)

4. **Request right reviewers**
   - Backend changes → backend dev
   - Frontend changes → frontend dev
   - Complex changes → senior dev

### Для Projects

1. **Regular grooming**
   - Weekly backlog review
   - Prioritize tasks
   - Remove outdated items

2. **Sprint planning**
   - Every 2 weeks
   - Pick tasks for sprint
   - Set realistic goals

3. **Daily standup** (async via comments)
   - Yesterday: что сделал
   - Today: что буду делать
   - Blockers: что мешает

4. **Sprint review**
   - Demo completed features
   - Discuss what went well
   - Identify improvements

## Metrics to Track

В GitHub Projects можно добавить Insights:

1. **Velocity**
   - Story points completed per sprint
   - Trend over time

2. **Cycle Time**
   - Time from In Progress → Done
   - Identify bottlenecks

3. **Lead Time**
   - Time from Backlog → Done
   - Overall efficiency

4. **WIP (Work in Progress)**
   - How many tasks in progress
   - Should be limited (2-3 per person)

5. **Bug Rate**
   - New bugs per week
   - Bug fix time

## Примеры использования

### Пример 1: Sprint Planning

```
1. Создайте Milestone "Sprint 5"
2. Фильтр в Project: No Milestone
3. Выберите 10-15 задач из Ready
4. Назначьте на Sprint 5
5. Назначьте на людей
6. Обсудите с командой
```

### Пример 2: Bug Triage

```
1. View: Bugs
2. Сортировка по Priority
3. Critical bugs → assign immediately
4. High bugs → add to current sprint
5. Medium/Low → backlog
```

### Пример 3: Feature Planning

```
1. Создайте epic issue для большой фичи
2. Разбейте на sub-tasks
3. Создайте отдельные issues для каждой задачи
4. Link к epic: "Part of #123"
5. Track progress в epic description
```

## Интеграция с фичами

У нас есть 36 фич в `backend/design-features/features/`. Для каждой фичи:

1. **Создайте Epic Issue**
   - Title: [Epic] Feature Name
   - Link to design doc
   - List of sub-tasks

2. **Создайте sub-tasks**
   - Backend API
   - Database migrations
   - Frontend UI
   - Tests
   - Documentation

3. **Назначьте ветку**
   - Branch: feature/001-auth-system
   - Все PR для этой фичи идут в эту ветку
   - Когда готово - merge в develop

4. **Track progress**
   - В Project Board
   - В Epic issue description
   - В Sprint milestone

## Полезные команды gh CLI

```bash
# Установите gh CLI
brew install gh

# Создать issue
gh issue create --title "Bug: Video upload fails" --body "..." --label bug

# Создать PR
gh pr create --title "feat: add comments" --body "..." --assignee @me

# Посмотреть issues
gh issue list --label bug --state open

# Посмотреть PRs
gh pr list --author @me

# Merge PR
gh pr merge 123 --squash --delete-branch
```

---

## Quick Start Checklist

- [ ] Создан GitHub Project "Nuum Development"
- [ ] Настроены колонки (Backlog, Ready, In Progress, In Review, Done)
- [ ] Добавлены custom fields (Priority, Size, Type, Component)
- [ ] Созданы views (Board, By Component, Current Sprint, Bugs)
- [ ] Добавлены issue templates
- [ ] Добавлен PR template
- [ ] Настроена автоматизация (GitHub Actions)
- [ ] Созданы первые issues для фич
- [ ] Назначены первые задачи команде

Теперь ваш GitHub Project готов к продуктивной работе! 🚀
