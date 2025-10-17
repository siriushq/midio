import React from "react"
import { shallow } from "enzyme"
import { PrefixActions } from "../PrefixActions"

describe("PrefixActions", () => {
  it("should render without crashing", () => {
    shallow(<PrefixActions object={{ name: "abc/" }} currentPrefix={"pre1/"} />)
  })

  it("should show DeleteObjectConfirmModal when delete action is clicked", () => {
    const wrapper = shallow(
      <PrefixActions object={{ name: "abc/" }} currentPrefix={"pre1/"} />
    )
    wrapper
      .find("a")
      .last()
      .simulate("click", { preventDefault: jest.fn() })
    expect(wrapper.state("showDeleteConfirmation")).toBeTruthy()
    expect(wrapper.find("DeleteObjectConfirmModal").length).toBe(1)
  })

  it("should hide DeleteObjectConfirmModal when Cancel button is clicked", () => {
    const wrapper = shallow(
      <PrefixActions object={{ name: "abc/" }} currentPrefix={"pre1/"} />
    )
    wrapper
      .find("a")
      .last()
      .simulate("click", { preventDefault: jest.fn() })
    wrapper.find("DeleteObjectConfirmModal").prop("hideDeleteConfirmModal")()
    wrapper.update()
    expect(wrapper.state("showDeleteConfirmation")).toBeFalsy()
    expect(wrapper.find("DeleteObjectConfirmModal").length).toBe(0)
  })

  it("should call deleteObject with object name", () => {
    const deleteObject = jest.fn()
    const wrapper = shallow(
      <PrefixActions
        object={{ name: "abc/" }}
        currentPrefix={"pre1/"}
        deleteObject={deleteObject}
      />
    )
    wrapper
      .find("a")
      .last()
      .simulate("click", { preventDefault: jest.fn() })
    wrapper.find("DeleteObjectConfirmModal").prop("deleteObject")()
    expect(deleteObject).toHaveBeenCalledWith("abc/")
  })


  it("should call downloadPrefix when single object is selected and download button is clicked", () => {
    const downloadPrefix = jest.fn()
    const wrapper = shallow(
      <PrefixActions
        object={{ name: "abc/" }}
        currentPrefix={"pre1/"}
        downloadPrefix={downloadPrefix} />
    )
    wrapper
      .find("a")
      .first()
      .simulate("click", { preventDefault: jest.fn() })
    expect(downloadPrefix).toHaveBeenCalled()
  })
})
