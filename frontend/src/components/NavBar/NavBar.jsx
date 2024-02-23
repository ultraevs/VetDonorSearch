import React from "react";
import styles from "./NavBar.module.css";

import { NavLink } from "react-router-dom"

import Danger from "../../assets/img/Danger.svg"

import classNames from "class-names";

const NavBar = () => {
  const clickForDanger = () => {
    window.localStorage.setItem("isDanger")
  }
  return (
    <div className={styles.navbar_block}>
      {/* <div className={styles.navbar_danger}>
        <img src={Danger} alt="danger" />
      </div> */}
      <div className={styles.navbar}>
      <NavLink className={({isActive}) => isActive ? classNames(styles.navbar_item, styles.active) : styles.navbar_item} to="/Main">главная</NavLink>
      <NavLink className={({isActive}) => isActive ? classNames(styles.navbar_item, styles.active) : styles.navbar_item} to="/Where">где сдать?</NavLink>
      {
        window.localStorage.getItem("isAuth")
        ?
        <NavLink className={({isActive}) => isActive ? classNames(styles.navbar_item, styles.active) : styles.navbar_item} to="/Profile">профиль</NavLink>
        :
        <NavLink className={({isActive}) => isActive ? classNames(styles.navbar_item, styles.active) : styles.navbar_item} to="/Auth">войти</NavLink>
      }
      </div>
    </div>
  );
};

export default NavBar;
