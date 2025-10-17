import React from "react"
import { shallow, mount } from "enzyme"
import { AlertContainer } from "../AlertContainer"

describe("Alert", () => {
  it("should render without crashing", () => {
    shallow(
      <AlertContainer alert={{ show: true, type: "danger", message: "Test" }} />
    )
  })

  it("should render nothing if message is empty", () => {
    const wrapper = shallow(
      <AlertContainer alert={{ show: true, type: "danger", message: "" }} />
    )
    expect(wrapper.find("Alert").length).toBe(0)
  })
})
