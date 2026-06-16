# 009 - Система уведомлений

**Приоритет:** MVP (009)
**Категория:** Взаимодействие и коммуникация
**Зависимости:** 001 (Авторизация), 008 (Каналы)

## Описание

Центр уведомлений для информирования пользователей о новых видео, комментариях, лайках и других важных событиях.

## Функциональные требования

### 1. Типы уведомлений

#### Категории:
1. **Новый контент** - новые видео от подписок
2. **Социальные** - лайки, комментарии, упоминания
3. **Система** - обновления платформы, новости
4. **Персональные** - достижения, рекомендации

### 2. Иконка уведомлений

```
[🔔 3]  ← Badge с количеством непрочитанных
```

**Клик открывает dropdown:**
```
┌─────────────────────────────────────┐
│ Уведомления        [⚙️] [✓ Все прочитано]│
├─────────────────────────────────────┤
│ 🎥 Новое видео                      │
│ [Thumbnail] "Название видео"        │
│ Канал Name • 2 часа назад      [×] │
├─────────────────────────────────────┤
│ 💬 Ответ на комментарий             │
│ @username ответил на ваш...         │
│ 5 часов назад                  [×] │
├─────────────────────────────────────┤
│ 👍 Новый лайк                       │
│ Ваш комментарий понравился...       │
│ Вчера                          [×] │
├─────────────────────────────────────┤
│ [Показать все уведомления]          │
└─────────────────────────────────────┘
```

### 3. Страница уведомлений

**URL:** `/notifications`

**Фильтры:**
```
[ Все ] [ Непрочитанные ] [ Упоминания ]
```

**Действия:**
- Отметить все как прочитанные
- Удалить все уведомления
- Настройки уведомлений

### 4. Push уведомления

- **Браузерные push** (Web Push API)
- **Запрос разрешения** при первом входе
- **Настройки** - выбор типов push уведомлений

### 5. Email уведомления

**Типы email:**
- Дайджест новых видео (дневной/недельный)
- Важные уведомления (упоминания, ответы)
- Новости платформы
- Отписка от каждого типа отдельно

## UI/UX требования

**Особенности:**
- Real-time обновление (WebSocket)
- Группировка похожих уведомлений
- Умные уведомления (не спамить)
- Quiet hours (не беспокоить ночью)

## Технические детали

### API endpoints
```typescript
GET /api/v1/notifications
PATCH /api/v1/notifications/{id}/read
PATCH /api/v1/notifications/read-all
DELETE /api/v1/notifications/{id}

GET /api/v1/notifications/settings
PATCH /api/v1/notifications/settings

POST /api/v1/notifications/subscribe-push
DELETE /api/v1/notifications/unsubscribe-push
```

### WebSocket
```typescript
// Подписка на уведомления
ws://platform.ru/ws/notifications
{
  "event": "new_notification",
  "data": { ... }
}
```

### Типы данных
```typescript
interface Notification {
  id: string;
  type: "new_video" | "comment_reply" | "mention" | "like" | "system";
  title: string;
  body: string;
  thumbnail?: string;
  link: string;
  is_read: boolean;
  created_at: string;
  actor?: {
    id: string;
    username: string;
    avatar: string;
  };
}

interface NotificationSettings {
  email: {
    new_videos: "instant" | "daily" | "weekly" | "never";
    comments: boolean;
    mentions: boolean;
    system: boolean;
  };
  push: {
    new_videos: boolean;
    comments: boolean;
    mentions: boolean;
    live: boolean;
  };
  quiet_hours: {
    enabled: boolean;
    start: string; // "22:00"
    end: string; // "08:00"
  };
}
```

## Метрики успеха
- CTR на уведомления > 30%
- Opt-in rate push > 40%
- Unsubscribe rate < 5%
- Среднее время до прочтения < 1 час

## Особенности для русскоязычной аудитории
- Локализованные тексты
- Учет часовых поясов
- Русские форматы времени ("2 часа назад", "вчера")
