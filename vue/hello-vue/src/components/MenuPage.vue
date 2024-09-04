<template>

  <el-badge :value="99" :max="110" class="item">
    <el-button>评论</el-button>
  </el-badge>

  <el-badge :value="3" class="item">
    <el-button size="small">回复</el-button>
  </el-badge>

  <el-badge :value="2" class="item" type="warning">
    <el-button size="small">回复</el-button>
  </el-badge>
  <br>
  <el-badge is-dot class="item">数据查询</el-badge>
  <br>

  <el-dropdown trigger="click">
    <span class="el-dropdown-link">点我查看</span>
  </el-dropdown>
  <hr>

  <el-row :sm="12" :lg="6">
    <el-result icon="success" title="提示成功" sub-title="请根据提示操作">
      <template #extra>
        <el-button type="primary" size="default">返回</el-button>
      </template>
    </el-result>

    <el-result icon="warning" title="警告成功" sub-title="请根据提示操作">
      <template #extra>
        <el-button type="primary" size="default">返回</el-button>
      </template>
    </el-result>

    <el-result icon="error" title="错误成功" sub-title="请根据提示操作">
      <template #extra>
        <el-button type="primary" size="default">返回</el-button>
      </template>
    </el-result>

    <el-result icon="info" title="信息提示" sub-title="请根据提示操作">
      <template #extra>
        <el-button type="primary" size="small">返回</el-button>
      </template>
    </el-result>
  </el-row>

  <el-result title="404" sub-title="抱歉请求错误">
    <template #icon>
      <el-button type="primary" size="small">返回</el-button>
    </template>
  </el-result>
  <hr>

  <el-radio-group>
    <el-radio label="ltr">从左往右开</el-radio>
  </el-radio-group>
  <br>
  <el-button @click="drawer = true" type="primary" style="margin-left: 16px">点我打开</el-button>
  <el-drawer title="登飞来峰"></el-drawer>

  <hr>
  <el-row class="tac">
    <el-col :span="12">
      <h5>默认颜色</h5>
      <el-menu
          default-active="2"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
      >
        <el-sub-menu index="1">
          <template #title>
            <i class="el-icon-location"></i>
            <span>导航一</span>
          </template>
          <el-menu-item-group>
            <template #title>分组一</template>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="分组2">
            <el-menu-item index="1-3">选项3</el-menu-item>
          </el-menu-item-group>
          <el-sub-menu index="1-4">
            <template #title>选项4</template>
            <el-menu-item index="1-4-1">选项1</el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
        <el-menu-item index="2">
          <i class="el-icon-menu"></i>
          <template #title>导航二</template>
        </el-menu-item>
        <el-menu-item index="3" disabled>
          <i class="el-icon-document"></i>
          <template #title>导航三</template>
        </el-menu-item>
        <el-menu-item index="4">
          <i class="el-icon-setting"></i>
          <template #title>导航四</template>
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :span="12">
      <h5>自定义颜色</h5>
      <el-menu
          :uniqueOpened="true"
          default-active="2"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b"
      >
        <el-sub-menu index="1">
          <template #title>
            <i class="el-icon-location"></i>
            <span>导航一</span>
          </template>
          <el-menu-item-group>
            <template #title>分组一</template>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="分组2">
            <el-menu-item index="1-3">选项3</el-menu-item>
          </el-menu-item-group>
          <el-sub-menu index="1-4">
            <template #title>选项4</template>
            <el-menu-item index="1-4-1">选项1</el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
        <el-menu-item index="2">
          <i class="el-icon-menu"></i>
          <template #title>导航二</template>
        </el-menu-item>
        <el-menu-item index="3" disabled>
          <i class="el-icon-document"></i>
          <template #title>导航三</template>
        </el-menu-item>
        <el-menu-item index="4">
          <i class="el-icon-setting"></i>
          <template #title>导航四</template>
        </el-menu-item>
        <el-sub-menu index="5">
          <template #title>
            <i class="el-icon-location"></i>
            <span>导航一</span>
          </template>
          <el-menu-item-group>
            <template #title>分组一</template>
            <el-menu-item index="5-1">选项1</el-menu-item>
            <el-menu-item index="5-2">选项2</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="分组2">
            <el-menu-item index="5-3">选项3</el-menu-item>
          </el-menu-item-group>
        </el-sub-menu>
      </el-menu>
    </el-col>
  </el-row>


</template>


<script setup>

import {ref} from "vue";


const handleOpen =(key, keyPath) => {
  console.log(key, keyPath)
}

const handleClose =(key, keyPath) => {
  console.log(key, keyPath)
}

const drawer = ref(true)

</script>


<style>

</style>