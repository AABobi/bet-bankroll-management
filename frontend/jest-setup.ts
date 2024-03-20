import { config } from "@vue/test-utils";
import { createTestingPinia } from "@pinia/testing";
import ElementPlus from "element-plus";

config.global = {
  ...config.global,
  plugins: [createTestingPinia({ stubActions: false }), ElementPlus],
};

if (!window.global.activeFeatureTogglz) {
  window.global.activeFeatureTogglz = [];
}
