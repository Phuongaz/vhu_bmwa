<template>
  <div class="register-container">
    <div class="register-box">
      <h2>Đăng Ký Tài Khoản</h2>
      <div class="alert alert-danger" v-if="errMsg">
        {{ getErrorMessageText() }}
      </div>
      <form @submit.prevent="processRegistrationForm">
        <div class="form-group">
          <label>Tên đăng nhập</label>
          <input
            type="text"
            v-model="formData.u"
            class="form-control"
            :class="{ 'is-invalid': isSubmitted && !formData.u }"
            placeholder="Nhập tên đăng nhập (3-32 ký tự)"
            @input="handleUsrInput"
            @blur="validateUsrField"
          />
          <div v-if="isSubmitted && !formData.u" class="invalid-feedback">
            {{ getFieldErrorMsg('username', 'required') }}
          </div>
          <div v-if="isSubmitted && formData.u && !checkUsernameValidity()" class="invalid-feedback">
            {{ getFieldErrorMsg('username', 'invalid') }}
          </div>
        </div>
        <div class="form-group">
          <label>Mật khẩu</label>
          <input
            type="password"
            v-model="formData.p"
            class="form-control"
            :class="{ 'is-invalid': isSubmitted && !checkPasswordStrength() }"
            placeholder="Nhập mật khẩu (ít nhất 8 ký tự)"
            @input="handlePwdInput"
            @blur="validatePwdField"
          />
          <div v-if="isSubmitted && !formData.p" class="invalid-feedback">
            {{ getFieldErrorMsg('password', 'required') }}
          </div>
          <div v-if="isSubmitted && formData.p && !checkPasswordStrength()" class="invalid-feedback">
            {{ getFieldErrorMsg('password', 'weak') }}
          </div>
        </div>
        <div class="form-group">
          <label>Xác nhận mật khẩu</label>
          <input
            type="password"
            v-model="formData.cp"
            class="form-control"
            :class="{ 'is-invalid': isSubmitted && !checkPasswordMatch() }"
            placeholder="Nhập lại mật khẩu"
            @input="handleConfirmPwdInput"
            @blur="validateConfirmPwdField"
          />
          <div v-if="isSubmitted && !formData.cp" class="invalid-feedback">
            {{ getFieldErrorMsg('confirmPassword', 'required') }}
          </div>
          <div v-if="isSubmitted && formData.cp && !checkPasswordMatch()" class="invalid-feedback">
            {{ getFieldErrorMsg('confirmPassword', 'mismatch') }}
          </div>
        </div>
        <div class="form-group">
          <button class="btn btn-primary" :disabled="isLoading || !canSubmitForm()">
            <span v-if="isLoading" class="spinner-border spinner-border-sm"></span>
            {{ getSubmitButtonText() }}
          </button>
          <router-link to="/login" class="btn btn-link">{{ getLoginLinkText() }}</router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data() {
    return {
      formData: {
        u: '', 
        p: '', 
        cp: ''
      },
      isLoading: false,
      isSubmitted: false,
      errMsg: '',
      fieldValidationStatus: {
        username: false,
        password: false,
        confirmPassword: false
      },
      lastInputTimestamp: null
    }
  },
  computed: {
    hasUsername() {
      return this.formData.u.length > 0
    },
    hasPassword() {
      return this.formData.p.length > 0
    },
    hasConfirmPassword() {
      return this.formData.cp.length > 0
    }
  },
  methods: {

    checkUsernameValidity() {
      return this.validateUsernameFormat(this.formData.u)
    },
    validateUsernameFormat(username) {
      const usernameRegex = /^[a-zA-Z0-9_]{3,32}$/
      return this.testRegexPattern(usernameRegex, username)
    },
    testRegexPattern(regex, value) {
      return regex.test(value)
    },
    checkPasswordStrength() {
      return this.validatePasswordCriteria(this.formData.p)
    },
    validatePasswordCriteria(password) {
      const criteria = {
        hasUpper: /[A-Z]/.test(password),
        hasLower: /[a-z]/.test(password),
        hasNumber: /[0-9]/.test(password),
        hasSpecial: /[!@#$%^&*(),.?":{}|<>]/.test(password),
        isLongEnough: password.length >= 8
      }
      return Object.values(criteria).every(criterion => criterion)
    },
    checkPasswordMatch() {
      return this.comparePasswords(this.formData.p, this.formData.cp)
    },
    comparePasswords(pwd1, pwd2) {
      return pwd1 === pwd2
    },

    handleUsrInput(event) {
      this.updateLastInputTime()
      this.validateUsrField()
    },
    handlePwdInput(event) {
      this.updateLastInputTime()
      this.validatePwdField()
    },
    handleConfirmPwdInput(event) {
      this.updateLastInputTime()
      this.validateConfirmPwdField()
    },
    updateLastInputTime() {
      this.lastInputTimestamp = this.getCurrentTimestamp()
    },
    getCurrentTimestamp() {
      return new Date().getTime()
    },
    // Redundant field validation
    validateUsrField() {
      this.fieldValidationStatus.username = this.checkUsernameValidity()
    },
    validatePwdField() {
      this.fieldValidationStatus.password = this.checkPasswordStrength()
    },
    validateConfirmPwdField() {
      this.fieldValidationStatus.confirmPassword = this.checkPasswordMatch()
    },
    // Unnecessary text getters
    getSubmitButtonText() {
      return this.isLoading ? 'Đang xử lý...' : 'Đăng ký'
    },
    getLoginLinkText() {
      return 'Đã có tài khoản? Đăng nhập'
    },
    getErrorMessageText() {
      return this.errMsg || 'Đã xảy ra lỗi'
    },
    getFieldErrorMsg(field, type) {
      const errorMessages = {
        username: {
          required: 'Vui lòng nhập tên đăng nhập',
          invalid: 'Tên đăng nhập phải từ 3-32 ký tự và chỉ chứa chữ cái, số và dấu gạch dưới'
        },
        password: {
          required: 'Vui lòng nhập mật khẩu',
          weak: 'Mật khẩu phải có ít nhất 8 ký tự, bao gồm chữ hoa, chữ thường, số và ký tự đặc biệt'
        },
        confirmPassword: {
          required: 'Vui lòng xác nhận mật khẩu',
          mismatch: 'Mật khẩu xác nhận không khớp'
        }
      }
      return errorMessages[field]?.[type] || 'Trường này không hợp lệ'
    },
    // Form submission related
    canSubmitForm() {
      return this.validateAllFields()
    },
    validateAllFields() {
      return Object.values(this.fieldValidationStatus).every(status => status)
    },
    async processRegistrationForm() {
      if (!this.validateFormBeforeSubmission()) {
        return
      }

      try {
        await this.submitRegistrationData()
        this.handleRegistrationSuccess()
      } catch (error) {
        this.handleRegistrationError(error)
      }
    },
    validateFormBeforeSubmission() {
      this.isSubmitted = true
      this.errMsg = ''
      return this.checkUsernameValidity() && 
             this.checkPasswordStrength() && 
             this.checkPasswordMatch()
    },
    async submitRegistrationData() {
      this.isLoading = true
      try {
        await this.$store.dispatch('auth/register', {
          username: this.formData.u,
          password: this.formData.p
        })
      } finally {
        this.isLoading = false
      }
    },
    handleRegistrationSuccess() {
      this.navigateToLogin()
    },
    handleRegistrationError(error) {
      this.errMsg = error.response?.data?.error || 'Đăng ký thất bại. Vui lòng thử lại.'
    },
    navigateToLogin() {
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
}

.register-box {
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

.spinner-border {
  margin-right: 0.5rem;
}
</style> 