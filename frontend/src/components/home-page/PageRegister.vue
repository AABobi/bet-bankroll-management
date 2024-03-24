<template>
  <section>
    <BetForm class="PageLogin__loginBox" label-position="top">
      <header class="PageLogin__header">
        <p>Welcome</p>
        <span class="PageLogin__header--description"
          >Enter your credentials to acces your account</span
        >
      </header>
      <BetFormItem label="email ">
        <BetInput v-model="formRegisterAccessibility.email"></BetInput>
      </BetFormItem>
      <BetFormItem label="login ">
        <BetInput v-model="formRegisterAccessibility.login"></BetInput>
      </BetFormItem>
      <BetFormItem label="password">
        <BetInput v-model="formRegisterAccessibility.password"></BetInput>
      </BetFormItem>
      <BetFormItem label="password x2">
        <BetInput v-model="formRegisterAccessibility.password"></BetInput>
      </BetFormItem>
      <BetFormItem>
        <BetButton
          type="primary"
          class="PageLogin__loginButton"
          @click="register"
        >
          Register
        </BetButton>
      </BetFormItem>
    </BetForm>
  </section>
</template>
<script lang="ts" setup>
import { reactive } from "vue";

const formRegisterAccessibility = reactive({
  email: "",
  login: "",
  password: "",
});

const register = () => {};
</script>
<style lang="scss"></style>
