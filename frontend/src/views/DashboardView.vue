<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <nav class="bg-white shadow">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <h1 class="text-xl font-semibold text-gray-900">
              VHU Security Test Dashboard
            </h1>
          </div>

          <div class="flex items-center space-x-4">
            <!-- User Info -->
            <div class="flex items-center space-x-2">
              <span class="text-sm text-gray-700">{{ user?.username }}</span>
              <span
                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                :class="roleClass"
              >
                {{ user?.role }}
              </span>
            </div>

            <!-- Navigation Links -->
            <router-link
              to="/profile"
              class="text-gray-500 hover:text-gray-700 px-3 py-2 rounded-md text-sm font-medium"
            >
              Profile
            </router-link>

            <router-link
              to="/products"
              class="text-gray-500 hover:text-gray-700 px-3 py-2 rounded-md text-sm font-medium"
            >
              Products
            </router-link>

            <!-- Logout Button -->
            <button
              @click="handleLogout"
              :disabled="isLoading"
              class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-md text-sm font-medium disabled:opacity-50"
            >
              {{ isLoading ? "Logging out..." : "Logout" }}
            </button>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div class="px-4 py-6 sm:px-0">
        <!-- Welcome Section -->
        <div class="bg-white overflow-hidden shadow rounded-lg mb-6">
          <div class="px-4 py-5 sm:p-6">
            <h2 class="text-2xl font-bold text-gray-900 mb-4">
              Chào mừng, {{ user?.username }}!
            </h2>
            <p class="text-gray-600 mb-4">
              Đây là hệ thống kiểm thử bảo mật API được xây dựng bằng Vue 3 và Golang.
            </p>

            <!-- User Details -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div class="bg-gray-50 p-4 rounded-lg">
                <h3 class="text-sm font-medium text-gray-500">User ID</h3>
                <p class="text-lg font-semibold text-gray-900">{{ user?.id }}</p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg">
                <h3 class="text-sm font-medium text-gray-500">Username</h3>
                <p class="text-lg font-semibold text-gray-900">{{ user?.username }}</p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg">
                <h3 class="text-sm font-medium text-gray-500">Role</h3>
                <p class="text-lg font-semibold" :class="roleTextClass">
                  {{ user?.role }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Features Grid -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <!-- Security Features -->
          <div class="bg-white overflow-hidden shadow rounded-lg">
            <div class="p-6">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <svg
                    class="h-8 w-8 text-green-500"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                    ></path>
                  </svg>
                </div>
                <div class="ml-5 w-0 flex-1">
                  <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                      JWT Authentication
                    </dt>
                    <dd class="text-lg font-medium text-gray-900">Hoạt động</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <!-- Role-based Access -->
          <div class="bg-white overflow-hidden shadow rounded-lg">
            <div class="p-6">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <svg
                    class="h-8 w-8 text-blue-500"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                    ></path>
                  </svg>
                </div>
                <div class="ml-5 w-0 flex-1">
                  <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                      Role-based Access
                    </dt>
                    <dd class="text-lg font-medium text-gray-900">{{ user?.role }}</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <!-- Rate Limiting -->
          <div class="bg-white overflow-hidden shadow rounded-lg">
            <div class="p-6">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <svg
                    class="h-8 w-8 text-yellow-500"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M13 10V3L4 14h7v7l9-11h-7z"
                    ></path>
                  </svg>
                </div>
                <div class="ml-5 w-0 flex-1">
                  <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                      Rate Limiting
                    </dt>
                    <dd class="text-lg font-medium text-gray-900">10 req/sec</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Quick Actions -->
        <div class="mt-8 bg-white shadow rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
              Quick Actions
            </h3>
            <div class="flex flex-wrap gap-4">
              <router-link
                to="/profile"
                class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
              >
                Xem Profile
              </router-link>

              <router-link
                to="/products"
                class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
              >
                Quản lý Products
              </router-link>

              <router-link
                to="/security-test"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700"
              >
                Kiểm Thử Bảo Mật
              </router-link>

              <button
                v-if="isAdmin"
                @click="testRateLimit"
                class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
              >
                Test Rate Limit
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { productAPI } from "@/services/api";

const router = useRouter();
const authStore = useAuthStore();
const isLoading = ref(false);

const user = computed(() => authStore.user);
const isAdmin = computed(() => authStore.isAdmin);

const roleClass = computed(() => {
  switch (user.value?.role) {
    case "admin":
      return "bg-red-100 text-red-800";
    case "moderator":
      return "bg-yellow-100 text-yellow-800";
    default:
      return "bg-green-100 text-green-800";
  }
});

const roleTextClass = computed(() => {
  switch (user.value?.role) {
    case "admin":
      return "text-red-600";
    case "moderator":
      return "text-yellow-600";
    default:
      return "text-green-600";
  }
});

async function handleLogout() {
  isLoading.value = true;
  try {
    await authStore.logout();
    router.push("/login");
  } catch (error) {
    console.error("Logout failed:", error);
  } finally {
    isLoading.value = false;
  }
}

async function testRateLimit() {
  // Test rate limiting by making multiple rapid requests
  const promises = Array.from({ length: 20 }, () => productAPI.list(1, 1));

  try {
    const results = await Promise.allSettled(promises);
    const successful = results.filter((r) => r.status === "fulfilled").length;
    const failed = results.filter((r) => r.status === "rejected").length;

    alert(`Rate limit test completed.\nSuccessful: ${successful}\nFailed: ${failed}`);
  } catch (error) {
    console.error("Rate limit test failed:", error);
  }
}
</script>
