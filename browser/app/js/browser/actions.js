import web from "../web"

export const TOGGLE_SIDEBAR = "common/TOGGLE_SIDEBAR"
export const CLOSE_SIDEBAR = "common/CLOSE_SIDEBAR"
export const SET_STORAGE_INFO = "common/SET_STORAGE_INFO"
export const SET_SERVER_INFO = "common/SET_SERVER_INFO"

export const toggleSidebar = () => ({
  type: TOGGLE_SIDEBAR
})

export const closeSidebar = () => ({
  type: CLOSE_SIDEBAR
})

export const fetchStorageInfo = () => {
  return function(dispatch) {
    return web.StorageInfo().then(res => {
      const storageInfo = {
        used: res.used
      }
      dispatch(setStorageInfo(storageInfo))
    })
  }
}

export const setStorageInfo = storageInfo => ({
  type: SET_STORAGE_INFO,
  storageInfo
})

export const fetchServerInfo = () => {
  return function(dispatch) {
    return web.ServerInfo().then(res => {
      const serverInfo = {
        version: res.MinioVersion,
        platform: res.MinioPlatform,
        runtime: res.MinioRuntime,
        info: res.MinioGlobalInfo,
        userInfo: res.MinioUserInfo
      }
      dispatch(setServerInfo(serverInfo))
    })
  }
}

export const setServerInfo = serverInfo => ({
  type: SET_SERVER_INFO,
  serverInfo
})
