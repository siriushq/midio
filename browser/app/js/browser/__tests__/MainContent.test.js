import React from "react"
import { shallow } from "enzyme"
import MainContent from "../MainContent"

describe("MainContent", () => {
  it("should render without crashing", () => {
    shallow(<MainContent />)
  })
})
