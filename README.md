# Transaction Management System

A full-stack web application for managing financial transactions through CSV file uploads. The system processes transaction data, calculates balances, and displays unsuccessful transactions.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Getting Started](#getting-started)

## Features

- **CSV Upload**: Upload transaction data via CSV files (max 10MB)
- **Balance Calculation**: Real-time balance calculation based on successful transactions
- **Issue Tracking**: Display failed and pending transactions
- **Transaction List**: Sortable table view of unsuccessful transactions

## Tech Stack

### Backend
- **Language**: Go 1.24+
- **Web Framework**: Standard `net/http`

### Frontend
- **Language**: TypeScript
- **Framework**: React 19
- **Build Tool**: Vite
- **State Management**: TanStack Query

## Architecture

### Backend Architecture

```
backend/
├── cmd/app/              # Application entry point
├── internal/
│   ├── domain/           # Business entities
│   ├── repository/       # Data storage layer (in-memory)
│   ├── service/          # Business logic
│   ├── handler/          # HTTP handlers
│   ├── parser/           # CSV parsing logic
│   └── middleware/       # HTTP middleware
```

**Layer Responsibilities:**

1. **Domain Layer**: Defines core business entities and constants
   - Transaction types (DEBIT, CREDIT)
   - Transaction statuses (SUCCESS, PENDING, FAILED)

2. **Repository Layer**: Manages data persistence
   - In-memory storage with concurrent access (RWMutex)
   - Thread-safe operations

3. **Service Layer**: Contains business logic
   - Transaction validation
   - Balance calculation
   - Filtering unsuccessful transactions

4. **Handler Layer**: HTTP request/response handling
   - File upload processing
   - JSON response formatting
   - Error handling


### Frontend Architecture

```
frontend/
├── src/
│   ├── components/          # Reusable UI components
│   ├── modules/
│   │   └── Transaction/     # Transaction feature module
│   ├── libs/
│   │   ├── hooks/           # Custom React hooks
│   │   ├── services/        # API service layer
│   │   ├── types/           # TypeScript types
│   │   └── helpers/         # Utility functions
│   └── styles/              # Global styles
```

## Getting Started

### Prerequisites

- **Go**: 1.24 or higher
- **Node.js**: 22 or higher
- **pnpm**: 9 or higher

### Backend Setup

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the server:**
   ```bash
   go run cmd/app/main.go
   ```

   The server will start on `http://localhost:8080`

### Frontend Setup

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install dependencies:**
   ```bash
   pnpm install
   ```

3. **Run development server:**
   ```bash
   pnpm dev
   ```

   The app will start on `http://localhost:5173`


Created for Flip Technical Assessment
