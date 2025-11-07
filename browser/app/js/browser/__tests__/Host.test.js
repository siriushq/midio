// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
import React from "react"
import { shallow } from "enzyme"
import Host from "../Host"

describe("Host", () => {
  it("should render without crashing", () => {
    shallow(<Host />)
  })
})
