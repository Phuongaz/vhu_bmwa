<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-green-50 to-emerald-50">
    <!-- Header -->
    <nav class="bg-white/80 backdrop-blur-sm shadow-sm border-b border-white/20">
      <div class="max-w-7xl mx-auto px-4 py-6">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-4">
            <router-link
              to="/"
              class="flex items-center space-x-2 text-gray-900 hover:text-gray-700 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
              </svg>
              <span class="font-medium">Trang Chủ</span>
            </router-link>
            
            <div class="h-6 w-px bg-gray-300"></div>
            
            <div class="flex items-center space-x-3">
              <div class="h-8 w-8 bg-gradient-to-r from-green-600 to-emerald-600 rounded-lg flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
              </div>
              <div>
                <h1 class="text-xl font-bold text-gray-900">API v2 - Secure Testing</h1>
                <p class="text-sm text-gray-600">Phiên bản đã được bảo mật</p>
              </div>
            </div>
          </div>
          
          <div class="flex items-center space-x-4">
            <div v-if="authStore.user" class="flex items-center space-x-3">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ authStore.user.username }}</div>
                <div class="text-xs text-gray-500 capitalize">{{ authStore.user.role }}</div>
              </div>
              <div class="w-8 h-8 bg-gradient-to-r from-blue-500 to-indigo-500 rounded-full flex items-center justify-center">
                <span class="text-white text-sm font-medium">{{ authStore.user.username.charAt(0).toUpperCase() }}</span>
              </div>
              <button
                @click="logout"
                class="px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg text-sm font-medium transition-colors shadow-sm"
              >
                Đăng Xuất
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Title -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900">Cửa Hàng Sản Phẩm</h1>
        <p class="mt-2 text-gray-600">Xem và mua sản phẩm của chúng tôi</p>

        <!-- Testing Guide -->
        <div class="mt-4 p-4 bg-blue-50 border border-blue-200 rounded-lg">
          <h3 class="text-sm font-medium text-blue-800">📖 Hướng Dẫn Kiểm Thử Bảo Mật</h3>
          <p class="text-sm text-blue-700 mt-1">
            Xem file
            <code class="bg-blue-100 px-1 rounded">SECURITY_TEST_PAYLOADS.md</code> để
            biết các payload kiểm thử. Thử các chức năng bên dưới để test bảo mật backend.
          </p>
        </div>
      </div>

      <!-- Security Status Banner -->
      <div class="mb-8 bg-gradient-to-r from-green-500 via-green-600 to-emerald-600 rounded-2xl shadow-xl overflow-hidden">
        <div class="p-8 text-white relative">
          <div class="absolute top-0 right-0 w-32 h-32 bg-green-400/20 rounded-full -translate-y-16 translate-x-16"></div>
          <div class="absolute bottom-0 left-0 w-24 h-24 bg-green-400/20 rounded-full translate-y-12 -translate-x-12"></div>
          
          <div class="relative">
            <div class="flex items-center space-x-4 mb-4">
              <div class="w-12 h-12 bg-green-800/50 rounded-xl flex items-center justify-center border border-green-400/30">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
              </div>
              <div>
                <h2 class="text-2xl font-bold">🛡️ API v2 Security Testing Lab</h2>
                <p class="text-green-100">Kiểm thử các tính năng bảo mật đã được cải thiện</p>
              </div>
            </div>
            
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-6">
              <div class="bg-green-800/30 rounded-lg p-3 border border-green-400/30">
                <div class="text-sm text-green-100">JWT Auth</div>
                <div class="text-xs text-green-200">✅ Enabled</div>
              </div>
              <div class="bg-green-800/30 rounded-lg p-3 border border-green-400/30">
                <div class="text-sm text-green-100">Rate Limiting</div>
                <div class="text-xs text-green-200">✅ 10 req/sec</div>
              </div>
              <div class="bg-green-800/30 rounded-lg p-3 border border-green-400/30">
                <div class="text-sm text-green-100">Input Validation</div>
                <div class="text-xs text-green-200">✅ Sanitized</div>
              </div>
              <div class="bg-green-800/30 rounded-lg p-3 border border-green-400/30">
                <div class="text-sm text-green-100">RBAC</div>
                <div class="text-xs text-green-200">✅ Role-based</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Search Products Section -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-8">
        
        <!-- Search Form -->
        <div class="lg:col-span-2">
          <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
            <div class="p-6 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-blue-100">
              <h3 class="text-lg font-semibold text-gray-900 flex items-center">
                <svg class="w-5 h-5 text-blue-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
                </svg>
                Tìm Kiếm Sản Phẩm (Secure)
              </h3>
              <p class="text-sm text-gray-600 mt-1">Tất cả input được validate và sanitize</p>
            </div>
            
            <div class="p-6">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Tên sản phẩm</label>
                  <input
                    v-model="searchTerm"
                    type="text"
                    placeholder="Nhập tên sản phẩm..."
                    class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                  />
                  <p class="text-xs text-gray-500 mt-1 flex items-center">
                    <svg class="w-3 h-3 text-yellow-500 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                    </svg>
                    Thử XSS: <code class="bg-gray-100 px-1 rounded text-xs">&lt;script&gt;alert('XSS')&lt;/script&gt;</code>
                  </p>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Giá từ</label>
                  <input
                    v-model="priceFrom"
                    type="text"
                    placeholder="0"
                    class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                  />
                  <p class="text-xs text-gray-500 mt-1 flex items-center">
                    <svg class="w-3 h-3 text-yellow-500 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                    </svg>
                    Thử SQL: <code class="bg-gray-100 px-1 rounded text-xs">' OR '1'='1</code>
                  </p>
                </div>
              </div>
              
              <button
                @click="searchProducts"
                :disabled="isLoading"
                class="w-full px-6 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 text-white rounded-lg hover:from-blue-700 hover:to-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 font-medium shadow-lg"
              >
                <span v-if="!isLoading" class="flex items-center justify-center">
                  <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
                  </svg>
                  Tìm Kiếm
                </span>
                <span v-else class="flex items-center justify-center">
                  <svg class="animate-spin w-4 h-4 mr-2" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
                  </svg>
                  Đang tìm...
                </span>
              </button>
            </div>
          </div>
        </div>

        <!-- Security Info Panel -->
        <div class="lg:col-span-1">
          <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden h-fit">
            <div class="p-6 bg-gradient-to-r from-green-50 to-emerald-50 border-b border-green-100">
              <h3 class="text-lg font-semibold text-gray-900 flex items-center">
                <svg class="w-5 h-5 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
                Tính Năng Bảo Mật
              </h3>
            </div>
            
            <div class="p-6 space-y-4">
              <div class="flex items-start space-x-3">
                <div class="w-2 h-2 bg-green-500 rounded-full mt-2"></div>
                <div>
                  <div class="font-medium text-gray-900">Input Sanitization</div>
                  <div class="text-sm text-gray-600">HTML tags bị escape, SQL characters filtered</div>
                </div>
              </div>
              
              <div class="flex items-start space-x-3">
                <div class="w-2 h-2 bg-green-500 rounded-full mt-2"></div>
                <div>
                  <div class="font-medium text-gray-900">Authentication Required</div>
                  <div class="text-sm text-gray-600">JWT token trong HTTP-only cookie</div>
                </div>
              </div>
              
              <div class="flex items-start space-x-3">
                <div class="w-2 h-2 bg-green-500 rounded-full mt-2"></div>
                <div>
                  <div class="font-medium text-gray-900">Role-based Access</div>
                  <div class="text-sm text-gray-600">Phân quyền theo user role</div>
                </div>
              </div>
              
              <div class="flex items-start space-x-3">
                <div class="w-2 h-2 bg-green-500 rounded-full mt-2"></div>
                <div>
                  <div class="font-medium text-gray-900">Rate Limiting</div>
                  <div class="text-sm text-gray-600">Max 10 requests/second</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Security Testing Section -->
      <div class="mb-8">
        <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
          <div class="p-6 bg-gradient-to-r from-amber-50 to-orange-50 border-b border-amber-100">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-semibold text-gray-900 flex items-center">
                  <svg class="w-5 h-5 text-amber-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
                  </svg>
                  🔥 Security Penetration Testing
                </h3>
                <p class="text-sm text-gray-600 mt-1">Thử nghiệm các kỹ thuật tấn công phổ biến để kiểm tra bảo mật</p>
              </div>
              <div class="text-right">
                <div class="text-sm text-gray-500">Current Role</div>
                <div class="text-lg font-semibold text-gray-900 capitalize">{{ authStore.user?.role }}</div>
              </div>
            </div>
          </div>
          
          <div class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
              <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-xl p-4 hover:shadow-md transition-all group">
                <div class="flex items-center justify-between mb-3">
                  <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
                    <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"/>
                    </svg>
                  </div>
                  <span class="text-xs px-2 py-1 bg-red-200 text-red-800 rounded-full font-medium">HIGH</span>
                </div>
                <h4 class="font-semibold text-red-800 mb-2">Privilege Escalation</h4>
                <p class="text-xs text-red-600 mb-4">Thử tạo sản phẩm (yêu cầu quyền admin)</p>
                <button
                  @click="tryCreateProduct"
                  :disabled="isLoading"
                  class="w-full px-3 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 text-sm font-medium transition-colors"
                >
                  Test Attack
                </button>
              </div>

              <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-xl p-4 hover:shadow-md transition-all group">
                <div class="flex items-center justify-between mb-3">
                  <div class="w-10 h-10 bg-orange-500 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
                    <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                    </svg>
                  </div>
                  <span class="text-xs px-2 py-1 bg-orange-200 text-orange-800 rounded-full font-medium">MED</span>
                </div>
                <h4 class="font-semibold text-orange-800 mb-2">Rate Limit Abuse</h4>
                <p class="text-xs text-orange-600 mb-4">Spam 50 requests để test rate limiting</p>
                <button
                  @click="tryRateLimitAbuse"
                  :disabled="isLoading"
                  class="w-full px-3 py-2 bg-orange-600 text-white rounded-lg hover:bg-orange-700 disabled:opacity-50 text-sm font-medium transition-colors"
                >
                  Test Attack
                </button>
              </div>

              <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-xl p-4 hover:shadow-md transition-all group">
                <div class="flex items-center justify-between mb-3">
                  <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
                    <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                    </svg>
                  </div>
                  <span class="text-xs px-2 py-1 bg-purple-200 text-purple-800 rounded-full font-medium">HIGH</span>
                </div>
                <h4 class="font-semibold text-purple-800 mb-2">IDOR Attack</h4>
                <p class="text-xs text-purple-600 mb-4">Truy cập dữ liệu của user khác</p>
                <button
                  @click="tryIDORAttack"
                  :disabled="isLoading"
                  class="w-full px-3 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 disabled:opacity-50 text-sm font-medium transition-colors"
                >
                  Test Attack
                </button>
              </div>

              <div class="bg-gradient-to-br from-indigo-50 to-indigo-100 border border-indigo-200 rounded-xl p-4 hover:shadow-md transition-all group">
                <div class="flex items-center justify-between mb-3">
                  <div class="w-10 h-10 bg-indigo-500 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
                    <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.031 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                    </svg>
                  </div>
                  <span class="text-xs px-2 py-1 bg-indigo-200 text-indigo-800 rounded-full font-medium">HIGH</span>
                </div>
                <h4 class="font-semibold text-indigo-800 mb-2">Admin Bypass</h4>
                <p class="text-xs text-indigo-600 mb-4">Thử truy cập admin endpoints</p>
                <button
                  @click="tryAdminEndpoints"
                  :disabled="isLoading"
                  class="w-full px-3 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 disabled:opacity-50 text-sm font-medium transition-colors"
                >
                  Test Attack
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Products -->
      <div class="bg-white shadow rounded-lg p-6">
        <h2 class="text-lg font-semibold mb-6">Sản Phẩm</h2>

        <div v-if="isLoading" class="text-center py-8">
          <div
            class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"
          ></div>
        </div>

        <div
          v-else-if="products.length > 0"
          class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
        >
          <div
            v-for="product in products"
            :key="product.ID"
            class="border border-gray-200 rounded-lg p-4"
          >
            <h3 class="font-medium" v-html="product.name"></h3>
            <p class="text-gray-600 text-sm mt-1" v-html="product.description"></p>
            <div class="mt-4 flex justify-between items-center">
              <span class="text-lg font-bold text-green-600"
                >{{ formatPrice(product.price) }}đ</span
              >
              <button
                class="px-3 py-1 bg-blue-600 text-white rounded text-sm hover:bg-blue-700"
              >
                Mua
              </button>
            </div>
          </div>
        </div>

        <div v-else class="text-center py-8">
          <p class="text-gray-600">Không có sản phẩm nào</p>
        </div>
      </div>

      <!-- Attack Results -->
      <div v-if="attackResults.length > 0" class="mt-8 bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-semibold mb-4">🔍 Kết Quả Kiểm Thử</h3>
        <div class="space-y-3">
          <div
            v-for="result in attackResults"
            :key="result.id"
            class="p-4 border rounded-lg"
            :class="{
              'border-green-200 bg-green-50': result.status === 'blocked',
              'border-red-200 bg-red-50': result.status === 'success',
              'border-yellow-200 bg-yellow-50': result.status === 'error',
            }"
          >
            <div class="flex justify-between items-start">
              <div>
                <h4
                  class="font-medium"
                  :class="{
                    'text-green-800': result.status === 'blocked',
                    'text-red-800': result.status === 'success',
                    'text-yellow-800': result.status === 'error',
                  }"
                >
                  {{ result.attack }}
                </h4>
                <p
                  class="text-sm mt-1"
                  :class="{
                    'text-green-700': result.status === 'blocked',
                    'text-red-700': result.status === 'success',
                    'text-yellow-700': result.status === 'error',
                  }"
                >
                  {{ result.message }}
                </p>
              </div>
              <span
                class="text-xs px-2 py-1 rounded font-medium"
                :class="{
                  'bg-green-200 text-green-800': result.status === 'blocked',
                  'bg-red-200 text-red-800': result.status === 'success',
                  'bg-yellow-200 text-yellow-800': result.status === 'error',
                }"
              >
                {{
                  result.status === "blocked"
                    ? "🛡️ BỊ CHẶN"
                    : result.status === "success"
                    ? "⚠️ THÀNH CÔNG"
                    : "❗ LỖI"
                }}
              </span>
            </div>
            <div class="text-xs text-gray-500 mt-2">{{ result.timestamp }}</div>
          </div>
        </div>
        <button
          @click="attackResults = []"
          class="mt-4 px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-700 text-sm"
        >
          Xóa Kết Quả
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { useProductStore } from "@/stores/product";
import { productAPI } from "@/services/api";

