import React from "react";
import styles from "./Layout.module.css"
import { NavLink } from 'react-router-dom'

import Close_layout from './../../assets/icon/layout_close.svg'

const Layout = ({ children }) => {
  return (
    <div className={styles.start_page}>
      <div className={styles.page_left}>
        <div className={styles.logo}>
          <NavLink to='/' className={styles.logo_icon}>
            <img src={Close_layout} alt="logo" />
          </NavLink>
        </div>
        {children}
      </div>
    </div>
  );
};

export default Layout;
