<template>

  <div class="app-container">
    <router-link :to="{name: 'vue_template'}">02.Vue.js 模版应用</router-link>
    <router-link :to="{name: 'method_cal'}">03.组件的方法和计算属性</router-link>
    <router-link :to="{name: 'v_model'}">04.表单双向绑定</router-link>
    <router-link :to="{name: 'user_interactive'}">05.处理用户交互</router-link>
    <router-link :to="{name: 'component_api'}">06.精通组件和组合API</router-link>
    <router-link :to="{name: 'home'}" >element plus 1</router-link>
    <router-link :to="{name: 'about'}">element plus 2</router-link>
    <router-link :to="{name: 'menu'}">element plus 3</router-link>
    <router-link :to="{name: 'axios'}">axios</router-link>

    <hr>
    <router-view></router-view>
  </div>

</template>

<script>

</script>

<style scoped>

.app-container {
  //text-align: center;
  padding-left: 5px;
  color: #000;
}

.app-container a {
  padding: 10px;
  color: #000;
}

.app-container a.router-link-active {
  color: #fff;
  background-color: #000;
}

</style>
