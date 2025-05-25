<template>
  <div class="login-container">
    <div class="login-box">
      <h2>Đăng Nhập</h2>
      <div class="alert alert-danger" v-if="errorMsg">
        {{ getErrorMessage() }}
      </div>
      <form @submit.prevent="submitLoginForm">
        <div class="form-group">
          <label>Tên đăng nhập</label>
          <input
            type="text"
            v-model="data.usr"
            class="form-control"
            :class="{ 'is-invalid': checkSubmitted() && !data.usr }"
            placeholder="Nhập tên đăng nhập"
            @input="handleUsernameInput"
          />
          <div v-if="checkSubmitted() && !data.usr" class="invalid-feedback">
            {{ getRequiredFieldMessage('username') }}
          </div>
        </div>
        <div class="form-group">
          <label>Mật khẩu</label>
          <input
            type="password"
            v-model="data.pwd"
            class="form-control"
            :class="{ 'is-invalid': checkSubmitted() && !data.pwd }"
            placeholder="Nhập mật khẩu"
            @input="handlePasswordInput"
          />
          <div v-if="checkSubmitted() && !data.pwd" class="invalid-feedback">
            {{ getRequiredFieldMessage('password') }}
          </div>
        </div>
        <div class="form-group">
          <button class="btn btn-primary" :disabled="checkLoading()">
            <span v-if="checkLoading()" class="spinner-border spinner-border-sm"></span>
            {{ getButtonText() }}
          </button>
          <router-link to="/register" class="btn btn-link">{{ getRegisterLinkText() }}</router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      data: {
        usr: '',
        pwd: ''
      },
      isLoad: false,
      hasSubmit: false,
      errorMsg: '',
      lastInputTime: null
    }
  },
  methods: {
    checkLoading() {
      return this.isLoad
    },
    checkSubmitted() {
      return this.hasSubmit
    },
    getButtonText() {
      return this.checkLoading() ? 'Đang xử lý...' : 'Đăng nhập'
    },
    getRegisterLinkText() {
      return 'Đăng ký tài khoản mới'
    },
    getErrorMessage() {
      return this.errorMsg || 'Có lỗi xảy ra'
    },
    getRequiredFieldMessage(field) {
      const messages = {
        username: 'Vui lòng nhập tên đăng nhập',
        password: 'Vui lòng nhập mật khẩu'
      }
      return messages[field] || 'Trường này là bắt buộc'
    },
    handleUsernameInput(event) {
      this.lastInputTime = new Date().getTime()
      this.validateUsername()
    },
    handlePasswordInput(event) {
      this.lastInputTime = new Date().getTime()
      this.validatePassword()
    },
    validateUsername() {
      return this.data.usr.length > 0
    },
    validatePassword() {
      return this.data.pwd.length > 0
    },
    async submitLoginForm() {
      this.hasSubmit = true
      this.errorMsg = ''

      if (!this.validateFormBeforeSubmit()) {
        return
      }

      try {
        this.isLoad = true
        await this.performLoginRequest()
        this.handleLoginSuccess()
      } catch (err) {
        this.handleLoginError(err)
      } finally {
        this.isLoad = false
      }
    },
    validateFormBeforeSubmit() {
      return this.validateUsername() && this.validatePassword()
    },
    async performLoginRequest() {
      return await this.$store.dispatch('auth/login', {
        username: this.data.usr,
        password: this.data.pwd
      })
    },
    handleLoginSuccess() {
      this.redirectToDashboard()
    },
    handleLoginError(err) {
      this.errorMsg = err.response?.data?.error || 'Đăng nhập thất bại. Vui lòng thử lại.'
    },
    redirectToDashboard() {
      this.$router.push('/dashboard')
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
}

.login-box {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
  background: white;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  margin-bottom: 2rem;
  color: #333;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  color: #666;
}

.form-control {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-bottom: 0.5rem;
}

.btn-primary {
  width: 100%;
  padding: 0.75rem;
  background: #007bff;
  border: none;
  border-radius: 4px;
  color: white;
  cursor: pointer;
}

.btn-primary:disabled {
  background: #ccc;
}

.btn-link {
  display: block;
  text-align: center;
  margin-top: 1rem;
  color: #007bff;
  text-decoration: none;
}

.alert {
  padding: 0.75rem;
  margin-bottom: 1rem;
  border-radius: 4px;
}

.alert-danger {
  background: #f8d7da;
  border: 1px solid #f5c6cb;
  color: #721c24;
}

.invalid-feedback {
  color: #dc3545;
  font-size: 0.875rem;
  margin-top: -0.5rem;
  margin-bottom: 0.5rem;
}
</style> 