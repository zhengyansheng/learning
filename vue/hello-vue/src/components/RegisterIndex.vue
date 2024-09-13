<template>
  <div class="admin-login">

    <el-form v-model="form" label-width="auto" style="max-width: 600px">
      <el-form-item label="账户: ">
        <el-input v-model="form.username" placeholder="用户名" />
      </el-form-item>

      <el-form-item label="密码: ">
        <el-input type="password" v-model="form.password" placeholder="密码"/>
      </el-form-item>

      <el-button @click="onRegister()" type="primary" class="login-btn">注册</el-button>

    </el-form>

  </div>
</template>

<script setup>

import {reactive} from "vue";

const form = reactive({
  username: "",
  password: "",
})

const onRegister = () => {
  console.log(form.username, form.password)
}

</script>

<style scoped>

.admin-login {
  height: 100vh; /* 设置视窗高度 */
  margin: 0;
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
}

.login-btn {
  width: 300px;
}

</style>