import React from "react"
import { Route, Switch, Redirect } from "react-router-dom"
import Browser from "./browser/Browser"
import Login from "./browser/Login"
import OpenIDLogin from "./browser/OpenIDLogin"
import web from "./web"

export const App = () => {
  return (
    <Switch>
      <Route path={"/login/openid"} component={OpenIDLogin} />
      <Route path={"/login"} component={Login} />
      <Route path={"/:bucket?/*"} component={Browser} />
    </Switch>
  )
}

export default App
