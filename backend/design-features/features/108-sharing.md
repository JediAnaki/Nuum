# 108 - Поделиться (Sharing)

**Приоритет:** Высокий (108)
**Категория:** Виральность и распространение
**Зависимости:** 004 (Страница видео)

## Функции

### Кнопка "Поделиться"
Модальное окно с опциями:
```
─────── Поделиться ───────
📱 Социальные сети:
[VK] [Telegram] [WhatsApp] [Viber]
[OK] [Twitter] [Facebook]

🔗 Ссылка:
https://platform.ru/watch?v=abc123
☐ Начать с 1:23
[Копировать]

📧 Email:
[Email]

📋 Embed код:
[Получить код]

📱 QR-код:
[███████]
```

### Timestamp sharing
- Поделиться с текущего момента
- Автодобавление ?t=83

### Социальные сети
- Нативный sharing (Web Share API)
- Предзаполненный текст
- Open Graph метаданные

## API
```typescript
POST /api/v1/videos/{id}/share
Body: { platform: string, timestamp?: number }
```

## Метрики
- Share rate > 5%
- Топ платформы: Telegram, VK, WhatsApp
