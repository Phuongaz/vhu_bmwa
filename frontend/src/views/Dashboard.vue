<template>
  <div class="dashboard-container">
    <header class="dashboard-header">
      <h1>Bảng Điều Khiển</h1>
      <div class="user-info">
        <span>Xin chào, {{ username }}</span>
        <button @click="logout" class="btn-logout">Đăng xuất</button>
      </div>
    </header>

    <div class="dashboard-content">
      <div class="stats-grid">
        <div class="stat-card">
          <h3>Tổng Sản Phẩm</h3>
          <div class="stat-value">{{ totalProducts }}</div>
        </div>
        <div class="stat-card">
          <h3>Sản Phẩm Hoạt Động</h3>
          <div class="stat-value">{{ activeProducts }}</div>
        </div>
        <div class="stat-card">
          <h3>Đơn Hàng Hôm Nay</h3>
          <div class="stat-value">{{ todayOrders }}</div>
        </div>
        <div class="stat-card">
          <h3>Doanh Thu Tháng</h3>
          <div class="stat-value">{{ monthlyRevenue }}đ</div>
        </div>
      </div>

      <div class="recent-activity">
        <h2>Hoạt Động Gần Đây</h2>
        <div v-if="loading" class="loading">
          <div class="spinner"></div>
          <span>Đang tải...</span>
        </div>
        <div v-else-if="activities.length === 0" class="no-data">
          Chưa có hoạt động nào
        </div>
        <ul v-else class="activity-list">
          <li v-for="activity in activities" :key="activity.id" class="activity-item">
            <div class="activity-icon" :class="activity.type">
              <i :class="getActivityIcon(activity.type)"></i>
            </div>
            <div class="activity-details">
              <div class="activity-message">{{ activity.message }}</div>
              <div class="activity-time">{{ formatTime(activity.timestamp) }}</div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      username: '',
      totalProducts: 0,
      activeProducts: 0,
      todayOrders: 0,
      monthlyRevenue: 0,
      activities: [],
      loading: true
    }
  },
  async created() {
    try {
      await this.loadDashboardData()
    } catch (error) {
      console.error('Failed to load dashboard data:', error)
    }
  },
  methods: {
    async loadDashboardData() {
      try {
        const response = await this.$store.dispatch('dashboard/getData')
        this.username = response.username
        this.totalProducts = response.totalProducts
        this.activeProducts = response.activeProducts
        this.todayOrders = response.todayOrders
        this.monthlyRevenue = response.monthlyRevenue
        this.activities = response.activities
      } catch (error) {
        console.error('Error loading dashboard data:', error)
      } finally {
        this.loading = false
      }
    },
    getActivityIcon(type) {
      const icons = {
        order: 'fas fa-shopping-cart',
        product: 'fas fa-box',
        user: 'fas fa-user',
        system: 'fas fa-cog'
      }
      return icons[type] || 'fas fa-info-circle'
    },
    formatTime(timestamp) {
      const date = new Date(timestamp)
      return date.toLocaleString('vi-VN')
    },
    async logout() {
      try {
        await this.$store.dispatch('auth/logout')
        this.$router.push('/login')
      } catch (error) {
        console.error('Logout failed:', error)
      }
    }
  }
}
</script>

<style scoped>
.dashboard-container {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.btn-logout {
  padding: 0.5rem 1rem;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.stat-card {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.stat-card h3 {
  margin: 0;
  color: #666;
  font-size: 1rem;
}

.stat-value {
  font-size: 2rem;
  font-weight: bold;
  color: #333;
  margin-top: 0.5rem;
}

.recent-activity {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.activity-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.activity-item {
  display: flex;
  align-items: center;
  padding: 1rem 0;
  border-bottom: 1px solid #eee;
}

.activity-item:last-child {
  border-bottom: none;
}

.activity-icon {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 1rem;
}

.activity-icon.order {
  background: #e3f2fd;
  color: #1976d2;
}

.activity-icon.product {
  background: #f3e5f5;
  color: #7b1fa2;
}

.activity-icon.user {
  background: #e8f5e9;
  color: #388e3c;
}

.activity-icon.system {
  background: #fff3e0;
  color: #f57c00;
}

.activity-details {
  flex: 1;
}

.activity-message {
  color: #333;
  margin-bottom: 0.25rem;
}

.activity-time {
  color: #666;
  font-size: 0.875rem;
}

.loading {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
}

.spinner {
  width: 24px;
  height: 24px;
  border: 3px solid #f3f3f3;
  border-top: 3px solid #3498db;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-right: 1rem;
}

.no-data {
  text-align: center;
  padding: 2rem;
  color: #666;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style> 