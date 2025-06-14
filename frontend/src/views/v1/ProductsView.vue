<template>
  <div class="min-h-screen bg-gray-50 font-sans">
     <Header></Header>

    <div class="max-w-7xl mx-auto px-4 py-6">
      <div class="mb-8 p-4 bg-red-50 border-l-4 border-red-500 rounded-r-md">
        <div class="flex">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor">
              <path
                fill-rule="evenodd"
                d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">
              Các lỗ hổng bảo mật
            </h3>
            <div class="mt-2 text-sm text-red-700">
              <ul class="list-disc pl-5 space-y-1">
                <li>Không cần xác thực cho các endpoint API</li>
                <li>Không có xác thực</li>
                <li>Dữ liệu nhạy cảm được lưu trữ trong localStorage</li>
                <li>Không có giới hạn tốc độ cho các yêu cầu API</li>
                <li>XSS Trong Thêm Sản Phẩm</li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content -->
      <div class="bg-white shadow-sm rounded-lg">
        <!-- Search Form -->
        <div class="p-6 border-b">
          <h2 class="text-lg font-medium text-gray-900 mb-4">Tìm kiếm sản phẩm</h2>
          <form
            @submit.prevent="searchProducts"
            class="grid grid-cols-1 md:grid-cols-3 gap-4"
          >
            <div>
              <label class="block text-sm font-medium text-gray-700">Tên sản phẩm</label>
              <input
                v-model="searchForm.name"
                type="text"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Mô tả</label>
              <input
                v-model="searchForm.description"
                type="text"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Giá</label>
              <input
                v-model="searchForm.price"
                type="number"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-orange-500 focus:ring-orange-500 sm:text-sm"
              />
            </div>
            <div class="md:col-span-3">
              <button
                type="submit"
                :disabled="isLoading"
                class="w-full inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-orange-600 hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500"
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
                Tìm kiếm
              </button>
            </div>
          </form>
        </div>

        <!-- Products List -->
        <div class="p-6">
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-lg font-medium text-gray-900">Sản phẩm</h2>
            <button
              v-if="authStore.isAuthenticated"
              @click="showCreateForm = true"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-orange-600 hover:bg-orange-700"
            >
              Thêm sản phẩm
            </button>
          </div>

          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Tên sản phẩm
                  </th>
                  <th
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Mô tả
                  </th>
                  <th
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Giá
                  </th>
                  <th
                    class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Hành động
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="product in products" :key="product.ID">
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500" v-html="product.name">
                  </td>
                  <!-- <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {{ product.description }}
                  </td> -->
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500" v-html="product.description">
                  </td>
                  
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    ${{ product.price }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                    <button
                      @click="deleteProduct(product.id)"
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
    </div>

    <!-- Create Product Modal -->
    <div
      v-if="showCreateForm"
      class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center p-4"
    >
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full">
        <div class="p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Thêm sản phẩm</h3>
          <form @submit.prevent="createProduct" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">Tên sản phẩm</label>
              <input
                v-model="createForm.name"
                type="text"
                required
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Mô tả</label>
              <input
                v-model="createForm.description"
                type="text"
                required
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Giá</label>
              <input
                v-model="createForm.price"
                type="number"
                required
                step="0.01"
                min="0"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
            <div class="flex justify-end space-x-3">
              <button
                type="button"
                @click="showCreateForm = false"
                class="inline-flex justify-center py-2 px-4 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
              >
                Hủy
              </button>
              <button
                type="submit"
                class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-orange-600 hover:bg-orange-700"
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
import { useRouter } from "vue-router";
import { useInsecureAuthStore } from "@/stores/insecureAuth";
import { insecureAPI } from "@/services/api";
import Header from "@/components/v1/Header.vue";

const router = useRouter();
const authStore = useInsecureAuthStore();

const isLoading = ref(false);
const products = ref<any[]>([]);
const showCreateForm = ref(false);

const searchForm = ref({
  name: "",
  description: "",
  price: "",
});

const createForm = ref({
  name: "",
  description: "",
  price: 0,
});

onMounted(async () => {
  await searchProducts();
  await getProducts();
});

async function getProducts() {
  const response = await insecureAPI.listProducts();
  products.value = response.data.products;
  console.log(products.value);
}

async function searchProducts() {
  isLoading.value = true;
  try {
    const response = await insecureAPI.listProducts({
      name: searchForm.value.name,
      description: searchForm.value.description,
      price: searchForm.value.price,
    });
    products.value = response.data.products;
  } catch (error: any) {
    alert(error.response?.data?.error || "Failed to fetch products");
  } finally {
    isLoading.value = false;
  }
}

async function createProduct() {
  try {
    await insecureAPI.createProduct(createForm.value);
    showCreateForm.value = false;
    await searchProducts();
  } catch (error: any) {
    alert(error.response?.data?.error || "Failed to create product");
  }
}

async function deleteProduct(id: number) {
  if (!confirm("Are you sure you want to delete this product?")) return;
  console.log(id);
  try {
    await insecureAPI.deleteProduct(id);
    //await searchProducts();
    await loadProducts();
  } catch (error: any) {
    alert(error.response?.data?.error || "Failed to delete product");
  }
}

async function loadProducts() {
  const response = await insecureAPI.listProducts();
  products.value = response.data?.products || [];
  console.log(products.value);
}

async function handleLogout() {
  try {
    await authStore.logout();
    router.push("/v1/login");
  } catch (error: any) {
    alert(error.response?.data?.error || "Failed to logout");
  }
}
</script>
