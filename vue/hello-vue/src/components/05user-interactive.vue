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
      商品名称: <input type="text" v-model="data.name"><br>
      <button @click="sub">减一个</button> 购买数量: {{ data.number }}
      <button @click="add2">加一个</button>
      <button @click="addpush">加入购物车</button>
      <br>
      {{data.name}} X{{ data.number}}
      <hr>

      <h2>练习2</h2>
      <p>{{ data2.msg }}</p>
      <button v-on:click="handleClick()">单击按钮</button>
      <button @click="handleClick()">单击按钮</button>
      <br>
      <select>
        <option>Python</option>
        <option>Go</option>
        <option>Java</option>
      </select>

      <p>表单提交</p>

      <input type="checkbox" v-model="data2.isAgree" value="data2.isAgree"> 同意本站协议 <br>
      <button :disabled="data2.isDisabled">注册</button>

      {{ data2 }}



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
  name: "",
  number: 0,
  isMax: false,
  list: [],

})

const add2 = () => {
  data.number++
}

const sub = () => {
  data.number--
}

const addpush = () => {
  data.list.push({name: data.name, number: data.number})
  console.log(data.list)
}

watch(data, (newValue, oldValue) => {
  console.log(newValue.number, oldValue.number)
  if (newValue.number > 10) {
    data.isMax = true
  }
  if (newValue.number < 0) {
    data.number = 0
  }
})

// trains2
const data2 = reactive({
  msg: "注册用户",
  isDisabled: true,
  isAgree: false,
})

watch(data2, (newValue, oldValue) => {
  console.log(oldValue)
  if(newValue.isAgree) {
    data2.isDisabled = false
  } else {
    data2.isDisabled = true
  }
})

const handleClick = () => {
  console.log("btn is clicked")
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