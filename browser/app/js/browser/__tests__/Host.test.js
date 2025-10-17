import React from "react"
import { shallow } from "enzyme"
import Host from "../Host"

describe("Host", () => {
  it("should render without crashing", () => {
    shallow(<Host />)
  })
})
