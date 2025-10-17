import React from "react"
import { shallow } from "enzyme"
import { ObjectsSection } from "../ObjectsSection"

describe("ObjectsSection", () => {
  it("should render without crashing", () => {
    shallow(<ObjectsSection />)
  })
})