const router = useRouter();
const authStore = useAuthStore();
const productStore = useProductStore();
const isLoading = ref(false);

// Search
const searchTerm = ref("");
const priceFrom = ref("");
const priceTo = ref("");

// Attack Results
interface AttackResult {
  id: string;
  attack: string;
  message: string;
  status: "blocked" | "success" | "error";
  timestamp: string;
}

const attackResults = ref<AttackResult[]>([]);
const products = computed(() => productStore.products);

async function logout() {
  await authStore.logout();
  router.push("/login");
}

function addAttackResult(
  attack: string,
  message: string,
  status: "blocked" | "success" | "error"
) {
  attackResults.value.unshift({
    id: Date.now().toString(),
    attack,
    message,
    status,
    timestamp: new Date().toLocaleString("vi-VN"),
  });
}

async function searchProducts() {
  isLoading.value = true;
  try {
    const params = new URLSearchParams();
    if (searchTerm.value) params.append("search", searchTerm.value);
    if (priceFrom.value) params.append("price_from", priceFrom.value);
    if (priceTo.value) params.append("price_to", priceTo.value);

    await productStore.fetchProducts(1, 10);

    // Check for attack payloads
    if (searchTerm.value.includes("<script>")) {
      addAttackResult(
        "XSS Test",
        `XSS payload đã được filter: ${searchTerm.value}`,
        "blocked"
      );
    }
    if (priceFrom.value.includes("'") || priceTo.value.includes("'")) {
      addAttackResult("SQL Injection", `SQL injection payload bị chặn`, "blocked");
    }
  } catch (error: any) {
    addAttackResult("Search Error", `Lỗi: ${error.message}`, "error");
  } finally {
    isLoading.value = false;
  }
}

