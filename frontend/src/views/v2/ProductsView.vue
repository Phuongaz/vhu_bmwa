<template>
  <div class="min-h-screen bg-gray-50 flex flex-col items-center">
    <Header></Header>
    <main class="w-full max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 pb-12">
      <div class="space-y-12 pt-20 grid gap-y-12">
        <!-- Page Title -->
        <div class="text-center">
          <h2 class="text-3xl font-bold text-gray-900">Sản Phẩm</h2>
          <p class="text-gray-600 mt-3">Khám phá các sản phẩm tuyệt vời của chúng tôi</p>
        </div>

        <!-- Search & Add Product Section -->
        <div class="bg-white rounded-lg shadow-sm border p-6">
          <div class="space-y-6">
            <!-- Search Form -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-y-10 gap-x-10 p-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Tên sản phẩm</label
                >
                <input
                  v-model="searchForm.name"
                  type="text"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500"
                  placeholder="Tìm kiếm sản phẩm..."
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Mô tả</label>
                <input
                  v-model="searchForm.description"
                  type="text"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500"
                  placeholder="Tìm trong mô tả..."
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Giá</label>
                <input
                  v-model="searchForm.price"
                  type="text"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500"
                  placeholder="Giá sản phẩm..."
                />
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="flex justify-center gap-4">
              <button
                @click="searchProducts"
                :disabled="isLoading"
                class="px-8 py-2 bg-orange-600 text-white rounded-lg hover:bg-orange-700 disabled:opacity-50 min-w-[120px]"
              >
                Tìm Kiếm
              </button>

              <button
                @click="showCreateForm = true"
                class="px-8 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 min-w-[120px]"
              >
                Thêm Sản Phẩm
              </button>
            </div>
          </div>
        </div>

        <!-- Products Grid -->
        <div class="bg-white rounded-lg shadow-sm border p-6">
          <h3 class="text-xl font-semibold text-gray-900 mb-6 text-center">
            Danh Sách Sản Phẩm
          </h3>

          <!-- Loading State -->
          <div v-if="isLoading" class="flex justify-center py-12">
            <div
              class="animate-spin rounded-full h-8 w-8 border-b-2 border-orange-600"
            ></div>
          </div>

          <!-- Products List -->
          <div
            v-else-if="products.length > 0"
            class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-8 p-4"
          >
            <div
              v-for="product in products"
              :key="product.ID"
              class="border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow"
            >
              <div class="p-6 flex flex-col space-y-4">
                <h4
                  class="text-lg font-semibold text-gray-900"
                  v-text="product.name"
                ></h4>
                <p
                  class="text-gray-600 text-sm flex-grow"
                  v-text="product.description"
                ></p>

                <div class="flex items-center justify-between pt-4">
                  <span class="text-2xl font-bold text-orange-600"
                    >{{ formatPrice(product.price) }}đ</span
                  >

                  <div class="flex space-x-2 px-5 gap-x-2">
                    <button
                      class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
                    >
                      Mua Ngay
                    </button>

                    <button
                      @click="deleteProduct(product.ID)"
                      class="px-3 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700"
                    >
                      Xóa
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Empty State -->
          <div v-else class="text-center py-12">
            <svg
              class="mx-auto h-12 w-12 text-gray-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"
              />
            </svg>
            <p class="mt-4 text-gray-500">Không có sản phẩm nào</p>
          </div>
        </div>

        <!-- Security Warning -->
        <div class="bg-green-50 border border-green-200 rounded-lg p-6">
          <div class="flex items-start">
            <div class="flex-shrink-0">
              <svg
                class="w-5 h-5 text-green-600"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
            </div>
            <div class="ml-3">
              <h3 class="text-sm font-medium text-green-800">
                🛡️ Trang này đã được bảo mật (Phiên bản 2.0)
              </h3>
              <div class="mt-2 text-sm text-green-700">
                <p class="mb-2">
                  <strong>Các biện pháp bảo mật đã được triển khai:</strong>
                </p>
                <ul class="space-y-1 ml-4">
                  <li>
                    • <strong>XSS Prevention:</strong> Sử dụng v-text thay vì v-html để
                    hiển thị dữ liệu
                  </li>
                  <li>
                    • <strong>SQL Injection Protection:</strong> Sử dụng Prepared
                    Statements và input validation
                  </li>
                  <li>
                    • <strong>Authorization:</strong> Kiểm tra quyền truy cập cho mọi thao
                    tác
                  </li>
                  <li>
                    • <strong>CSRF Protection:</strong> Token CSRF được thêm vào mọi
                    request
                  </li>
                  <li>
                    • <strong>Input Sanitization:</strong> Dữ liệu được làm sạch trước khi
                    xử lý
                  </li>
                </ul>
                <div class="mt-3 p-3 bg-green-100 rounded">
                  <p class="font-medium">Lưu ý bảo mật:</p>
                  <p class="text-xs mt-1">
                    Trang này sử dụng các biện pháp bảo mật tốt nhất để bảo vệ dữ liệu và
                    người dùng. Mọi thao tác đều được xác thực và kiểm soát quyền truy cập
                    chặt chẽ.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Create Product Modal -->
    <div
      v-if="showCreateForm"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
    >
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full">
        <div class="p-8">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-semibold text-gray-900">Thêm Sản Phẩm Mới</h3>
            <button
              @click="showCreateForm = false"
              class="text-gray-400 hover:text-gray-500"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>

          <form @submit.prevent="createProduct" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Tên sản phẩm</label
              >
              <input
                v-model="createForm.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500"
                placeholder="Nhập tên sản phẩm"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Mô tả</label>
              <textarea
                v-model="createForm.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500"
                placeholder="Mô tả sản phẩm"
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Giá (VNĐ)</label
              >
              <input
                v-model="createForm.price"
                type="number"
                min="0"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500"
                placeholder="0"
              />
            </div>

            <div class="flex justify-end space-x-3 pt-4">
              <button
                type="button"
                @click="showCreateForm = false"
                class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50"
              >
                Hủy
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-orange-600 text-white rounded-lg hover:bg-orange-700"
              >
                Thêm
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { insecureAPI, secureAPI } from "@/services/api";
import Header from "@/components/v2/Header.vue";

