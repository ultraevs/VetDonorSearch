import React from "react"; 
import styles from "./NavBar.module.css"; 
 
import { NavLink } from "react-router-dom"; 

import classNames from "class-names"; 
 
import useProfileData from "../../pages/Main/http"; 
 
const NavBar = (title) => { 
  const profileData = useProfileData(); 
  return ( 
    <div className={styles.navbar_block}>  
      <div className={styles.navbar}> 
        <NavLink 
          className={({ isActive }) => 
            isActive 
              ? classNames(styles.navbar_item, styles.active) 
              : styles.navbar_item 
          } 
          to="/" 
        > 
          главная 
        </NavLink> 
        <NavLink 
          className={({ isActive }) => 
            isActive 
              ? classNames(styles.navbar_item, styles.active) 
              : styles.navbar_item 
          } 
          to="/Where" 
        > 
          где сдать? 
        </NavLink> 
        {window.localStorage.getItem("isAuth") ? ( 
          profileData?.type === "user" ? ( 
            <NavLink 
              className={({ isActive }) => 
                isActive 
                  ? classNames(styles.navbar_item, styles.active) 
                  : styles.navbar_item 
              } 
              to="/Profile" 
            > 
              {profileData?.name ? profileData?.name : "Профиль"} 
            </NavLink> 
          ) : ( 
            <NavLink 
              className={({ isActive }) => 
                isActive 
                  ? classNames(styles.navbar_item, styles.active) 
                  : styles.navbar_item 
              } 
              to="/ProfileClinic" 
            > 
              {profileData?.name ? profileData?.name : "Профиль"} 
            </NavLink> 
          ) 
        ) : ( 
          <NavLink 
            className={({ isActive }) => 
              isActive 
                ? classNames(styles.navbar_item, styles.active) 
                : styles.navbar_item 
            } 
            to="/Auth" 
          > 
            войти 
          </NavLink> 
        )} 
      </div> 
    </div> 
  ); 
}; 
 
export default NavBar;