// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
import { createSelector } from "reselect"

export const getServerInfo = state => state.browser.serverInfo

export const hasServerPublicDomain = createSelector(
  getServerInfo,
  serverInfo => Boolean(serverInfo.info && serverInfo.info.domains && serverInfo.info.domains.length),
)
