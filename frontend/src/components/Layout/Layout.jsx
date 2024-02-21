import React from "react";
import styles from "./Layout.module.css"

import Logo from "../../assets/icon/logo.svg"
import Cat from "../../assets/img/Cat.png"

const Layout = ({ children }) => {
  return (
    <div className={styles.start_page}>
      <div className={styles.page_left}>
        <div className={styles.logo}>
          <img src={Logo} alt="logo" />
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
