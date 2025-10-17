import React from "react"
import { connect } from "react-redux"
import { Dropdown } from "react-bootstrap"
import DeleteObjectConfirmModal from "./DeleteObjectConfirmModal"
import * as actions from "./actions"

export class PrefixActions extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      showDeleteConfirmation: false,
    }
  }
  handleDownload(e) {
    e.preventDefault()
    const { object, downloadPrefix } = this.props
    downloadPrefix(object.name)
  }
  deleteObject() {
    const { object, deleteObject } = this.props
    deleteObject(object.name)
  }
  showDeleteConfirmModal(e) {
    e.preventDefault()
    this.setState({ showDeleteConfirmation: true })
  }
  hideDeleteConfirmModal() {
    this.setState({
      showDeleteConfirmation: false,
    })
  }
  render() {
    const { object, showShareObjectModal, shareObjectName } = this.props
    return (
      <Dropdown id={`obj-actions-${object.name}`}>
        <Dropdown.Toggle noCaret className="fia-toggle" />
        <Dropdown.Menu>
          <a
            href=""
            className="fiad-action"
            title="Download as zip"
            onClick={this.handleDownload.bind(this)}
          >
            <i className="fas fa-cloud-download-alt" />
          </a>
          <a
            href=""
            className="fiad-action"
            title="Delete"
            onClick={this.showDeleteConfirmModal.bind(this)}
          >
            <i className="fas fa-trash-alt" />
          </a>
        </Dropdown.Menu>
        {this.state.showDeleteConfirmation && (
          <DeleteObjectConfirmModal
            deleteObject={this.deleteObject.bind(this)}
            hideDeleteConfirmModal={this.hideDeleteConfirmModal.bind(this)}
          />
        )}
      </Dropdown>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    object: ownProps.object,
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    downloadPrefix: object => dispatch(actions.downloadPrefix(object)),
    deleteObject: (object) => dispatch(actions.deleteObject(object)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(PrefixActions)
