import React from "react"
import { shallow } from "enzyme"
import { AboutModal } from "../AboutModal"

describe("AboutModal", () => {
  const serverInfo = {
    version: "test",
    platform: "test",
    runtime: "test"
  }

  it("should render without crashing", () => {
    shallow(<AboutModal serverInfo={serverInfo} />)
  })

  it("should call hideAbout when close button is clicked", () => {
    const hideAbout = jest.fn()
    const wrapper = shallow(
      <AboutModal serverInfo={serverInfo} hideAbout={hideAbout} />
    )
    wrapper.find("button").simulate("click")
    expect(hideAbout).toHaveBeenCalled()
  })
})
