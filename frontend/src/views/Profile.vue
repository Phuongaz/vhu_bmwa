<template>
  <div class="profile">
    <div class="profile-card">
      <h1>Profile</h1>
      <div class="profile-info">
        <div class="info-group">
          <label>Username</label>
          <p>{{ currentUser?.username }}</p>
        </div>
        <div class="info-group">
          <label>Role</label>
          <p>{{ currentUser?.role }}</p>
        </div>
        <div class="info-group">
          <label>Member Since</label>
          <p>{{ memberSince }}</p>
        </div>
      </div>

      <div class="actions">
        <button @click="showChangePasswordDialog" class="btn-primary">
          Change Password
        </button>
      </div>
    </div>

    <div v-if="showDialog" class="dialog-overlay" @click="closeDialog">
      <div class="dialog" @click.stop>
        <h2>Change Password</h2>
        <form @submit.prevent="handleChangePassword">
          <div class="form-group">
            <label for="currentPassword">Current Password</label>
            <input
              type="password"
              id="currentPassword"
              v-model="form.currentPassword"
              required
              class="form-control"
            />
          </div>
          <div class="form-group">
            <label for="newPassword">New Password</label>
            <input
              type="password"
              id="newPassword"
              v-model="form.newPassword"
              required
              class="form-control"
            />
          </div>
          <div class="form-group">
            <label for="confirmPassword">Confirm New Password</label>
            <input
              type="password"
              id="confirmPassword"
              v-model="form.confirmPassword"
              required
              class="form-control"
            />
          </div>
          <div class="error" v-if="error">{{ error }}</div>
          <div class="dialog-actions">
            <button type="button" @click="closeDialog" class="btn-secondary">
              Cancel
            </button>
            <button type="submit" :disabled="!isValid" class="btn-primary">
              Change Password
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Profile',
  data() {
    return {
      showDialog: false,
      error: null,
      form: {
        currentPassword: '',
        newPassword: '',
        confirmPassword: ''
      }
    }
  },
  computed: {
    ...mapGetters(['currentUser']),
    isValid() {
      return (
        this.form.currentPassword &&
        this.form.newPassword &&
        this.form.newPassword === this.form.confirmPassword
      )
    },
    memberSince() {
      if (!this.currentUser?.createdAt) return 'N/A'
      return new Date(this.currentUser.createdAt).toLocaleDateString()
    }
  },
  methods: {
    showChangePasswordDialog() {
      this.showDialog = true
      this.error = null
      this.form = {
        currentPassword: '',
        newPassword: '',
        confirmPassword: ''
      }
    },
    closeDialog() {
      this.showDialog = false
    },
    async handleChangePassword() {
      if (!this.isValid) {
        this.error = 'Please check your input'
        return
      }

      try {
        await this.$store.dispatch('changePassword', {
          currentPassword: this.form.currentPassword,
          newPassword: this.form.newPassword
        })
        this.closeDialog()
        // Show success message
        alert('Password changed successfully')
      } catch (error) {
        this.error = error.response?.data?.message || 'Failed to change password'
      }
    }
  },
  async created() {
    try {
      await this.$store.dispatch('fetchProfile')
    } catch (error) {
      console.error('Failed to fetch profile:', error)
    }
  }
}
</script>

<style scoped>
.profile {
  padding: 2rem;
  display: flex;
  justify-content: center;
}

.profile-card {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 600px;
}

.profile-card h1 {
  margin: 0 0 2rem;
  color: #2c3e50;
}

.profile-info {
  margin-bottom: 2rem;
}

.info-group {
  margin-bottom: 1.5rem;
}

.info-group label {
  display: block;
  color: #666;
  font-size: 0.875rem;
  margin-bottom: 0.5rem;
}

.info-group p {
  margin: 0;
  font-size: 1.1rem;
  color: #2c3e50;
}

.actions {
  border-top: 1px solid #eee;
  padding-top: 1.5rem;
}

.btn-primary {
  background-color: #42b983;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.2s;
}

.btn-primary:hover:not(:disabled) {
  background-color: #3aa876;
}

.btn-primary:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}

.dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.dialog {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  width: 100%;
  max-width: 400px;
}

.dialog h2 {
  margin: 0 0 1.5rem;
  color: #2c3e50;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #333;
}

.form-control {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

.error {
  color: #dc3545;
  margin-bottom: 1rem;
  font-size: 0.875rem;
}

.dialog-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}
</style> 