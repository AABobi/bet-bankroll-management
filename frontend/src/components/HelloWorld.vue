<template>
  <button data-test="data-test" @click="testFun">TEST</button>
</template>
<script setup lang="ts">
const testFun = () => {
  console.log("TESTTEST");
};
</script>
