<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <!-- Background Pattern -->
    <div class="absolute inset-0 bg-grid-pattern opacity-5"></div>
    
    <div class="max-w-md w-full space-y-8 relative">
      <!-- Logo and Header -->
      <div class="text-center">
        <div class="mx-auto h-16 w-16 bg-gradient-to-r from-blue-600 to-indigo-600 rounded-2xl flex items-center justify-center mb-6 shadow-lg">
          <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
          </svg>
        </div>
        
        <h2 class="text-3xl font-bold text-gray-900 mb-2">
          {{ showRegister ? 'Tạo Tài Khoản' : 'Đăng Nhập' }}
        </h2>
        <p class="text-gray-600 mb-2">
          {{ showRegister ? 'Tham gia VHU Security Lab' : 'Chào mừng trở lại' }}
        </p>
        
        <!-- API Version Badge -->
        <div class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800 border border-green-200">
          <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
          </svg>
          API v2 - Secure Authentication
        </div>
      </div>

      <!-- Main Login Card -->
      <div class="bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl border border-white/20 overflow-hidden">
        
        <!-- Error Messages -->
        <div v-if="error" class="bg-red-50 border-l-4 border-red-400 p-4 m-6 mb-0 rounded-lg">
          <div class="flex">
            <div class="flex-shrink-0">
              <svg class="h-5 w-5 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm text-red-700">{{ error }}</p>
            </div>
          </div>
        </div>

        <div v-if="oauthError" class="bg-red-50 border-l-4 border-red-400 p-4 m-6 mb-0 rounded-lg">
          <div class="flex">
            <div class="flex-shrink-0">
              <svg class="h-5 w-5 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm text-red-700">Google OAuth Error: {{ oauthError }}</p>
            </div>
          </div>
        </div>

        <!-- Form Content -->
        <div class="p-8">
          <form class="space-y-6" @submit.prevent="handleSubmit">
            
            <!-- Username Field -->
            <div>
              <label for="username" class="block text-sm font-semibold text-gray-700 mb-2">
                Tên đăng nhập
              </label>
              <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                  </svg>
                </div>
                <input
                  id="username"
                  v-model="form.username"
                  name="username"
                  type="text"
                  required
                  class="w-full pl-10 pr-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-gray-50 focus:bg-white"
                  placeholder="Nhập tên đăng nhập"
                />
              </div>
            </div>

            <!-- Password Field -->
            <div>
              <label for="password" class="block text-sm font-semibold text-gray-700 mb-2">
                Mật khẩu
              </label>
              <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                  </svg>
                </div>
                <input
                  id="password"
                  v-model="form.password"
                  name="password"
                  type="password"
                  required
                  class="w-full pl-10 pr-4 py-3 border border-gray-200 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-gray-50 focus:bg-white"
                  placeholder="Nhập mật khẩu"
                />
              </div>
              
              <!-- Password Requirements (for register) -->
              <div v-if="showRegister" class="mt-2 p-3 bg-blue-50 rounded-lg border border-blue-200">
                <p class="text-xs font-medium text-blue-800 mb-2">Yêu cầu mật khẩu:</p>
                <ul class="text-xs text-blue-700 space-y-1">
                  <li class="flex items-center">
                    <svg class="w-3 h-3 text-blue-500 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                    </svg>
                    Ít nhất 8 ký tự
                  </li>
                  <li class="flex items-center">
                    <svg class="w-3 h-3 text-blue-500 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                    </svg>
                    Chữ hoa, chữ thường, số và ký tự đặc biệt
                  </li>
                </ul>
              </div>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="isLoading"
              class="w-full flex justify-center items-center py-3 px-4 border border-transparent rounded-xl text-sm font-semibold text-white bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 shadow-lg"
            >
              <svg v-if="isLoading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              {{ isLoading ? 'Đang xử lý...' : (showRegister ? 'Tạo Tài Khoản' : 'Đăng Nhập') }}
            </button>

            <!-- Divider -->
            <div class="relative my-6">
              <div class="absolute inset-0 flex items-center">
                <div class="w-full border-t border-gray-200"/>
              </div>
              <div class="relative flex justify-center text-sm">
                <span class="px-4 bg-white text-gray-500 font-medium">Hoặc</span>
              </div>
            </div>

            <!-- Google OAuth Button -->
            <button
              type="button"
              @click="handleGoogleLogin"
              class="w-full flex justify-center items-center py-3 px-4 border border-gray-200 rounded-xl text-sm font-semibold text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-200 shadow-sm"
            >
              <svg class="w-5 h-5 mr-3" viewBox="0 0 24 24">
                <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
              </svg>
              Đăng nhập với Google
            </button>
          </form>
        </div>
      </div>

      <!-- Toggle Register/Login -->
      <div class="text-center">
        <p class="text-gray-600">
          {{ showRegister ? 'Đã có tài khoản?' : 'Chưa có tài khoản?' }}
          <button
            @click="showRegister = !showRegister"
            class="font-semibold text-blue-600 hover:text-blue-700 transition-colors ml-1"
          >
            {{ showRegister ? 'Đăng nhập ngay' : 'Đăng ký miễn phí' }}
          </button>
        </p>
      </div>

      <!-- Security Features Info -->
      <div class="bg-white/60 backdrop-blur-sm rounded-2xl border border-white/20 p-6">
        <h3 class="text-sm font-semibold text-gray-900 mb-3 flex items-center">
          <svg class="w-4 h-4 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
          </svg>
          Tính Năng Bảo Mật
        </h3>
        <div class="grid grid-cols-2 gap-3 text-xs text-gray-600">
          <div class="flex items-center">
            <div class="w-1.5 h-1.5 bg-green-500 rounded-full mr-2"></div>
            Password hashing (bcrypt)
          </div>
          <div class="flex items-center">
            <div class="w-1.5 h-1.5 bg-green-500 rounded-full mr-2"></div>
            JWT trong HTTP-only cookies
          </div>
          <div class="flex items-center">
            <div class="w-1.5 h-1.5 bg-green-500 rounded-full mr-2"></div>
            Rate limiting
          </div>
          <div class="flex items-center">
            <div class="w-1.5 h-1.5 bg-green-500 rounded-full mr-2"></div>
            Input validation
          </div>
        </div>
      </div>

      <!-- Back to Home -->
      <div class="text-center">
        <router-link
          to="/"
          class="text-sm text-gray-500 hover:text-gray-700 transition-colors inline-flex items-center"
        >
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
          </svg>
          Quay lại trang chủ
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { secureAPI } from '@/services/api'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const showRegister = ref(false)
const form = ref({
  username: '',
  password: ''
})
const error = ref<string | null>(null)
const oauthError = ref<string | null>(null)
const isLoading = ref(false)

