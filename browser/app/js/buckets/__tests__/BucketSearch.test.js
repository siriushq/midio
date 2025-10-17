import React from "react"
import { shallow } from "enzyme"
import { BucketSearch } from "../BucketSearch"

describe("BucketSearch", () => {
  it("should render without crashing", () => {
    shallow(<BucketSearch />)
  })

  it("should call onChange with search text", () => {
    const onChange = jest.fn()
    const wrapper = shallow(<BucketSearch onChange={onChange} />)
    wrapper.find("input").simulate("change", { target: { value: "test" } })
    expect(onChange).toHaveBeenCalledWith("test")
  })
})
