import { createStore } from 'vuex'
import axios from '../utils/axios'

export default createStore({
  state: {
    token: localStorage.getItem('token') || null,
    user: null,
    products: []
  },
  getters: {
    isLoggedIn: state => !!state.token,
    currentUser: state => state.user,
    products: state => state.products
  },
  mutations: {
    setToken(state, token) {
      state.token = token
      if (token) {
        localStorage.setItem('token', token)
      } else {
        localStorage.removeItem('token')
      }
    },
    setUser(state, user) {
      state.user = user
    },
    setProducts(state, products) {
      state.products = products
    }
  },
  actions: {
    async login({ commit }, credentials) {
      try {
        const response = await axios.post('/auth/login', credentials)
        const { token, user } = response.data
        commit('setToken', token)
        commit('setUser', user)
        return response
      } catch (error) {
        console.error('Login error:', error)
        throw error
      }
    },
    async register({ commit }, userData) {
      try {
        const response = await axios.post('/auth/register', userData)
        const { token, user } = response.data
        commit('setToken', token)
        commit('setUser', user)
        return response
      } catch (error) {
        console.error('Registration error:', error)
        throw error
      }
    },
    logout({ commit }) {
      commit('setToken', null)
      commit('setUser', null)
      commit('setProducts', [])
    },
    async fetchProfile({ commit }) {
      try {
        const response = await axios.get('/profile')
        commit('setUser', response.data)
        return response
      } catch (error) {
        console.error('Fetch profile error:', error)
        throw error
      }
    },
    async changePassword(_, { currentPassword, newPassword }) {
      try {
        const response = await axios.post('/profile/change-password', {
          currentPassword,
          newPassword
        })
        return response
      } catch (error) {
        console.error('Change password error:', error)
        throw error
      }
    },
    async fetchProducts({ commit }) {
      try {
        const response = await axios.get('/products')
        commit('setProducts', response.data)
        return response
      } catch (error) {
        console.error('Fetch products error:', error)
        throw error
      }
    },
    async createProduct({ dispatch }, productData) {
      try {
        const response = await axios.post('/products', productData)
        await dispatch('fetchProducts')
        return response
      } catch (error) {
        console.error('Create product error:', error)
        throw error
      }
    },
    async updateProduct({ dispatch }, { id, ...productData }) {
      try {
        const response = await axios.put(`/products/${id}`, productData)
        await dispatch('fetchProducts')
        return response
      } catch (error) {
        console.error('Update product error:', error)
        throw error
      }
    },
    async deleteProduct({ dispatch }, id) {
      try {
        const response = await axios.delete(`/products/${id}`)
        await dispatch('fetchProducts')
        return response
      } catch (error) {
        console.error('Delete product error:', error)
        throw error
      }
    }
  }
}) 