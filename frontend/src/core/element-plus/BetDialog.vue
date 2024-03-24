<template>
  <ElDialog
    v-model="visible"
    title="Warning"
    width="30%"
    center
    class="BetDialog"
    :show-close="false"
  >
    <ElIcon :size="30" color="green" name="el-icon-check">
      <Check />
    </ElIcon>
    <span>
      {{ message }}
    </span>

    <template #footer>
      <span>
        <ElButton v-if="showCancel" @click="closeDialog">Cancel</ElButton>
        <ElButton v-if="showConfirm" type="primary" @click="confirmDialog">
          Confirm
        </ElButton>
      </span>
    </template>
  </ElDialog>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { ElButton, ElDialog, ElIcon } from "element-plus";
import { CircleCloseFilled } from "@element-plus/icons-vue";
import { defineProps, defineEmits } from "vue";

const props = defineProps({
  showDialog: { type: Boolean, default: false },
  message: { type: String, required: true },
  showConfirm: { type: Boolean, default: false },
  showCancel: { type: Boolean, default: false },
});

const emit = defineEmits(["close", "confirm"]);

const visible = computed(() => props.showDialog);

const closeDialog = () => {
  emit("close");
};

const confirmDialog = () => {
  emit("confirm");
};
</script>
<style scoped>
.BetDialog {
  display: flex;
}
</style>
