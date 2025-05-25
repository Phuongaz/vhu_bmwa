import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'

const app = createApp(App)

// Configure axios for API calls
import axios from 'axios'
axios.defaults.baseURL = process.env.VUE_APP_API_URL || 'https://localhost/api'
axios.interceptors.request.use(config => {
  const token = store.state.token
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Add PrimeVue
app.use(PrimeVue)

// Add router and store
app.use(router)
app.use(store)

app.mount('#app') 