import React from "react"
import { shallow } from "enzyme"
import { ObjectsList } from "../ObjectsList"

describe("ObjectsList", () => {
  it("should render without crashing", () => {
    shallow(<ObjectsList objects={[]} />)
  })

  it("should render ObjectContainer for every object", () => {
    const wrapper = shallow(
      <ObjectsList objects={[{ name: "test1.jpg" }, { name: "test2.jpg" }]} />
    )
    expect(wrapper.find("Connect(ObjectContainer)").length).toBe(2)
  })

  it("should render PrefixContainer for every prefix", () => {
    const wrapper = shallow(
      <ObjectsList objects={[{ name: "abc/" }, { name: "xyz/" }]} />
    )
    expect(wrapper.find("Connect(PrefixContainer)").length).toBe(2)
  })
})
