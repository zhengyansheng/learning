<template>
  <div class="container">

    <div class="left">

      <h2>01) 监听事件 v-on:click | @click | @dblclick </h2>
      <button v-on:click="num--">减少1分</button>
      <button v-on:click="num++">增加1分</button><br>

      单击事件
      <button @click="reduce(1)">减少1分</button>
      <button @click="add(1)">增加1分</button>
      <p>期末考试总成绩是: {{num}}</p>

      双击事件
      <button @dblclick="reduce(10)">减少10分</button>
      <button @dblclick="add(10)">增加10分</button>
      <p>期末考试总成绩是: {{num}}</p>

      <h2>02) 使用监听器 watch | watchEffect</h2>

      时：<input type="text" v-model="time">
      分：<input type="text" v-model="minute">


      <h2>03) 监听对象</h2>
      商品价格: <input type="text" v-model="goods.price">
      {{ goods.desc }}


    </div>

    <div class="divider">
    </div>

    <div class="right">
      <h2>练习1</h2>
      商品名称: <input type="text"><br>
      <button>减一个</button> 购买数量: {{ data.number }}
      <button>加一个</button>
      <button>加入购物车</button>
      <br>
      洗衣机

    </div>
  </div>
</template>

<script setup>

import {reactive, ref, watch, watchEffect} from "vue";

const num = ref(360)

const add = (param) => {
  num.value+=param
}

const reduce = (param) => {
  num.value-=param
}

// watch
const time = ref(0)
const minute = ref(0)

watch(time, (newValue, oldValue) => {
  console.log(`time changed from ${oldValue} to ${newValue}`)
  // minute.value = time.value * 60
  minute.value = newValue * 60
})

watch(minute, (newValue, oldValue) => {
  console.log(`minute changed from ${oldValue} to ${newValue}`)
  // minute.value = time.value * 60
  time.value = newValue / 60
})

watchEffect(() => {
  console.log(`time is: ${time.value}, minute is: ${minute.value}`)
})

// 监听对象
const goods = reactive({
  name: "洗衣机",
  price : 0,
  desc: "",
})

watch(goods, (newValue, oldValue) => {
  console.log(oldValue, newValue)
  if (newValue.price > 8000) {
    goods.desc = "价格太贵了，不可以采购"
  } else {
    goods.desc = "价格合适，可以采购"
  }
}, { immediate: false, deep: true })

// trains
const data = reactive({
  number: 0,

})


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