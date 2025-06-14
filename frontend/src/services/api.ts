import axios from 'axios'

interface APIResponse<T = any> {
  success: boolean
  error?: string
  data?: T
}

const baseURL = 'http://192.168.2.13'

// Secure API instance with proper configuration
export const api = axios.create({
  baseURL: `${baseURL}/api/v2`,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Insecure API instance demonstrating security vulnerabilities
export const insecureAPI = {
  // VULNERABILITY: No CSRF protection
  // VULNERABILITY: No request validation
  // VULNERABILITY: No error handling
  // VULNERABILITY: Storing sensitive data in localStorage
  async register(data: { username: string; password: string; role?: string }): Promise<APIResponse> {
    try {
      const response = await axios.post(`${baseURL}/api/v1/auth/register`, data)
      // VULNERABILITY: Storing sensitive data in localStorage
      localStorage.setItem('username', response.data.username)
      localStorage.setItem('role', response.data.role)
      //redirect to home page
      window.location.href = '/'
      return { success: true, data: response.data }
    } catch (error: any) {
      return {
        success: false,
        error: error.response?.data?.error || 'Registration failed'
      }
    }
  },

  // VULNERABILITY: No CSRF protection
  // VULNERABILITY: Plain text password transmission
  // VULNERABILITY: No request validation
  async login(data: { username: string; password: string }): Promise<APIResponse> {
    try {
      const response = await axios.post(`${baseURL}/api/v1/auth/login`, data)
      // VULNERABILITY: Storing sensitive data in localStorage
      localStorage.setItem('user', JSON.stringify(response.data.user))
      return { success: true, data: response.data }
    } catch (error: any) {
      return { 
        success: false,
        error: error.response?.data?.error || 'Login failed'
      }
    }
  },

  // VULNERABILITY: No token invalidation
  async logout(): Promise<APIResponse> {
    try {
      await axios.post(`${baseURL}/api/v1/auth/logout`)
      // VULNERABILITY: Not clearing all sensitive data
      localStorage.removeItem("user")
      return { success: true }
    } catch (error: any) {
      return {
        success: false,
        error: error.response?.data?.error || 'Logout failed'
      }
    }
  },

  // VULNERABILITY: No authentication check
  // VULNERABILITY: SQL injection possible
  async listProducts(params?: { name?: string; description?: string; price?: string | number }) {
    try {
      //send X-User header
      const user = JSON.parse(localStorage.getItem('user') || '{}')
      const headers = {
        'X-User': user.username
      }
      // VULNERABILITY: Not using proper query parameters
      let url = `${baseURL}/api/v1/products`
      if (params) {
        const queryString = Object.entries(params)
          .filter(([_, value]) => value)
          .map(([key, value]) => `${key}=${value}`)
          .join('&')
        if (queryString) {
          url += '?' + queryString
        }
      }
      const response = await axios.get(url, { headers })
      return response
    } catch (error: any) {
      console.error('Failed to fetch products:', error)
      throw error
    }
  },

  // VULNERABILITY: No authentication check
  // VULNERABILITY: No input validation
  async createProduct(data: { name: string; description: string; price: number }) {
    try {
      const storgeData = localStorage.getItem('user')
      const userData = JSON.parse(storgeData || '{}')
      const headers = {
        "X-User": userData.username,
        "X-Role": userData.role
      }
      const response = await axios.post(`${baseURL}/api/v1/products`, data, { headers } )
      return response
    } catch (error: any) {
      console.error('Failed to create product:', error)
      throw error
    }
  },

  // VULNERABILITY: No authentication check
  // VULNERABILITY: No confirmation required
  async deleteProduct(id: number) {
    try {
      const storgeData = localStorage.getItem('user')
      const data = JSON.parse(storgeData || '{}')
      const headers = {
        "X-User": data.username,
        "X-Role": data.role
      }
      const response = await axios.delete(`${baseURL}/api/v1/products/${id}`, { headers } )
      return { success: true, data: response.data }
    } catch (error: any) {
      console.error('Failed to delete product:', error)
      throw error
    }
  },

  // VULNERABILITY: No authentication check
  // VULNERABILITY: SQL injection possible
  async getProfile(username: string) {
    try {
      const response = await axios.get(`${baseURL}/api/v1/users/${username}`)
      return response
    } catch (error: any) {
      console.error('Failed to fetch profile:', error)
      throw error
    }
  },

  // VULNERABILITY: No authentication check
  // VULNERABILITY: Plain text password transmission
  async changePassword(data: { username: string; newPassword: string }) {
    try {
      const storgeData = localStorage.getItem('user')
      const userData = JSON.parse(storgeData || '{}')
      const headers = {
        "X-User": userData.username
      }
      const response = await axios.post(`${baseURL}/api/v1/auth/change-password`, data, { headers } )
      return { success: true, data: response.data }
    } catch (error: any) {
      console.error('Failed to change password:', error)
      throw error
    }
  }
}

// API v2 - Đã được bảo mật (Secure)
export const secureAPI = {
  // Auth - Password hashing, JWT cookies, proper validation
  register: (username: string, password: string) =>
    api.post('/auth/register', { username, password }, { withCredentials: true }),
  
  login: (username: string, password: string) =>
    api.post('/auth/login', { username, password }, { withCredentials: true }),
  
  logout: () => api.post('/auth/logout', {}, { withCredentials: true }),
  
  changePassword: (data: { currentPassword: string; newPassword: string }) =>
    api.post('/profile/change-password', 
      {current_password: data.currentPassword, new_password: data.newPassword},
      { withCredentials: true }
    ),
  
  getProfile: () => api.get('/profile', { withCredentials: true }),

  // OAuth
  googleLogin: () => api.get('/oauth2/authorize'),

  // Products - Requires authentication, role-based access
  listProducts: (page: number = 1, limit: number = 10) =>
    api.get('/products', {
      params: { page, limit },
      withCredentials: true
    }),
  
  createProduct: (data: any) =>
    api.post('/products', data, { withCredentials: true }),
  
  updateProduct: (id: number, data: any) =>
    api.put(`/products/${id}`, data, { withCredentials: true }),
  
  deleteProduct: (id: number) =>
    api.delete(`/products/${id}`, { withCredentials: true })
}

// Re-export for backward compatibility
export const authAPI = secureAPI
export const productAPI = secureAPI

// Types cho API responses
export interface User {
  id: number
  username: string
  role: string
}

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  password: string
}

export interface ChangePasswordRequest {
  currentPassword: string
  newPassword: string
}

export interface Product {
  ID: number
  name: string
  description: string
  price: number
  CreatedAt: string
  UpdatedAt: string
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  total_pages: number
  limit: number
}

export interface ApiResponse<T = any> {
  user?: User
  message?: string
  error?: string
  data?: T
}

// Response interceptor để xử lý lỗi
api.interceptors.response.use(
  (response) => response,
  (error) => {
    // Xử lý các lỗi phổ biến
    if (error.response?.status === 401) {
      // Token hết hạn hoặc không hợp lệ
      console.log('Authentication failed')
    } else if (error.response?.status === 403) {
      // Không có quyền truy cập
      console.log('Access denied')
    } else if (error.response?.status === 429) {
      // Rate limit exceeded
      console.log('Rate limit exceeded')
    }
    return Promise.reject(error)
  }
)

export default api 