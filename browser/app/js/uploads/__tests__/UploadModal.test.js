import React from "react"
import { shallow } from "enzyme"
import { UploadModal } from "../UploadModal"

describe("UploadModal", () => {
  it("should render without crashing", () => {
    shallow(<UploadModal uploads={{}} />)
  })

  it("should render AbortConfirmModal when showAbort is true", () => {
    const wrapper = shallow(<UploadModal uploads={{}} showAbort={true} />)
    expect(wrapper.find("Connect(AbortConfirmModal)").length).toBe(1)
  })

  it("should render nothing when there are no files being uploaded", () => {
    const wrapper = shallow(<UploadModal uploads={{}} />)
    expect(wrapper.find("noscript").length).toBe(1)
  })

  it("should show upload progress when one or more files are being uploaded", () => {
    const wrapper = shallow(
      <UploadModal
        uploads={{ "a-b/-test": { size: 100, loaded: 50, name: "test" } }}
      />
    )
    expect(wrapper.find("ProgressBar").length).toBe(1)
  })

  it("should call showAbortModal when close button is clicked", () => {
    const showAbortModal = jest.fn()
    const wrapper = shallow(
      <UploadModal
        uploads={{ "a-b/-test": { size: 100, loaded: 50, name: "test" } }}
        showAbortModal={showAbortModal}
      />
    )
    wrapper.find("button").simulate("click")
    expect(showAbortModal).toHaveBeenCalled()
  })
})
