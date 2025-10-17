import React from "react"
import { shallow } from "enzyme"
import { ObjectContainer } from "../ObjectContainer"

describe("ObjectContainer", () => {
  it("should render without crashing", () => {
    shallow(<ObjectContainer object={{ name: "test1.jpg" }} />)
  })

  it("should render ObjectItem with props", () => {
    const wrapper = shallow(<ObjectContainer object={{ name: "test1.jpg" }} />)
    expect(wrapper.find("Connect(ObjectItem)").length).toBe(1)
    expect(wrapper.find("Connect(ObjectItem)").prop("name")).toBe("test1.jpg")
  })

  it("should pass actions to ObjectItem", () => {
    const wrapper = shallow(
      <ObjectContainer object={{ name: "test1.jpg" }} checkedObjectsCount={0} />
    )
    expect(wrapper.find("Connect(ObjectItem)").prop("actionButtons")).not.toBe(
      undefined
    )
  })

  it("should pass empty actions to ObjectItem when checkedObjectCount is more than 0", () => {
    const wrapper = shallow(
      <ObjectContainer object={{ name: "test1.jpg" }} checkedObjectsCount={1} />
    )
    expect(wrapper.find("Connect(ObjectItem)").prop("actionButtons")).toBe(
      undefined
    )
  })
})
