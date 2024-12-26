# YouTube Video Fetching Service

A scalable service that fetches and stores the latest YouTube videos based on search queries, featuring real-time updates and a modern dashboard interface.

![Dashboard Preview](screenshots/dashboard.png)

## Live Demo
- Frontend: [https://unlimited-cricket.vercel.app](https://unlimited-cricket.vercel.app/)
- Backend API: [https://unlimited-cricket.onrender.com](https://unlimited-cricket.onrender.com)

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
- [API Documentation](#api-documentation)
- [Dashboard Screenshots](#dashboard-screenshots)

## Features

### Core Functionality
- ✅ Real-time YouTube video fetching
- ✅ Advanced caching with Redis
- ✅ Intelligent rate limiting
- ✅ Multiple API key rotation
- ✅ Responsive dashboard interface
- ✅ Paginated API responses

### Dashboard Features
![Search Interface](screenshots/search.png)
- Real-time video updates
- Responsive design
- Infinite Scrolling 


## Tech Stack

### Backend
- **Language**: Go 1.21
- **Framework**: Gin Web Framework
- **Databases**: 
  - PostgreSQL (Primary storage)
  - Redis (Caching layer)
- **API Integration**: YouTube Data API v3

### Frontend
- **Framework**: Next.js 15
- **Styling**: Tailwind CSS

## Architecture

### Data Flow Architecture


### Rate Limiting Architecture
```mermaid
flowchart TD
    A[Request] --> B[Extract IP]
    B --> C{Check Rate Limiter}
    C -->|Limit Not Exceeded| D[Process Request]
    C -->|Limit Exceeded| E[429 Too Many Requests]
    D --> F[Return Response]
    
    subgraph "Rate Limiter Details"
    G[Token Bucket] -->|Refill Rate| H[5 tokens/second]
    G -->|Burst Size| I[10 tokens max]
    end
```

## Project Structure
```
├── server/
│   ├── config/           # Configuration management
│   ├── models/           # Data models
│   ├── repository/       # Database interactions
│   ├── service/          # Business logic
│   ├── controller/          # HTTP handlers
│   ├── client/           # External API clients
│   ├── cache/            # Cache management
│   ├── middleware/       # HTTP middlewares
│   └── router/           # Route definitions
│   └── main.go           # Application entry point
```

## Setup Instructions

### Prerequisites
- Go 1.21+
- PostgreSQL 16+
- Redis 6+
- Node.js 20+

### Environment Setup
```bash
# Clone the repository
git clone https://github.com/TejasThombare20/Unlimited-Cricket.git


# Copy environment file
cp .env.example .env

# Update environment variables
nano .env
```

### Environment Variables
```env
# YouTube API Keys
YOUTUBE_API_KEY_1=your_key_1
YOUTUBE_API_KEY_2=your_key_2
YOUTUBE_API_KEY_3=your_key_3
YOUTUBE_API_KEY_4=your_key_4
YOUTUBE_API_KEY_5=your_key_5

# time in minutes to fetch the data from youtube api 
FETCH_TIME = 4


# Database
DATABASE_URL =postgres://user_name:password@127.0.0.1:5432/fampay

# Redis
REDIS_URL=redis://localhost:6379/0

# Server
PORT=8080
GO_ENVIRONMENT=production

#Rate Limiting
RPS = 5
BURST_TIME = 10
```

### Running the Application

#### Backend
```bash
# Navigate to backend directory
cd server

# Install dependencies
go mod download

# Start the server
go run main.go
```

#### Frontend
```bash
# Navigate to frontend directory
cd client

# Install dependencies
npm install

# Start development server
npm run dev
```


```

## API Documentation

### GET /api/v1/videos
Fetch stored videos with pagination

**Parameters:**
- `page` (optional, default: 1): Page number
- `page_size` (optional, default: 10): Items per page

**Response:**
```json
{
    "page": 1,
    "page_size": 10,
    "data": [
        {
            "id": "video_id",
            "title": "Video Title",
            "description": "Description",
            "published_at": "2024-12-26T10:00:00Z",
            "thumbnail_url": "https://example.com/thumbnail.jpg",
            "created_at": "2024-12-26T10:01:00Z"
        }
    ]
}
```

## Dashboard Screenshots

### Home Page
![Home Page](screenshots/home.png)
The main dashboard interface showing recent videos and statistics.

### Video List
![Video List](screenshots/videos.png)
Paginated list of videos with search and filter options.

### Dark Mode
![Dark Mode](screenshots/dark-mode.png)
Dark mode interface for better visibility.

## Performance Optimizations

### Caching Strategy
- Redis caching with 5-minute TTL
- Automatic cache invalidation
- Paginated results caching

### Rate Limiting
- IP-based rate limiting
- Token bucket algorithm
- 5 requests per second per IP
- Burst allowance of 10 requests

### Database Optimizations
- Indexed queries
- Connection pooling
- Batch processing
