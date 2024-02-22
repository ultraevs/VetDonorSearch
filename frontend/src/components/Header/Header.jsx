import React from 'react'
import styles from "./Header.module.css";

import NavBar from "../NavBar/NavBar";
import Logo from "../../assets/img/LOGO.svg"

const Header = () => {
  return (
    <div className={styles.header}>
      <div className={styles.logo}>
        <img src={Logo} alt="Logo" />
      </div>
      <NavBar />
    </div>
  )
}

export default Header
