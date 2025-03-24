<template>
  <div class="artistic-container">
    <!-- 保持导航栏风格统一 -->
    <NavBar :initial-nav="1" @nav-change="handleNav" />

    <!-- 个人信息主体 -->
    <main class="art-grid">
      <el-card shadow="hover" class="profile-card">
        <!-- 头像上传区域 -->
        <div class="avatar-section">
          <el-upload
            class="avatar-uploader"
            action="/api/upload"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
          >
            <el-badge :value="userStatus" class="avatar-badge">
              <el-avatar 
                :src="userInfo.avatar" 
                :size="120"
                class="profile-avatar"
              />
            </el-badge>
          </el-upload>
          <h2 class="user-name">{{ userInfo.username }}</h2>
        </div>

        <!-- 个人信息详情 -->
        <div class="info-grid">
          <div class="info-item" v-for="(item, index) in dynamicDetails" :key="index">
            <div class="info-header">
              <div class="color-block" :style="{background: item.color}"></div>
              <h3>{{ item.label }}</h3>
              <el-button 
                type="text" 
                @click="openEditDialog(item)"
                class="edit-btn"
              >
                <i class="el-icon-edit"></i>
              </el-button>
            </div>
            <div class="info-content">{{ item.value }}</div>
          </div>
        </div>
      </el-card>

      <!-- 编辑对话框 -->
      <el-dialog
        :title="`编辑${editingField.label}`"
        v-model:visible="editDialogVisible"
        width="30%"
      >
        <el-form>
          <el-form-item :label="editingField.label">
            <el-input v-model="editingField.value"></el-input>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="editDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmEdit">确定</el-button>
        </template>
      </el-dialog>
    </main>
  </div>
</template>

<script>
import axios from 'axios'
import defaultAvatar from '@/assets/momo.png'
import NavBar from './NavBar.vue'

export default {
  components: {
    NavBar
  },
  data() {
    return {
      defaultAvatar,
      userInfo: {
        avatar: defaultAvatar, // 初始使用默认头像
        username: 'Shawn',
        email: 'shawn@example.com',
        phone: '138-1234-5678',
        status: 'online'
      },
      editDialogVisible: false,
      editingField: {},
      // 修改为只保留结构定义
      userDetails: [
        { label: '用户名', key: 'username', color: '#6a11cb' },
        { label: '电子邮箱', key: 'email', color: '#2ebf91' },
        { label: '手机号码', key: 'phone', color: '#ff6b6b' },
        { label: '账户状态', key: 'status', color: '#4b6cb7' }
      ]
    }
  },
  created() {
    this.loadUserInfo()
    this.loadUserAvatar()
  },
  computed: {
    // 新增计算属性
    dynamicDetails() {
      return this.userDetails.map(item => ({
        ...item,
        value: this.userInfo[item.key] || '加载中...'
      }))
    }
  },
  methods: {
    async loadUserInfo() {
      try {
        const response = await axios.get('http://127.0.0.1:8088/shawnOJ/info', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('jwtToken')}`
          }
        })
        if (response.data.status === 200) {
          // 修改为合并更新而不是直接赋值
          Object.assign(this.userInfo, {
            username: response.data.username,
            email: response.data.email,
            phone: response.data.phone,
            status: this.mapStatus(response.data.state)
          })
        }
      } catch (error) {
        console.error('获取用户信息失败:', error)
      }
    },
    mapStatus(state) {
      return {
        active: '活跃',
        inactive: '未激活',
        frozen: '冻结'
      }[state] || '未知状态'
    },
    // 新增独立头像获取方法
    async loadUserAvatar() {
      await axios.get('http://127.0.0.1:8088/shawnOJ/avatar', {  // 修改为直接使用axios
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`
        }
      }).then(response => {
        if (response.data.status === 200) {
          this.userInfo.avatar = response.data.message;
          console.log('头像地址:', this.userInfo.avatar);
        } else {
          this.setDefaultAvatar();
        }
      }).catch(() => {
        this.setDefaultAvatar();
      })
    },
    setDefaultAvatar() {
      console.log('使用默认头像');
      this.userInfo.avatar = require('@/assets/momo.png');
    },
    async confirmEdit() {
      try {
        await axios.post('/api/update', {
          [this.editingField.key]: this.editingField.value
        }, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('jwtToken')}`
          }
        })
        this.editDialogVisible = false
        this.loadUserInfo()
      } catch (error) {
        console.error('更新失败:', error)
      }
    }
  }
}
</script>

<style scoped>
.profile-card {
  border-radius: 16px;
  background: linear-gradient(145deg, #ffffff, #f8f9fa);
  padding: 2rem;
  margin-top: 80px;
}

.avatar-section {
  text-align: center;
  margin-bottom: 2rem;
}

.profile-avatar {
  border: 3px solid #fff;
  box-shadow: 0 8px 16px rgba(0,0,0,0.1);
  transition: transform 0.3s ease;
}

.profile-avatar:hover {
  transform: scale(1.05);
}

.user-name {
  font-family: 'Cormorant Garamond', serif;
  font-size: 2.2rem;
  margin-top: 1rem;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 2rem;
}

.info-item {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 4px 12px rgba(0,0,0,0.08);
}

.info-header {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}

.info-content {
  font-size: 1.2rem;
  color: #2d3436;
  padding-left: 2rem;
}

.edit-btn {
  margin-left: auto;
  font-size: 1.1rem;
  color: #6a11cb;
  transition: color 0.3s ease;
}

.edit-btn:hover {
  color: #2575fc;
}
</style>
