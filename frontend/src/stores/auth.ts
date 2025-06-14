import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { secureAPI } from '@/services/api'

interface User {
  ID: number
  username: string
  role: string
}

interface LoginRequest {
  username: string
  password: string
}

interface RegisterRequest {
  username: string
  password: string
}

interface ChangePasswordRequest {
  currentPassword: string
  newPassword: string
}

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref<User | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // Getters
  const isAuthenticated = computed(() => !!user.value)
  const isAdmin = computed(() => user.value?.role === 'admin')
  const isModerator = computed(() => user.value?.role === 'moderator')
  const hasRole = computed(() => (role: string) => user.value?.role === role)

  // Actions
  async function login(credentials: LoginRequest) {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await secureAPI.login(credentials.username, credentials.password)
      
      if (response.data.user) {
        user.value = response.data.user
        return { success: true, user: response.data.user }
      } else {
        throw new Error('Invalid response format')
      }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Login failed'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  async function register(data: RegisterRequest) {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await secureAPI.register(data.username, data.password)
      
      if (response.data.user) {
        user.value = response.data.user
        return { success: true, user: response.data.user }
      } else {
        throw new Error('Invalid response format')
      }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Registration failed'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  async function logout() {
    isLoading.value = true
    error.value = null
    
    try {
      await secureAPI.logout()
      user.value = null
      return { success: true }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Logout failed'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  async function loadUserFromCookie() {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await secureAPI.getProfile()
      
      if (response.data.user) {
        user.value = response.data.user
        return { success: true, user: response.data.user }
      } else {
        throw new Error('Invalid response format')
      }
    } catch (err: any) {
      // Không có token hoặc token hết hạn - đây là trường hợp bình thường
      user.value = null
      return { success: false, error: 'Not authenticated' }
    } finally {
      isLoading.value = false
    }
  }

  async function changePassword(data: ChangePasswordRequest) {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await secureAPI.changePassword(data.currentPassword, data.newPassword)
      return { success: true, message: response.data.message || 'Password changed successfully' }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Password change failed'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  function clearError() {
    error.value = null
  }

  function clearUser() {
    user.value = null
  }

  return {
    // State
    user,
    isLoading,
    error,
    
    // Getters
    isAuthenticated,
    isAdmin,
    isModerator,
    hasRole,
    
    // Actions
    login,
    register,
    logout,
    loadUserFromCookie,
    changePassword,
    clearError,
    clearUser,
  }
}) 