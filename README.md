# expense-api

API de gastos personales en Go (Gin + PostgreSQL + Cognito), hecha como ejercicio de aprendizaje.

## Requisitos

- Docker Desktop y VS Code con la extension Dev Containers
- Un user pool y un app client de Cognito propios de esta practica (flujo `ALLOW_USER_PASSWORD_AUTH`, sin client secret)

## Primer arranque

1. `cp .env.example .env` y completa `COGNITO_*`.
2. VS Code: **Dev Containers: Reopen in Container**.
3. Dentro del contenedor:

```bash
make migrate-up
make run        # o `make dev` para recarga con air
curl -i localhost:8080/healthz
```

## Variables de entorno

| Variable | Obligatoria | Default |
|---|---|---|
| `DATABASE_URL` | si | |
| `COGNITO_REGION` | si | |
| `COGNITO_USER_POOL_ID` | si | |
| `COGNITO_APP_CLIENT_ID` | si | |
| `PORT` | no | `8080` |
| `DB_MAX_CONNS` / `DB_MIN_CONNS` | no | `10` / `2` |
| `DB_MAX_CONN_LIFETIME` / `DB_MAX_CONN_IDLE_TIME` | no | `30m` / `5m` |

Los cambios en `.env` requieren recrear el contenedor del devcontainer.

## Comandos

| Comando | Que hace |
|---|---|
| `make run` / `make dev` | Ejecuta la API / con recarga (air) |
| `make fmt` / `make vet` / `make lint` | Formato, vet y golangci-lint |
| `make migrate-up` / `migrate-down` / `migrate-status` | Migraciones con goose |
| `make migrate-create name=add_algo` | Crea una migracion nueva |
| `make migrate-reset` | Revierte todo y reaplica (solo DB local) |

## Probar la autenticacion

```bash
TOKEN=$(aws cognito-idp initiate-auth \
  --auth-flow USER_PASSWORD_AUTH \
  --client-id "$COGNITO_APP_CLIENT_ID" \
  --auth-parameters USERNAME=tu@email.com,PASSWORD='tu-password' \
  --query 'AuthenticationResult.AccessToken' --output text)

curl -i -H "Authorization: Bearer $TOKEN" localhost:8080/api/v1/me
```

La API valida el access token de Cognito (firma, issuer, expiracion, `token_use` y `client_id`).

## Estructura

```
cmd/api/            punto de entrada y cableado
internal/config/    variables de entorno
internal/database/  pool de pgx
internal/model/     entidades
internal/repository/ SQL
internal/service/   reglas de negocio
internal/handler/   HTTP (Gin)
internal/middleware/ autenticacion
internal/apperr/    errores de dominio
migrations/         migraciones goose (un archivo por migracion)
infra/              Terraform
```

Dependencias: `handler -> service -> repository -> database`.

## Problemas conocidos

**El devcontainer falla con `additional privileges requested: pass "--allow=fs.read=..."`**
Es el chequeo de permisos de archivos de `docker buildx bake`, que choca con el Dockerfile temporal que genera VS Code. En Windows:

```powershell
setx BUILDX_BAKE_ENTITLEMENTS_FS 0
```

Cierra VS Code por completo (y reinicia Docker Desktop) y vuelve a abrir el contenedor.
