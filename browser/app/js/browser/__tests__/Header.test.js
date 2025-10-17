import React from "react"
import { shallow } from "enzyme"
import Header from "../Header"

jest.mock("../../web", () => ({
  LoggedIn: jest
    .fn(() => true)
    .mockReturnValueOnce(true)
    .mockReturnValueOnce(false)
}))
describe("Header", () => {
  it("should render without crashing", () => {
    shallow(<Header />)
  })

  it("should render Login button when the user has not LoggedIn", () => {
    const wrapper = shallow(<Header />)
    expect(wrapper.find("a").text()).toBe("Login")
  })

  it("should render StorageInfo and BrowserDropdown when the user has LoggedIn", () => {
    const wrapper = shallow(<Header />)
    expect(wrapper.find("Connect(BrowserDropdown)").length).toBe(1)
    expect(wrapper.find("Connect(StorageInfo)").length).toBe(1)
  })
})
