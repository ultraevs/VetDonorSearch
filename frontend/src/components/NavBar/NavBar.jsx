import React from "react";
import styles from "./NavBar.module.css";

import { NavLink } from "react-router-dom"

import classNames from "class-names";

const NavBar = () => {
  return (
    <div className={styles.navbar}>
      <NavLink className={({isActive}) => isActive ? classNames(styles.navbar_item, styles.active) : styles.navbar_item} to="/Main">главная</NavLink>
      <NavLink className={({isActive}) => isActive ? classNames(styles.navbar_item, styles.active) : styles.navbar_item} to="/Where">где сдать?</NavLink>
      <NavLink className={({isActive}) => isActive ? classNames(styles.navbar_item, styles.active) : styles.navbar_item} to="/Auth">войти</NavLink>
    </div>
  );
};

export default NavBar;
