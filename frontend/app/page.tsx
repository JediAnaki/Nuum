import Link from 'next/link';

export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-b from-gray-900 to-black">
      <div className="container mx-auto px-4 py-16">
        <div className="text-center">
          <h1 className="text-6xl font-bold text-white mb-6">
            Видеоплатформа
          </h1>
          <p className="text-xl text-gray-300 mb-12 max-w-2xl mx-auto">
            Современная платформа для размещения и просмотра видео.
            Создана специально для русскоязычной аудитории.
          </p>

          <div className="flex gap-4 justify-center">
            <Link
              href="/videos"
              className="bg-primary-600 hover:bg-primary-700 text-white px-8 py-3 rounded-lg font-semibold transition"
            >
              Смотреть видео
            </Link>
            <Link
              href="/auth/register"
              className="bg-gray-700 hover:bg-gray-600 text-white px-8 py-3 rounded-lg font-semibold transition"
            >
              Зарегистрироваться
            </Link>
          </div>
        </div>

        <div className="mt-24 grid md:grid-cols-3 gap-8">
          <div className="bg-gray-800 p-6 rounded-lg">
            <h3 className="text-2xl font-bold text-white mb-3">
              Быстрая загрузка
            </h3>
            <p className="text-gray-300">
              Оптимизированная инфраструктура для быстрой загрузки и обработки видео
            </p>
          </div>

          <div className="bg-gray-800 p-6 rounded-lg">
            <h3 className="text-2xl font-bold text-white mb-3">
              Высокое качество
            </h3>
            <p className="text-gray-300">
              Поддержка разрешений до 4K с адаптивным стримингом
            </p>
          </div>

          <div className="bg-gray-800 p-6 rounded-lg">
            <h3 className="text-2xl font-bold text-white mb-3">
              Монетизация
            </h3>
            <p className="text-gray-300">
              Справедливые условия для авторов - больше заработка
            </p>
          </div>
        </div>
      </div>
    </main>
  );
}
