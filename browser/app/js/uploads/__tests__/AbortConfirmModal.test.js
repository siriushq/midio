import React from "react"
import { shallow } from "enzyme"
import { AbortConfirmModal } from "../AbortConfirmModal"

describe("AbortConfirmModal", () => {
  it("should render without crashing", () => {
    shallow(<AbortConfirmModal />)
  })

  it("should call abort for every upload when Abort is clicked", () => {
    const abort = jest.fn()
    const wrapper = shallow(
      <AbortConfirmModal
        uploads={{
          "a-b/-test1": { size: 100, loaded: 50, name: "test1" },
          "a-b/-test2": { size: 100, loaded: 50, name: "test2" }
        }}
        abort={abort}
      />
    )
    wrapper.instance().abortUploads()
    expect(abort.mock.calls.length).toBe(2)
    expect(abort.mock.calls[0][0]).toBe("a-b/-test1")
    expect(abort.mock.calls[1][0]).toBe("a-b/-test2")
  })

  it("should call hideAbort when cancel is clicked", () => {
    const hideAbort = jest.fn()
    const wrapper = shallow(<AbortConfirmModal hideAbort={hideAbort} />)
    wrapper.find("ConfirmModal").prop("cancelHandler")()
    expect(hideAbort).toHaveBeenCalled()
  })
})