// Check for OAuth errors in URL
onMounted(() => {
  const urlError = route.query.error as string
  if (urlError) {
    switch (urlError) {
      case 'invalid_state':
        oauthError.value = 'Invalid OAuth state'
        break
      case 'token_exchange_failed':
        oauthError.value = 'Failed to exchange OAuth token'
        break
      case 'userinfo_failed':
        oauthError.value = 'Failed to get user information'
        break
      case 'user_creation_failed':
        oauthError.value = 'Failed to create user account'
        break
      default:
        oauthError.value = 'OAuth authentication failed'
    }
  }
})

async function handleSubmit() {
  if (!form.value.username || !form.value.password) {
    error.value = 'Vui lòng nhập đầy đủ thông tin'
    return
  }

  error.value = null
  isLoading.value = true

  try {
    let result
    if (showRegister.value) {
      result = await authStore.register({
        username: form.value.username,
        password: form.value.password
      })
    } else {
      result = await authStore.login({
        username: form.value.username,
        password: form.value.password
      })
    }

    if (result.success) {
      // Redirect to intended page or dashboard
      const redirect = route.query.redirect as string || '/v2'
      router.push(redirect)
    } else {
      error.value = result.error || 'Authentication failed'
    }
  } catch (err) {
    error.value = 'Network error occurred'
  } finally {
    isLoading.value = false
  }
}

function handleGoogleLogin() {
  secureAPI.googleLogin()
}
</script>

<style scoped>
.bg-grid-pattern {
  background-image: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23000000' fill-opacity='0.1'%3E%3Ccircle cx='7' cy='7' r='1'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
}
</style> 