import createHistory from "history/createBrowserHistory"
import { minioBrowserPrefix } from "./constants"

const history = createHistory({
  basename: minioBrowserPrefix
})

export default history
