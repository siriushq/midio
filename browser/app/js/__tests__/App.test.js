import React from "react"
import { shallow, mount } from "enzyme"
import { MemoryRouter } from "react-router-dom"
import App from "../App"

jest.mock("../browser/Login", () => () => <div>Login</div>)
jest.mock("../browser/Browser", () => () => <div>Browser</div>)

describe("App", () => {
  it("should render without crashing", () => {
    shallow(<App />)
  })

  it("should render Login component for '/login' route", () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={["/login"]}>
        <App />
      </MemoryRouter>
    )
    expect(wrapper.text()).toBe("Login")
  })

  it("should render Browser component for '/' route", () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={["/"]}>
        <App />
      </MemoryRouter>
    )
    expect(wrapper.text()).toBe("Browser")
  })

  it("should render Browser component for '/bucket' route", () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={["/bucket"]}>
        <App />
      </MemoryRouter>
    )
    expect(wrapper.text()).toBe("Browser")
  })

  it("should render Browser component for '/bucket/a/b/c' route", () => {
    const wrapper = mount(
      <MemoryRouter initialEntries={["/bucket/a/b/c"]}>
        <App />
      </MemoryRouter>
    )
    expect(wrapper.text()).toBe("Browser")
  })
})
