# VK-Service

Здравствуйте! Этот проект представляет собой систему мониторинга состояния Docker-контейнеров. Он позволяет отслеживать доступность контейнеров через пинг и сохранять результаты в базе данных PostgreSQL. Данные о состоянии контейнеров отображаются в виде таблицы на динамически формируемой веб-странице.

Проект состоит из четырех основных сервисов:

**Backend** : RESTful API, обеспечивающее взаимодействие с базой данных.

**Frontend** : Веб-интерфейс, написанный на React + TypeScript, для отображения данных.

**Database (PostgreSQL)** : Хранение информации о контейнерах.

**Pinger** : Сервис, который периодически пингует контейнеры и отправляет данные в Backend.

# Функциональность
**Постоянный мониторинг контейнеров** :

Pinger-сервис получает список всех Docker-контейнеров.
Каждые 10 секунд он пингует каждый контейнер и отправляет результаты в Backend.

**Хранение данных** :
Информация о контейнерах (IP-адрес, время пинга, дата последнего успешного пинга) сохраняется в PostgreSQL.

**Веб-интерфейс** :
Frontend-приложение отображает данные из Backend в виде таблицы.

**Таблица содержит следующие столбцы**:

IP-адрес контейнера.

Время пинга (в миллисекундах).

Дата последнего успешного пинга.
  
# Требования

Для запуска проекта необходимы следующие инструменты:

**Docker** : Версия 20.x или выше.

**Docker Compose** : Версия 2.x или выше.

**Git** : Для клонирования репозитория.

# Установка и запуск

**Клонируйте репозиторий** :

Copy git clone https://github.com/Digtya/VK-Service.git

cd container-monitoring

**Соберите и запустите все сервисы** :

Copy docker-compose up --build

**Откройте веб-интерфейс** :

После успешной сборки откройте браузер и перейдите по адресу http://localhost:3000 . Вы должны увидеть таблицу со статусом контейнеров.


