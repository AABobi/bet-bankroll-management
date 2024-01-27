import axios from "axios";
import { defineStore } from "pinia";
import { ref } from "vue";
//import { ref } from "vue";
//import { useRouter } from "vue-router";

export const useAuthenticationStore = defineStore("authentication", () => {
  const login = async (email: string, password: string) => {
    console.log("DUPA - login");

    const response = await axios.post("http://localhost:8080/login", {
      email: email,
      password: password,
    });

    if (response.data.token != undefined) {
      localStorage.setItem("token", response.data.token);
      return true;
    }

    return false;
  };

  const register = async (
    email: string,
    password: string
  ): Promise<boolean> => {
    const resposne = await axios.post("http://localhost:8080/signup", {
      email: email,
      login: email,
      password: password,
    });

    // eslint-disable-next-line
    if (resposne.data.message === "User created") {
      return true;
    }
    return false;
  };

  const test1 = () => {
    console.log("ASDJHASJKFGBASLFGBASBGJASBGBASBJGAS");
  };

  return {
    login,
    register,
    test1,
  };
});
