import reducer from "../reducer"
import * as actionsAlert from "../actions"

describe("alert reducer", () => {
  it("should return the initial state", () => {
    expect(reducer(undefined, {})).toEqual({
      show: false,
      type: "danger"
    })
  })

  it("should handle SET_ALERT", () => {
    expect(
      reducer(undefined, {
        type: actionsAlert.SET,
        alert: { id: 1, type: "danger", message: "Test message" }
      })
    ).toEqual({
      show: true,
      id: 1,
      type: "danger",
      message: "Test message"
    })
  })

  it("should clear alert if id not passed", () => {
    expect(
      reducer(
        { show: true, type: "danger", message: "Test message" },
        {
          type: actionsAlert.CLEAR
        }
      )
    ).toEqual({
      show: false,
      type: "danger"
    })
  })

  it("should clear alert if id is matching", () => {
    expect(
      reducer(
        { show: true, id: 1, type: "danger", message: "Test message" },
        {
          type: actionsAlert.CLEAR,
          alert: { id: 1 }
        }
      )
    ).toEqual({
      show: false,
      type: "danger"
    })
  })

  it("should not clear alert if id is not matching", () => {
    expect(
      reducer(
        { show: true, id: 1, type: "danger", message: "Test message" },
        {
          type: actionsAlert.CLEAR,
          alert: { id: 2 }
        }
      )
    ).toEqual({
      show: true,
      id: 1,
      type: "danger",
      message: "Test message"
    })
  })
})
