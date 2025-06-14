<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b">
      <div class="max-w-7xl mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <router-link to="/" class="flex items-center space-x-2 text-gray-600 hover:text-gray-900">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
              </svg>
              <span class="text-sm">Trang chủ</span>
            </router-link>
            <div class="w-10 h-10 bg-green-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
            </div>
            <div>
              <h1 class="text-xl font-bold text-gray-900">VHU Shop</h1>
              <p class="text-xs text-green-600">Phiên bản 2.0 - Bảo mật</p>
            </div>
          </div>
          
          <div class="flex items-center space-x-4">
            <span class="text-sm text-gray-600">Không có tài khoản?</span>
            <button 
              @click="isRegister = true" 
              v-if="!isRegister"
              class="text-green-600 hover:text-green-700 text-sm font-medium"
            >
              Đăng ký ngay
            </button>
          </div>
        </div>
      </div>
    </header>

    <div class="flex items-center justify-center py-12 px-4">
      <div class="max-w-md w-full">
        
        <!-- Login Form -->
        <div class="bg-white rounded-lg shadow-md border p-8">
          <div class="text-center mb-8">
            <h2 class="text-2xl font-bold text-gray-900">
              {{ isRegister ? 'Tạo tài khoản mới' : 'Đăng nhập vào VHU Shop' }}
            </h2>
            <p class="text-gray-600 mt-2">
              {{ isRegister ? 'Tham gia cộng đồng mua sắm của chúng tôi' : 'Chào mừng bạn quay trở lại!' }}
            </p>
          </div>

          <!-- Error Message -->
          <div v-if="error" class="mb-4 p-4 bg-red-50 border border-red-200 rounded-lg">
            <p class="text-red-700 text-sm">{{ error }}</p>
          </div>

          <form @submit.prevent="handleSubmit" class="space-y-6">
            
            <!-- Username -->
            <div>
              <label for="username" class="block text-sm font-medium text-gray-700 mb-2">
                Tên đăng nhập
              </label>
              <input
                id="username"
                v-model="form.username"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500"
                placeholder="Nhập tên đăng nhập của bạn"
              />
              <p class="text-xs text-gray-500 mt-1">3-32 ký tự, chỉ chứa chữ, số và dấu gạch dưới</p>
            </div>

            <!-- Password -->
            <div>
              <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
                Mật khẩu
              </label>
              <input
                id="password"
                v-model="form.password"
                type="password"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500"
                placeholder="Nhập mật khẩu của bạn"
              />
              <div v-if="isRegister" class="mt-2 p-3 bg-blue-50 rounded-lg border border-blue-200">
                <p class="text-xs font-medium text-blue-800 mb-1">Yêu cầu mật khẩu mạnh:</p>
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

            <!-- Note: No role selection in secure version -->

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="isLoading"
              class="w-full py-3 px-4 bg-green-600 hover:bg-green-700 text-white font-medium rounded-lg transition-colors disabled:opacity-50"
            >
              {{ isLoading ? 'Đang xử lý...' : (isRegister ? 'Tạo tài khoản' : 'Đăng nhập') }}
            </button>

            <!-- Toggle Register/Login -->
            <div class="text-center">
              <button
                type="button"
                @click="toggleMode"
                class="text-green-600 hover:text-green-700 text-sm"
              >
                {{ isRegister ? 'Đã có tài khoản? Đăng nhập' : 'Chưa có tài khoản? Đăng ký ngay' }}
              </button>
            </div>
          </form>
        </div>

        <!-- Security Features -->
        <div class="mt-8 bg-green-50 border border-green-200 rounded-lg p-6">
          <div class="flex items-start">
            <div class="flex-shrink-0">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
            </div>
            <div class="ml-3">
              <h3 class="text-sm font-medium text-green-800">
                🛡️ Tính năng bảo mật (Phiên bản 2.0)
              </h3>
              <div class="mt-2 text-sm text-green-700">
                <p class="mb-2"><strong>Trang này được bảo vệ bởi:</strong></p>
                <ul class="space-y-1 ml-4">
                  <li>• <strong>Password Hashing:</strong> Mật khẩu được mã hóa bằng bcrypt</li>
                  <li>• <strong>Input Validation:</strong> Kiểm tra và làm sạch tất cả dữ liệu đầu vào</li>
                  <li>• <strong>SQL Injection Prevention:</strong> Sử dụng parameterized queries</li>
                  <li>• <strong>Rate Limiting:</strong> Giới hạn số lần thử đăng nhập (10 req/sec)</li>
                  <li>• <strong>Secure Sessions:</strong> JWT token trong HTTP-only cookies</li>
                  <li>• <strong>Role Management:</strong> Phân quyền được quản lý bởi admin</li>
                </ul>
                <p class="mt-2 text-green-600">
                  <strong>Thử nghiệm:</strong> Thử các payload tấn công - chúng sẽ bị chặn và làm sạch!
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const isRegister = ref(false)
const isLoading = ref(false)
const error = ref('')

const form = ref({
  username: '',
  password: ''
})

function toggleMode() {
  isRegister.value = !isRegister.value
  error.value = ''
  form.value = {
    username: '',
    password: ''
  }
}

async function handleSubmit() {
  if (!form.value.username || !form.value.password) {
    error.value = 'Vui lòng nhập đầy đủ thông tin'
    return
  }

  isLoading.value = true
  error.value = ''

  try {
    let result
    if (isRegister.value) {
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
      // Successful authentication - go to products
      router.push('/v2/products')
    } else {
      error.value = result.error || 'Xác thực thất bại'
    }
  } catch (err: any) {
    error.value = 'Đã xảy ra lỗi kết nối'
  } finally {
    isLoading.value = false
  }
}
</script> 