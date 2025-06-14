<script setup lang="ts">
// Static page - no script needed
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b">
      <div class="flex items-center justify-center w-full">
        <div class="w-full max-w-7xl mx-auto px-4 py-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <div
                class="w-12 h-12 bg-blue-600 rounded-lg flex items-center justify-center"
              >
                <svg
                  class="w-7 h-7 text-white"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M16 11V7a4 4 0 00-8 0v4M8 11v6h8v-6M8 11h8"
                  />
                </svg>
              </div>
              <div>
                <h1 class="text-3xl font-bold text-gray-900">VHU Shop</h1>
                <p class="text-base text-gray-600">Demo Thương Mại Điện Tử</p>
              </div>
            </div>

            <div class="text-right">
              <p class="text-base text-gray-600 font-medium">Nhóm 10</p>
              <p class="text-sm text-gray-500">Triển Khai & Đánh giá RestfulAPI</p>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <div class="flex flex-col min-h-[calc(100vh-80px)] items-center justify-center">
      <main class="w-full max-w-4xl mx-auto px-4 py-16">
        <div class="text-center mb-12">
          <h2 class="text-4xl font-bold text-gray-900 mb-4">Chọn Phiên Bản Demo</h2>
          <p class="text-xl text-gray-600 max-w-2xl mx-auto">
            Chúng tôi có 2 phiên bản của cùng một trang web để bạn so sánh và học hỏi về
            bảo mật web
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <!-- Version 1 - Insecure -->
          <div
            class="bg-white rounded-2xl shadow-lg border border-gray-200 overflow-hidden group hover:shadow-xl transition-all duration-300"
          >
            <div class="p-8">
              <div class="flex items-center justify-between mb-6">
                <h3 class="text-2xl font-bold text-gray-900">Phiên Bản 1.0</h3>
                <span
                  class="px-3 py-1 bg-yellow-100 text-yellow-800 text-sm font-medium rounded-full"
                >
                  Phiên Bản Cũ
                </span>
              </div>

              <div class="mb-6">
                <h4 class="text-lg font-semibold text-gray-900 mb-3">Tính năng:</h4>
                <ul class="space-y-2 text-gray-600">
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Đăng nhập/Đăng ký đơn giản
                  </li>
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Xem và tìm kiếm sản phẩm
                  </li>
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Thêm/Sửa/Xóa sản phẩm
                  </li>
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Xem profile người dùng
                  </li>
                </ul>
              </div>

              <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4 mb-6">
                <h5 class="text-sm font-semibold text-yellow-800 mb-2">
                  💡 Lưu ý về phiên bản này:
                </h5>
                <p class="text-sm text-yellow-700">
                  Đây là phiên bản được phát triển mà không chú ý đến bảo mật. Bạn sẽ thấy
                  được cách các lỗ hổng bảo mật có thể xuất hiện trong ứng dụng thực tế.
                </p>
              </div>

              <router-link
                to="/v1/login"
                class="block w-full px-6 py-3 text-center font-semibold text-white bg-gradient-to-r from-yellow-500 to-orange-500 hover:from-yellow-600 hover:to-orange-600 rounded-lg transition-all duration-200 shadow-lg group-hover:shadow-xl"
              >
                Vào VHU Shop v1.0
              </router-link>
            </div>
          </div>

          <!-- Version 2 - Secure -->
          <div
            class="bg-white rounded-2xl shadow-lg border border-gray-200 overflow-hidden group hover:shadow-xl transition-all duration-300"
          >
            <div class="p-8">
              <div class="flex items-center justify-between mb-6">
                <h3 class="text-2xl font-bold text-gray-900">Phiên Bản 2.0</h3>
                <span
                  class="px-3 py-1 bg-green-100 text-green-800 text-sm font-medium rounded-full"
                >
                  Phiên Bản Mới
                </span>
              </div>

              <div class="mb-6">
                <h4 class="text-lg font-semibold text-gray-900 mb-3">Tính năng:</h4>
                <ul class="space-y-2 text-gray-600">
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Đăng nhập/Đăng ký bảo mật
                  </li>
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Xem và tìm kiếm sản phẩm
                  </li>
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Thêm/Sửa/Xóa sản phẩm (với phân quyền)
                  </li>
                  <li class="flex items-center">
                    <svg
                      class="w-4 h-4 text-green-500 mr-2"
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
                    Xem profile bảo mật
                  </li>
                </ul>
              </div>

              <div class="bg-green-50 border border-green-200 rounded-lg p-4 mb-6">
                <h5 class="text-sm font-semibold text-green-800 mb-2">
                  🛡️ Lưu ý về phiên bản này:
                </h5>
                <p class="text-sm text-green-700">
                  Đây là phiên bản được phát triển theo các chuẩn bảo mật hiện đại. Bạn sẽ
                  thấy được cách bảo vệ ứng dụng khỏi các cuộc tấn công phổ biến.
                </p>
              </div>

              <router-link
                to="/v2/login"
                class="block w-full px-6 py-3 text-center font-semibold text-white bg-gradient-to-r from-green-500 to-blue-500 hover:from-green-600 hover:to-blue-600 rounded-lg transition-all duration-200 shadow-lg group-hover:shadow-xl"
              >
                Vào VHU Shop v2.0
              </router-link>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped>
.bg-grid-pattern {
  background-image: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23000000' fill-opacity='0.1'%3E%3Ccircle cx='7' cy='7' r='1'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
}
</style>
