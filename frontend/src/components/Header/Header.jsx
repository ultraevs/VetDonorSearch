import React from "react";
import { NavLink } from "react-router-dom"

import styles from "./Header.module.css"

import NavBar from "../NavBar/NavBar";
import Logo from './../../assets/img/LOGO.svg'

const Header = () => {
    return (
        <div className={styles.header_block}>
            <NavLink to='/Main'>
                <img src={Logo} alt="Логотип" />
            </NavLink>
            <NavBar />
        </div>
    );
}

export default Header;