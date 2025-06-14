<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-gray-100">
    <!-- Header -->
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 py-4">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-8">
            <router-link to="/" class="text-gray-900 hover:text-gray-700">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
              </svg>
            </router-link>
            <h1 class="text-2xl font-bold text-gray-900">API v1 <span class="text-red-600">(Không Bảo Mật)</span></h1>
          </div>
          <div class="flex items-center space-x-4">
            <div v-if="currentUser" class="flex items-center space-x-3">
              <span class="text-sm text-gray-700">{{ currentUser.username }}</span>
              <span class="px-2 py-1 text-xs font-medium bg-red-100 text-red-800 rounded-full">
                {{ currentUser.role }}
              </span>
              <button @click="logout" class="text-red-600 hover:text-red-700 text-sm font-medium">
                Đăng Xuất
              </button>
            </div>
            <button v-else @click="showLoginForm = true" class="text-red-600 hover:text-red-700 text-sm font-medium">
              Đăng Nhập
            </button>
          </div>
        </div>
      </div>
    </nav>

    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Warning Banner -->
      <div class="mb-8 bg-red-50 border-l-4 border-red-500 p-4 rounded-r-lg">
        <div class="flex">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
            </svg>
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">Cảnh Báo Bảo Mật</h3>
            <p class="mt-1 text-sm text-red-700">
              API v1 cố tình có các lỗ hổng bảo mật để thực hành tấn công. Xem file <code class="bg-red-100 px-1 rounded">SECURITY_TEST_PAYLOADS.md</code> để biết chi tiết.
            </p>
          </div>
        </div>
      </div>

      <!-- Login Form Modal -->
      <div v-if="showLoginForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4 transform transition-all">
          <div class="p-6">
            <div class="flex justify-between items-center mb-4">
              <h2 class="text-xl font-semibold text-gray-900">
                {{ isRegistering ? 'Đăng Ký Tài Khoản' : 'Đăng Nhập' }}
              </h2>
              <button @click="showLoginForm = false" class="text-gray-400 hover:text-gray-500">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
                </svg>
              </button>
            </div>
            
            <form @submit.prevent="handleAuth" class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">Username</label>
                <div class="mt-1 relative">
                  <input
                    v-model="loginForm.username"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                    :placeholder="isRegistering ? 'Chọn username' : 'Thử SQL injection: admin\' OR \'1\'=\'1'"
                  />
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                    </svg>
                  </div>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700">Password</label>
                <div class="mt-1 relative">
                  <input
                    v-model="loginForm.password"
                    type="password"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                    :placeholder="isRegistering ? 'Chọn password' : 'Thử: \' OR \'1\'=\'1'"
                  />
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                    </svg>
                  </div>
                </div>
              </div>

              <div v-if="isRegistering">
                <label class="block text-sm font-medium text-gray-700">Role</label>
                <div class="mt-1">
                  <select
                    v-model="loginForm.role"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                  >
                    <option value="user">User</option>
                    <option value="admin">Admin</option>
                    <option value="moderator">Moderator</option>
                  </select>
                  <p class="mt-1 text-sm text-red-600">💡 API v1 cho phép chọn role khi đăng ký!</p>
                </div>
              </div>

              <div class="mt-6 flex items-center justify-between">
                <button
                  type="button"
                  @click="isRegistering = !isRegistering"
                  class="text-sm text-red-600 hover:text-red-700"
                >
                  {{ isRegistering ? 'Đã có tài khoản? Đăng nhập' : 'Chưa có tài khoản? Đăng ký' }}
                </button>
                
                <button
                  type="submit"
                  class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                >
                  {{ isRegistering ? 'Đăng Ký' : 'Đăng Nhập' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <!-- Main Content -->
      <div class="space-y-8">
        <!-- Search Form -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h2 class="text-lg font-medium text-gray-900 mb-4">Tìm Kiếm Sản Phẩm</h2>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Tên sản phẩm</label>
              <div class="relative">
                <input
                  v-model="searchForm.name"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                  placeholder="Thử XSS: <script>alert('XSS')</script>"
                />
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                  </svg>
                </div>
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Mô tả</label>
              <div class="relative">
                <input
                  v-model="searchForm.description"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                  placeholder="Thử SQL injection: %' OR '1'='1"
                />
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                  </svg>
                </div>
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Giá</label>
              <div class="relative">
                <input
                  v-model="searchForm.price"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                  placeholder="Thử SQL injection: 0 OR 1=1"
                />
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                  </svg>
                </div>
              </div>
            </div>
          </div>
          <div class="mt-4">
            <button
              @click="searchProducts"
              class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
            >
              Tìm Kiếm
            </button>
          </div>
        </div>

        <!-- Products Grid -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <div class="flex justify-between items-center mb-6">
            <h2 class="text-lg font-medium text-gray-900">Danh Sách Sản Phẩm</h2>
            <button
              v-if="currentUser"
              @click="showCreateForm = true"
              class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
            >
              Thêm Sản Phẩm
            </button>
          </div>

          <!-- Loading State -->
          <div v-if="isLoading" class="py-12">
            <div class="flex justify-center">
              <svg class="animate-spin h-8 w-8 text-red-600" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
            </div>
            <p class="text-center mt-4 text-sm text-gray-500">Đang tải sản phẩm...</p>
          </div>

          <!-- Products List -->
          <div v-else-if="products.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
              v-for="product in products"
              :key="product.ID"
              class="border border-gray-200 rounded-lg p-4 hover:shadow-md transition-shadow"
            >
              <div class="flex flex-col h-full">
                <h3 class="text-lg font-medium text-gray-900" v-html="product.name"></h3>
                <p class="mt-1 text-sm text-gray-500 flex-grow" v-html="product.description"></p>
                <div class="mt-4 flex items-center justify-between">
                  <span class="text-lg font-bold text-red-600">{{ formatPrice(product.price) }}đ</span>
                  <div class="flex space-x-2">
                    <button
                      v-if="currentUser"
                      @click="deleteProduct(product.ID)"
                      class="px-3 py-1 bg-red-100 text-red-700 text-sm font-medium rounded-md hover:bg-red-200"
                    >
                      Xóa
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Empty State -->
          <div v-else class="py-12 text-center">
            <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"/>
            </svg>
            <p class="mt-4 text-sm text-gray-500">Không có sản phẩm nào</p>
          </div>
        </div>

        <!-- Create Product Modal -->
        <div v-if="showCreateForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4">
            <div class="p-6">
              <div class="flex justify-between items-center mb-4">
                <h2 class="text-xl font-semibold text-gray-900">Thêm Sản Phẩm Mới</h2>
                <button @click="showCreateForm = false" class="text-gray-400 hover:text-gray-500">
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
                  </svg>
                </button>
              </div>
              
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700">Tên sản phẩm</label>
                  <div class="mt-1 relative">
                    <input
                      v-model="createForm.name"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                      placeholder="Thử XSS: <script>alert('hack')</script>"
                    />
                    <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                      <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                      </svg>
                    </div>
                  </div>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">Mô tả</label>
                  <div class="mt-1 relative">
                    <textarea
                      v-model="createForm.description"
                      rows="3"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                      placeholder="Thử XSS: <img src=x onerror=alert('hack')>"
                    ></textarea>
                    <div class="absolute top-0 right-0 pr-3 pt-2 flex items-center pointer-events-none">
                      <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                      </svg>
                    </div>
                  </div>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700">Giá</label>
                  <div class="mt-1 relative">
                    <input
                      v-model="createForm.price"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-red-500 focus:border-red-500"
                      placeholder="Thử giá âm: -100000"
                    />
                    <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                      <svg class="h-5 w-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                      </svg>
                    </div>
                  </div>
                </div>
              </div>

              <div class="mt-6 flex justify-end space-x-3">
                <button
                  @click="showCreateForm = false"
                  class="px-4 py-2 border border-gray-300 text-gray-700 text-sm font-medium rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                >
                  Hủy
                </button>
                <button
                  @click="createProduct"
                  class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                >
                  Thêm
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { insecureAPI } from '@/services/api'

const router = useRouter()
const isLoading = ref(false)

// Auth
const currentUser = ref<any>(null)
const showLoginForm = ref(false)
const isRegistering = ref(false)
const loginForm = ref({
  username: '',
  password: '',
  role: 'user'
})

// Products
const products = ref<any[]>([])
const showCreateForm = ref(false)
const searchForm = ref({
  name: '',
  description: '',
  price: ''
})
const createForm = ref({
  name: '',
  description: '',
  price: ''
})

async function handleAuth() {
  try {
    if (isRegistering.value) {
      const response = await insecureAPI.register(
        loginForm.value.username,
        loginForm.value.password,
        loginForm.value.role
      )
      currentUser.value = response.data
    } else {
      const response = await insecureAPI.login(
        loginForm.value.username,
        loginForm.value.password
      )
      currentUser.value = response.data.user
    }
    showLoginForm.value = false
  } catch (error: any) {
    alert(error.response?.data?.error || 'Authentication failed')
  }
}

async function logout() {
  try {
    await insecureAPI.logout()
    currentUser.value = null
  } catch (error: any) {
    alert(error.response?.data?.error || 'Logout failed')
  }
}

async function searchProducts() {
  isLoading.value = true
  try {
    const response = await insecureAPI.listProducts({
      name: searchForm.value.name,
      description: searchForm.value.description,
      price: searchForm.value.price
    })
    products.value = response.data
  } catch (error: any) {
    alert(error.response?.data?.error || 'Failed to fetch products')
  } finally {
    isLoading.value = false
  }
}

async function createProduct() {
  try {
    await insecureAPI.createProduct({
      name: createForm.value.name,
      description: createForm.value.description,
      price: parseFloat(createForm.value.price) || 0
    })
    showCreateForm.value = false
    searchProducts()
  } catch (error: any) {
    alert(error.response?.data?.error || 'Failed to create product')
  }
}

async function deleteProduct(id: number) {
  if (!confirm('Bạn có chắc muốn xóa sản phẩm này?')) return
  
  try {
    await insecureAPI.deleteProduct(id)
    searchProducts()
  } catch (error: any) {
    alert(error.response?.data?.error || 'Failed to delete product')
  }
}

function formatPrice(price: number) {
  return new Intl.NumberFormat('vi-VN').format(price)
}

onMounted(() => {
  searchProducts()
})
</script> 