import configureStore from "redux-mock-store"
import thunk from "redux-thunk"
import * as actionsAlert from "../actions"

const middlewares = [thunk]
const mockStore = configureStore(middlewares)

jest.useFakeTimers()

describe("Alert actions", () => {
  it("creates alert/SET action", () => {
    const store = mockStore()
    const expectedActions = [
      {
        type: "alert/SET",
        alert: { id: 0, message: "Test alert", type: "danger" }
      }
    ]
    store.dispatch(actionsAlert.set({ message: "Test alert", type: "danger" }))
    const actions = store.getActions()
    expect(actions).toEqual(expectedActions)
  })

  it("creates alert/CLEAR action for non danger alerts", () => {
    const store = mockStore()
    const expectedActions = [
      {
        type: "alert/SET",
        alert: { id: 1, message: "Test alert" }
      },
      {
        type: "alert/CLEAR",
        alert: { id: 1 }
      }
    ]
    store.dispatch(actionsAlert.set({ message: "Test alert" }))
    jest.runAllTimers()
    const actions = store.getActions()
    expect(actions).toEqual(expectedActions)
  })

  it("creates alert/CLEAR action directly", () => {
    const store = mockStore()
    const expectedActions = [
      {
        type: "alert/CLEAR"
      }
    ]
    store.dispatch(actionsAlert.clear())
    const actions = store.getActions()
    expect(actions).toEqual(expectedActions)
  })
})
