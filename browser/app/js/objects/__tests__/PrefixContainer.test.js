import React from "react"
import { shallow } from "enzyme"
import { PrefixContainer } from "../PrefixContainer"

describe("PrefixContainer", () => {
  it("should render without crashing", () => {
    shallow(<PrefixContainer object={{ name: "abc/" }} />)
  })

  it("should render ObjectItem with props", () => {
    const wrapper = shallow(<PrefixContainer object={{ name: "abc/" }} />)
    expect(wrapper.find("Connect(ObjectItem)").length).toBe(1)
    expect(wrapper.find("Connect(ObjectItem)").prop("name")).toBe("abc/")
  })

  it("should call selectPrefix when the prefix is clicked", () => {
    const selectPrefix = jest.fn()
    const wrapper = shallow(
      <PrefixContainer
        object={{ name: "abc/" }}
        currentPrefix={"xyz/"}
        selectPrefix={selectPrefix}
      />
    )
    wrapper.find("Connect(ObjectItem)").prop("onClick")()
    expect(selectPrefix).toHaveBeenCalledWith("xyz/abc/")
  })

  it("should pass actions to ObjectItem", () => {
    const wrapper = shallow(
      <PrefixContainer object={{ name: "abc/" }} checkedObjectsCount={0} />
    )
    expect(wrapper.find("Connect(ObjectItem)").prop("actionButtons")).not.toBe(
      undefined
    )
  })

  it("should pass empty actions to ObjectItem when checkedObjectCount is more than 0", () => {
    const wrapper = shallow(
      <PrefixContainer object={{ name: "abc/" }} checkedObjectsCount={1} />
    )
    expect(wrapper.find("Connect(ObjectItem)").prop("actionButtons")).toBe(
      undefined
    )
  })
})
