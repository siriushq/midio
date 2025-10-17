import React from "react"
import { shallow } from "enzyme"
import { ObjectsSearch } from "../ObjectsSearch"

describe("ObjectsSearch", () => {
  it("should render without crashing", () => {
    shallow(<ObjectsSearch />)
  })

  it("should call onChange with search text", () => {
    const onChange = jest.fn()
    const wrapper = shallow(<ObjectsSearch onChange={onChange} />)
    wrapper.find("input").simulate("change", { target: { value: "test" } })
    expect(onChange).toHaveBeenCalledWith("test")
  })
})
