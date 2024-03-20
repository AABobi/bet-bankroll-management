import AuthenticationPage from "../components/home-page/AuthenticationPage.vue";
import BetManagement from "../components/BetManagement.vue";
import BetManagementPage from "../components/bet-page/BetManagementPage.vue";

import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "AuthenticationPage",
    component: AuthenticationPage,
  },
  { path: "/add", name: "BetManagement", component: BetManagement },
  {
    path: "/man",
    name: "BetManagementPage",
    component: BetManagementPage,
  },
];

export { routes };
