// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
import createHistory from "history/createBrowserHistory"
import { minioBrowserPrefix } from "./constants"

const history = createHistory({
  basename: minioBrowserPrefix
})

export default history