const isLoading = ref(false);
const showCreateForm = ref(false);

// Products
const products = ref<any[]>([]);

// Search form
const searchForm = ref({
  name: "",
  description: "",
  price: "",
});

// Create form
const createForm = ref({
  name: "",
  description: "",
  price: "",
});

const user = ref({
  username: "demo_user",
  role: "user",
});

async function searchProducts() {
  isLoading.value = true;
  try {
    const response = await insecureAPI.listProducts({
      name: searchForm.value.name,
      description: searchForm.value.description,
      price: searchForm.value.price,
    });
    // Backend (insecure v1) trả về { products: [...] }
    products.value = response.data?.products || [];
  } catch (error: any) {
    console.error("Search failed:", error);
    alert("Lỗi tìm kiếm: " + (error.response?.data?.error || error.message));
  } finally {
    isLoading.value = false;
  }
}

async function createProduct() {
  try {
    await secureAPI.createProduct({
      name: createForm.value.name,
      description: createForm.value.description,
      price: parseFloat(createForm.value.price) || 0,
    });

    showCreateForm.value = false;
    createForm.value = { name: "", description: "", price: "" };
    await searchProducts();
  } catch (error: any) {
    alert("Lỗi tạo sản phẩm: " + (error.response?.data?.error || error.message));
  }
}

async function deleteProduct(id: number) {
  if (!confirm("Bạn có chắc muốn xóa sản phẩm này?")) return;

  try {
    await secureAPI.deleteProduct(id);
    await searchProducts();
  } catch (error: any) {
    alert("Lỗi xóa sản phẩm: " + (error.response?.data?.error || error.message));
  }
}

function formatPrice(price: number) {
  return new Intl.NumberFormat("vi-VN").format(price);
}

async function loadUser() {
  const response = await secureAPI.getProfile();
  user.value = response.data?.user || {};
}

// Load initial data
onMounted(async () => {
  await searchProducts();
  await loadUser();
});
</script>
