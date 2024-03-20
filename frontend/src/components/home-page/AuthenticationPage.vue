<template>
  <section class="AuthenticationPage">
    <BetDialog
      :show-dialog="showDialog"
      :message="dialogMessage"
      :show-confirm="true"
      @confirm="dialogConfirmation"
    />
    <BetRadioGroup v-model="radio" size="large" @click="clearForm">
      <BetRadioButton :label="ApiType.LOGIN" />
      <BetRadioButton :label="ApiType.REGISTER" />
    </BetRadioGroup>

    <BetForm :class="classes" label-position="top">
      <header class="AuthenticationPage__header">
        <p>Welcome</p>
        <span class="AuthenticationPage__header--description">{{
          headerDescription
        }}</span>
      </header>
      <BetFormItem :error="validationError.getError('email')" label="email ">
        <BetInput
          v-model="formLoginAccessibility.email"
          @input="validationError.removeError('email')"
          data-test="auth-page-email"
        />
      </BetFormItem>
      <BetFormItem
        :error="validationError.getError('password')"
        label="password"
      >
        <BetInput
          v-model="formLoginAccessibility.password"
          type="password"
          @input="validationError.removeError('password')"
          data-test="auth-page-password"
        />
      </BetFormItem>
      <BetFormItem
        :error="validationError.getError('repeatedPassword')"
        v-if="isRegister"
        label="password x2"
      >
        <BetInput
          v-model="formLoginAccessibility.repeatedPassword"
          type="password"
          @input="validationError.removeError('repeatedPassword')"
        />
      </BetFormItem>
      <BetFormItem>
        <BetButton
          type="primary"
          class="AuthenticationPage__menuButton"
          @click="request"
        >
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
  BetRadioGroup,
  BetRadioButton,
  BetDialog,
} from "@/core/element-plus/index";
import { computed, onMounted, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthenticationStore } from "@/store/authentication-store";
import { ValidationError } from "../domain/validation/validation-error";
import { LoginType } from "../domain/login-type";
import { RegisterType } from "../domain/register-type";

/*const statusTypesName: { [key in StatusTypeEnum]: string } = {
    [StatusTypeEnum.CONTAINER]: "Container",
    [StatusTypeEnum.SHIPMENT]: "Shipment",
  };

  const statusString = statusTypesName[store.booking.status.type];*/
enum ApiType {
  REGISTER = "register",
  LOGIN = "login",
}

enum MenuFields {
  email = "email",
  password = "password",
  repeatedPassword = "repeatedPassword",
}

onMounted(() => {
  store.test1();
});

const clearForm = () => {
  for (const v in MenuFields) {
    console.log(v);
    formLoginAccessibility[v] = "";
    validationError.value.removeError(v);
  }
};

const formLoginAccessibility = reactive({
  email: "",
  password: "",
  repeatedPassword: "",
});

const router = useRouter();
const store = useAuthenticationStore();

const radio = ref<ApiType>(ApiType.LOGIN);
const showDialog = ref(false);

const headerDescription = ref("Enter your credentials to acces your account");
const dialogMessage = ref("User is created successfully!");

const isAuthenticated = ref<boolean>(false);
const validationError = ref(new ValidationError());

const dialogConfirmation = () => {
  showDialog.value = false;
};

const request = async () => {
  console.log("REQUEST");
  const email = formLoginAccessibility.email;
  const password = formLoginAccessibility.password;
  const loginType: LoginType = {
    email: email,
    password: password,
  };

  if (!isRegister.value) {
    console.log("ifLogin");
    console.log(email);
    console.log(password);
    const wasBlank = validationError.value.reportBlank(loginType);
    if (wasBlank) {
      return;
    }

    const errorMessage = "Email or password is incorrect";
    await callBackend(ApiType.LOGIN, loginType, errorMessage);
    if (isAuthenticated.value) {
      router.push("/add");
    }
  } else {
    console.log("ifRegister");
    const register: RegisterType = {
      email: email,
      password: password,
      repeatedPassword: formLoginAccessibility.repeatedPassword,
    };

    const wasBlank = validationError.value.reportBlank(register);
    if (wasBlank) {
      return;
    }

    if (!email.includes("@") || !email.includes(".")) {
      const emailMessage = "Email have to contain '@' and '.'";
      validationError.value.addError("email", emailMessage);

      if (password.length < 3) {
        const passwordMessage = "Password have to be longer that 3 chars";
        const repeatedPasswordMessage =
          "Password have to be longer that 3 chars";

        validationError.value.addError("password", passwordMessage);
        validationError.value.addError(
          "repeatedPassword",
          repeatedPasswordMessage
        );
      }

      return;
    }

    if (password.length < 3) {
      const passwordMessage = "Password have to be longer that 3 chars";
      const repeatedPasswordMessage = "Password have to be longer that 3 chars";

      validationError.value.addError("password", passwordMessage);
      validationError.value.addError(
        "repeatedPassword",
        repeatedPasswordMessage
      );
      return;
    }

    if (!comparePassword.value) {
      validationError.value.addError(
        "password",
        "The passwords must be identical"
      );
      validationError.value.addError(
        "repeatedPassword",
        "The passwords must be identical"
      );
      return;
    }

    const message = "Cannot create user";
    await callBackend(ApiType.REGISTER, register, message);
    if (isAuthenticated.value) {
      console.log("TUTAJ");
      showDialog.value = true;
    }
  }
};

const callBackend = async (
  api: ApiType,
  data: LoginType | RegisterType,
  message: string
) => {
  try {
    isAuthenticated.value = await store[api](data.email, data.password);
    console.log("isAuthenticated.value", isAuthenticated.value);
  } catch (e) {
    for (const key in data) {
      validationError.value.addError(key, message);
    }
  }
};

const classes = computed(() => ({
  AuthenticationPage__container: true,
  "AuthenticationPage__container--register": radio.value === ApiType.REGISTER,
}));

const isRegister = computed(() => radio.value === ApiType.REGISTER);

const comparePassword = computed(
  () =>
    formLoginAccessibility.password === formLoginAccessibility.repeatedPassword
);
</script>
<style lang="scss">
.AuthenticationPage {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: absolute;
  background-position: center;
  background-size: cover;
  height: 100%;
  width: 100%;
  background-color: $colorBackBright;

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
      color: $colorDesBright;
    }
  }

  &__container {
    background: $colorWhite;
    width: 300px;
    height: 300px;
    padding: 0 30px 0 30px;
    border-radius: 5px;

    &--register {
      height: 360px;
    }
  }

  &__menuButton {
    margin-top: 20px;
    width: 100%;
  }
}
</style>
