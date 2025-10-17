import React from "react"
import { shallow } from "enzyme"
import { SideBar } from "../SideBar"

jest.mock("../../web", () => ({
  LoggedIn: jest.fn(() => false).mockReturnValueOnce(true)
}))

describe("SideBar", () => {
  it("should render without crashing", () => {
    shallow(<SideBar />)
  })

  it("should not render BucketSearch for non LoggedIn users", () => {
    const wrapper = shallow(<SideBar />)
    expect(wrapper.find("Connect(BucketSearch)").length).toBe(0)
  })

  it("should call clickOutside when the user clicks outside the sidebar", () => {
    const clickOutside = jest.fn()
    const wrapper = shallow(<SideBar clickOutside={clickOutside} />)
    wrapper.simulate("clickOut", {
      preventDefault: jest.fn(),
      target: { classList: { contains: jest.fn(() => false) } }
    })
    expect(clickOutside).toHaveBeenCalled()
  })

  it("should not call clickOutside when user clicks on sidebar toggle", () => {
    const clickOutside = jest.fn()
    const wrapper = shallow(<SideBar clickOutside={clickOutside} />)
    wrapper.simulate("clickOut", {
      preventDefault: jest.fn(),
      target: { classList: { contains: jest.fn(() => true) } }
    })
    expect(clickOutside).not.toHaveBeenCalled()
  })
})