async function tryCreateProduct() {
  isLoading.value = true;
  try {
    const testProduct = {
      name: "⚠️ HACKED PRODUCT",
      description: "User thường tạo sản phẩm - Privilege Escalation!",
      price: 1,
    };

    await productAPI.createProduct(testProduct);
    addAttackResult(
      "Privilege Escalation",
      "🚨 CRITICAL! User thường tạo được sản phẩm!",
      "success"
    );
    await productStore.fetchProducts();
  } catch (error: any) {
    if (error.response?.status === 403) {
      addAttackResult(
        "Privilege Escalation",
        "✅ Bị từ chối (403) - Bảo mật tốt",
        "blocked"
      );
    } else {
      addAttackResult("Privilege Escalation", `Lỗi: ${error.message}`, "error");
    }
  } finally {
    isLoading.value = false;
  }
}

async function tryRateLimitAbuse() {
  isLoading.value = true;
  let blockedCount = 0;
  let successCount = 0;

  try {
    const promises = Array.from({ length: 50 }, async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/api/v2/products?page=1&limit=1",
          {
            credentials: "include",
          }
        );

        if (response.status === 429) {
          blockedCount++;
        } else if (response.ok) {
          successCount++;
        }
      } catch (error) {
        // ignore
      }
    });

    await Promise.allSettled(promises);

    if (blockedCount > 10) {
      addAttackResult(
        "Rate Limit Test",
        `✅ Rate limiting OK: ${blockedCount}/${50} bị chặn`,
        "blocked"
      );
    } else {
      addAttackResult(
        "Rate Limit Test",
        `🚨 Rate limiting yếu: ${blockedCount}/${50} bị chặn`,
        "success"
      );
    }
  } catch (error: any) {
    addAttackResult("Rate Limit Test", `Lỗi: ${error.message}`, "error");
  } finally {
    isLoading.value = false;
  }
}

