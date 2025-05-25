<template>
  <div class="products-container">
    <div class="products-header">
      <h1>Quản Lý Sản Phẩm</h1>
      <button v-if="isAdmin" @click="showAddModal" class="btn-add">
        <i class="fas fa-plus"></i> Thêm Sản Phẩm
      </button>
    </div>

    <div class="filters">
      <div class="search-box">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Tìm kiếm sản phẩm..."
          @input="handleSearch"
        />
        <i class="fas fa-search"></i>
      </div>
      <div class="filter-options">
        <select v-model="sortBy" @change="handleSort">
          <option value="name">Tên A-Z</option>
          <option value="-name">Tên Z-A</option>
          <option value="price">Giá Thấp-Cao</option>
          <option value="-price">Giá Cao-Thấp</option>
        </select>
      </div>
    </div>

    <div v-if="loading" class="loading">
      <div class="spinner"></div>
      <span>Đang tải sản phẩm...</span>
    </div>
    <div v-else-if="filteredProducts.length === 0" class="no-products">
      <i class="fas fa-box-open"></i>
      <p>Không tìm thấy sản phẩm nào</p>
    </div>
    <div v-else class="products-grid">
      <div v-for="product in filteredProducts" :key="product.id" class="product-card">
        <div class="product-image">
          <img :src="product.image || '/placeholder.png'" :alt="product.name">
        </div>
        <div class="product-info">
          <h3>{{ product.name }}</h3>
          <p class="description">{{ product.description }}</p>
          <div class="price">{{ formatPrice(product.price) }}đ</div>
        </div>
        <div v-if="isAdmin" class="product-actions">
          <button @click="editProduct(product)" class="btn-edit">
            <i class="fas fa-edit"></i>
          </button>
          <button @click="confirmDelete(product)" class="btn-delete">
            <i class="fas fa-trash"></i>
          </button>
        </div>
      </div>
    </div>

    <!-- Modal thêm/sửa sản phẩm -->
    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <h2>{{ isEditing ? 'Sửa Sản Phẩm' : 'Thêm Sản Phẩm Mới' }}</h2>
        <form @submit.prevent="handleSubmit">
          <div class="form-group">
            <label>Tên sản phẩm</label>
            <input
              type="text"
              v-model="productForm.name"
              required
              placeholder="Nhập tên sản phẩm"
            />
          </div>
          <div class="form-group">
            <label>Mô tả</label>
            <textarea
              v-model="productForm.description"
              rows="3"
              placeholder="Nhập mô tả sản phẩm"
            ></textarea>
          </div>
          <div class="form-group">
            <label>Giá</label>
            <input
              type="number"
              v-model="productForm.price"
              required
              min="0"
              step="1000"
              placeholder="Nhập giá sản phẩm"
            />
          </div>
          <div class="form-group">
            <label>Hình ảnh</label>
            <input
              type="file"
              @change="handleImageUpload"
              accept="image/*"
            />
          </div>
          <div class="modal-actions">
            <button type="button" @click="closeModal" class="btn-cancel">Hủy</button>
            <button type="submit" class="btn-submit" :disabled="submitting">
              {{ submitting ? 'Đang xử lý...' : (isEditing ? 'Cập nhật' : 'Thêm mới') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <div v-if="showDeleteModal" class="modal">
      <div class="modal-content">
        <h2>Xác Nhận Xóa</h2>
        <p>Bạn có chắc chắn muốn xóa sản phẩm "{{ selectedProduct?.name }}"?</p>
        <div class="modal-actions">
          <button @click="closeDeleteModal" class="btn-cancel">Hủy</button>
          <button @click="deleteProduct" class="btn-delete" :disabled="deleting">
            {{ deleting ? 'Đang xóa...' : 'Xóa' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Products',
  data() {
    return {
      products: [],
      loading: true,
      searchQuery: '',
      sortBy: 'name',
      showModal: false,
      showDeleteModal: false,
      isEditing: false,
      submitting: false,
      deleting: false,
      selectedProduct: null,
      productForm: {
        name: '',
        description: '',
        price: 0,
        image: null
      }
    }
  },
  computed: {
    isAdmin() {
      return this.$store.state.auth.user?.role === 'admin'
    },
    filteredProducts() {
      let result = [...this.products]
      
      // Tìm kiếm
      if (this.searchQuery) {
        const query = this.searchQuery.toLowerCase()
        result = result.filter(product => 
          product.name.toLowerCase().includes(query) ||
          product.description.toLowerCase().includes(query)
        )
      }
      
      const [field, order] = this.sortBy.startsWith('-') 
        ? [this.sortBy.slice(1), 'desc']
        : [this.sortBy, 'asc']
      
      result.sort((a, b) => {
        if (order === 'asc') {
          return a[field] > b[field] ? 1 : -1
        }
        return a[field] < b[field] ? 1 : -1
      })
      
      return result
    }
  },
  async created() {
    await this.loadProducts()
  },
  methods: {
    async loadProducts() {
      try {
        this.loading = true
        const response = await this.$store.dispatch('products/getAll')
        this.products = response.data
      } catch (error) {
        console.error('Failed to load products:', error)
      } finally {
        this.loading = false
      }
    },
    formatPrice(price) {
      return new Intl.NumberFormat('vi-VN').format(price)
    },
    handleSearch() {
      // Debounce implementation can be added here
    },
    handleSort() {
      // Additional sort logic if needed
    },
    showAddModal() {
      this.isEditing = false
      this.productForm = {
        name: '',
        description: '',
        price: 0,
        image: null
      }
      this.showModal = true
    },
    editProduct(product) {
      this.isEditing = true
      this.selectedProduct = product
      this.productForm = { ...product }
      this.showModal = true
    },
    async handleSubmit() {
      try {
        this.submitting = true
        if (this.isEditing) {
          await this.$store.dispatch('products/update', {
            id: this.selectedProduct.id,
            ...this.productForm
          })
        } else {
          await this.$store.dispatch('products/create', this.productForm)
        }
        await this.loadProducts()
        this.closeModal()
      } catch (error) {
        console.error('Failed to save product:', error)
      } finally {
        this.submitting = false
      }
    },
    confirmDelete(product) {
      this.selectedProduct = product
      this.showDeleteModal = true
    },
    async deleteProduct() {
      try {
        this.deleting = true
        await this.$store.dispatch('products/delete', this.selectedProduct.id)
        await this.loadProducts()
        this.closeDeleteModal()
      } catch (error) {
        console.error('Failed to delete product:', error)
      } finally {
        this.deleting = false
      }
    },
    handleImageUpload(event) {
      const file = event.target.files[0]
      if (file) {
        this.productForm.image = URL.createObjectURL(file)
      }
    },
    closeModal() {
      this.showModal = false
      this.isEditing = false
      this.selectedProduct = null
      this.productForm = {
        name: '',
        description: '',
        price: 0,
        image: null
      }
    },
    closeDeleteModal() {
      this.showDeleteModal = false
      this.selectedProduct = null
      this.deleting = false
    }
  }
}
</script>

<style scoped>
.products-container {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.products-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.btn-add {
  padding: 0.75rem 1.5rem;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.filters {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
}

.search-box {
  flex: 1;
  position: relative;
}

.search-box input {
  width: 100%;
  padding: 0.75rem;
  padding-left: 2.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.search-box i {
  position: absolute;
  left: 1rem;
  top: 50%;
  transform: translateY(-50%);
  color: #666;
}

.filter-options select {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 1.5rem;
}

.product-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.product-image {
  height: 200px;
  overflow: hidden;
}

.product-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.product-info {
  padding: 1rem;
}

.product-info h3 {
  margin: 0;
  color: #333;
}

.description {
  color: #666;
  margin: 0.5rem 0;
  font-size: 0.875rem;
}

.price {
  font-weight: bold;
  color: #28a745;
  font-size: 1.25rem;
}

.product-actions {
  display: flex;
  padding: 1rem;
  gap: 0.5rem;
  border-top: 1px solid #eee;
}

.btn-edit, .btn-delete {
  flex: 1;
  padding: 0.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-edit {
  background: #007bff;
  color: white;
}

.btn-delete {
  background: #dc3545;
  color: white;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  width: 100%;
  max-width: 500px;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #666;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 2rem;
}

.btn-cancel {
  padding: 0.75rem 1.5rem;
  background: #6c757d;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-submit {
  padding: 0.75rem 1.5rem;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-submit:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.loading {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 4rem;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid #3498db;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-right: 1rem;
}

.no-products {
  text-align: center;
  padding: 4rem;
  color: #666;
}

.no-products i {
  font-size: 4rem;
  margin-bottom: 1rem;
  color: #ddd;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style> 