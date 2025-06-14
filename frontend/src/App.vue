<script setup lang="ts">
import { RouterView } from "vue-router";
import { onMounted } from "vue";
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();

// Khôi phục trạng thái đăng nhập từ cookie khi app khởi động
onMounted(async () => {
  await authStore.loadUserFromCookie();
});
</script>

<template>
  <div class="min-h-screen bg-gray-100">
    <main class="container mx-auto px-4 py-8">
      <RouterView />
    </main>
  </div>
</template>

<style>
@import '@/assets/base.css';
@import '@/assets/main.css';

/* Reset và global styles */
* {
  box-sizing: border-box;
}

body {
  margin: 0;
  font-family: "Inter", -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
    "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

#app {
  min-height: 100vh;
}
</style>
