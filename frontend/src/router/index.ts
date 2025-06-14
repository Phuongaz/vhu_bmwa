import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import { secureAPI } from '@/services/api'

// V1 Auth Guard - Check localStorage
const v1AuthGuard = () => {
  const token = localStorage.getItem('user')
  if (!token) {
    return { name: 'v1-login' }
  }
  return true
}

// V2 Auth Guard - Check cookie via API
const v2AuthGuard = async () => {
  try {
    const response = await secureAPI.getProfile()
    if (!response.data.user) {
      return { name: 'v2-login' }
    }
    return true
  } catch {
    return { name: 'v2-login' }
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    
    // Version 1 Routes (Insecure)
    {
      path: '/v1/login',
      name: 'v1-login',
      component: () => import('../views/v1/LoginView.vue')
    },
    {
      path: '/v1/products',
      name: 'v1-products',
      component: () => import('../views/v1/ProductsView.vue'),
      beforeEnter: v1AuthGuard
    },
    {
      path: '/v1/profile',
      name: 'v1-profile',
      component: () => import('../views/v1/ProfileView.vue'),
      beforeEnter: v1AuthGuard
    },
    
    // Version 2 Routes (Secure) 
    {
      path: '/v2/login',
      name: 'v2-login',
      component: () => import('../views/v2/LoginView.vue')
    },
    {
      path: '/v2/products',
      name: 'v2-products',
      component: () => import('../views/v2/ProductsView.vue'),
      beforeEnter: v2AuthGuard
    },
    {
      path: '/v2/profile',
      name: 'v2-profile',
      component: () => import('../views/v2/ProfileView.vue'),
      beforeEnter: v2AuthGuard
    },

    // Legacy routes (redirect to new structure)
    {
      path: '/v1',
      redirect: '/v1/login'
    },
    {
      path: '/v2',
      redirect: '/v2/login'
    },
    
    // Catch all route
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      redirect: '/'
    }
  ]
})

export default router



