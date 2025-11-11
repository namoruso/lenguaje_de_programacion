# ☕ Coffee House - Sistema de Gestión de Cafetería

Proyecto de aprendizaje Full Stack con Vue.js y Express.js

## 🚀 Características

- Sistema de autenticación (Login/Registro)
- Menú interactivo de cafés con imágenes
- Backend con Express.js y almacenamiento en JSON
- Frontend con Vue.js 3 (Composition API)
- Diseño responsive con temática de cafetería

## 🛠️ Tecnologías

### Frontend

- Vue.js 3
- Vite
- Vue Router
- Axios

### Backend

- Node.js
- Express.js
- CORS
- Body Parser

## 📦 Instalación

### 1. Backend

```bash
cd backend
npm install
npm start
```

El servidor correrá en `http://localhost:3000`

### 2. Frontend

```bash
cd frontend
npm install
npm run dev
```

La aplicación correrá en `http://localhost:5173`

## 📁 Estructura del Proyecto

```
coffee-shop/
├── backend/
│   ├── server.js
│   ├── routes/
│   │   ├── auth.js
│   │   └── coffees.js
│   └── data/
│       ├── users.json
│       └── coffees.json
└── frontend/
    ├── src/
    │   ├── views/
    │   │   ├── Login.vue
    │   │   ├── Register.vue
    │   │   └── CoffeeMenu.vue
    │   ├── router/
    │   ├── services/
    │   └── main.js
    └── package.json
```

## 🎨 Características del Diseño

- Paleta de colores temática de café (marrones, cremas, dorados)
- Animaciones suaves y transiciones
- Sistema de calificación con estrellas
- Badges "Popular" en productos destacados
- Diseño responsive para móviles

## 🔐 Notas de Seguridad

⚠️ Este es un proyecto de aprendizaje. Las contraseñas NO están encriptadas.

Para producción, se recomienda:

- Usar bcrypt para encriptar contraseñas
- Implementar JWT para autenticación
- Usar una base de datos real (MongoDB, PostgreSQL)
- Variables de entorno para configuración sensible

## 📝 Flujo de Usuario

1. Registro con email y contraseña
2. Inicio de sesión
3. Visualización del menú de cafés
4. Selección y orden de productos

## 📄 Licencia

MIT
