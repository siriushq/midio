import React from "react"
import { shallow } from "enzyme"
import BucketContainer from "../BucketContainer"
import configureStore from "redux-mock-store"

const mockStore = configureStore()

describe("BucketContainer", () => {
  let store
  beforeEach(() => {
    store = mockStore({
      buckets: {
        currentBucket: "Test"
      }
    })
    store.dispatch = jest.fn()
  })

  it("should render without crashing", () => {
    shallow(<BucketContainer store={store}/>)
  })

  it('maps state and dispatch to props', () => {
    const wrapper = shallow(<BucketContainer store={store}/>)
    expect(wrapper.props()).toEqual(expect.objectContaining({
      isActive: expect.any(Boolean),
      selectBucket: expect.any(Function)
    }))
  })

  it('maps selectBucket to dispatch action', () => {
    const wrapper = shallow(<BucketContainer store={store}/>)
    wrapper.props().selectBucket()
    expect(store.dispatch).toHaveBeenCalled()
  })
})
