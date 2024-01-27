import { defineStore } from "pinia";
import { ref } from "vue";
//import { ref } from "vue";
//import { useRouter } from "vue-router";

export const useTestStore = defineStore("test", () => {
  const t = ref(1);
  const test1 = (email: string) => {};
  return {
    t,
    test1,
  };
});
