<template>

  <div>
    {{ data }}
    <hr>
    <h3> {{ data1.city }} 天气预报</h3>
    <h4> {{ data1.date }} {{ data1.week }}</h4>
    <h4> {{ data1.message }}</h4>
    <ul >
      <li v-for="item in data1.obj" :key="item.date">
        <div>
          <h3>{{item.week}}</h3>

          <h3>{{item.wea}}</h3>
        </div>
      </li>
    </ul>
    {{ tqData }}
  </div>

</template>


<script setup>

import { reactive, ref, onMounted } from 'vue';
import axios from 'axios';

const data = ref(null);
const tqData = ref(null);

const data1 = reactive({
  city: "",
  obj: [],
  date: "",
  week: "",
  message: "",
})

// 在组件挂载时调用函数
onMounted(() => {
  fetchData();
  fetchTQiData()
});

const fetchData = async () => {
  try {
    const response = await axios.get('http://localhost:8080/data/user.json');
    console.log(response.data)
    data.value = response.data
  } catch (error) {
    console.error('请求出错', error);
  }
};

const fetchTQiData = async () => {
  try {
    const response = await axios.get('http://v1.yiketianqi.com/api?unescape=1&version=v91&appid=64926583&appsecret=OEfgkR6A&city=北京');
    // console.log(response.data)
    data1.city = response.data.city
    data1.obj = response.data.data
    data1.date = response.data.data[0].date
    data1.week = response.data.data[0].week
    data1.message = response.data.aqi.air_tips
  } catch (error) {
    console.error('请求出错', error);
  }
};

</script>


<style>

li {
  float: left;
  list-style-type: none;
  width: 200px;
  text-align: center;
  border: 1px solid red;
}

</style>