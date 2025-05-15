# ZycoNauth 🤖🧠

ZycoNauth es un bot automatizado escrito en Go que utiliza la API de Gemini de Google para generar imágenes artísticas con inteligencia artificial y publica una imagen junto con un prompt creativo cada día en X (anteriormente Twitter).

---

## ✨ Características

- ✅ Generación diaria de prompts visuales usando Gemini.
- 🖼️ Creación de imágenes AI con prompts artísticos.
- 📤 Publicación automática del contenido en X/Twitter.
- 🔒 Autenticación segura con OAuth 1.0a (HMAC-SHA1).
- 🕐 Ejecutado diariamente mediante `cron`.

---

## ⚙️ Requisitos

- Go 1.20 o superior
- Cuenta de desarrollador en [X (Twitter)](https://developer.x.com/)
- Acceso a [Gemini API de Google](https://ai.google.dev/)
- Archivo `.env` con tus credenciales

---

## 🔐 Variables de entorno

Crea un archivo `.env` con el siguiente contenido:

```env
GOOGLE_API_KEY=tu_api_key_de_gemini
CONSUMER_KEY=tu_consumer_key_x
CONSUMER_SECRET=tu_consumer_secret_x
ACCESS_TOKEN=tu_access_token_x
TOKEN_SECRET=tu_token_secret_x
```

---

## 🚀 Uso

1. Clona el repositorio:

```bash
git clone https://github.com/atrox39/ZycoNauth
cd zyconauth
```

2. Instala las dependencias:

```bash
go mod tidy
```

3. Ejecuta manualmente para probar:

```bash
go run main.go
```

4. Compila el ejecutable:

```bash
go build -o zyconauth main.go
```

---

## 🕒 Automatización con Cron

1. Crea un script de ejecución:

```bash
#!/bin/bash
export $(grep -v '^#' /ruta/absoluta/.env | xargs)
cd /ruta/absoluta
./zyconauth
```

Guarda esto como `run_zyconauth.sh` y dale permisos:

```bash
chmod +x run_zyconauth.sh
```

2. Edita tu `crontab`:

```bash
crontab -e
```

3. Agrega esta línea para ejecutar el bot diariamente a las 10 AM:

```bash
0 10 * * * /ruta/absoluta/run_zyconauth.sh >> /ruta/absoluta/cron.log 2>&1
```

---

## 📂 Estructura del proyecto

```
zyconauth/
├── main.go
├── util/
│   ├── gemini.go      // Lógica para generación de prompts e imágenes
├── .env
├── go.mod
├── README.md
```

---

## 📸 Resultado

El bot publica diariamente algo como esto en X:

> _“Una ciudad flotante en el cielo al atardecer, arte digital estilo Studio Ghibli”_  
> (imagen adjunta generada por IA)

---

## 🧑‍💻 Autor

Desarrollado por [atrox39](https://github.com/atrox39)

---

## 📝 Licencia

Este proyecto está licenciado bajo la MIT License. Consulta el archivo [LICENSE](LICENSE) para más información.
