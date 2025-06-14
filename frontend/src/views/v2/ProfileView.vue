<template>
    <Header :username="user?.username" :role="user?.role" />
    <div class="max-w-2xl mx-auto p-6">
        <!-- Profile Information Card -->
        <div class="bg-white rounded-lg shadow mb-6">
            <div class="px-6 py-4 border-b border-gray-200">
                <h2 class="text-xl font-semibold text-gray-800">Thông tin tài khoản</h2>
            </div>
            <div v-if="user" class="p-6">
                <div class="grid grid-cols-2 gap-4">
                    <div>
                        <label class="block text-sm font-medium text-gray-700">Tên đăng nhập</label>
                        <div class="mt-1 text-gray-900">{{ user.username }}</div>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-gray-700">Vai trò</label>
                        <div class="mt-1 text-gray-900">{{ user.role }}</div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Change Password Card -->
        <div class="bg-white rounded-lg shadow">
            <div class="px-6 py-4 border-b border-gray-200">
                <h2 class="text-xl font-semibold text-gray-800">Đổi mật khẩu</h2>
            </div>
            <form @submit.prevent="handleSubmit" class="p-6 space-y-4">
                <div>
                    <label for="currentPassword" class="block text-sm font-medium text-gray-700">
                        Mật khẩu hiện tại
                    </label>
                    <div class="mt-1">
                        <input
                            id="currentPassword"
                            v-model="passwordForm.currentPassword"
                            type="password"
                            required
                            class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                            :class="{ 'border-red-300': errors.currentPassword }"
                        />
                    </div>
                </div>
                <div>
                    <label for="newPassword" class="block text-sm font-medium text-gray-700">
                        Mật khẩu mới
                    </label>
                    <div class="mt-1">
                        <input
                            id="newPassword"
                            v-model="passwordForm.newPassword"
                            type="password"
                            required
                            class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                            :class="{ 'border-red-300': errors.newPassword }"
                        />
                        <p v-if="errors.newPassword" class="mt-1 text-sm text-red-600">
                            {{ errors.newPassword }}
                        </p>
                    </div>
                </div>

                <div>
                    <label for="confirmPassword" class="block text-sm font-medium text-gray-700">
                        Nhập lại mật khẩu mới
                    </label>
                    <div class="mt-1">
                        <input
                            id="confirmPassword"
                            v-model="passwordForm.confirmPassword"
                            type="password"
                            required
                            class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                            :class="{ 'border-red-300': errors.confirmPassword }"
                        />
                        <p v-if="errors.confirmPassword" class="mt-1 text-sm text-red-600">
                            {{ errors.confirmPassword }}
                        </p>
                    </div>
                </div>

                <div>
                    <button
                        type="submit"
                        :disabled="isLoading"
                        class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 disabled:opacity-50"
                    >
                        <span v-if="isLoading">Đổi mật khẩu...</span>
                        <span v-else>Đổi mật khẩu</span>
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { insecureAPI, secureAPI } from '@/services/api'
import Header from '@/components/v2/Header.vue'
import router from '@/router'
interface User {
    id: number
    username: string
    role: string
}

interface PasswordForm {
    currentPassword: string
    newPassword: string
    confirmPassword: string
}

interface FormErrors {
    currentPassword?: string
    newPassword?: string
    confirmPassword?: string
}

const user = ref<User | null>(null)
const isLoading = ref(false)
const errors = ref<FormErrors>({})

const passwordForm = ref<PasswordForm>({
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
})

const validateForm = (): boolean => {
    errors.value = {}
    let isValid = true

    if (!passwordForm.value.currentPassword) {
        errors.value.currentPassword = 'Nhập mật khẩu hiện tại'
        isValid = false
    }

    if (!passwordForm.value.newPassword) {
        errors.value.newPassword = 'Nhập mật khẩu mới'
        isValid = false
    }

    if (!passwordForm.value.confirmPassword) {
        errors.value.confirmPassword = 'Nhập lại mật khẩu mới'
        isValid = false
    } else if (passwordForm.value.confirmPassword !== passwordForm.value.newPassword) {
        errors.value.confirmPassword = 'Mật khẩu mới không khớp'
        isValid = false
    }

    return isValid
}

const fetchUserProfile = async () => {
    try {
        const response = await secureAPI.getProfile()
        user.value = response.data.user
    } catch (error) {
        console.error('Failed to fetch profile:', error)
    }
}

const handleSubmit = async () => {
    if (!validateForm()) return

    isLoading.value = true
    try {
        await secureAPI.changePassword({
            currentPassword: passwordForm.value.currentPassword,
            newPassword: passwordForm.value.newPassword
        })
        
        passwordForm.value = {
            currentPassword: '',
            newPassword: '',
            confirmPassword: ''
        }
        errors.value = {}
        
        alert('Đổi mật khẩu thành công')
        secureAPI.logout()
        router.push('/v2/login')
    } catch (error: any) {
        const errorMessage = error.response?.data?.error || 'Failed to change password'
        errors.value.newPassword = errorMessage
    } finally {
        isLoading.value = false
    }
}

onMounted(() => {
    fetchUserProfile()
})
</script>