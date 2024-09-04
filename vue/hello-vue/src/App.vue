<template>

  <div class="app-container">
<!--    <h1>App根组件</h1>-->
<!--    <router-link to="/home" >首页</router-link>-->
    <router-link :to="{name: 'root'}" >首页</router-link>
<!--    <router-link to="/about" v-slot="{naviqate}">-->
<!--      <button @click="naviqate" @keypress.enter="naviqate">关于</button>-->
<!--    </router-link>-->
    <router-link :to="{name: 'about'}">关于</router-link>
    <router-link :to="{name: 'menu'}">菜单</router-link>
    <router-link :to="{name: 'axios'}">axios</router-link>
    <hr>
    <router-view></router-view>
  </div>


<!--  <img alt="Vue logo" src="./assets/logo.png">-->
<!--  <HelloWorld msg="Welcome to Your Vue.js App"/>-->
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'


export default {
  name: 'App',
  components: {
    // HelloWorld,
  }
}
</script>

<style scoped>

.app-container {
  text-align: center;
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
