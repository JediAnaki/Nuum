# План реализации: Видеоплеер (002)

## Frontend компоненты

### Компонент плеера
```
/frontend/components/player/
├── VideoPlayer.tsx               # Главный компонент
├── PlayerControls.tsx            # Панель управления
├── ProgressBar.tsx               # Прогресс-бар с превью
├── QualitySelector.tsx           # Выбор качества
├── SubtitlesMenu.tsx             # Меню субтитров
├── SettingsMenu.tsx              # Настройки плеера
├── VolumeControl.tsx             # Громкость
└── PlayerOverlay.tsx             # Overlay элементы
```

### Hooks
```
/frontend/hooks/player/
├── useVideoPlayer.ts             # Основная логика плеера
├── usePlaybackProgress.ts        # Отслеживание прогресса
├── useQuality.ts                 # Управление качеством
├── useSubtitles.ts               # Субтитры
└── useKeyboardControls.ts        # Горячие клавиши
```

## Библиотеки

```json
{
  "video.js": "^8.10.0",
  "videojs-contrib-quality-levels": "^3.0.0",
  "hls.js": "^1.5.0",
  "@videojs/http-streaming": "^3.10.0"
}
```

## Backend обработка видео

### Video processing pipeline
```
/backend/internal/services/video/
├── processor.go                  # Главный процессор
├── transcoder.go                 # FFmpeg транскодинг
├── thumbnail_generator.go        # Генерация превью
└── hls_packager.go              # HLS packaging
```

### FFmpeg команды
```bash
# Транскодинг в разные качества
ffmpeg -i input.mp4 \
  -c:v libx264 -crf 23 -preset medium \
  -s 1920x1080 -b:v 5000k output_1080p.mp4

# HLS packaging
ffmpeg -i input.mp4 \
  -codec: copy -start_number 0 \
  -hls_time 10 -hls_list_size 0 \
  -f hls output.m3u8
```

## Время разработки: 10-14 дней

### Этап 1: Базовый плеер (3-4 дня)
- Интеграция Video.js
- Play/Pause, seek, volume
- Fullscreen

### Этап 2: Адаптивное качество (3-4 дня)
- HLS.js интеграция
- Переключение качества
- Backend транскодинг

### Этап 3: Продвинутые фичи (2-3 дня)
- Субтитры
- Главы
- Скорость воспроизведения
- PiP

### Этап 4: UX polish (2-3 дня)
- Горячие клавиши
- Превью кадров
- Анимации
- Accessibility
