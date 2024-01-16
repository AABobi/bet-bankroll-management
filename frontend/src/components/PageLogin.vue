<template>
    <section class="PageLogin">
        <BetForm class="PageLogin__loginBox" label-position="top">
            <BetFormItem label="email " >
               <BetInput v-model="formLoginAccessibility.email"></BetInput>
            </BetFormItem>
            <BetFormItem label="login" >
               <BetInput v-model="formLoginAccessibility.login"></BetInput>
            </BetFormItem>
            <BetFormItem label="password" >
                <BetInput v-model="formLoginAccessibility.password"></BetInput>
            </BetFormItem>
            <BetFormItem>
                <BetButton type="primary" class="PageLogin__loginButton" @click="login">Login</BetButton>
            </BetFormItem>
        </BetForm>
    </section>
</template>
   
<script lang="ts" setup>
import { BetForm, BetFormItem, BetInput, BetButton } from "@/core/element-plus/index"
import { reactive } from 'vue';
import {useRouter} from "vue-router";
import  { useAuthenticationStore } from "@/store/authentication-store";
//import { useAuthenticationStore } from "src/store/authentication-store";

const formLoginAccessibility = reactive({
    email:'',
    login: '',
    password: '',
})

const router = useRouter()
const store = useAuthenticationStore()
const login = async () => {
    const t = await store.login()
    if(t) {
        router.push("/add")
    }
//router.push("/add")
}

</script>
   
<style lang="scss">
  .PageLogin {
       display: flex;
       justify-content: center;
       align-items: center;
       border: solid 5px  black;
   
     &__loginBox {
       background: white;
       width: 300px;
       height: 220px;
       padding: 50px;
       border-radius: 5px;
     }
   
     &__loginButton {
       margin-top: 20px;
     }
   }
</style>