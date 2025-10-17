import React from "react"
import { shallow } from "enzyme"
import { Bucket } from "../Bucket"

describe("Bucket", () => {
  it("should render without crashing", () => {
    shallow(<Bucket />)
  })

  it("should call selectBucket when clicked", () => {
    const selectBucket = jest.fn()
    const wrapper = shallow(
      <Bucket bucket={"test"} selectBucket={selectBucket} />
    )
    wrapper.find("li").simulate("click", { preventDefault: jest.fn() })
    expect(selectBucket).toHaveBeenCalledWith("test")
  })

  it("should highlight the selected bucket", () => {
    const wrapper = shallow(<Bucket bucket={"test"} isActive={true} />)
    expect(wrapper.find("li").hasClass("active")).toBeTruthy()
  })
})
