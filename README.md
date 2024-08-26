

![Warehaus Logo](https://github.com/raeeceip/werehaus/raw/main/logo.png)

Warehaus is a cutting-edge warehouse management system built with Go and Astro, designed to streamline your inventory operations and boost efficiency. This project combines the power of Go's backend performance with Astro's modern frontend capabilities to deliver a robust, scalable solution for warehouse management.

## 🚀 Features

- **Package Tracking**: Real-time tracking of all packages within your warehouse, including location, status, and history.
- **Location Management**: Easily add, edit, and manage location information for optimal space utilization and quick item retrieval.
- **Inventory Control**: Keep track of stock levels, set reorder points, and receive low stock alerts to maintain optimal inventory levels.
- **User-friendly Interface**: Built with Astro for a smooth, fast user experience with server-side rendering capabilities.
- **Robust Backend**: Powered by Go for high performance, concurrency, and reliability in handling large-scale warehouse operations.
- **API Integration**: RESTful API for seamless integration with other systems, including ERP and e-commerce platforms.
- **Reporting**: Generate comprehensive reports on inventory levels, item movements, pending issues, and more for data-driven decision making.
- **User Management**: Admin panel for creating and managing user accounts with role-based access control.
- **Barcode/QR Code Support**: Scan and generate barcodes or QR codes for efficient item tracking and processing.
- **Mobile Responsive**: Access the system on-the-go from any device for real-time warehouse management.

## 🛠️ Tech Stack

- **Backend**: Go
- **Frontend**: Astro
- **Documentation**: MkDocs
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **API**: RESTful with Swagger documentation

## 🏗️ Project Structure
```bash
warehaus/
├── backend/
│ ├── cmd/
│ ├── internal/
│ ├── pkg/
│ └── tests/
├── frontend/
│ ├── src/
│ ├── public/
│ └── astro.config.mjs
├── docs/
└── docker-compose.yml
``` 


## 🚀 Getting Started

### Prerequisites

- Go 1.16+
- Node.js 14+
- PostgreSQL 13+
- Docker (optional)

### Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/warehaus.git
   cd warehaus
   ```

2. Set up the backend:
   ```
   cd backend
   go mod tidy
   cp .env.example .env  # Edit .env with your configuration
   ```

3. Set up the frontend:
   ```
   cd frontend
   npm install
   ```

4. Start the development servers:
   ```
   # In the backend directory
   go run cmd/server/main.go

   # In the frontend directory
   npm run dev
   ```

5. Access the application at `http://localhost:3000`

## 🧪 Testing

### Backend Testing

Run Go tests:
```bash
cd backend
go test ./...
```


## 🤝 Contributing

We welcome contributions to Warehaus! Please follow these steps:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

Please read our [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## 🌟 Astro Frontend

Warehaus leverages Astro's powerful features:

- **Static Site Generation (SSG)**: Pre-render pages for lightning-fast load times
- **Partial Hydration**: Selectively add interactivity only where needed
- **Component Islands**: Isolate dynamic components for optimal performance
- **Built-in Markdown Support**: Easily create content-driven pages
- **TypeScript Integration**: Enjoy type safety throughout your frontend code

Learn more about our Astro setup in the [frontend README](frontend/README.md).

## 🚀 Go Backend

Our Go backend is designed for high performance and scalability:

- **Clean Architecture**: Follows domain-driven design principles
- **Middleware Chain**: Custom middleware for logging, authentication, and more
- **Graceful Shutdown**: Implements proper shutdown procedures for all services

Explore our backend architecture in the [backend README](backend/README.md).

## 📚 MkDocs Documentation

We use MkDocs for comprehensive project documentation:

- **Material Theme**: Beautiful, responsive documentation layout
- **Automatic Navigation**: Generated from your directory structure
- **Markdown Support**: Write docs in easy-to-read Markdown
- **Code Highlighting**: Syntax highlighting for code snippets
- **Search Functionality**: Built-in search for easy navigation

To build and serve the documentation:
```bash 
cd docs
mkdocs serve
```
