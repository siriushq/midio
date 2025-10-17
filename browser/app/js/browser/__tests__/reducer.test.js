import reducer from "../reducer"
import * as actionsCommon from "../actions"

describe("common reducer", () => {
  it("should return the initial state", () => {
    expect(reducer(undefined, {})).toEqual({
      sidebarOpen: false,
      storageInfo: {used: 0},
      serverInfo: {}
    })
  })

  it("should handle TOGGLE_SIDEBAR", () => {
    expect(
      reducer(
        { sidebarOpen: false },
        {
          type: actionsCommon.TOGGLE_SIDEBAR
        }
      )
    ).toEqual({
      sidebarOpen: true
    })
  })

  it("should handle CLOSE_SIDEBAR", () => {
    expect(
      reducer(
        { sidebarOpen: true },
        {
          type: actionsCommon.CLOSE_SIDEBAR
        }
      )
    ).toEqual({
      sidebarOpen: false
    })
  })

  it("should handle SET_STORAGE_INFO", () => {
    expect(
      reducer(
        {},
        {
          type: actionsCommon.SET_STORAGE_INFO,
          storageInfo: { }
        }
      )
    ).toEqual({
      storageInfo: { }
    })
  })

  it("should handle SET_SERVER_INFO", () => {
    expect(
      reducer(undefined, {
        type: actionsCommon.SET_SERVER_INFO,
        serverInfo: {
          version: "test",
          platform: "test",
          runtime: "test",
          info: "test"
        }
      }).serverInfo
    ).toEqual({
      version: "test",
      platform: "test",
      runtime: "test",
      info: "test"
    })
  })
})
