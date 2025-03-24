<template>
  <div class="artistic-container">
    <!-- 使用导航栏组件 -->
    <NavBar @nav-change="handleNavChange" />

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
import NavBar from './NavBar.vue'

export default {
  components: {
    NavBar
  },
  data() {
    return {
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
  methods: {
    handleNavChange(index) {
      console.log('当前导航索引:', index);
      // 这里可以添加导航变化后的处理逻辑
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

.art-grid {
  margin-top: 40px;
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
</style>
