FROM mcr.microsoft.com/devcontainers/go:1.21

RUN apt-get update && apt-get install -y \
    postgresql postgresql-contrib libpq-dev

RUN service postgresql start && \
    sudo -u postgres psql -c "CREATE USER postgres WITH SUPERUSER PASSWORD 'secret';" && \
    sudo -u postgres createdb students_db

WORKDIR /workspace