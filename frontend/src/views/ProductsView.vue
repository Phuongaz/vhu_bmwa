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
            <span class="text-sm text-gray-700">{{ authStore.user?.username }}</span>
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

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="sm:flex sm:items-center">
        <div class="sm:flex-auto">
          <h1 class="text-2xl font-semibold text-gray-900">Quản lý Sản phẩm</h1>
          <p class="mt-2 text-sm text-gray-700">
            Danh sách tất cả sản phẩm trong hệ thống. 
            {{ authStore.isAdmin ? 'Bạn có thể thêm, sửa, xóa sản phẩm.' : 'Bạn chỉ có thể xem danh sách.' }}
          </p>
        </div>
        <div class="mt-4 sm:mt-0 sm:ml-16 sm:flex-none">
          <button
            v-if="authStore.isAdmin"
            @click="showCreateModal = true"
            type="button"
            class="inline-flex items-center justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:w-auto"
          >
            Thêm sản phẩm
          </button>
        </div>
      </div>

      <!-- Error Display -->
      <div v-if="error" class="mt-4 bg-red-50 border border-red-200 rounded-md p-4">
        <div class="flex">
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">{{ error }}</h3>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="productStore.isLoading" class="mt-8 text-center">
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
        <p class="mt-2 text-sm text-gray-600">Đang tải...</p>
      </div>

      <!-- Products Table -->
      <div v-else class="mt-8 flex flex-col">
        <div class="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
            <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
              <table class="min-w-full divide-y divide-gray-300">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      ID
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Tên sản phẩm
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Mô tả
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Giá
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Ngày tạo
                    </th>
                    <th v-if="authStore.isAdmin" class="relative px-6 py-3">
                      <span class="sr-only">Hành động</span>
                    </th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="product in productStore.products" :key="product.ID">
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      {{ product.ID }}
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      {{ product.name }}
                    </td>
                    <td class="px-6 py-4 text-sm text-gray-900">
                      {{ product.description || 'Không có mô tả' }}
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      {{ formatPrice(product.price) }}
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      {{ formatDate(product.CreatedAt) }}
                    </td>
                    <td v-if="authStore.isAdmin" class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                      <button
                        @click="editProduct(product)"
                        class="text-indigo-600 hover:text-indigo-900 mr-4"
                      >
                        Sửa
                      </button>
                      <button
                        @click="deleteProduct(product.ID)"
                        class="text-red-600 hover:text-red-900"
                      >
                        Xóa
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Pagination -->
        <div v-if="productStore.totalPages > 1" class="mt-6 flex items-center justify-between">
          <div class="flex-1 flex justify-between sm:hidden">
            <button
              @click="prevPage"
              :disabled="productStore.currentPage <= 1"
              class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
            >
              Trước
            </button>
            <button
              @click="nextPage"
              :disabled="!productStore.hasMorePages"
              class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
            >
              Tiếp
            </button>
          </div>
          <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div>
              <p class="text-sm text-gray-700">
                Hiển thị
                <span class="font-medium">{{ (productStore.currentPage - 1) * productStore.limit + 1 }}</span>
                đến
                <span class="font-medium">{{ Math.min(productStore.currentPage * productStore.limit, productStore.totalProducts) }}</span>
                trong
                <span class="font-medium">{{ productStore.totalProducts }}</span>
                kết quả
              </p>
            </div>
            <div>
              <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                <button
                  @click="prevPage"
                  :disabled="productStore.currentPage <= 1"
                  class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
                >
                  Trước
                </button>
                <button
                  @click="nextPage"
                  :disabled="!productStore.hasMorePages"
                  class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
                >
                  Tiếp
                </button>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Create/Edit Product Modal -->
    <div v-if="showCreateModal || showEditModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ showCreateModal ? 'Thêm sản phẩm mới' : 'Sửa sản phẩm' }}
          </h3>
          <form @submit.prevent="saveProduct" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">Tên sản phẩm</label>
              <input
                v-model="productForm.name"
                type="text"
                required
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Mô tả</label>
              <textarea
                v-model="productForm.description"
                rows="3"
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
              ></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Giá (VND)</label>
              <input
                v-model.number="productForm.price"
                type="number"
                step="0.01"
                min="0"
                required
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
              >
            </div>
            <div class="flex justify-end space-x-3 pt-4">
              <button
                type="button"
                @click="closeModal"
                class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-200 border border-gray-300 rounded-md hover:bg-gray-300"
              >
                Hủy
              </button>
              <button
                type="submit"
                :disabled="productStore.isLoading"
                class="px-4 py-2 text-sm font-medium text-white bg-indigo-600 border border-transparent rounded-md hover:bg-indigo-700 disabled:opacity-50"
              >
                {{ productStore.isLoading ? 'Đang lưu...' : (showCreateModal ? 'Thêm' : 'Cập nhật') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useProductStore } from '@/stores/product'
import type { Product } from '@/services/api'

const router = useRouter()
const authStore = useAuthStore()
const productStore = useProductStore()

const error = ref<string | null>(null)
const showCreateModal = ref(false)
const showEditModal = ref(false)
const editingProduct = ref<Product | null>(null)

const productForm = ref({
  name: '',
  description: '',
  price: 0
})

onMounted(async () => {
  await loadProducts()
})

async function loadProducts() {
  const result = await productStore.fetchProducts()
  if (!result.success) {
    error.value = result.error || 'Failed to load products'
  }
}

async function prevPage() {
  if (productStore.currentPage > 1) {
    await productStore.fetchProducts(productStore.currentPage - 1)
  }
}

async function nextPage() {
  if (productStore.hasMorePages) {
    await productStore.fetchProducts(productStore.currentPage + 1)
  }
}

function editProduct(product: Product) {
  editingProduct.value = product
  productForm.value = {
    name: product.name,
    description: product.description,
    price: product.price
  }
  showEditModal.value = true
}

async function deleteProduct(id: number) {
  if (confirm('Bạn có chắc chắn muốn xóa sản phẩm này?')) {
    const result = await productStore.deleteProduct(id)
    if (!result.success) {
      error.value = result.error || 'Failed to delete product'
    }
  }
}

async function saveProduct() {
  let result
  if (showCreateModal.value) {
    result = await productStore.createProduct(productForm.value)
  } else if (editingProduct.value) {
    result = await productStore.updateProduct(editingProduct.value.ID, productForm.value)
  }

  if (result?.success) {
    closeModal()
  } else {
    error.value = result?.error || 'Failed to save product'
  }
}

function closeModal() {
  showCreateModal.value = false
  showEditModal.value = false
  editingProduct.value = null
  productForm.value = {
    name: '',
    description: '',
    price: 0
  }
}

function formatPrice(price: number): string {
  return new Intl.NumberFormat('vi-VN', {
    style: 'currency',
    currency: 'VND'
  }).format(price)
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('vi-VN')
}

async function logout() {
  await authStore.logout()
  router.push('/login')
}
</script> 