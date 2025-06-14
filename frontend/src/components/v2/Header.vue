<template>
    <header class="bg-white shadow-sm border-b w-full mb-8">
        <div class="container mx-auto px-4 sm:px-6 lg:px-8 max-w-5xl py-6">
            <div class="flex items-center justify-between">
                <!-- Left: Logo -->
                <div class="flex items-center gap-3">
                    <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
                        <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M16 11V7a4 4 0 00-8 0v4M8 11v6h8v-6M8 11h8" />
                        </svg>
                    </div>
                    <div>
                        <h1 class="text-xl font-bold text-gray-900">VHU Shop</h1>
                        <p class="text-xs text-green-600">Phiên bản 2.0</p>
                    </div>
                </div>

                <!-- Center: Navigation -->
                <nav class="flex items-center gap-6 font-sans mx-8">
                    <router-link 
                        to="/v2/products" 
                        class="text-gray-600 hover:text-green-600 text-sm font-medium whitespace-nowrap"
                        active-class="text-green-600"
                    >
                        Sản phẩm
                    </router-link>
                    <router-link 
                        to="/v2/profile" 
                        class="text-gray-600 hover:text-green-600 text-sm font-medium whitespace-nowrap"
                        active-class="text-green-600"
                    >
                        Tài khoản
                    </router-link>
                </nav>

                <!-- Right: User Menu -->
                <div class="flex items-center gap-4 min-w-[200px] justify-end">
                    <div class="text-right flex items-center gap-2">
                        <p class="text-sm font-medium text-gray-900 truncate max-w-[120px]">{{ authStore.user?.username }}</p>
                        <span class="px-2 py-1 text-xs font-medium bg-green-100 text-green-800 rounded-full">
                            {{ authStore.user?.role }}
                        </span>
                    </div>
                    <button 
                        @click="logout" 
                        class="text-green-600 hover:text-green-700 text-sm font-medium flex items-center gap-1 whitespace-nowrap"
                    >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                                d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                        </svg>
                        Đăng xuất
                    </button>
                </div>
            </div>
        </div>
    </header>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

const router = useRouter()

const authStore = useAuthStore();

const logout = () => {
    localStorage.removeItem('user')
    router.push('/v2/login')
}
</script>
