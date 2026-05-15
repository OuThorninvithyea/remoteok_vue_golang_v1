<div align="center">
  <img src="./frontend/public/images/logo.png" alt="RemoteOK Logo" width="120" />
  <h1>RemoteOK Vue + Go</h1>
  <p><strong>Full-Stack Remote Job Board Platform</strong></p>

  <p>
    <a href="https://cv-ou-thorninvithyea-fvkg.vercel.app/">
      <img src="https://img.shields.io/badge/Portfolio-Ou%20Thorninvithyea-blue?style=flat-square" alt="Portfolio" />
    </a>
    <img src="https://img.shields.io/badge/Vue-3.5-4FC08D?style=flat-square&logo=vue.js" alt="Vue 3" />
    <img src="https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go" alt="Go" />
    <img src="https://img.shields.io/badge/Vite-8.0-646CFF?style=flat-square&logo=vite" alt="Vite 8" />
    <img src="https://img.shields.io/badge/Status-Development-yellow?style=flat-square" alt="Status" />
  </p>
</div>

---

## Overview

**RemoteOK Vue + Go** is a full-stack remote job board platform featuring a Vue 3 frontend and a Go (Gin) API backend. It provides job seekers with a premium browsing experience with advanced filtering, while offering employers a complete job posting workflow with payment integration.

The backend serves structured API endpoints with mock data, and the frontend consumes them via composable data hooks for a clean separation of concerns.

## Key Features

### Frontend
- **Job Discovery** — Browse jobs with multi-criteria filtering (benefits, location, salary, sort)
- **Advanced Post-a-Job Flow** — Multi-step wizard: company info, job design, details, preview, payment
- **Trust & Social Proof** — Testimonials, trusted companies, statistical data
- **Responsive UI** — Mobile-optimized with reusable component architecture

### Backend (Go)
- **RESTful API** — Organized route handlers with Gin framework
- **Data Stores** — In-memory data stores for all entities (jobs, categories, dropdowns, etc.)
- **Structured Models** — Clean separation of models, handlers, stores, and routes
- **Login Support** — Basic authentication flow

## Tech Stack

| Layer | Technology |
|-------|-----------|
| **Frontend** | Vue 3 (Composition API) + Vite 8 + JavaScript |
| **Backend** | Go 1.23 + Gin |
| **Architecture** | RESTful API with in-memory data stores |

## Project Structure

```
remoteok_vue_golang_v1/
├── frontend/                    # Vue 3 SPA
│   ├── src/
│   │   ├── components/          # UI components organized by page
│   │   │   ├── homePage/        # Job listing, filters, header
│   │   │   ├── postJob/         # Post-a-job flow components
│   │   │   ├── login/           # Login form
│   │   │   └── icons/           # SVG icon components
│   │   ├── composables/         # API data hooks
│   │   ├── layouts/             # Main layout
│   │   ├── router/              # Vue Router config
│   │   ├── styles/              # Page-level CSS
│   │   └── views/               # Page-level view components
│   └── ...
│
└── backend/                     # Go API server
    ├── main.go                  # Entry point
    ├── go.mod / go.sum          # Go dependencies
    ├── handler/                 # HTTP request handlers
    │   ├── home/                # Job listings, categories, dropdowns
    │   ├── login/               # Authentication
    │   └── postJob/             # Job posting flow
    ├── model/                   # Data models
    ├── store/                   # In-memory data stores
    └── route/                   # Route registrations
```

## Getting Started

### Backend

```bash
cd backend
go run main.go

# The API server starts on http://localhost:8080
```

### Frontend

```bash
cd frontend
npm install
npm run dev

# Dev server starts on http://localhost:5173
```

### Building for Production

```bash
cd frontend
npm run build    # Output in ./frontend/dist
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/jobs` | Fetch job listings |
| `GET` | `/api/categories` | Job categories |
| `GET` | `/api/home/header` | Header data |
| `GET` | `/api/post-job/*` | Post-a-job flow data |
| `POST` | `/api/login` | User authentication |

## Author

**Ou Thorninvithyea**

- 🌐 [Portfolio](https://cv-ou-thorninvithyea-fvkg.vercel.app/)
- 🐙 [GitHub](https://github.com/OuThorninvithyea)
- 📧 Vithyeasa@gmail.com
- 📍 Phnom Penh, Cambodia

> Software Engineer with expertise in full-stack web development, performance optimization, and building scalable applications with modern frameworks.

---

<div align="center">
  <sub>Built with ❤️ by Ou Thorninvithyea</sub>
</div>
