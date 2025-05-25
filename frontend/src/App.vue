<template>
  <div id="app">
    <nav class="navbar" v-if="isLoggedIn">
      <router-link to="/dashboard">Dashboard</router-link> |
      <router-link to="/profile">Profile</router-link> |
      <router-link to="/products">Products</router-link> |
      <a href="#" @click.prevent="logout">Logout</a>
    </nav>
    <router-view></router-view>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'App',
  computed: {
    ...mapGetters(['isLoggedIn'])
  },
  methods: {
    logout() {
      this.$store.dispatch('logout')
      this.$router.push('/login')
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.navbar {
  padding: 20px;
  background-color: #f8f9fa;
  margin-bottom: 20px;
}

.navbar a {
  font-weight: bold;
  color: #2c3e50;
  text-decoration: none;
  margin: 0 10px;
}

.navbar a.router-link-exact-active {
  color: #42b983;
}
</style> 