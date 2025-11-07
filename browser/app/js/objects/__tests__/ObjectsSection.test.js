// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
import React from "react"
import { shallow } from "enzyme"
import { ObjectsSection } from "../ObjectsSection"

describe("ObjectsSection", () => {
  it("should render without crashing", () => {
    shallow(<ObjectsSection />)
  })
})
