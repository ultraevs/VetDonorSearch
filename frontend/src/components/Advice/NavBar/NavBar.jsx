
import styles from './NavBar.module.css'

const NavBar = () => {
    return (
        <div className={styles.nav}>
            <a href='#' >главная</a>
            <a href='#' >где сдать?</a>
            <a href='#'>войти</a>
        </div>
    );
}

export default NavBar;