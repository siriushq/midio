import React from "react"
import { Modal, ModalHeader, ModalBody } from "react-bootstrap"

class PreviewObjectModal extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      url: "",
    }
  }

  componentDidMount() {
    this.props.getObjectURL(this.props.object.name, (url) => {
      this.setState({
        url: url,
      })
    })
  }

  render() {
    const { hidePreviewModal } = this.props
    return (
      <Modal
        show={true}
        animation={false}
        onHide={hidePreviewModal}
        bsSize="large"
      >
        <ModalHeader>Preview</ModalHeader>
        <ModalBody>
          <div className="input-group">
            {this.state.url && (
              <object data={this.state.url} style={{ display: "block", width: "100%" }}>
                <h3 style={{ textAlign: "center", display: "block", width: "100%" }}>
                  Do not have read permissions to preview "{this.props.object.name}"
                </h3>
              </object>
            )}
          </div>
        </ModalBody>
        <div className="modal-footer">
          {
            <button className="btn btn-link" onClick={hidePreviewModal}>
              Cancel
            </button>
          }
        </div>
      </Modal>
    )
  }
}
export default PreviewObjectModal
