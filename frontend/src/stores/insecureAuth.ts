import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { insecureAPI } from '@/services/api'

interface User {
  ID: number
  username: string
  role: string
  password?: string // VULNERABILITY: Password exposed in user object
  token?: string // VULNERABILITY: Token exposed in user object
}

interface LoginRequest {
  username: string
  password: string
}

interface RegisterRequest {
  username: string
  password: string
  role?: string
}

// VULNERABILITY: Using localStorage for sensitive data
const STORAGE_KEY = 'insecure_user_data'

export const useInsecureAuthStore = defineStore('insecureAuth', () => {
  // State
  const user = ref<User | null>(loadUserFromStorage())
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // Getters
  const isAuthenticated = computed(() => !!user.value)
  const isAdmin = computed(() => user.value?.role === 'admin')
  const isModerator = computed(() => user.value?.role === 'moderator')

  // Actions
  async function register(data: RegisterRequest) {
    isLoading.value = true
    error.value = null

    try {
      const result = await insecureAPI.register(data)
      if (result.success) {
        // VULNERABILITY: Loading sensitive data from localStorage
        const userData = JSON.parse(localStorage.getItem('user') || '{}')
        user.value = userData
        return { success: true }
      } else {
        error.value = result.error || 'Registration failed'
        return { success: false, error: result.error || 'Registration failed' }
      }
    } catch (err: any) {
      error.value = err.message
      return { success: false, error: err.message }
    } finally {
      isLoading.value = false
    }
  }

  async function login(data: LoginRequest) {
    isLoading.value = true
    error.value = null

    try {
      const result = await insecureAPI.login(data)
      if (result.success) {
        // VULNERABILITY: Loading sensitive data from localStorage
        const userData = JSON.parse(localStorage.getItem('user') || '{}')
        user.value = userData
        console.log(user.value)
        return { success: true }
      } else {
        error.value = result.error || 'Login failed'
        return { success: false, error: result.error || 'Login failed' }
      }
    } catch (err: any) {
      error.value = err.message
      return { success: false, error: err.message }
    } finally {
      isLoading.value = false
    }
  }

  async function logout() {
    isLoading.value = true
    error.value = null

    try {
      const result = await insecureAPI.logout()
      if (result.success) {
        user.value = null
        // VULNERABILITY: Not clearing all sensitive data
        localStorage.removeItem('user')
        localStorage.removeItem('token')
        return { success: true }
      } else {
        error.value = result.error || 'Logout failed'
        return { success: false, error: result.error || 'Logout failed' }
      }
    } catch (err: any) {
      error.value = err.message
      return { success: false, error: err.message }
    } finally {
      isLoading.value = false
    }
  }

  // Helper function to load user data from localStorage
  function loadUserFromStorage(): User | null {
    try {
      // VULNERABILITY: Loading sensitive data from localStorage
      const userData = localStorage.getItem('user')
      return userData ? JSON.parse(userData) : null
    } catch {
      return null
    }
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

    // Actions
    register,
    login,
    logout
  }
}) 