import React from "react"
import { shallow } from "enzyme"
import { BrowserDropdown } from "../BrowserDropdown"

describe("BrowserDropdown", () => {
  const serverInfo = {
    version: "test",
    platform: "test",
    runtime: "test"
  }

  it("should render without crashing", () => {
    shallow(
      <BrowserDropdown serverInfo={serverInfo} fetchServerInfo={jest.fn()} />
    )
  })

  it("should call fetchServerInfo after its mounted", () => {
    const fetchServerInfo = jest.fn()
    const wrapper = shallow(
      <BrowserDropdown
        serverInfo={serverInfo}
        fetchServerInfo={fetchServerInfo}
      />
    )
    expect(fetchServerInfo).toHaveBeenCalled()
  })

  it("should show AboutModal when About link is clicked", () => {
    const wrapper = shallow(
      <BrowserDropdown serverInfo={serverInfo} fetchServerInfo={jest.fn()} />
    )
    wrapper.find("#show-about").simulate("click", { preventDefault: jest.fn() })
    wrapper.update()
    expect(wrapper.state("showAboutModal")).toBeTruthy()
    expect(wrapper.find("AboutModal").length).toBe(1)
  })

  it("should logout and redirect to /login when logout is clicked", () => {
    const wrapper = shallow(
      <BrowserDropdown serverInfo={serverInfo} fetchServerInfo={jest.fn()} />
    )
    wrapper.find("#logout").simulate("click", { preventDefault: jest.fn() })
    expect(window.location.pathname.endsWith("/login")).toBeTruthy()
  })
})
