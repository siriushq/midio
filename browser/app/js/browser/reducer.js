import * as actionsCommon from "./actions"

export default (
  state = {
    sidebarOpen: false,
    storageInfo: {used: 0},
    serverInfo: {}
  },
  action
) => {
  switch (action.type) {
    case actionsCommon.TOGGLE_SIDEBAR:
      return Object.assign({}, state, {
        sidebarOpen: !state.sidebarOpen
      })
    case actionsCommon.CLOSE_SIDEBAR:
      return Object.assign({}, state, {
        sidebarOpen: false
      })
    case actionsCommon.SET_STORAGE_INFO:
      return Object.assign({}, state, {
        storageInfo: action.storageInfo
      })
    case actionsCommon.SET_SERVER_INFO:
      return { ...state, serverInfo: action.serverInfo }
    default:
      return state
  }
}
