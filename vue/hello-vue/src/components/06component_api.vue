<template>
  <div class="container">

    <h1>1) prop 父向子传递数据</h1>

    <Child1ComponentApi date-title="天下第一" v-bind:date-title2="data.content"></Child1ComponentApi>
    <br>

    <Child2ComponentApi :name="data.name" :price="data.price" :num="data.num"></Child2ComponentApi>

    <hr>
    <h1>2) emit 子向父传递数据</h1>
    <div>

      <Child3ComponentApi @updateData="handleUpdateData"></Child3ComponentApi>
      <p>接收到的数据: {{ receivedData }}</p>

    </div>

    <hr>
    <h1>3) slot 插槽</h1>
    <Child4ComponentApi>天下第三</Child4ComponentApi>

    <h1>3.2) slot 插槽 v-slot:slot_name</h1>
    <Child5ComponentApi>
      <template v-slot:header>
        <p style="background-color: #42b983">这里有一个页面标题</p>
      </template>

      <template v-slot:default>
        <p>这里是正文</p>
        <p>这里是正文</p>
        <p>这里是正文</p>
        <p>这里是正文</p>
       <p>这里是正文</p>
      </template>

      <template v-slot:footer>
        <p style="background-color: red">这里有一些联系方式</p>
      </template>
    </Child5ComponentApi>

  </div>
</template>

<script setup>

import Child1ComponentApi from '@/components/06_1child.vue'
import Child2ComponentApi from '@/components/06_2child.vue'
import Child3ComponentApi from '@/components/06_3child.vue'
import Child4ComponentApi from '@/components/06_4child.vue'
import Child5ComponentApi from '@/components/06_5child.vue'
import {reactive, ref} from "vue";


const data = reactive({
  content: "天下第二",
  name: "苹果",
  price: 5,
  num: "100",
})

// 定义接收子组件数据的变量
const receivedData = ref({})

const handleUpdateData = (info) => {
  console.log(info)
  receivedData.value = info
}



</script>

<style scoped>
</style>