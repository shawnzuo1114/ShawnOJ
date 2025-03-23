<template>
  <div class="artistic-container">
    <!-- 高级质感导航栏 -->
    <header class="luxury-nav">
      <div class="nav-brand" @mouseover="brandHover = true" @mouseleave="brandHover = false">
        <div class="logo-container" :class="{'logo-hover': brandHover}">
          <i class="el-icon-notebook-1 brand-icon"></i>
        </div>
        <span class="logo-text" style="font-size: 25px;">ShawnOJ</span>
        <nav class="nav-links">
          <a v-for="(link, index) in navLinks" 
             :key="index" 
             :class="['nav-item', { 'active': activeNav === index }]"
             @click="activeNav = index">
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
            />
          </el-badge>
        </div>
      </div>
    </header>

    <!-- 艺术化主体 -->
    <main class="art-grid">
      <!-- 左侧创作区 -->
      <el-card shadow="hover" class="creation-card">
        <template #header>
          <div class="card-header">
            <span class="deco-line"></span>
            <h2 class="art-title">数据实验室</h2>
            <i class="el-icon-edit"></i>
          </div>
        </template>
        <div class="palette-grid">
          <el-button 
            v-for="(btn, index) in creativeBtns"
            :key="index"
            class="art-btn"
            :style="{background: btn.color}"
          >
            <i :class="btn.icon"></i>
            <span>{{ btn.label }}</span>
          </el-button>
        </div>
      </el-card>

      <!-- 右侧数据画布 -->
      <div class="data-canvas">
        <div class="stats-card" v-for="stat in visualStats" :key="stat.title">
          <div class="stat-header">
            <div class="color-block" :style="{background: stat.color}"></div>
            <h3>{{ stat.title }}</h3>
          </div>
          <div class="stat-content">
            <div class="art-number">{{ stat.value }}</div>
            <el-progress 
              :percentage="stat.percentage" 
              :color="stat.color"
              :show-text="false"
              :stroke-width="4"
            />
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import axios from 'axios'  // 新增axios引入

export default {
  data() {
    return {
      // 新增用户信息数据
      userInfo: {
        avatar: '' // 初始为空，将从后端获取
      },
      brandHover: false,
      activeNav: 0,
      navLinks: ['主页', '排行榜', '算法题库', '知识库'],
      userStatus: 'online',
      creativeBtns: [
        { icon: 'el-icon-magic-stick', label: '智能生成', color: 'linear-gradient(45deg, #6a11cb, #2575fc)' },
        { icon: 'el-icon-cpu', label: '沙盒实验', color: 'linear-gradient(45deg, #2ebf91, #8360c3)' },
        { icon: 'el-icon-data-analysis', label: '效能分析', color: 'linear-gradient(45deg, #ff6b6b, #ff8e53)' },
        { icon: 'el-icon-chat-line-round', label: '灵感社区', color: 'linear-gradient(45deg, #4b6cb7, #182848)' }
      ],
      visualStats: [
        { title: '创作能量', value: '∞', color: '#6a11cb', percentage: 85 },
        { title: '代码纯度', value: '98%', color: '#2ebf91', percentage: 98 },
        { title: '思维活跃', value: '1.2k', color: '#ff6b6b', percentage: 72 }
      ]
    }
  },
  created() {
    this.loadUserAvatar();
  },
  methods: {
    async loadUserAvatar() {
      await axios.get('http://127.0.0.1:8088/shawnOJ/avatar', {  // 修改为直接使用axios
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`
        }
      }).then(response => {
        if (response.status === 200) {
          this.userInfo.avatar = response.data.message;
        } else {
          this.setDefaultAvatar();
        }
      }).catch(() => {
        this.setDefaultAvatar();
      });
    },
    setDefaultAvatar() {
      console.log('使用默认头像');
      this.userInfo.avatar = require('@/assets/momo.png');
    }
  }
}
</script>

<style scoped>
.artistic-container {
  background: #f8f9fa;
  min-height: 100vh;
  padding: 2rem;
}

/* 新增main元素的上边距 */
.art-grid {
  margin-top: 40px;  /* 根据导航栏高度调整 */
}

.creation-card {
  border: none;
  border-radius: 16px;
  background: linear-gradient(145deg, #ffffff, #f8f9fa);
  box-shadow: 0 12px 24px rgba(0,0,0,0.06);
}

.art-title {
  font-family: 'Cormorant Garamond', serif;
  font-size: 2.4rem;
  color: #2d3436;
  margin: 0 1.5rem;
}

.palette-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1.5rem;
  padding: 1rem;
}

.art-btn {
  height: 120px;
  border: none;
  border-radius: 12px;
  color: white;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.art-btn:hover {
  transform: translateY(-5px);
}

.data-canvas {
  display: grid;
  gap: 1.5rem;
}

.stats-card {
  background: white;
  border-radius: 16px;
  margin-top: 40px;
  padding: 2rem;
  box-shadow: 0 8px 16px rgba(0,0,0,0.05);
}

.art-number {
  font-family: 'Roboto Condensed', sans-serif;
  font-size: 3rem;
  color: #2d3436;
  margin: 1rem 0;
}

.deco-line {
  width: 40px;
  height: 3px;
  background: #6a11cb;
  margin-bottom: 1rem;
}

.color-block {
  width: 24px;
  height: 24px;
  border-radius: 6px;
  margin-right: 1rem;
}

@keyframes float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-8px); }
}

.brand-icon {
  animation: float 3s ease-in-out infinite;
}
</style>

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
  background-color: #67C23A;  /* 新增绿色背景 */
  border: 2px solid white;
}
</style>
