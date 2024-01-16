import PageLogin from "../components/PageLogin.vue";
import BetManagement from "../components/BetManagement.vue";
import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "PageLogin",
    component: PageLogin,
  },
  { path: "/add", name: "BetManagement", component: BetManagement },
];

export { routes };
