<template>

  <div class="app-container">
    <h1>App根组件</h1>
    <router-link to="/home">首页</router-link>
    <router-link to="/about">关于</router-link>
    <hr>
    <router-view></router-view>
  </div>

<!--  <HomePage/>-->
<!--  <AboutPage/>-->
<!--  <img alt="Vue logo" src="./assets/logo.png">-->
<!--  <HelloWorld msg="Welcome to Your Vue.js App"/>-->
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
// import HomePage from './components/HomePage.vue'
// import AboutPage from './components/AboutPage.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld,
    // HomePage,
    // AboutPage,
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
