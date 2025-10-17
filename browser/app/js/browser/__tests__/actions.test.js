import configureStore from "redux-mock-store"
import thunk from "redux-thunk"
import * as actionsCommon from "../actions"

jest.mock("../../web", () => ({
  StorageInfo: jest.fn(() => {
    return Promise.resolve({
      used: 60
    })
  }),
  ServerInfo: jest.fn(() => {
    return Promise.resolve({
      MinioVersion: "test",
      MinioPlatform: "test",
      MinioRuntime: "test",
      MinioGlobalInfo: "test"
    })
  })
}))

const middlewares = [thunk]
const mockStore = configureStore(middlewares)

describe("Common actions", () => {
  it("creates common/SET_STORAGE_INFO after fetching the storage details ", () => {
    const store = mockStore()
    const expectedActions = [
      { type: "common/SET_STORAGE_INFO", storageInfo: { used: 60 } }
    ]
    return store.dispatch(actionsCommon.fetchStorageInfo()).then(() => {
      const actions = store.getActions()
      expect(actions).toEqual(expectedActions)
    })
  })

  it("creates common/SET_SERVER_INFO after fetching the server details", () => {
    const store = mockStore()
    const expectedActions = [
      {
        type: "common/SET_SERVER_INFO",
        serverInfo: {
          version: "test",
          platform: "test",
          runtime: "test",
          info: "test"
        }
      }
    ]
    return store.dispatch(actionsCommon.fetchServerInfo()).then(() => {
      const actions = store.getActions()
      expect(actions).toEqual(expectedActions)
    })
  })
})
