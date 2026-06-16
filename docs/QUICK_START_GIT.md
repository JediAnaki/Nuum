# Quick Start - Git Setup

## Запуск всех веток на GitHub

У вас уже созданы все ветки локально. Теперь нужно запушить их на GitHub.

### Вариант 1: Запушить все ветки сразу

```bash
# Запушить все локальные ветки на remote
git push origin --all
```

Это запушит:
- main
- develop
- 36 веток фич (feature/001-auth-system, feature/002-video-player, и т.д.)

### Вариант 2: Запушить выборочно

Если хотите запушить только нужные ветки:

```bash
# Только main и develop
git push -u origin main
git push -u origin develop

# Конкретную фичу
git push -u origin feature/001-auth-system

# Несколько фич
git push -u origin feature/001-auth-system feature/002-video-player
```

### Вариант 3: Запушить по группам

```bash
# Базовые фичи (001-010)
for branch in $(git branch | grep 'feature/00[0-9]'); do
  git push -u origin $branch
done

# Продвинутые фичи (100-110)
for branch in $(git branch | grep 'feature/10[0-9]'); do
  git push -u origin $branch
done

# И так далее...
```

## Проверка

После push проверьте на GitHub:

```bash
# Посмотреть все ветки
git branch -a

# Или в браузере
open https://github.com/JediAnaki/Nuum/branches
```

Вы должны увидеть 38 веток:
- main
- develop
- 36 feature веток

## Настройка branch protection

После push настройте защиту веток:

### Для main:

1. Идите в Settings → Branches → Add rule
2. Branch name pattern: `main`
3. Включите:
   - ✅ Require a pull request before merging
   - ✅ Require approvals (2)
   - ✅ Dismiss stale pull request approvals
   - ✅ Require status checks to pass
   - ✅ Require branches to be up to date
   - ✅ Require conversation resolution
   - ✅ Include administrators

### Для develop:

1. Branch name pattern: `develop`
2. Включите:
   - ✅ Require a pull request before merging
   - ✅ Require approvals (1)
   - ✅ Require status checks to pass
   - ✅ Require conversation resolution

## Следующие шаги

После push веток:

1. **Создайте GitHub Project** (см. `docs/GITHUB_PROJECTS_SETUP.md`)
2. **Создайте первые issues** для каждой фичи
3. **Настройте CI/CD** (GitHub Actions)
4. **Назначьте задачи** команде

## Troubleshooting

### Ошибка: "Connection closed by remote host"

Это может быть проблема с SSH ключами:

```bash
# Проверьте SSH ключ
ssh -T git@github.com

# Должно вывести:
# Hi JediAnaki! You've successfully authenticated...

# Если ошибка, добавьте SSH ключ:
ssh-keygen -t ed25519 -C "your_email@example.com"
cat ~/.ssh/id_ed25519.pub
# Добавьте ключ в GitHub Settings → SSH Keys
```

### Ошибка: "Repository not found"

```bash
# Проверьте remote URL
git remote -v

# Должно быть:
# origin  git@github.com:JediAnaki/Nuum.git (fetch)
# origin  git@github.com:JediAnaki/Nuum.git (push)

# Если нет, исправьте:
git remote set-url origin git@github.com:JediAnaki/Nuum.git
```

### Ошибка: "Permission denied"

Убедитесь что у вас есть права на push в репозиторий.

1. Вы должны быть owner или collaborator
2. SSH ключ должен быть добавлен в ваш GitHub аккаунт
3. Репозиторий должен существовать на GitHub

### Push занимает слишком много времени

Если у вас много больших файлов:

```bash
# Проверьте размер репозитория
du -sh .git

# Если больше 100MB, рассмотрите:
# - Добавьте большие файлы в .gitignore
# - Используйте Git LFS для больших файлов
# - Очистите историю от случайно закоммиченных файлов
```

## Полезные команды

```bash
# Посмотреть все ветки (локальные и remote)
git branch -a

# Посмотреть только remote ветки
git branch -r

# Удалить локальную ветку
git branch -d feature/branch-name

# Удалить remote ветку
git push origin --delete feature/branch-name

# Переименовать ветку
git branch -m old-name new-name
git push origin :old-name new-name
git push origin -u new-name

# Синхронизировать все ветки
git fetch --all --prune
```

## Best Practices

1. **Регулярно push изменения**
   - Не ждите долго
   - Делайте backup на remote
   - Команда видит ваш прогресс

2. **Используйте meaningful commit messages**
   - Следуйте Conventional Commits
   - Облегчает понимание изменений
   - Помогает в code review

3. **Синхронизируйтесь с develop**
   ```bash
   git fetch origin develop
   git rebase origin/develop
   ```

4. **Не push в main напрямую**
   - Всегда через PR
   - Code review обязателен
   - CI checks должны пройти

5. **Удаляйте merged ветки**
   ```bash
   # После merge в GitHub, локально:
   git checkout develop
   git pull
   git branch -d feature/merged-branch
   ```

---

Готово! Теперь ваш репозиторий полностью настроен на GitHub. 🚀
