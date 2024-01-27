import HelloWorld from "./HelloWorld.vue";
import { mount } from "@vue/test-utils";
import { useAuthenticationStore } from "@/store/authentication-store";
import { setActivePinia, createPinia } from "pinia";
import { createTestingPinia } from "@pinia/testing";
import { useTestStore } from "@/store/test-store";

describe("HelloWorld", () => {
  beforeEach(() => {
    // creates a fresh pinia and makes it active
    // so it's automatically picked up by any useStore() call
    // without having to pass it to it: `useStore(pinia)`
    const pinia = createTestingPinia();
    const queryStore = useAuthenticationStore(pinia);
    //setActivePinia(createPinia());
  });
  it("renders properly", () => {
    /* const wrapper = mount(HelloWorld, {
      global: {
        plugins: [createTestingPinia()],
      },
    });*/
    const wrapper = mount(HelloWorld);
    expect(wrapper.text()).toContain("TEST");
  });
});
