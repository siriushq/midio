import React from "react"
import { shallow } from "enzyme"
import { MobileHeader } from "../MobileHeader"

describe("Bucket", () => {
  it("should render without crashing", () => {
    shallow(<MobileHeader sidebarOpen={false} />)
  })

  it("should toggleSidebar when trigger is clicked", () => {
    const toggleSidebar = jest.fn()
    const wrapper = shallow(
      <MobileHeader sidebarOpen={false} toggleSidebar={toggleSidebar} />
    )
    wrapper
      .find("#sidebar-toggle")
      .simulate("click", { stopPropagation: jest.fn() })
    expect(toggleSidebar).toHaveBeenCalled()
  })
})
