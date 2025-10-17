import { combineReducers } from "redux"
import browser from "./browser/reducer"
import alert from "./alert/reducer"
import buckets from "./buckets/reducer"
import objects from "./objects/reducer"
import uploads from "./uploads/reducer"

const rootReducer = combineReducers({
  browser,
  alert,
  buckets,
  objects,
  uploads
})

export default rootReducer
