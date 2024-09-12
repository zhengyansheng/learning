<template>
  <div class="container">

    <div class="left">
      <h2>1) v-model 双向绑定 反转字符串</h2>
      输入内容: <input type="text" v-model="message"><br>
      反转内容: {{ reverseMessage() }}<br>
      <hr>

      <h2>2) @click</h2>
      {{ num }}
      <br>
      <button @click="add(1)">+1</button><br>
      <button @click="add(2)">+2</button>
    </div>

    <div class="divider">
    </div>

    <div class="right">
      world
    </div>
  </div>
</template>

<script setup>

import {ref} from "vue";

// v-model
const message = ref("")
const reverseMessage = () => {
  return message.value.split("").reverse().join("")
}

// @click
const num = ref(1)
const add = (param) => {
  if(param===""){
    num.value++
  } else{
    num.value+=param
    debugLog(param)
    debugLog(num.value)
  }
}

const debugLog = (param) => {
  console.log(param)
}


</script>

<style scoped>
.container {
  display: flex;
  height: 100vh; /* 页面高度 */
}

.left {
  flex: 1; /* 左侧占据 50% */
  //background-color: lightblue;
}

.right {
  padding-left: 20px;
  flex: 1; /* 右侧占据 50% */
  //background-color: lightgreen;
}

/* 竖线 */
.divider {
  width: 2px;
  background-color: black; /* 竖线颜色 */
}
</style>