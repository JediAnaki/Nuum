# План реализации: Система аутентификации (001)

## Frontend компоненты

### Pages (Next.js App Router)
```
/frontend/app/
├── auth/
│   ├── login/
│   │   └── page.tsx              # Страница входа
│   ├── register/
│   │   └── page.tsx              # Страница регистрации
│   ├── forgot-password/
│   │   └── page.tsx              # Восстановление пароля
│   ├── reset-password/
│   │   └── page.tsx              # Сброс пароля (с токеном)
│   └── verify-email/
│       └── page.tsx              # Подтверждение email
```

### Компоненты
```
/frontend/components/auth/
├── LoginForm.tsx                 # Форма входа
├── RegisterForm.tsx              # Форма регистрации
├── ForgotPasswordForm.tsx        # Форма восстановления
├── OAuthButtons.tsx              # Кнопки OAuth (Google, VK, Yandex)
├── TwoFactorModal.tsx            # Модальное окно 2FA
├── PasswordStrengthIndicator.tsx # Индикатор силы пароля
└── SessionManager.tsx            # Управление сессиями
```

### Hooks
```
/frontend/hooks/
├── useAuth.ts                    # Главный хук аутентификации
├── useLogin.ts                   # Логика входа
├── useRegister.ts                # Логика регистрации
├── useSession.ts                 # Управление сессией
└── useOAuth.ts                   # OAuth логика
```

### State Management (Zustand)
```
/frontend/lib/stores/
└── authStore.ts                  # Глобальное состояние auth
```

```typescript
interface AuthStore {
  user: User | null;
  accessToken: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (credentials: LoginCredentials) => Promise<void>;
  logout: () => Promise<void>;
  register: (data: RegisterData) => Promise<void>;
  refreshToken: () => Promise<void>;
}
```

### API Layer (React Query)
```
/frontend/lib/api/
└── auth.ts                       # API функции для auth
```

```typescript
export const authApi = {
  login: (data: LoginCredentials) => axios.post('/auth/login', data),
  register: (data: RegisterData) => axios.post('/auth/register', data),
  logout: () => axios.post('/auth/logout'),
  refreshToken: () => axios.post('/auth/refresh'),
  verifyEmail: (token: string) => axios.get(`/auth/verify-email?token=${token}`),
  // ... остальные
};
```

### Утилиты
```
/frontend/lib/utils/
├── validation.ts                 # Валидация форм (Zod schemas)
├── tokens.ts                     # Работа с JWT
└── passwordStrength.ts           # Проверка силы пароля
```

## Backend endpoints (Go)

### HTTP handlers
```
/backend/internal/handlers/auth/
├── login.go                      # POST /api/v1/auth/login
├── register.go                   # POST /api/v1/auth/register
├── logout.go                     # POST /api/v1/auth/logout
├── refresh.go                    # POST /api/v1/auth/refresh
├── verify_email.go               # GET /api/v1/auth/verify-email
├── forgot_password.go            # POST /api/v1/auth/forgot-password
├── reset_password.go             # POST /api/v1/auth/reset-password
├── oauth.go                      # OAuth handlers
└── sessions.go                   # Управление сессиями
```

### Services
```
/backend/internal/services/auth/
├── auth_service.go               # Основной auth сервис
├── jwt_service.go                # Работа с JWT
├── oauth_service.go              # OAuth интеграции
├── email_service.go              # Отправка email
└── session_service.go            # Управление сессиями
```

### Models
```
/backend/internal/models/
├── user.go                       # User модель
├── session.go                    # Session модель
└── oauth_account.go              # OAuth account модель
```

### Middleware
```
/backend/internal/middleware/
├── auth.go                       # Auth middleware
├── rate_limit.go                 # Rate limiting
└── cors.go                       # CORS настройки
```

## База данных (PostgreSQL)

### Migrations
```sql
-- users table
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(30) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  email_verified BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

-- sessions table
CREATE TABLE sessions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  refresh_token VARCHAR(255) UNIQUE NOT NULL,
  device_info JSONB,
  ip_address INET,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

-- oauth_accounts table
CREATE TABLE oauth_accounts (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  provider VARCHAR(50) NOT NULL,
  provider_user_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(provider, provider_user_id)
);
```

## Библиотеки

### Frontend
```json
{
  "dependencies": {
    "@tanstack/react-query": "^5.28.0",
    "zustand": "^4.5.0",
    "axios": "^1.6.0",
    "zod": "^3.22.0",
    "react-hook-form": "^7.50.0",
    "@hookform/resolvers": "^3.3.0"
  }
}
```

### Backend
```go
require (
    github.com/golang-jwt/jwt/v5
    github.com/go-chi/chi/v5
    github.com/go-chi/cors
    golang.org/x/crypto/bcrypt
    golang.org/x/oauth2
)
```

## Порядок разработки

### Этап 1: Backend основа (3-5 дней)
1. Создать модели и миграции БД
2. Реализовать JWT сервис
3. Реализовать auth сервис (register, login)
4. Настроить middleware
5. Тестирование

### Этап 2: Frontend базовая auth (3-4 дня)
1. Создать store и hooks
2. Реализовать формы (Login, Register)
3. Интеграция с API
4. Валидация форм
5. Error handling

### Этап 3: OAuth (2-3 дня)
1. Backend OAuth flow
2. Frontend OAuth кнопки
3. Интеграция провайдеров
4. Тестирование

### Этап 4: Дополнительные функции (2-3 дня)
1. Email verification
2. Password reset
3. 2FA (опционально в MVP)
4. Session management

### Этап 5: Тестирование и доработка (2 дня)
1. E2E тесты
2. Security audit
3. UX улучшения
4. Документация

**Общее время: 12-17 дней**

## Метрики готовности
- [ ] Unit тесты покрытие > 80%
- [ ] E2E тесты для критичных путей
- [ ] Security audit пройден
- [ ] Документация API готова
- [ ] UX review пройден
