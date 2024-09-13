<template>
  <div class="admin-login">

    <el-form v-model="form" label-width="auto" style="max-width: 600px">
      <el-form-item label="账户: ">
        <el-input v-model="form.username" />
      </el-form-item>

      <el-form-item label="密码: ">
        <el-input type="password" v-model="form.password" />
      </el-form-item>

      <el-button @click="onSubmit()" type="primary" class="login-btn">登录</el-button>
      <br>
      <div>
        <a href="/register">没有账号？去注册></a>
      </div>
    </el-form>

  </div>
</template>

<script setup>

import {reactive} from "vue";

const form = reactive({
  username: "",
  password: "",
})

const onSubmit = () => {
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