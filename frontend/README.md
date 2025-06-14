# VHU Security Test Frontend

Frontend application cho hệ thống kiểm thử bảo mật API được xây dựng bằng Vue 3, Vite, Pinia, và Tailwind CSS.

## 🚀 Tính năng

### Xác thực & Bảo mật
- ✅ **JWT Authentication** với HTTP-only cookies
- ✅ **Google OAuth2** integration
- ✅ **Role-based Access Control** (user, admin, moderator)
- ✅ **Route Guards** bảo vệ các trang cần xác thực
- ✅ **Rate Limiting** testing và monitoring
- ✅ **Password Security** với validation mạnh

### Quản lý Sản phẩm
- ✅ **CRUD Operations** cho sản phẩm (chỉ admin)
- ✅ **Pagination** cho danh sách sản phẩm
- ✅ **Real-time Error Handling** cho API responses

### Giao diện
- ✅ **Responsive Design** với Tailwind CSS
- ✅ **Modern UI/UX** đơn giản và thân thiện
- ✅ **Loading States** và error handling
- ✅ **Vietnamese Localization**

## 🛠️ Technology Stack

- **Vue 3** với Composition API
- **TypeScript** cho type safety
- **Vite** cho development và build
- **Pinia** cho state management
- **Vue Router** cho routing
- **Axios** cho HTTP requests
- **Tailwind CSS** cho styling

## 📦 Cài đặt

1. **Cài đặt dependencies:**
```bash
npm install
```

2. **Tạo file .env.local:**
```bash
# Copy from template
cp .env.example .env.local

# Cấu hình URL backend
VITE_API_URL=http://localhost:8080/api
VITE_FRONTEND_URL=http://localhost:5173
```

3. **Chạy development server:**
```bash
npm run dev
```

4. **Build cho production:**
```bash
npm run build
```

## 🏗️ Cấu trúc thư mục

```
src/
├── components/          # Vue components
├── services/           
│   └── api.ts          # API service với Axios
├── stores/             
│   ├── auth.ts         # Authentication store
│   └── product.ts      # Product management store
├── views/              
│   ├── LoginView.vue   # Trang đăng nhập/đăng ký
│   ├── DashboardView.vue # Trang chính
│   ├── ProductsView.vue  # Quản lý sản phẩm
│   └── ProfileView.vue   # Trang profile
├── router/
│   └── index.ts        # Route definitions và guards
└── assets/
    └── main.css        # Global styles với Tailwind
```

## 🔐 Tính năng Bảo mật

### 1. JWT Authentication
- Token được lưu trong **HTTP-only cookies**
- Automatic cookie transmission với `withCredentials: true`
- Token expiry: 24 giờ
- Auto-refresh user state từ cookies

### 2. Role-based Access Control
- **User**: Xem products, quản lý profile
- **Admin**: Full CRUD operations cho products
- **Moderator**: Quyền trung gian (tùy backend implementation)

### 3. Route Protection
```typescript
// Các route được bảo vệ
meta: { requiresAuth: true }        // Cần đăng nhập
meta: { requiresAdmin: true }       // Cần quyền admin
meta: { requiresGuest: true }       // Chỉ cho guest
```

### 4. API Error Handling
- **401 Unauthorized**: Auto redirect to login
- **403 Forbidden**: Access denied notification
- **429 Rate Limit**: Rate limit exceeded warning
- **Network errors**: Connection error handling

## 🎨 UI Components

### Authentication
- **LoginView**: Form đăng nhập/đăng ký với Google OAuth
- **Route Guards**: Auto redirect based on auth status

### Dashboard
- **User Info Display**: ID, username, role
- **Security Status**: JWT, RBAC, Rate Limiting status
- **Quick Actions**: Navigate to key features
- **Rate Limit Testing**: Admin-only feature

### Product Management
- **Product Table**: Paginated list với CRUD actions
- **Modal Forms**: Create/Edit product với validation
- **Admin-only Actions**: Create, Update, Delete permissions

### Profile
- **User Information**: Display account details
- **Change Password**: Secure password update form
- **Security Overview**: Auth và security features status

## 🔧 Configuration

### Environment Variables
```bash
# API Backend URL
VITE_API_URL=http://localhost:8080/api

# Frontend URL (for OAuth callbacks)
VITE_FRONTEND_URL=http://localhost:5173
```

### Axios Configuration
```typescript
// API service với cookie-based auth
const api = axios.create({
  baseURL: process.env.VITE_API_URL,
  withCredentials: true,  // Quan trọng!
  headers: {
    'Content-Type': 'application/json',
  },
})
```

## 🧪 Testing Security Features

### 1. Authentication Flow
1. Đăng ký account mới hoặc đăng nhập
2. Test Google OAuth2 flow
3. Verify JWT token trong cookies (DevTools)
4. Test auto-logout khi token expires

### 2. Role-based Access
1. Login với user thường → không thấy admin features
2. Login với admin → full access to product CRUD
3. Test unauthorized access attempts

### 3. Rate Limiting
1. Sử dụng "Test Rate Limit" button (admin only)
2. Monitor network tab cho 429 errors
3. Verify rate limit messages

### 4. Security Headers
1. Check cookies: `httpOnly=true`, `secure=true`
2. Verify CORS configuration
3. Test CSRF protection

## 📋 API Endpoints

```typescript
// Authentication
POST /api/auth/login        // Đăng nhập
POST /api/auth/register     // Đăng ký
POST /api/auth/logout       // Đăng xuất

// OAuth2
GET  /api/oauth2/authorize  // Google OAuth redirect
GET  /api/oauth2/callback   // OAuth callback

// User Profile
GET  /api/profile           // Lấy thông tin user
POST /api/profile/change-password // Đổi mật khẩu

// Products (requires auth)
GET    /api/products        // Danh sách sản phẩm
POST   /api/products        // Tạo sản phẩm (admin only)
PUT    /api/products/:id    // Cập nhật sản phẩm (admin only)
DELETE /api/products/:id    // Xóa sản phẩm (admin only)
```

## 🐛 Troubleshooting

### CORS Issues
```bash
# Đảm bảo backend cho phép credentials
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: http://localhost:5173
```

### Cookie Issues
```bash
# Kiểm tra cookie trong DevTools
Application > Cookies > localhost:5173
# Tìm cookie 'auth_token' với httpOnly=true
```

### Rate Limiting
```bash
# Nếu gặp 429 errors thường xuyên
# Giảm tần suất requests hoặc tăng rate limit trong backend
```

## 📞 Support

Dự án này được tạo ra để **kiểm thử bảo mật API**. Các tính năng bảo mật được thiết kế để demo và test các vulnerability patterns phổ biến.

**Lưu ý**: Đây là hệ thống test, không dùng cho production mà không có thêm security hardening.
