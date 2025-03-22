<template>
  <div class="simple-container">
    <el-form 
      :model="registerForm" 
      @submit.prevent="handleRegister"
      class="simple-form"
    >
      <h2 class="form-title">用户注册</h2>

      <!-- 用户名 -->
      <el-form-item prop="username">
        <el-input
          v-model="registerForm.username"
          placeholder="用户名"
          prefix-icon="el-icon-user"
        />
      </el-form-item>

      <!-- 密码 -->
      <el-form-item prop="password">
        <el-input
          v-model="registerForm.password"
          type="password"
          placeholder="密码"
          prefix-icon="el-icon-lock"
          show-password
        />
      </el-form-item>

      <!-- 手机号 -->
      <el-form-item prop="phone">
        <el-input
          v-model="registerForm.phone"
          placeholder="手机号"
          prefix-icon="el-icon-mobile-phone"
        />
      </el-form-item>

      <!-- 邮箱和验证码 -->
      <el-form-item class="captcha-container">
        <el-input
          v-model="registerForm.email"
          placeholder="邮箱"
          prefix-icon="el-icon-message"
          style="flex: 2; margin-right: 10px;"
        />
        <el-button
          type="primary"
          :disabled="isSending"
          @click="sendVerificationCode"
          style="flex: 1;"
        >
          {{ isSending ? `${countdown}秒后重发` : '获取验证码' }}
        </el-button>
      </el-form-item>

      <!-- 验证码 -->
      <el-form-item prop="captcha">
        <el-input
          v-model="registerForm.code"
          placeholder="邮箱验证码"
          prefix-icon="el-icon-key"
        />
      </el-form-item>

      <el-form-item>
        <el-button 
          type="primary" 
          native-type="submit"
          class="submit-btn"
        >
          立即注册
        </el-button>
      </el-form-item>

      <div class="form-footer">
        <router-link to="/login" class="register-link">已有账号？立即登录</router-link>
      </div>
    </el-form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      registerForm: {
        username: '',
        password: '',
        phone: '',
        email: '',
        code: ''
      },
      isSending: false,
      countdown: 60
    }
  },
  methods: {
    async sendVerificationCode() {
      // Add email validation before sending
      if (!this.registerForm.email || !this.registerForm.email.includes('@')) {
        this.$message.error('请输入有效的邮箱地址');
        return;
      }

      try {
        const formData = new FormData();
        formData.append('email', this.registerForm.email);
        
        // 修改为POST请求并添加API前缀
        await axios.post('http://127.0.0.1:8088/shawnOJ/email', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });

        // 启动倒计时
        this.isSending = true;
        const timer = setInterval(() => {
          if (this.countdown <= 0) {
            clearInterval(timer);
            this.isSending = false;
            this.countdown = 60;
            return;
          }
          this.countdown--;
        }, 1000);

        this.$message.success('验证码已发送');

      } catch (error) {
        const errMsg = error.response?.data?.message || '验证码发送失败';
        this.$message.error(errMsg);
      }
    },

    async handleRegister() {
      // 验证码校验接口
      const verifyForm = new FormData();
      verifyForm.append('email', this.registerForm.email);
      verifyForm.append('code', this.registerForm.code);
      
      const verifyResponse = await axios.put('http://127.0.0.1:8088/shawnOJ/email', verifyForm, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
      });
  
      if (verifyResponse.status !== 200) {
        const errMsg = verifyResponse.data?.message || '验证码校验失败';
        return this.$message.error(errMsg);
      }
  
      // 注册接口
      const registerForm = new FormData();
      registerForm.append('username', this.registerForm.username);
      registerForm.append('password', this.registerForm.password);
      registerForm.append('phone', this.registerForm.phone);
      registerForm.append('email', this.registerForm.email);
  
      const registerResponse = await axios.put('http://127.0.0.1:8088/shawnOJ', registerForm, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
      });
  
      if (registerResponse.status === 200) {
        this.$message.success('注册成功');
        this.$router.push('/login');
      } else {
        const errMsg = registerResponse.data?.message || '注册流程失败';
        this.$message.error(errMsg);
      }
    }
  }
}
</script>

<style scoped>
.captcha-container {
  display: flex;
  justify-content: space-between;
}

/* 复用登录页面的样式 */
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
