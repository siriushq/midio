import React from 'react'
import connect from 'react-redux/lib/components/connect'

import Tooltip from 'react-bootstrap/lib/Tooltip'
import OverlayTrigger from 'react-bootstrap/lib/OverlayTrigger'

let BrowserUpdate = ({latestUiVersion}) => {
  // Don't show an update if we're already updated!
  if (latestUiVersion === currentUiVersion) return ( <noscript></noscript> )

  return (
    <li className="hidden-xs hidden-sm">
      <a href="">
        <OverlayTrigger placement="left" overlay={ <Tooltip id="tt-version-update">
                                                     New update available. Click to refresh.
                                                   </Tooltip> }> <i className="fas fa-sync"></i> </OverlayTrigger>
      </a>
    </li>
  )
}

export default connect(state => {
  return {
    latestUiVersion: state.latestUiVersion
  }
})(BrowserUpdate)
