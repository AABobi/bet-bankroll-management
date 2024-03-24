import HelloWorld from "./HelloWorld.vue";
import { mount } from "@vue/test-utils"; //flushPromises

describe("HelloWorld", () => {
  it("renders properly", async () => {
    const wrapper = mount(HelloWorld);
    const testCall = jest.spyOn(wrapper.vm, "testFun");

    await wrapper.find("button").trigger("click");

    expect(testCall).toHaveBeenCalled();
  });
});
