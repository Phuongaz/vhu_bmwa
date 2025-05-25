import { createRouter, createWebHistory } from 'vue-router'
import store from '../store'

const Login = () => import('../views/Login.vue')
const Register = () => import('../views/Register.vue')
const Dashboard = () => import('../views/Dashboard.vue')
const Profile = () => import('../views/Profile.vue')
const Products = () => import('../views/Products.vue')

const routes = [
  {
    path: '/',
    redirect: '/dashboard'
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { guest: true }
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: { guest: true }
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
    meta: { requiresAuth: true }
  },
  {
    path: '/products',
    name: 'Products',
    component: Products,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.getters.isLoggedIn) {
      next('/login')
      return
    }
    next()
  } else if (to.matched.some(record => record.meta.guest)) {
    if (store.getters.isLoggedIn) {
      next('/dashboard')
      return
    }
    next()
  } else {
    next()
  }
})

export default router 