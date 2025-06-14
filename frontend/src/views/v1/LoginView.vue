<template>
  <div class="font-sans min-h-screen bg-gradient-to-b from-gray-50 to-gray-100 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md mx-auto">
      <!-- Security Warning Banner -->
      <div class="mb-8 p-4 bg-red-50 border-l-4 border-red-500 rounded-r-md">
        <div class="flex">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">Các lỗ hổng bảo mật</h3>
            <div class="mt-2 text-sm text-red-700">
              <ul class="list-disc pl-5 space-y-1">
                <li>Mật khẩu được lưu trữ dưới dạng văn bản thuần túy</li>
                <li>Dữ liệu người dùng được lưu trữ trong localStorage</li>
                <li>Không yêu cầu mật khẩu phức tạp khi đăng ký</li>
                <li>Lỗ hổng SQL injection</li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <div class="sm:mx-auto sm:w-full sm:max-w-md">
          <h2 class="text-center text-3xl font-extrabold text-gray-900">
            {{ isRegister ? 'Tạo tài khoản' : 'Đăng nhập' }}
          </h2>
        </div>

        <form class="space-y-6 mt-8" @submit.prevent="handleSubmit">
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">
              Tên đăng nhập
            </label>
            <div class="mt-1">
              <input
                id="username"
                v-model="form.username"
                name="username"
                type="text"
                required
                class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">
              Mật khẩu
            </label>
            <div class="mt-1">
              <input
                id="password"
                v-model="form.password"
                name="password"
                type="password"
                required
                class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
          </div>

          <!-- Role Selection (Security Issue) -->
          <div v-if="isRegister">
            <label for="role" class="block text-sm font-medium text-gray-700">
              Vai trò
              <span class="text-red-500 text-xs ml-1">(Lỗ hổng: Vai trò có thể được chọn)</span>
            </label>
            <div class="mt-1">
              <select
                id="role"
                v-model="form.role"
                name="role"
                class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              >
                <option value="user">User</option>
                <option value="moderator">Moderator</option>
                <option value="admin">Admin</option>
              </select>
            </div>
          </div>

          <div v-if="error" class="rounded-md bg-red-50 p-4">
            <div class="flex">
              <div class="ml-3">
                <h3 class="text-sm font-medium text-red-800">
                  {{ error }}
                </h3>
              </div>
            </div>
          </div>

          <div>
            <button
              type="submit"
              :disabled="isLoading"
              class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-orange-600 hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500 disabled:opacity-50"
            >
              <svg
                v-if="isLoading"
                class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                ></circle>
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              {{ isRegister ? 'Tạo tài khoản' : 'Đăng nhập' }}
            </button>
          </div>
        </form>

        <div class="mt-6">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300"></div>
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-2 bg-white text-gray-500">
                {{ isRegister ? 'Đã có tài khoản?' : "Chưa có tài khoản?" }}
              </span>
            </div>
          </div>

          <div class="mt-6">
            <button
              @click="toggleMode"
              class="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
            >
              {{ isRegister ? 'Đăng nhập' : 'Tạo tài khoản' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useInsecureAuthStore } from '@/stores/insecureAuth'

const router = useRouter()
const authStore = useInsecureAuthStore()

const isRegister = ref(false)
const isLoading = ref(false)
const error = ref('')

const form = ref({
  username: '',
  password: '',
  role: 'user'
})

function toggleMode() {
  isRegister.value = !isRegister.value
  error.value = ''
  form.value = {
    username: '',
    password: '',
    role: 'user'
  }
}

async function handleSubmit() {
  if (!form.value.username || !form.value.password) {
    error.value = 'Please fill in all fields'
    return
  }

  isLoading.value = true
  error.value = ''

  try {
    let result
    if (isRegister.value) {
      result = await authStore.register({
        username: form.value.username,
        password: form.value.password,
        role: form.value.role
      })
    } else {
      result = await authStore.login({
        username: form.value.username,
        password: form.value.password
      })
    }

    if (result.success) {
      // Successful authentication - go to products
      router.push('/v1/products')
    } else {
      error.value = result.error || 'Authentication failed'
    }
  } catch (err: any) {
    error.value = 'Connection error occurred'
  } finally {
    isLoading.value = false
  }
}
</script> 