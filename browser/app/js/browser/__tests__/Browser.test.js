import React from "react"
import { shallow } from "enzyme"
import Browser from "../Browser"
import configureStore from "redux-mock-store"

const mockStore = configureStore()

describe("Browser", () => {
  it("should render without crashing", () => {
    const store = mockStore()
    shallow(<Browser store={store}/>)
  })
})
