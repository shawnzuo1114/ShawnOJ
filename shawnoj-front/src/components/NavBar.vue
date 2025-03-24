<template>
  <header class="luxury-nav">
    <div class="nav-brand" @mouseover="brandHover = true" @mouseleave="brandHover = false">
      <div class="logo-container" :class="{'logo-hover': brandHover}">
        <i class="el-icon-notebook-1 brand-icon"></i>
      </div>
      <span class="logo-text" style="font-size: 25px;" @click="handleHomePage">ShawnOJ</span>
      <nav class="nav-links">
        <a 
          v-for="(link, index) in navLinks" 
          :key="index" 
          :class="['nav-item', { 'active': activeNav === index }]"
          @click="handleNavClick(index)"
        >
          {{ link }}
        </a>
      </nav>
    </div>
    <div class="user-menu">
      <div class="avatar-container">
        <el-badge :value="userStatus" class="avatar-badge">
          <el-avatar 
            :src="userInfo.avatar"
            icon="el-icon-user" 
            class="user-avatar"
            @click="handleAvatarClick"
          />
        </el-badge>
      </div>
    </div>
  </header>
</template>

<script>
import axios from 'axios'

export default {
  name: 'NavBar',
  props: {
    initialNav: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      brandHover: false,
      activeNav: this.initialNav,
      userStatus: 'online',
      userInfo: { avatar: '' },
      navLinks: ['主页', '排行榜', '算法题库', '知识库']
    }
  },
  created() {
    this.loadUserAvatar();
  },
  methods: {
    async loadUserAvatar() {
      try {
        const response = await axios.get('http://127.0.0.1:8088/shawnOJ/avatar', {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`
          }
        });
        if (response.status === 200) {
          this.userInfo.avatar = response.data.message;
        } else {
          this.setDefaultAvatar();
        }
      } catch {
        this.setDefaultAvatar();
      }
    },
    setDefaultAvatar() {
      this.userInfo.avatar = require('@/assets/momo.png');
    },
    handleNavClick(index) {
      this.activeNav = index;
      this.$emit('nav-change', index);
    },
    handleAvatarClick() {
      this.$router.push('/info')// 处理头像点击事件 
    },
    handleHomePage() {
      this.$router.push('/home') 
    }
  }
}
</script>

<style scoped>
.luxury-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 5%;
  background: rgba(255, 255, 255, 0.86);
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.nav-brand {
  display: flex;
  align-items: center;
  gap: 24px;
}

.logo-container {
  display: flex;
  padding: 8px;
  background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
  border-radius: 12px;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.logo-hover {
  transform: rotate(15deg) scale(1.1);
}

.nav-links {
  display: flex;
  gap: 32px;
  margin-left: 40px;
}

.nav-item {
  position: relative;
  color: #2d3436;
  font-weight: 500;
  cursor: pointer;
  padding: 8px 0;
  transition: color 0.3s ease;
}

.nav-item::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 2px;
  background: #6a11cb;
  transition: width 0.3s ease;
}

.nav-item:hover::after {
  width: 100%;
}

.nav-item.active {
  color: #6a11cb;
  font-weight: 600;
}

.user-avatar {
  transition: transform 0.3s ease;
}

.user-avatar:hover {
  transform: scale(1.1);
}

.avatar-badge ::v-deep .el-badge__content {
  height: 18px;
  padding: 0 4px;
  line-height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #67C23A;
  border: 2px solid white;
}

@keyframes float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-8px); }
}

.brand-icon {
  animation: float 3s ease-in-out infinite;
}
</style>
