# План реализации: Загрузка видео (005)

## Frontend компоненты

### Pages
```
/frontend/app/
└── studio/
    └── upload/
        └── page.tsx              # Страница загрузки
```

### Компоненты
```
/frontend/components/upload/
├── UploadZone.tsx                # Drag & drop зона
├── UploadProgress.tsx            # Индикатор прогресса
├── VideoDetailsForm.tsx          # Форма метаданных
├── ThumbnailSelector.tsx         # Выбор превью
├── PrivacySettings.tsx           # Настройки приватности
└── ChapterEditor.tsx             # Редактор глав
```

### Hooks
```
/frontend/hooks/upload/
├── useVideoUpload.ts             # Chunked upload
├── useFormPersistence.ts         # Автосохранение
└── useProcessingStatus.ts        # Статус обработки
```

## Chunked upload

### Библиотека
```json
{
  "tus-js-client": "^3.1.0"
}
```

### Реализация
```typescript
const upload = new tus.Upload(file, {
  endpoint: '/api/v1/videos/upload',
  chunkSize: 5 * 1024 * 1024, // 5 MB
  retryDelays: [0, 3000, 5000, 10000],
  metadata: {
    filename: file.name,
    filetype: file.type,
  },
  onProgress: (uploaded, total) => {
    setProgress((uploaded / total) * 100);
  },
  onSuccess: () => {
    // Переход к редактированию метаданных
  },
});
```

## Backend обработка

### Handlers
```
/backend/internal/handlers/upload/
├── init.go                       # Инициация загрузки
├── chunk.go                      # Прием чанков
├── complete.go                   # Завершение загрузки
└── metadata.go                   # Обновление метаданных
```

### Worker для обработки
```
/backend/worker/
└── video_processor.go            # Фоновая обработка
```

### Очередь задач (Redis)
```
video:processing:{video_id}
```

## S3 Storage

### Структура хранения
```
/videos/{video_id}/
  ├── original/
  │   └── video.mp4
  ├── processed/
  │   ├── 1080p/
  │   ├── 720p/
  │   ├── 480p/
  │   └── 360p/
  ├── thumbnails/
  │   ├── default.jpg
  │   ├── auto_1.jpg
  │   └── auto_2.jpg
  └── subtitles/
      ├── ru.vtt
      └── en.vtt
```

## Время разработки: 12-18 дней

### Этап 1: Базовая загрузка (3-4 дня)
- Tus клиент
- Chunked upload backend
- S3 интеграция

### Этап 2: Форма метаданных (3-4 дня)
- Все поля формы
- Валидация
- Автосохранение

### Этап 3: Обработка видео (4-6 дней)
- FFmpeg pipeline
- Worker система
- Транскодинг

### Этап 4: UI/UX (2-4 дня)
- Превью редактор
- Прогресс индикаторы
- Error handling
