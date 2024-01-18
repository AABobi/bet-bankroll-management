import axios from "axios";
import { defineStore } from "pinia";
//import { ref } from "vue";
//import { useRouter } from "vue-router";

export const useAuthenticationStore = defineStore("authentication", () => {
  const login = async (): Promise<boolean> => {
    //const token = localStorage.getItem("token");
    //console.log(userData);
    const response = await axios.post("http://backend-go2-service:80/login", {
      email: "test3",
      login: "test3",
      password: "test3",
    });
    /*.then((res) => {
        console.log(res.data.token);
      })
      .catch((err) => {
        console.log(err.response);
      });*/

    console.log(response.data.token);
    if (response.data.token != undefined) {
      return true;
    } else {
      return false;
    }
  };
  return {
    login,
  };
});
