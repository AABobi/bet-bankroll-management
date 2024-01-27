import AuthenticationPage from "@/components/home-page/AuthenticationPage.vue";
import { useAuthenticationStore } from "@/store/authentication-store";
import * as testStore from "@/store/authentication-store";

import { createTestingPinia } from "@pinia/testing";
import { mount } from "@vue/test-utils";
import { createRouter, createWebHistory } from "vue-router";

jest.mock("@/store/authentication-store", () => ({
  useAuthenticationStore: jest.fn(() => ({
    // Mock your store methods here
    test1: jest.fn(() => {
      console.log("DUPA");
    }),
    login: jest.fn(() => {
      console.log("DUPA - loginJest");
    }),
    register: jest.fn(() => {
      console.log("DUPA - register");
    }),
    // ... other mocked methods
  })),
}));

describe("AuthenticationPage.vue", () => {
  const pinia = createTestingPinia();
  afterEach(() => {
    // const pinia = createTestingPinia();
    // queryStore = useAuthenticationStore(pinia);
  });

  it("renders properly", () => {
    const wrapper = mount(AuthenticationPage, {
      Check: true,
    });
    // console.log(wrapper.exists()).toBe(true);
    expect(wrapper.exists()).toBe(true);
  });

  it("renders properly2", async () => {
    const wrapper = mount(AuthenticationPage, {
      Check: true,
      global: {
        plugins: [pinia],
      },
    });

    const emailInput = wrapper.find("[data-test='auth-page-email']");
    await wrapper.vm.$nextTick();
    emailInput.setValue("test@gmail.com");

    const passwordInput = wrapper.find("[data-test='auth-page-password']");
    await wrapper.vm.$nextTick();
    passwordInput.setValue("qwerty");
    await wrapper.vm.$nextTick();

    const loginSpy = jest
      .spyOn(useAuthenticationStore, "login")
      .mockImplementation();

    //console.log(emailInput.isVisible());
    const confirmButton = wrapper
      .find(".AuthenticationPage__menuButton")
      .trigger("click");

    await wrapper.vm.$nextTick();
    // console.log(confirmButton.isVisible());
    //console.log(wrapper.html());
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick();

    expect(loginSpy).toHaveBeenCalled();
  });
});
