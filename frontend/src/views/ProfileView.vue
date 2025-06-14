<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <nav class="bg-white shadow mb-6">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <router-link to="/" class="text-xl font-semibold text-gray-900 hover:text-gray-700">
              ← Dashboard
            </router-link>
          </div>
          <div class="flex items-center space-x-4">
            <span class="text-sm text-gray-700">{{ user?.username }}</span>
            <button
              @click="logout"
              class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-md text-sm font-medium"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </nav>

    <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Page Header -->
      <div class="mb-8">
        <h1 class="text-2xl font-bold text-gray-900">Thông tin cá nhân</h1>
        <p class="mt-2 text-sm text-gray-600">
          Quản lý thông tin tài khoản và cài đặt bảo mật của bạn.
        </p>
      </div>

      <!-- Success Message -->
      <div v-if="successMessage" class="mb-6 bg-green-50 border border-green-200 rounded-md p-4">
        <div class="flex">
          <div class="ml-3">
            <h3 class="text-sm font-medium text-green-800">{{ successMessage }}</h3>
          </div>
        </div>
      </div>

      <!-- Error Message -->
      <div v-if="error" class="mb-6 bg-red-50 border border-red-200 rounded-md p-4">
        <div class="flex">
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">{{ error }}</h3>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        <!-- User Information Card -->
        <div class="bg-white shadow rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
              Thông tin tài khoản
            </h3>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700">User ID</label>
                <div class="mt-1 text-sm text-gray-900 bg-gray-50 px-3 py-2 rounded-md">
                  {{ user?.id }}
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">Tên đăng nhập</label>
                <div class="mt-1 text-sm text-gray-900 bg-gray-50 px-3 py-2 rounded-md">
                  {{ user?.username }}
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">Vai trò</label>
                <div class="mt-1">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="roleClass">
                    {{ user?.role }}
                  </span>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">Trạng thái xác thực</label>
                <div class="mt-1">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                    Đã xác thực
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Change Password Card -->
        <div class="bg-white shadow rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
              Đổi mật khẩu
            </h3>
            <p class="text-sm text-gray-600 mb-6">
              Đảm bảo tài khoản của bạn sử dụng mật khẩu mạnh để giữ an toàn.
            </p>
            
            <form @submit.prevent="handleChangePassword" class="space-y-4">
              <div>
                <label for="currentPassword" class="block text-sm font-medium text-gray-700">
                  Mật khẩu hiện tại
                </label>
                <input
                  id="currentPassword"
                  v-model="passwordForm.currentPassword"
                  type="password"
                  required
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
                  placeholder="Nhập mật khẩu hiện tại"
                />
              </div>
              
              <div>
                <label for="newPassword" class="block text-sm font-medium text-gray-700">
                  Mật khẩu mới
                </label>
                <input
                  id="newPassword"
                  v-model="passwordForm.newPassword"
                  type="password"
                  required
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
                  placeholder="Nhập mật khẩu mới"
                />
                <p class="mt-1 text-xs text-gray-500">
                  Mật khẩu phải có ít nhất 8 ký tự, bao gồm chữ hoa, chữ thường, số và ký tự đặc biệt
                </p>
              </div>
              
              <div>
                <label for="confirmPassword" class="block text-sm font-medium text-gray-700">
                  Xác nhận mật khẩu mới
                </label>
                <input
                  id="confirmPassword"
                  v-model="passwordForm.confirmPassword"
                  type="password"
                  required
                  class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
                  placeholder="Nhập lại mật khẩu mới"
                />
                <p v-if="passwordForm.newPassword && passwordForm.confirmPassword && !passwordsMatch" 
                   class="mt-1 text-xs text-red-600">
                  Mật khẩu xác nhận không khớp
                </p>
              </div>
              
              <div class="flex justify-end">
                <button
                  type="submit"
                  :disabled="isLoading || !passwordsMatch || !isFormValid"
                  class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
                >
                  <span v-if="isLoading" class="mr-2">
                    <svg class="animate-spin h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                  </span>
                  {{ isLoading ? 'Đang cập nhật...' : 'Đổi mật khẩu' }}
                </button>
              </div>
            </form>
          </div>
        </div>

        <!-- Security Information Card -->
        <div class="bg-white shadow rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
              Thông tin bảo mật
            </h3>
            
            <div class="space-y-4">
              <div class="flex items-center justify-between py-3 border-b border-gray-200">
                <div>
                  <h4 class="text-sm font-medium text-gray-900">JWT Authentication</h4>
                  <p class="text-sm text-gray-600">Token được lưu trong HTTP-only cookie</p>
                </div>
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  Hoạt động
                </span>
              </div>
              
              <div class="flex items-center justify-between py-3 border-b border-gray-200">
                <div>
                  <h4 class="text-sm font-medium text-gray-900">Role-based Access Control</h4>
                  <p class="text-sm text-gray-600">Phân quyền dựa trên vai trò người dùng</p>
                </div>
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  {{ user?.role }}
                </span>
              </div>
              
              <div class="flex items-center justify-between py-3 border-b border-gray-200">
                <div>
                  <h4 class="text-sm font-medium text-gray-900">Rate Limiting</h4>
                  <p class="text-sm text-gray-600">Giới hạn số lượng request (10 req/sec)</p>
                </div>
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-yellow-100 text-yellow-800">
                  Hoạt động
                </span>
              </div>
              
              <div class="flex items-center justify-between py-3">
                <div>
                  <h4 class="text-sm font-medium text-gray-900">Password Security</h4>
                  <p class="text-sm text-gray-600">Mật khẩu được mã hóa bằng bcrypt</p>
                </div>
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  Bảo mật
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const error = ref<string | null>(null)
const successMessage = ref<string | null>(null)
const isLoading = ref(false)

const passwordForm = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const user = computed(() => authStore.user)

const roleClass = computed(() => {
  switch (user.value?.role) {
    case 'admin':
      return 'bg-red-100 text-red-800'
    case 'moderator':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-green-100 text-green-800'
  }
})

const passwordsMatch = computed(() => {
  return passwordForm.value.newPassword === passwordForm.value.confirmPassword
})

const isFormValid = computed(() => {
  return passwordForm.value.currentPassword && 
         passwordForm.value.newPassword && 
         passwordForm.value.confirmPassword &&
         passwordsMatch.value
})

async function handleChangePassword() {
  if (!isFormValid.value) {
    error.value = 'Vui lòng điền đầy đủ thông tin'
    return
  }

  error.value = null
  successMessage.value = null
  isLoading.value = true

  try {
    const result = await authStore.changePassword({
      currentPassword: passwordForm.value.currentPassword,
      newPassword: passwordForm.value.newPassword
    })

    if (result.success) {
      successMessage.value = result.message || 'Đổi mật khẩu thành công'
      // Reset form
      passwordForm.value = {
        currentPassword: '',
        newPassword: '',
        confirmPassword: ''
      }
    } else {
      error.value = result.error || 'Đổi mật khẩu thất bại'
    }
  } catch (err) {
    error.value = 'Có lỗi xảy ra khi đổi mật khẩu'
  } finally {
    isLoading.value = false
  }
}

async function logout() {
  await authStore.logout()
  router.push('/login')
}
</script> 