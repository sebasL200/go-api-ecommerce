
## 📋 Requisitos Previos

Asegúrate de tener instalado lo siguiente en tu sistema:

* **Go** (Versión 1.22 o superior).
* **PostgreSQL** (Servidor corriendo localmente).
* **Postman** (Recomendado para realizar pruebas).

---

## 🛠️ Instalación de Herramientas Globales

Para que los comandos de generación y migración funcionen, debes instalar estas herramientas una sola vez:

1. **Instalar Goose (Migraciones):**
```powershell
go install github.com/pressly/goose/v3/cmd/goose@latest

```


*Verifica con:* `goose -version`
2. **Instalar SQLC (Generador de Código):**
```powershell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

```


*Verifica con:* `sqlc version`

---

## 🚀 Primeros Pasos

Sigue estos pasos en el orden exacto para preparar tu entorno local:

1. **Bajar dependencias del proyecto:**
```powershell
go mod tidy

```


2. **Sincronizar la Base de Datos:**
> **Nota:** Antes de este paso, crea manualmente la base de datos llamada `ecom` en pgAdmin.


```powershell
# Entrar a la carpeta de migraciones
cd internal/adapters/postgresql/migrations

# Ejecutar la migración
goose postgres "host=127.0.0.1 user=postgres password=123 dbname=ecom sslmode=disable" up

# Volver a la raíz del proyecto
cd ../../../..

```


3. **Cargar Datos Iniciales (Seed):**
Si la API regresa un valor `null`, es por falta de registros. Ejecuta este script SQL en el **Query Tool** de pgAdmin para poblar el catálogo:
```sql
INSERT INTO products (id, name, price_in_centers, quantity, created_at)
VALUES 
    (1, 'Mouse Gamer RGB', 4500, 15, NOW()),
    (2, 'Teclado Mecánico Pro', 12500, 8, NOW()),
    (3, 'Monitor 24" 144Hz', 35000, 5, NOW()),
    (4, 'Headset Wireless', 8900, 12, NOW()),
    (5, 'Alfombrilla XL', 2500, 50, NOW());

```


4. **Generar Archivos SQLC:**
```powershell
sqlc generate

```



---

## 🏃 Ejecución del Proyecto

Para iniciar el servidor, ejecuta el siguiente comando desde la raíz:

```powershell
go run ./cmd

```

---

## 📡 EndPoints Disponibles

| Método | Ruta | Descripción |
| --- | --- | --- |
| **GET** | `/products` | Listar todos los productos del catálogo. |
| **GET** | `/products/{id}` | Obtener detalle de un producto específico (ej. `/products/1`). |
| **POST** | `/orders` | Registrar una nueva orden de compra. |

### Ejemplo de Request para Orden (POST):

**URL:** `http://localhost:8080/orders`

**Body (JSON):**

```json
{
    "customer_id": 1,
    "items": [
        {
            "product_id": 1,
            "quantity": 2
        }
    ]
}

```

---

