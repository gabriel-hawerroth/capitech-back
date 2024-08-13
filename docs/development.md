# Configurando envs

Criar um arquivo `.env` dentro da pasta `cmd/capitech/` com as seguintes envs:

-   DB_DRIVER
-   DB_HOST
-   DB_PORT
-   DB_USER
-   DB_PASSWORD
-   DB_NAME
-   WEB_SERVER_PORT

# DB migrations

Lib [golang-migrate](https://github.com/golang-migrate/migrate)

Instalando o cli da lib no linux:

```bash
wget https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz

tar -xvzf migrate.linux-amd64.tar.gz

chmod +x migrate

sudo mv migrate /usr/local/bin/
```

Instalando o cli da lib no windows:

```ps1
curl -L -o migrate.zip https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.windows-amd64.zip

Expand-Archive -Path migrate.zip -DestinationPath C:\migrate

$env:Path += ";C:\migrate"

[System.Environment]::SetEnvironmentVariable("Path", $env:Path, [System.EnvironmentVariableTarget]::Machine)
```

Verifique a instalação com o comando `migrate -version`

Comando cli para criar uma migration:

```bash
migrate create -ext sql -dir db/migrations -seq migration_name
```

Serão criados dois arquivos, um `.up` que deve conter a migração e outro `.down` com o rollback em caso de falha.

A lib criará uma tabela nomeada `schema_migrations` e as migrações serão executadas automaticamente ao subir a aplicação.

! Não commitar migrações com erro, em caso de falha, altere o atributo `version` para o número da última migração aplicada com sucesso e marque o atributo `dirty` como false !
