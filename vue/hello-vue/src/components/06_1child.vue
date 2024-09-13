<template>
  <div class="container">

    <h3> {{ dateTitle }}</h3>
    <h3> {{ dateTitle2 }}</h3>

  </div>
</template>

<script setup>

defineProps({
  dateTitle: {
    type: String,
    required: true
  },
  dateTitle2: {
    type: String,
  },
})


</script>

<style scoped>

</style>