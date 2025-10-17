import reducer from "../reducer"
import * as actions from "../actions"

describe("uploads reducer", () => {
  it("should return the initial state", () => {
    const initialState = reducer(undefined, {})
    expect(initialState).toEqual({
      files: {},
      showAbortModal: false
    })
  })

  it("should handle ADD", () => {
    const newState = reducer(undefined, {
      type: actions.ADD,
      slug: "a-b-c",
      size: 100,
      name: "test"
    })
    expect(newState.files).toEqual({
      "a-b-c": { loaded: 0, size: 100, name: "test" }
    })
  })

  it("should handle UPDATE_PROGRESS", () => {
    const newState = reducer(
      {
        files: { "a-b-c": { loaded: 0, size: 100, name: "test" } }
      },
      {
        type: actions.UPDATE_PROGRESS,
        slug: "a-b-c",
        loaded: 50
      }
    )
    expect(newState.files).toEqual({
      "a-b-c": { loaded: 50, size: 100, name: "test" }
    })
  })

  it("should handle STOP", () => {
    const newState = reducer(
      {
        files: {
          "a-b-c": { loaded: 70, size: 100, name: "test1" },
          "x-y-z": { loaded: 50, size: 100, name: "test2" }
        }
      },
      {
        type: actions.STOP,
        slug: "a-b-c"
      }
    )
    expect(newState.files).toEqual({
      "x-y-z": { loaded: 50, size: 100, name: "test2" }
    })
  })

  it("should handle SHOW_ABORT_MODAL", () => {
    const newState = reducer(
      {
        showAbortModal: false
      },
      {
        type: actions.SHOW_ABORT_MODAL,
        show: true
      }
    )
    expect(newState.showAbortModal).toBeTruthy()
  })
})
