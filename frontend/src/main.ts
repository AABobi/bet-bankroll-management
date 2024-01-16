import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { createPinia } from "pinia";
import { createRouter, createWebHistory } from "vue-router";
import { routes } from "./routes/router";

function main() {
  const router = createRouter({
    history: createWebHistory(),
    routes: routes,
  });
  const pinia = createPinia();
  const app = createApp(App);
  app.use(pinia);
  app.use(router);
  app.use(ElementPlus);
  app.mount("#app");
}
main();
