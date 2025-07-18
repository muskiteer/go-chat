# 💬 Go Chat - Real-time Chat Application

A modern, full-stack real-time chat application built with Go (Golang) backend and React frontend. Features instant messaging, user authentication, online status tracking, and a beautiful responsive UI.

![Go](https://img.shields.io/badge/Go-1.24.2-00ADD8?style=flat-square&logo=go)
![React](https://img.shields.io/badge/React-19.1.0-61DAFB?style=flat-square&logo=react)
![MongoDB](https://img.shields.io/badge/MongoDB-Database-47A248?style=flat-square&logo=mongodb)
![WebSocket](https://img.shields.io/badge/WebSocket-Real--time-FF6B6B?style=flat-square)

## ✨ Features

### 🔐 Authentication & Security
- User registration and login system
- JWT-based authentication
- Password hashing with bcrypt
- Protected routes and middleware
- Input validation and sanitization

### 💬 Real-time Messaging
- Instant message delivery via WebSocket
- Online/offline user status tracking
- Real-time user presence indicators
- Message broadcasting to connected users
- Persistent message history

### 🎨 Modern UI/UX
- Responsive design with Tailwind CSS
- Beautiful and intuitive chat interface
- Dark/light theme support
- Loading skeletons and smooth animations
- Toast notifications for user feedback

### 🏗️ Architecture
- Clean MVC architecture
- Modular component structure
- State management with Zustand
- RESTful API design
- WebSocket integration

## 🛠️ Tech Stack

### Backend (Go)
- **Framework**: Gorilla Mux for HTTP routing
- **Database**: MongoDB with official Go driver
- **Real-time**: WebSocket connections
- **Authentication**: JWT tokens
- **Security**: Password hashing, input validation
- **Environment**: dotenv for configuration

### Frontend (React)
- **Framework**: React 19 with Vite
- **State Management**: Zustand
- **Styling**: Tailwind CSS + DaisyUI
- **HTTP Client**: Axios
- **Routing**: React Router DOM
- **Icons**: Lucide React
- **Notifications**: React Hot Toast

## 📁 Project Structure

```
go-chat/
├── backend/
│   ├── src/
│   │   ├── main.go                 # Application entry point
│   │   ├── controllers/            # HTTP request handlers
│   │   │   ├── auth.controller.go
│   │   │   └── messages.controller.go
│   │   ├── lib/                    # Core libraries
│   │   │   ├── db.go              # Database connection
│   │   │   └── socket.go          # WebSocket handling
│   │   ├── middleware/             # HTTP middleware
│   │   │   ├── authenticate.go
│   │   │   └── cors.go
│   │   ├── models/                 # Data models
│   │   │   ├── messages.go
│   │   │   └── users.go
│   │   └── routes/                 # API routes
│   │       ├── auth.route.go
│   │       └── messages.route.go
│   ├── internals/                  # Internal utilities
│   │   ├── security.go
│   │   └── validate.go
│   ├── utils/                      # Helper functions
│   │   ├── json.go
│   │   └── jwt.go
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── src/
│   │   ├── App.jsx                 # Main app component
│   │   ├── main.jsx               # App entry point
│   │   ├── components/            # Reusable components
│   │   │   ├── ChatContainer.jsx
│   │   │   ├── Sidebar.jsx
│   │   │   ├── MessageInput.jsx
│   │   │   └── ...
│   │   ├── pages/                 # Route components
│   │   │   ├── HomePage.jsx
│   │   │   ├── LoginPage.jsx
│   │   │   └── SignUpPage.jsx
│   │   ├── store/                 # State management
│   │   │   ├── useAuthStore.js
│   │   │   ├── useChatStore.js
│   │   │   └── useThemeStore.js
│   │   └── lib/                   # Utilities
│   │       ├── axios.js
│   │       └── utils.js
│   ├── package.json
│   └── vite.config.js
└── README.md
```

## 🚀 Getting Started

### Prerequisites
- Go 1.24.2 or higher
- Node.js 18+ and npm/yarn
- MongoDB database
- Git

### Backend Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/muskiteer/go-chat
   cd go-chat/backend
   ```

2. **Install Go dependencies**
   ```bash
   go mod download
   ```

3. **Environment Configuration**
   Create a `.env` file in the backend directory:
   ```env
   PORT=8000
   MONGODB_URI=mongodb://localhost:27017
   JWT_SECRET=your-super-secret-jwt-key
   NODE_ENV=development
   ```

4. **Start MongoDB**
   Make sure MongoDB is running on your system:
   ```bash
   # For local MongoDB installation
   mongod
   
   # Or using Docker
   docker run -d -p 27017:27017 --name mongodb mongo:latest
   ```

5. **Run the backend server**
   ```bash
   cd src
   go run main.go
   ```
   The server will start on `http://localhost:8000`

### Frontend Setup

1. **Navigate to frontend directory**
   ```bash
   cd ../frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   # or
   yarn install
   ```

3. **Environment Configuration**
   Create a `.env` file in the frontend directory:
   ```env
   VITE_API_URL=http://localhost:8000
   ```

4. **Start the development server**
   ```bash
   npm run dev
   # or
   yarn dev
   ```
   The frontend will start on `http://localhost:5173`

## 📡 API Endpoints

### Authentication
- `POST /api/auth/signup` - User registration
- `POST /api/auth/login` - User login
- `POST /api/auth/logout` - User logout
- `GET /api/auth/check` - Check authentication status

### Messages
- `GET /api/messages/users` - Get all users for chat
- `GET /api/messages/:id` - Get messages with specific user
- `POST /api/messages/send/:id` - Send message to user

### WebSocket
- `WS /ws?userId={userId}` - WebSocket connection for real-time features

## 🗄️ Database Schema

### Users Collection
```javascript
{
  _id: ObjectId,
  username: String,
  email: String,
  hashed_password: String,
  created_at: Date
}
```

### Messages Collection
```javascript
{
  _id: ObjectId,
  sender_id: ObjectId,
  receiver_id: ObjectId,
  content: String,
  created_at: Date
}
```

## 🔧 Key Dependencies

### Backend Dependencies
```go
github.com/golang-jwt/jwt/v5 v5.2.2
github.com/gorilla/mux v1.8.1
github.com/gorilla/websocket v1.5.3
github.com/joho/godotenv v1.5.1
github.com/microcosm-cc/bluemonday v1.0.27
go.mongodb.org/mongo-driver v1.17.4
golang.org/x/crypto v0.33.0
```

### Frontend Dependencies
```json
{
  "react": "^19.1.0",
  "zustand": "^5.0.6",
  "axios": "^1.10.0",
  "react-router-dom": "^7.6.3",
  "tailwindcss": "^3.4.17",
  "daisyui": "^4.12.23",
  "lucide-react": "^0.525.0",
  "react-hot-toast": "^2.5.2"
}
```

## 🌟 Features in Detail

### Real-time Communication
- WebSocket connections for instant messaging
- Online user status tracking
- Real-time message delivery
- Connection management and cleanup

### Security Features
- JWT-based authentication
- Password hashing with bcrypt
- Input validation and sanitization
- CORS protection
- Protected routes

### User Experience
- Responsive design for all devices
- Dark/light theme toggle
- Loading states and skeletons
- Error handling with toast notifications
- Smooth animations and transitions

## 🧪 Development

### Running Tests
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests
cd frontend
npm test
```

### Building for Production
```bash
# Backend
cd backend/src
go build -o ../bin/chat-app main.go

# Frontend
cd frontend
npm run build
```

## 🚀 Deployment

### Backend Deployment
1. Build the Go binary
2. Set up environment variables
3. Configure MongoDB connection
4. Deploy to your preferred platform (AWS, GCP, Heroku, etc.)

### Frontend Deployment
1. Build the React app
2. Deploy to static hosting (Vercel, Netlify, AWS S3, etc.)
3. Update API URLs for production

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Go community for excellent libraries
- React team for the amazing framework
- MongoDB for reliable database solutions
- All contributors and testers

---

**Made with ❤️ by [Muskiteer](https://github.com/muskiteer)**

For questions or support, please open an issue or contact the maintainers.