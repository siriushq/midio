import React from "react"
import { shallow, mount } from "enzyme"
import Alert from "../Alert"

describe("Alert", () => {
  it("should render without crashing", () => {
    shallow(<Alert />)
  })

  it("should call onDismiss when close button is clicked", () => {
    const onDismiss = jest.fn()
    const wrapper = mount(
      <Alert show={true} type="danger" message="test" onDismiss={onDismiss} />
    )
    wrapper.find("button").simulate("click", { preventDefault: jest.fn() })
    expect(onDismiss).toHaveBeenCalled()
  })
})
