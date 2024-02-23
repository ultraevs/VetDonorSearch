import React from "react";
import styles from "./Header.module.css"

import NavBar from "../NavBar/NavBar";
import Logo from './../../assets/img/LOGO.svg'

const Header = () => {
    return ( 
        <div className={styles.header_block}>
            <img src={Logo} alt="Логотип" />
            <NavBar />
        </div>
    );
}
 
export default Header;