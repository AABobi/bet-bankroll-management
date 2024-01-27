<template>
  <div>TEST</div>
</template>
<script setup lang="ts">
//import { useAuthenticationStore } from "@/store/authentication-store";
import { useTestStore } from "@/store/test-store";

import { ref } from "vue";
//const test = ref("test")
const store = useTestStore();
</script>