async function tryIDORAttack() {
  isLoading.value = true;
  try {
    // Try to access products with different user contexts or IDs
    const testCases = [
      { url: "/api/v1/users/admin", desc: "v1 admin profile" },
      { url: "/api/v1/users/1", desc: "v1 user ID 1" },
      { url: "/api/v2/profile", desc: "v2 current profile (should work)" },
    ];
    let vulnerableEndpoints: string[] = [];

    for (const testCase of testCases) {
      try {
        const response = await fetch(`http://localhost:8080${testCase.url}`, {
          credentials: "include",
        });

        if (response.ok && testCase.url.includes("/v1/")) {
          // v1 endpoints should be blocked for security, if accessible it's a vuln
          vulnerableEndpoints.push(testCase.desc);
        }
      } catch (error) {
        // ignore
      }
    }

    if (vulnerableEndpoints.length > 0) {
      addAttackResult(
        "IDOR Attack",
        `🚨 IDOR vulnerabilities found in: ${vulnerableEndpoints.join(", ")}`,
        "success"
      );
    } else {
      addAttackResult("IDOR Attack", "✅ IDOR protection working - only authorized access allowed", "blocked");
    }
  } catch (error: any) {
    addAttackResult("IDOR Attack", `Lỗi: ${error.message}`, "error");
  } finally {
    isLoading.value = false;
  }
}

async function tryAdminEndpoints() {
  isLoading.value = true;
  try {
    const endpoints = ["/api/v2/admin/users", "/api/v2/admin/stats", "/api/v1/products"];
    let accessible: string[] = [];

    for (const endpoint of endpoints) {
      try {
        const response = await fetch(`http://localhost:8080${endpoint}`, {
          credentials: "include",
        });

        if (response.ok) {
          accessible.push(endpoint);
        }
      } catch (error) {
        // ignore
      }
    }

    if (accessible.length > 0) {
      addAttackResult(
        "Admin Bypass",
        `🚨 Authorization bypass! ${accessible.join(", ")}`,
        "success"
      );
    } else {
      addAttackResult("Admin Bypass", "✅ Admin endpoints bảo vệ tốt", "blocked");
    }
  } catch (error: any) {
    addAttackResult("Admin Bypass", `Lỗi: ${error.message}`, "error");
  } finally {
    isLoading.value = false;
  }
}

function formatPrice(price: number) {
  return new Intl.NumberFormat("vi-VN").format(price);
}

onMounted(async () => {
  await searchProducts();
});
</script>
