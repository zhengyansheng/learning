<template>
  <div class="container">

    <header>
      <slot name="header"></slot>
    </header>
    <main>
      <slot></slot>
    </main>
    <footer>
      <slot name="footer"></slot>
    </footer>

  </div>
</template>

<script setup>

// import {defineEmits} from "vue";
//
// // 定义 emit 事件
// const emit = defineEmits(['updateData'])
//
// // 函数触发事件，并传递数据
// const sendDataToParent = () => {
//   const data = {
//     name: '哈密瓜',
//     price: "8.66",
//     num: "2600公斤"
//   }
//   // 触发 'updateData' 事件并传递 data.info
//   emit("updateData", data)
// }


</script>

<style scoped>

</style>