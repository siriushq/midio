import React from "react"
import { Modal, ModalBody } from "react-bootstrap"

let ConfirmModal = ({
  baseClass,
  icon,
  text,
  sub,
  okText,
  cancelText,
  okHandler,
  cancelHandler,
  show
}) => {
  return (
    <Modal
      bsSize="small"
      animation={false}
      show={show}
      className={"modal-confirm " + (baseClass || "")}
    >
      <ModalBody>
        <div className="mc-icon">
          <i className={icon} />
        </div>
        <div className="mc-text">{text}</div>
        <div className="mc-sub">{sub}</div>
      </ModalBody>
      <div className="modal-footer">
        <button className="btn btn-danger" onClick={okHandler}>
          {okText}
        </button>
        <button className="btn btn-link" onClick={cancelHandler}>
          {cancelText}
        </button>
      </div>
    </Modal>
  )
}

export default ConfirmModal
