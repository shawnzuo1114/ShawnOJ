<template>
  <router-view></router-view>
</template>

<script>
export default {
  name: 'App',
  // Add router configuration
  beforeCreate() {
    this.$router.push('/')
  }
}
</script>
