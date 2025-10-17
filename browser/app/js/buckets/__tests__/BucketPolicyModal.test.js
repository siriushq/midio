import React from "react"
import { shallow, mount } from "enzyme"
import { BucketPolicyModal } from "../BucketPolicyModal"
import { READ_ONLY, WRITE_ONLY, READ_WRITE } from "../../constants"

describe("BucketPolicyModal", () => {
  it("should render without crashing", () => {
    shallow(<BucketPolicyModal policies={[]}/>)
  })

  it("should call hideBucketPolicy when close button is clicked", () => {
    const hideBucketPolicy = jest.fn()
    const wrapper = shallow(
      <BucketPolicyModal hideBucketPolicy={hideBucketPolicy} policies={[]} />
    )
    wrapper.find("button").simulate("click")
    expect(hideBucketPolicy).toHaveBeenCalled()
  })

  it("should include the PolicyInput and Policy components when there are any policies", () => {
    const wrapper = shallow(
      <BucketPolicyModal policies={ [{prefix: "test", policy: READ_ONLY}] } />
    )
    expect(wrapper.find("Connect(PolicyInput)").length).toBe(1)
    expect(wrapper.find("Connect(Policy)").length).toBe(1)
  })
})
