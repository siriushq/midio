import React from "react"
import { shallow } from "enzyme"
import { ObjectItem } from "../ObjectItem"

describe("ObjectItem", () => {
  it("should render without crashing", () => {
    shallow(<ObjectItem name={"test"} />)
  })

  it("should render with content type", () => {
    const wrapper = shallow(<ObjectItem name={"test.jpg"} contentType={""} />)
    expect(wrapper.prop("data-type")).toBe("image")
  })

  it("shouldn't call onClick when the object isclicked", () => {
    const onClick = jest.fn()
    const checkObject = jest.fn()
    const wrapper = shallow(
      <ObjectItem name={"test"} checkObject={checkObject} />
    )
    wrapper.find("a").simulate("click", { preventDefault: jest.fn() })
    expect(onClick).not.toHaveBeenCalled()
  })

  it("should call onClick when the folder isclicked", () => {
    const onClick = jest.fn()
    const wrapper = shallow(<ObjectItem name={"test/"} onClick={onClick} />)
    wrapper.find("a").simulate("click", { preventDefault: jest.fn() })
    expect(onClick).toHaveBeenCalled()
  })

  it("should call checkObject when the object/prefix is checked", () => {
    const checkObject = jest.fn()
    const wrapper = shallow(
      <ObjectItem name={"test"} checked={false} checkObject={checkObject} />
    )
    wrapper.find("input[type='checkbox']").simulate("change")
    expect(checkObject).toHaveBeenCalledWith("test")
  })

  it("should render checked checkbox", () => {
    const wrapper = shallow(<ObjectItem name={"test"} checked={true} />)
    expect(wrapper.find("input[type='checkbox']").prop("checked")).toBeTruthy()
  })

  it("should call uncheckObject when the object/prefix is unchecked", () => {
    const checkObject = jest.fn()
    const uncheckObject = jest.fn()
    const wrapper = shallow(
      <ObjectItem
        name={"test"}
        checked={true}
        checkObject={checkObject}
        uncheckObject={uncheckObject}
      />
    )
    wrapper.find("input[type='checkbox']").simulate("change")
    expect(uncheckObject).toHaveBeenCalledWith("test")
  })
})
