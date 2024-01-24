<template>
  <section class="PageLogin">
    <BetForm class="PageLogin__loginBox" label-position="top">
      <header class="PageLogin__header">
        <p>Welcome</p>
        <span class="PageLogin__header--description"
          >Enter your credentials to acces your account</span
        >
      </header>
      <BetFormItem label="email ">
        <BetInput v-model="formLoginAccessibility.email"></BetInput>
      </BetFormItem>
      <BetFormItem label="password">
        <BetInput v-model="formLoginAccessibility.password"></BetInput>
      </BetFormItem>
      <BetFormItem>
        <BetButton type="primary" class="PageLogin__loginButton" @click="login">
          Login
        </BetButton>
      </BetFormItem>
    </BetForm>
  </section>
</template>

<script lang="ts" setup>
import {
  BetForm,
  BetFormItem,
  BetInput,
  BetButton,
} from "@/core/element-plus/index";
import { reactive } from "vue";
import { useRouter } from "vue-router";
import { useAuthenticationStore } from "@/store/authentication-store";

/*const statusTypesName: { [key in StatusTypeEnum]: string } = {
  [StatusTypeEnum.CONTAINER]: "Container",
  [StatusTypeEnum.SHIPMENT]: "Shipment",
};

const statusString = statusTypesName[store.booking.status.type];*/

const formLoginAccessibility = reactive({
  email: "",
  password: "",
});

const router = useRouter();
const store = useAuthenticationStore();
const login = async () => {
  const t = await store.login();
  if (t) {
    router.push("/add");
  }
};
</script>

<style lang="scss">
.PageLogin {
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  background-position: center;
  background-size: cover;
  height: 100%;
  width: 100%;
  background-color: #e7eaf6;

  &__header {
    p {
      margin-bottom: 5px;
    }

    text-align: center;
    font-family: "Open Sans", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    font-variation-settings: "wdth" 100;
    margin-bottom: 10px;

    &--description {
      font-size: 10px;
      color: #a2a8d3;
    }
  }

  &__loginBox {
    background: white;
    width: 300px;
    height: 300px;
    padding: 0 30px 0 30px;
    border-radius: 5px;
  }

  &__loginButton {
    margin-top: 20px;
    width: 100%;
  }
}
</style>
