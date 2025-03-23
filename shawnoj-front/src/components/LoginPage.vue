<template>
  <div class="simple-container">
    <el-form 
      :model="loginForm" 
      @submit.prevent="handleLogin"
      class="simple-form"
    >
      <h2 class="form-title">用户登录</h2>
      
      <el-form-item>
        <el-input
          v-model="loginForm.username"
          placeholder="用户名"
          prefix-icon="el-icon-user"
        />
      </el-form-item>

      <el-form-item>
        <el-input
          v-model="loginForm.password"
          type="password"
          placeholder="密码"
          prefix-icon="el-icon-lock"
          show-password
        />
      </el-form-item>

      <el-form-item>
        <el-button 
          type="primary" 
          native-type="submit"
          class="submit-btn"
          :loading="loading"
          :disabled="loading"
        >
          {{ loading ? '登录中...' : '立即登录' }}
        </el-button>
      </el-form-item>

      <div class="form-footer">
        <el-checkbox v-model="rememberMe">记住登录</el-checkbox>
        <router-link to="/register" class="register-link">注册账号</router-link>
      </div>
    </el-form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      rememberMe: false,
      loading: false
    }
  },
  methods: {
    async handleLogin() {
      const formData = new FormData();
      formData.append('username', this.loginForm.username);
      formData.append('password', this.loginForm.password);

      const response = await axios.post('http://127.0.0.1:8088/shawnOJ', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });

      if (response.data.status === 200) {
        localStorage.setItem('jwtToken', response.data.message);
        this.$router.push('/home');
      } else {
        this.$message.error(response.data.message);
      }
    }
  }
}
</script>

<style scoped>
.simple-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f5f7fa;
}

.simple-form {
  background: white;
  padding: 30px;
  border-radius: 4px;
  width: 400px;
  box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
}

.form-title {
  text-align: center;
  color: #409eff;
  margin-bottom: 25px;
}

.submit-btn {
  width: 100%;
  margin-top: 10px;
}

.form-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 15px;
}

.register-link {
  color: #409eff;
  text-decoration: none;
}
</style>