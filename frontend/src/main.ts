import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { createPinia } from "pinia";
import "@/core/css/global.scss";
import { createRouter, createWebHistory } from "vue-router";
import { routes } from "./routes/router";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

function main() {
  const router = createRouter({
    history: createWebHistory(),
    routes: routes,
  });
  const pinia = createPinia();
  const app = createApp(App);
  for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
  }
  app.use(pinia);
  app.use(router);
  app.use(ElementPlus);
  app.mount("#app");
}
main();
