<template>
  <div class="container">

    <div class="left">

      <h2>01) input 单行</h2>
      <input type="text" v-model="inputMessage" value="hello world"><br>
      {{ inputMessage }}

      <hr>
      <h2>02) textarea 多行文本</h2>
      <textarea v-model="ttMessage"></textarea>
      <p> {{ ttMessage }}</p>

      <hr>
      <h2>03) checkbox 多选</h2>
      <span>选择需要采购的商品</span><br>
      <input type="checkbox" v-model="data.deviceList" id="洗衣机" value="洗衣机">洗衣机
      <input type="checkbox" v-model="data.deviceList" id="冰箱" value="冰箱">冰箱
      <input type="checkbox" v-model="data.deviceList" id="电视机" value="电视机">电视机
      <input type="checkbox" v-model="data.deviceList" id="空调" value="空调">空调
      <br>
      选中的商品: {{ data.deviceList }}

      <hr>
      <h2>04) radio 单选</h2>
      <input type="radio" v-model="radioChecked" value="A">A 洗衣机<br>
      <input type="radio" v-model="radioChecked" value="B">B 冰箱<br>
      <input type="radio" v-model="radioChecked" value="C">C 空调<br>
      <input type="radio" v-model="radioChecked" value="D">D 电视机<br>
      选中: {{ radioChecked }}

      <hr>
      <h2>05) select 下拉列表单选</h2>
      <span>选择你喜欢的课程</span>
      <select v-model="selected" style="height: 30px">
        <option disabled value="">选择喜欢的课程</option>
        <option>Java开发班</option>
        <option>Python开发班</option>
        <option>前端开发班</option>
      </select>
      <br>
      <span>选择的课程: {{selected}}</span>

      <hr>
      <h2>06) select 下拉列表多选</h2>
      <span>选择你喜欢的课程</span>
      <select v-model="selected" multiple style="height: 100px">
        <option disabled value="">选择喜欢的课程</option>
        <option>Java开发班</option>
        <option>Python开发班</option>
        <option>前端开发班</option>
      </select>
      <br>
      <span>选择的课程: {{selected}}</span>
    </div>

    <div class="divider">
    </div>

    <div class="right">
      <h2>07) select 下拉列表 v-for实现</h2>
      <span>选择你喜欢的课程</span>
      <select v-model="selected">
        <option v-for="option in options" :value="option.value" :key="option.value">{{option.text}}</option>
      </select>
      <hr>
      <h2>08) 修饰符lazy</h2>
      <input v-model.lazy="message">
      {{ message }}
      <hr>
      <h2>09) 修饰符number</h2>
      <input type="number" v-model.number="val">
      数据: {{ val }} , 类型: {{ typeof(val)}}
      <hr>
      <h2>10) 修饰符trim</h2>
      <input type="text" v-model.trim="val2"><br>
      长度: {{ val2 }}

      <h2>11) 练习</h2>
      用户名称:
      <input type="text" v-model.trim="data.username">
      <br>
      用户密码: <input type="password" v-model.trim="data.password">
      <br>
      性别:
      <input type="radio" v-model="data.sex" value="女">女 <input type="radio" v-model="data.sex" value="男">男
      <br>
      喜欢的技术:
      <input type="checkbox" v-model="data.loves" id="1" value="Java开发">Java开发
      <input type="checkbox" v-model="data.loves" id="2" value="Python开发">Python开发
      <input type="checkbox" v-model="data.loves" id="3" value="Go开发">Go开发
      <br>
      就业城市:
      <select v-model="data.selectedCity">
        <option value="">未选择</option>
        <option v-for="city in data.citys" :value="city.name" :key="city.id">{{city.name}}</option>
      </select>
      <br>
      介绍:<br>
      <textarea v-model.trim="data.intra"></textarea>
      <br>
      <button @click="submitUserInfo()">注册</button>

    </div>
  </div>
</template>

<script setup>

import {reactive, ref} from "vue";

const inputMessage = ref("")
const ttMessage = ref("")

// checkbox
const data = reactive({
  deviceList: [],
  username: "",
  password: "",
  selectedCity: "",
  loves: [],
  intra: "",
  sex: "",
  citys: [{id:1, name:"北京"},{id:2, name:"上海"},{id:3, name:"广州"},],
})

// radio
const radioChecked = ref("")

// selected
const selected = ref("")

// v-for select
const options = reactive([
  {text: "课程1", "value": "Java开发班"},
  {text: "课程2", "value": "Python开发班"},
  {text: "课程3", "value": "前端开发班"},
])

// lazy
const message = ref("")

// number
const val2 = ref()

// trains
const submitUserInfo = () => {
  let optData = {
    "username": data.username,
    "password": data.password,
    "city": data.selectedCity,
    "intra": data.intra,
    "loves": data.loves,
    "sex": data.sex,
  }
  console.log(optData)
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