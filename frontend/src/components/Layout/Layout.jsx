import React from "react";
import styles from "./Layout.module.css"
import { NavLink } from 'react-router-dom'

import Logo from "../../assets/img/LOGO.svg"
import Cat from "../../assets/img/Cat.png"

const Layout = ({ children }) => {
  return (
    <div className={styles.start_page}>
      <div className={styles.page_left}>
        <div className={styles.logo}>
          <NavLink to='/Main'>
            <img src={Logo} alt="logo" />
          </NavLink>
        </div>
        {children}
      </div>
      <div className={styles.page_right}>
        <div className={styles.img_cat}>
          <img src={Cat} alt="cat" />
        </div>
      </div>
    </div>
  );
};

export default Layout;
