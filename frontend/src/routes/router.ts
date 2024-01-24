import AuthenticationPage from "../components/home-page/AuthenticationPage.vue";
import BetManagement from "../components/BetManagement.vue";
import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "AuthenticationPage",
    component: AuthenticationPage,
  },
  { path: "/add", name: "BetManagement", component: BetManagement },
];

export { routes };
