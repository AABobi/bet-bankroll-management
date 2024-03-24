//import { createTestingPinia } from "@pinia/testing";
import AuthenticationPage from "@/components/home-page/AuthenticationPage.vue";
import { shallowMount } from "@vue/test-utils";

describe("PageOrderBookingHeader", () => {
  beforeAll(() => {
    // const pinia = createTestingPinia();
  });
  it("Should initialize", () => {
    const wrapper = shallowMount(AuthenticationPage);
    expect(wrapper.exists()).toBe(true);
  });
});
