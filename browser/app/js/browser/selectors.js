import { createSelector } from "reselect"

export const getServerInfo = state => state.browser.serverInfo

export const hasServerPublicDomain = createSelector(
  getServerInfo,
  serverInfo => Boolean(serverInfo.info && serverInfo.info.domains && serverInfo.info.domains.length),
)
