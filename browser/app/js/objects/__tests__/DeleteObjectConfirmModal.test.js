import React from "react"
import { shallow } from "enzyme"
import { DeleteObjectConfirmModal } from "../DeleteObjectConfirmModal"

describe("DeleteObjectConfirmModal", () => {
  it("should render without crashing", () => {
    shallow(<DeleteObjectConfirmModal />)
  })

  it("should call deleteObject when Delete is clicked", () => {
    const deleteObject = jest.fn()
    const wrapper = shallow(
      <DeleteObjectConfirmModal deleteObject={deleteObject} />
    )
    wrapper.find("ConfirmModal").prop("okHandler")()
    expect(deleteObject).toHaveBeenCalled()
  })

  it("should call hideDeleteConfirmModal when Cancel is clicked", () => {
    const hideDeleteConfirmModal = jest.fn()
    const wrapper = shallow(
      <DeleteObjectConfirmModal
        hideDeleteConfirmModal={hideDeleteConfirmModal}
      />
    )
    wrapper.find("ConfirmModal").prop("cancelHandler")()
    expect(hideDeleteConfirmModal).toHaveBeenCalled()
  })
})
