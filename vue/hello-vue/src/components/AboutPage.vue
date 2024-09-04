<template>

  <div class="home-container">

    <el-alert title="提示的文案" type="success"></el-alert>
    <el-alert title="提示的文案" type="info" effect="dark"></el-alert>
    <el-alert title="提示的文案" type="warning" center></el-alert>
    <el-alert title="提示的文案" type="error" close-text="关闭"></el-alert>
    <el-alert title="古诗介绍" type="success" description="古诗是古代中国诗歌的环衬....."></el-alert>
    <hr>

    <el-button plain @click="open1">可自动关闭</el-button>
    <el-button plain @click="open2">不会自动关闭</el-button>

    <el-button plain @click="open3"> Top Right </el-button>
    <el-button plain @click="open4"> Bottom Right </el-button>
    <el-button plain @click="open5"> Bottom Left </el-button>
    <el-button plain @click="open6"> Top Left </el-button>
    <hr>

    <el-button :plain="true" @click="open10">成功</el-button>
    <el-button :plain="true" @click="open11">警告</el-button>
    <el-button :plain="true" @click="open12">消息</el-button>
    <el-button :plain="true" @click="open13">错误</el-button>
    <hr>

    <el-table :data="tableData.tableList" stripe  style="width: 100%">
      <el-table-column prop="date" label="日期" sortable width="150" />
      <el-table-column label="配送信息">
        <el-table-column prop="name" label="姓名" width="120" />
        <el-table-column label="地址信息">
          <el-table-column prop="state" label="省份" width="120" />
          <el-table-column prop="city" label="城市" width="120" />
          <el-table-column prop="address" label="地址" />
          <el-table-column prop="zip" label="邮编" width="120" />
        </el-table-column>
      </el-table-column>
    </el-table>
    <hr>

    


  </div>

</template>


<script setup>
import {h, reactive} from 'vue'
import { ElNotification, ElMessage  } from 'element-plus'

function open1() {
  ElNotification({
    title: '主题',
    message: h('i', { style: 'color: teal' }, '这是提示的内容'),
  })
}

function open2() {
  ElNotification({
    title: '主题',
    message: '这是一条不会关闭的消息',
    duration: 0,
  })
}

const open3 = () => {
  ElNotification({
    title: 'Custom Position',
    message: "I'm at the top right corner",
  })
}

const open4 = () => {
  ElNotification({
    title: 'Custom Position',
    message: "I'm at the bottom right corner",
    position: 'bottom-right',
  })
}

const open5 = () => {
  ElNotification({
    title: 'Custom Position',
    message: "I'm at the bottom left corner",
    position: 'bottom-left',
  })
}

const open6 = () => {
  ElNotification({
    title: 'Custom Position',
    message: "I'm at the top left corner",
    position: 'top-left',
  })
}

const open10 = () => {
  ElMessage.success(
      {
        message: "恭喜你，这是一条成功的消息",
        type: "success",
      }
  )
}

const open11 = () => {
  ElMessage.warning(
      {
        message: "恭喜你，这是一条警告的消息",
        type: "warning",
      }
  )
}

const open12 = () => {
  ElMessage("这是一条消息")
}

const open13 = () => {
  ElMessage.error("这是一条错误消息")
}

const tableData = reactive({
  tableList: [
      {date: "2023-05-03", name: "小明", state: "上海", city: "普陀区", address: "1518号", zip: 200333},
      {date: "2023-05-02", name: "小红", state: "北京", city: "海淀区", address: "1008号", zip: 100080},
      {date: "2023-05-05", name: "小红", state: "北京", city: "海淀区", address: "1008号", zip: 100080},
  ]
})



</script>


<style>

</style>