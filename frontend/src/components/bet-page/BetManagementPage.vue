<template>
  <section>
    <header class="BetManagementPage__header">
      <p class="BetManagementPage__header--description">Bet menager</p>
    </header>
    <div class="BetManagementPage__gridArea">
      <BetRow class="BetManagementPage__firstRow">
        <BetCol :span="12">
          asdsad
          <BetSelect :options="t" model-value="aaa" />
        </BetCol>
      </BetRow>
      <BetRow class="BetManagementPage__secondRow">
        <BetCol :span="12"> asdsad</BetCol>
      </BetRow>
    </div>
  </section>
</template>
<script setup lang="ts">
import { BetRow, BetCol, BetSelect } from "@/core/element-plus/index";
import { ref } from "vue";
const test = ref("testaa");

const t = ["aaa", "bbb"];
const modelValue = ref("");
</script>
<style lang="scss">
.BetManagementPage {
  &__header {
    height: $gutterXXXL;
    width: 100%;
    background-color: $colorLightblue;
    text-align: center;
    position: fixed;
    font-weight: 600;
    text-transform: uppercase;

    &--description {
      margin: $gutterM 0 0;
    }
  }

  &__gridArea {
    padding-top: $gutterXXXL;
    height: 100%;
  }

  &__firstRow {
    background-color: green;
    height: 150px;
  }

  &__secondRow {
    background-color: black;
    height: 400px;
  }
}
</style>
