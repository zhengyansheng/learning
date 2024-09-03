<template>

  <div class="home-container">
    <h3>Home组件</h3>

    <el-row>
    <el-button>按钮组件</el-button>
    <el-button type="primary">主要按钮</el-button>
    <el-button type="success">成功按钮</el-button>
    <el-button type="info">信息按钮</el-button>
    <el-button type="warning">警告按钮</el-button>
    <el-button type="danger">危险按钮</el-button>
    </el-row>

    <el-row>
      <el-button plain>按钮组件</el-button>
      <el-button type="primary" plain>主要按钮</el-button>
      <el-button type="success" plain>成功按钮</el-button>
      <el-button type="info" plain>信息按钮</el-button>
      <el-button type="warning" plain>警告按钮</el-button>
      <el-button type="danger" plain>危险按钮</el-button>
    </el-row>

    <el-row>
      <el-button round>按钮组件</el-button>
      <el-button type="primary" round>主要按钮</el-button>
      <el-button type="success" round>成功按钮</el-button>
      <el-button type="info" round>信息按钮</el-button>
      <el-button type="warning" round>警告按钮</el-button>
      <el-button type="danger" round>危险按钮</el-button>
    </el-row>

    <hr>

    <el-button-group>
      <el-button type="primary">上一页</el-button>
      <el-button type="primary">下一页</el-button>
    </el-button-group>

    <el-button-group>
      <el-button type="primary" :loading="true">正在加载中</el-button>
      <el-button type="primary" :loading="false">正在加载中</el-button>
    </el-button-group>

    <hr>

    <el-link>文字链接</el-link>
    <el-link type="primary">文字链接</el-link>
    <el-link type="danger" disabled>文字链接</el-link>

    <hr>
    <el-space>间距组件</el-space>

    <hr>
    <div>
      <el-radio v-model="data.radio" label="1">男</el-radio>
      <el-radio v-model="data.radio" label="2">女</el-radio>
      <el-radio v-model="data.radio1" label="选中且禁用">财务部</el-radio>
    </div>
    <div>
      <el-radio-group v-model="data.radio2">
        <el-radio label="1">销售部门</el-radio>
        <el-radio label="2">设计部门</el-radio>
        <el-radio label="3">财务部门</el-radio>
      </el-radio-group>
    </div>
    <div>
      <el-radio-group v-model="data.radio3">
        <el-radio-button label="销售部门"></el-radio-button>
        <el-radio-button label="设计部门"></el-radio-button>
        <el-radio-button label="财务部门"></el-radio-button>
      </el-radio-group>
    </div>
    <hr>

    <div>
      <el-checkbox v-model="data_checkbox.checked1" label="洗衣机"></el-checkbox>
      <el-checkbox v-model="data_checkbox.checked2" label="冰箱"></el-checkbox>
      <el-checkbox v-model="data_checkbox.checked3" label="电视机"></el-checkbox>
    </div>

    <el-checkbox-group v-model="data_checkbox.checkList">
      <el-checkbox label="洗衣机"></el-checkbox>
      <el-checkbox label="冰箱"></el-checkbox>
      <el-checkbox label="禁用" disabled></el-checkbox>
    </el-checkbox-group>

    <el-checkbox-group v-model="data_checkbox.checkList">
      <el-checkbox-button label="洗衣机"></el-checkbox-button>
      <el-checkbox-button label="冰箱"></el-checkbox-button>
      <el-checkbox-button label="电视机" disabled></el-checkbox-button>
    </el-checkbox-group>

    <el-checkbox v-model="data_checkbox.checkALL" @click="handleCheckAll">全选</el-checkbox>
  </div>

  <el-checkbox-group v-model="data_checkbox.checkedCities">
    <el-checkbox v-for="city in data_checkbox.cities" :label="city" :key="city">{{city}}</el-checkbox>
  </el-checkbox-group>

  <hr>

  <el-input v-model="input" placeholder="请输入内容"></el-input>
  <el-input v-model="input" placeholder="请输入内容" :disabled="true"></el-input>
  <el-input v-model="input" placeholder="请输入内容" clearable></el-input>
  <el-input v-model="input" placeholder="请输入密码" show-password></el-input>
  <el-input type="textarea" :rows="2" placeholder="请输入内容" v-model="textarea"></el-input>
  <hr>

  <el-select placeholder="请选择" v-model="data_select.value">
    <el-option v-for="item in data_select.options" :label="item.label" :value="item.value" :key="item.label" :disabled="item.disabled">
      {{item.label}}
    </el-option>
  </el-select>


</template>


<script setup>

import {reactive, ref} from "vue";

const data = {
  radio: "1",
  radio1: "选中且禁用",
  radio2: "2",
  radio3: "财务部门",
}

const cityOptions = ["上海", "北京", "广州","深圳"]

const data_checkbox = reactive({
  checked1: true,
  checked2: true,
  checked3: false,
  checkList: true,
  checkALL: false,
  checkedCities: ["上海", "北京"],
  cities: cityOptions,
})

function handleCheckAll() {
  console.log(data_checkbox.checkALL)
  if (data_checkbox.checkALL) {
    data_checkbox.checkedCities = []
  } else {
    data_checkbox.checkedCities = cityOptions
  }
}

const input = ref("你好")
const textarea = ref("......")


const data_select = reactive({
  options: [
    {value: "1", label: "洗衣机"},
    {value: "2", label: "冰箱"},
    {value: "3", label: "空调", disabled: true},
    {value: "4", label: "吹风机"},
  ],
  value: ""
})


</script>


<style>

</style>