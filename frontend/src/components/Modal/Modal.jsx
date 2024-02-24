import React from "react";
import classNames from 'class-names'

import styles from './Modal.module.css'

const Modal = ({active, setActive, children}) => {
    return ( 
        <div className={active ? classNames(styles.modal, styles.active) : styles.modal} onClick={() => setActive(false)}>
            <div className={active ? classNames(styles.modal_content, styles.active) : styles.modal_content} onClick={e => e.stopPropagation()}>
                {children}
            </div>
        </div>
    );
}
 
export default Modal;