<template>
  <ElSelect
    v-model="value"
    placeholder="Select"
    size="large"
    style="width: 240px"
  >
    <ElOption v-for="item in options" :key="item" :label="item" :value="item" />
  </ElSelect>
</template>
<script lang="ts" setup>
import { ElSelect, ElOption } from "element-plus";
import { computed, ref } from "vue";
import { defineProps, defineEmits } from "vue";

const props = defineProps({
  modelValue: String || undefined,
  options: {
    type: Array,
    required: true,
  },
});

const emit = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => {
    console.log(props.modelValue);
    return props.modelValue;
  },
  set: (v) => {
    console.log(v);
    emit("update:modelValue", v);
  },
});

const test1 = ref("");
</script>
<style lang="scss"></style>
