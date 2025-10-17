import React from "react"
import { shallow, mount } from "enzyme"
import { MakeBucketModal } from "../MakeBucketModal"

describe("MakeBucketModal", () => {
  it("should render without crashing", () => {
    shallow(<MakeBucketModal />)
  })

  it("should call hideMakeBucketModal when close button is clicked", () => {
    const hideMakeBucketModal = jest.fn()
    const wrapper = shallow(
      <MakeBucketModal hideMakeBucketModal={hideMakeBucketModal} />
    )
    wrapper.find("button").simulate("click")
    expect(hideMakeBucketModal).toHaveBeenCalled()
  })

  it("bucketName should be cleared before hiding the modal", () => {
    const hideMakeBucketModal = jest.fn()
    const wrapper = shallow(
      <MakeBucketModal hideMakeBucketModal={hideMakeBucketModal} />
    )
    wrapper.find("input").simulate("change", {
      target: { value: "test" }
    })
    expect(wrapper.state("bucketName")).toBe("test")
    wrapper.find("button").simulate("click")
    expect(wrapper.state("bucketName")).toBe("")
  })

  it("should call makeBucket when the form is submitted", () => {
    const makeBucket = jest.fn()
    const hideMakeBucketModal = jest.fn()
    const wrapper = shallow(
      <MakeBucketModal
        makeBucket={makeBucket}
        hideMakeBucketModal={hideMakeBucketModal}
      />
    )
    wrapper.find("input").simulate("change", {
      target: { value: "test" }
    })
    wrapper.find("form").simulate("submit", { preventDefault: jest.fn() })
    expect(makeBucket).toHaveBeenCalledWith("test")
  })

  it("should call hideMakeBucketModal and clear bucketName after the form is submited", () => {
    const makeBucket = jest.fn()
    const hideMakeBucketModal = jest.fn()
    const wrapper = shallow(
      <MakeBucketModal
        makeBucket={makeBucket}
        hideMakeBucketModal={hideMakeBucketModal}
      />
    )
    wrapper.find("input").simulate("change", {
      target: { value: "test" }
    })
    wrapper.find("form").simulate("submit", { preventDefault: jest.fn() })
    expect(hideMakeBucketModal).toHaveBeenCalled()
    expect(wrapper.state("bucketName")).toBe("")
  })
})
