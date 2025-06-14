import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { secureAPI } from '@/services/api'

interface Product {
  ID: number
  name: string
  description: string
  price: number
  CreatedAt: string
  UpdatedAt: string
}

interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  total_pages: number
  limit: number
}

export const useProductStore = defineStore('product', () => {
  // State
  const products = ref<Product[]>([])
  const currentProduct = ref<Product | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  
  // Pagination state
  const currentPage = ref(1)
  const totalPages = ref(0)
  const totalProducts = ref(0)
  const limit = ref(10)

  // Getters
  const hasProducts = computed(() => products.value.length > 0)
  const hasMorePages = computed(() => currentPage.value < totalPages.value)

  // Actions
  async function fetchProducts(page = 1, itemsPerPage = 10) {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await secureAPI.listProducts(page, itemsPerPage)
      
      products.value = response.data.data
      currentPage.value = response.data.page
      totalPages.value = response.data.total_pages
      totalProducts.value = response.data.total
      limit.value = response.data.limit
      
      return { success: true, data: response.data }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Failed to fetch products'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  async function createProduct(productData: Omit<Product, 'ID' | 'CreatedAt' | 'UpdatedAt'>) {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await secureAPI.createProduct(productData)
      
      // Thêm sản phẩm mới vào đầu danh sách
      products.value.unshift(response.data)
      totalProducts.value += 1
      
      return { success: true, product: response.data }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Failed to create product'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  async function updateProduct(id: number, productData: Omit<Product, 'ID' | 'CreatedAt' | 'UpdatedAt'>) {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await secureAPI.updateProduct(id, productData)
      
      // Cập nhật sản phẩm trong danh sách
      const index = products.value.findIndex(p => p.ID === id)
      if (index !== -1) {
        products.value[index] = response.data
      }
      
      // Cập nhật currentProduct nếu đang xem sản phẩm này
      if (currentProduct.value?.ID === id) {
        currentProduct.value = response.data
      }
      
      return { success: true, product: response.data }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Failed to update product'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  async function deleteProduct(id: number) {
    isLoading.value = true
    error.value = null
    
    try {
      await secureAPI.deleteProduct(id)
      
      // Xóa sản phẩm khỏi danh sách
      products.value = products.value.filter(p => p.ID !== id)
      totalProducts.value -= 1
      
      // Clear currentProduct nếu đang xem sản phẩm này
      if (currentProduct.value?.ID === id) {
        currentProduct.value = null
      }
      
      return { success: true }
    } catch (err: any) {
      const errorMessage = err.response?.data?.error || 'Failed to delete product'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      isLoading.value = false
    }
  }

  function setCurrentProduct(product: Product | null) {
    currentProduct.value = product
  }

  function clearError() {
    error.value = null
  }

  function clearProducts() {
    products.value = []
    currentProduct.value = null
    currentPage.value = 1
    totalPages.value = 0
    totalProducts.value = 0
  }

  return {
    // State
    products,
    currentProduct,
    isLoading,
    error,
    currentPage,
    totalPages,
    totalProducts,
    limit,
    
    // Getters
    hasProducts,
    hasMorePages,
    
    // Actions
    fetchProducts,
    createProduct,
    updateProduct,
    deleteProduct,
    setCurrentProduct,
    clearError,
    clearProducts,
  }
}) 