<template>
  <div class="container">

  <ul>
    <li>{{ name }}</li>
    <li>{{ price }}</li>
    <li>{{ num }}</li>
    <li>{{ addr }}</li>
  </ul>

  </div>
</template>

<script setup>

defineProps({
  name: String,
  price: {
    type: Number,
    default: 10,
  },
  num: {
    type: String,
    required: true,
  },
  addr: {
    type: String,
    default: "北京",
  },
})


</script>

<style scoped>

</style>